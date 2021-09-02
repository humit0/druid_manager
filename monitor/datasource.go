package monitor

import (
	"sync"
	"time"

	"github.com/humit0/druid_manager/internal/constants"
	"github.com/humit0/druid_manager/internal/models"
	"github.com/sirupsen/logrus"
)

var datasourceEntry = baseEntry.WithFields(logrus.Fields{"monitor": "datasource"})

type DatasourceLastUpdatedTimeResult struct {
	Datasource      string
	LastUpdatedTime time.Time
}

type NotUpdatedDatasourceResult struct {
	Datasource      string
	Owner           string
	LastUpdatedTime time.Time
}

type MonitorDatasource interface {
	// getLastUpdatedTime 함수는 입력으로 받은 데이터 소스 목록을 바탕으로 최종 ingestion 시간을 채널로 전달합니다.
	getLastUpdatedTime([]string, chan<- DatasourceLastUpdatedTimeResult)
	// GetNotUpdatedDatasources 함수는 데이터 소스 별로 기준 시간과 비교하여 업데이트가 되지 않은 데이터 소스 정보 목록을 반환합니다.
	GetNotUpdatedDatasources([]models.Datasource, time.Time) map[string][]string
}

func (monitor *MonitorImp) getLastUpdatedTime(datasources []string, result chan<- DatasourceLastUpdatedTimeResult) {
	wg := new(sync.WaitGroup)
	for _, datasource := range datasources {
		wg.Add(1)
		go (func(datasource string) {
			var queryResult []struct {
				Timestamp string `json:"timestamp"`
				Result    struct {
					MinTime time.Time `json:"minTime"`
					MaxTime time.Time `json:"maxTime"`
				} `json:"result"`
			}
			defer wg.Done()
			monitor.DruidClient.BrokerSvc.SendNativeQuery(constants.GetTimeBoundaryQuery(datasource, "maxTime"), &queryResult)
			datasourceEntry.Debugf("Last updated time of %s => %+v", datasource, queryResult[0].Result.MaxTime)
			result <- DatasourceLastUpdatedTimeResult{Datasource: datasource, LastUpdatedTime: queryResult[0].Result.MaxTime}
		})(datasource)
	}
	go (func() {
		wg.Wait()
		close(result)
	})()
}

func (monitor *MonitorImp) GetNotUpdatedDatasources(datasources []models.Datasource, baseTime time.Time) []NotUpdatedDatasourceResult {
	lastUpdatedInfos := make(chan DatasourceLastUpdatedTimeResult, 10)
	result := make([]NotUpdatedDatasourceResult, 0, 10)
	datasourceLastupdates := make(map[string]time.Time)

	datasourceNames := make([]string, 0, len(datasources))

	for _, datasource := range datasources {
		datasourceNames = append(datasourceNames, datasource.Name)
	}

	monitor.getLastUpdatedTime(datasourceNames, lastUpdatedInfos)

	for lastupdateInfo := range lastUpdatedInfos {
		datasourceLastupdates[lastupdateInfo.Datasource] = lastupdateInfo.LastUpdatedTime
	}

	for _, datasource := range datasources {
		ts, exists := datasourceLastupdates[datasource.Name]
		if !exists {
			datasourceEntry.Errorf("No datasources information from %s", datasource.Name)
			continue
		}
		if baseTime.Sub(ts.AddDate(0, 0, int(datasource.Interval))) > 0 {
			owner := datasource.Owner.String
			result = append(result, NotUpdatedDatasourceResult{Owner: owner, Datasource: datasource.Name, LastUpdatedTime: ts})
		}
	}

	return result
}
