package druid_coordinator

import (
	"github.com/sirupsen/logrus"
)

var (
	datasourceEntry = baseEntry.WithFields(logrus.Fields{"type": "datasource"})
)

// GetDatasourceList 함수는 해당 druid 서버에 있는 전체 데이터 소스 목록을 반환합니다.
func (coordinatorService *CoordinatorServiceImp) GetDatasourceList() []string {
	var result []string

	datasourceEntry.Debug("Get datasource list")
	coordinatorService.DruidClient.SendRequest("GET", "coordinator", "/druid/coordinator/v1/datasources", nil, &result)
	return result
}
