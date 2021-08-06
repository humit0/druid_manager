package druid_coordinator

import (
	"fmt"

	"github.com/humit0/druid_manager/internal/druid"
	"github.com/sirupsen/logrus"
)

var (
	compactionEntry = baseEntry.WithFields(logrus.Fields{"type": "compaction"})
)

// GetAllCompactionConfiguration 함수는 모든 데이터 소스에 대한 compaction 설정을 가져옵니다.
func GetAllCompactionConfiguration(druidClient *druid.DruidClient) interface{} {
	var result interface{}
	compactionEntry.Debug("Get all datasource compaction configuration")

	druidClient.SendRequest("GET", "coordinator", "/druid/coordinator/v1/config/compaction", nil, &result)

	return result
}

// GetCompactionConfigurationByDatasource 함수는 특정 데이터 소스에 대한 compaction 설정을 가져옵니다.
func GetCompactionConfigurationByDatasource(druidClient *druid.DruidClient, datasourceName string) interface{} {
	var result interface{}
	compactionEntry.Debugf("Get datasource(%s) compaction configuration", datasourceName)

	druidClient.SendRequest("GET", "coordinator", fmt.Sprintf("/druid/coordinator/v1/config/compaction/%s", datasourceName), nil, &result)

	return result
}
