package druid_broker

import (
	"fmt"

	"github.com/humit0/druid_manager/internal/druid"
)

type DatasourceResponseType struct {
	Dimensions []string
	Metrics    []string
}

// 해당 데이터 소스에 속한 컬럼명 정보를 반환합니다. (차원, 측정 값 순서)
func getDatasourceColumnsInfo(druidClient *druid.DruidClient, datasourceName string) ([]string, []string) {
	response := DatasourceResponseType{}

	druidClient.SendRequest("GET", "broker", fmt.Sprintf("/druid/v2/datasources/%s", datasourceName), nil, &response)

	return response.Dimensions, response.Metrics
}

// 해당 데이터 소스에 속한 차원 컬럼명을 반환합니다.
func GetDatasourceDimensions(druidClient *druid.DruidClient, datasourceName string) []string {
	result, _ := getDatasourceColumnsInfo(druidClient, datasourceName)

	return result
}

// 해당 데이터 소스에 속한 측정 값 컬럼명을 반환합니다.
func GetDatasourceMetrics(druidClient *druid.DruidClient, datasourceName string) []string {
	_, result := getDatasourceColumnsInfo(druidClient, datasourceName)

	return result
}
