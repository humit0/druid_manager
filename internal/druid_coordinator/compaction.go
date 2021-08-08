package druid_coordinator

import (
	"bytes"
	"net/url"

	"github.com/sirupsen/logrus"
)

var (
	compactionEntry = baseEntry.WithFields(logrus.Fields{"type": "compaction"})
)

// GetAllCompactionConfiguration 함수는 모든 데이터 소스에 대한 compaction 설정을 가져옵니다.
func (coordinatorService *CoordinatorServiceImp) GetAllCompactionConfiguration() interface{} {
	var result interface{}
	compactionEntry.Debug("Get all datasource compaction configuration")

	coordinatorService.DruidClient.SendRequest("GET", "coordinator", "/druid/coordinator/v1/config/compaction", nil, &result)

	return result
}

// GetCompactionConfigurationByDatasource 함수는 특정 데이터 소스에 대한 compaction 설정을 가져옵니다.
func (coordinatorService *CoordinatorServiceImp) GetCompactionConfigurationByDatasource(datasourceName string) interface{} {
	var result interface{}
	compactionEntry.Debugf("Get datasource(%s) compaction configuration", datasourceName)

	urlBuff := bytes.NewBufferString("/druid/coordinaotor/v1/config/compaction/")
	urlBuff.WriteString(url.PathEscape(datasourceName))

	coordinatorService.DruidClient.SendRequest("GET", "coordinator", urlBuff.String(), nil, &result)

	return result
}
