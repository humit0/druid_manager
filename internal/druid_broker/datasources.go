package druid_broker

import (
	"bytes"
	"net/url"
)

var (
	datasourceEntry = baseEntry.WithField("type", "datasource")
)

type DatasourceResponseType struct {
	Dimensions []string
	Metrics    []string
}

// getDatasourceColumnsInfo 함수는 해당 데이터 소스에 속한 컬럼명 정보를 반환합니다. (차원, 측정 값 순서)
func (brokerService *BrokerServiceImp) getDatasourceColumnsInfo(datasourceName string) ([]string, []string) {
	response := DatasourceResponseType{}

	urlBuff := bytes.NewBufferString("/druid/v2/datsources/")
	urlBuff.WriteString(url.PathEscape(datasourceName))

	brokerService.DruidClient.SendRequest("GET", "broker", urlBuff.String(), nil, &response)

	datasourceEntry.Debugf("dimension: %v", response.Dimensions)
	datasourceEntry.Debugf("metric: %v", response.Metrics)

	return response.Dimensions, response.Metrics
}

// GetDatasourceDimensions 함수는 해당 데이터 소스에 속한 차원 컬럼명을 반환합니다.
func (brokerService *BrokerServiceImp) GetDatasourceDimensions(datasourceName string) []string {
	result, _ := brokerService.getDatasourceColumnsInfo(datasourceName)

	return result
}

// GetDatasourceMetrics 함수는 해당 데이터 소스에 속한 측정 값 컬럼명을 반환합니다.
func (brokerService *BrokerServiceImp) GetDatasourceMetrics(datasourceName string) []string {
	_, result := brokerService.getDatasourceColumnsInfo(datasourceName)

	return result
}
