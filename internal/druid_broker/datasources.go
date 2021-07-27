package druid_broker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

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

// Druid SQL 쿼리를 실행하여 결과를 반환합니다.
func SendSQLQuery(druidClient *druid.DruidClient, sqlQuery string) []map[string]interface{} {
	var result []map[string]interface{}

	jsonBody := make(map[string]interface{})
	jsonBody["sql"] = sqlQuery
	jsonBody["resultFormat"] = "object"

	body, err := json.Marshal(jsonBody)
	if err != nil {
		log.Fatalf("Cannot serialize json (%v)", jsonBody)
	}

	druidClient.SendRequest("POST", "broker", "/druid/v2/sql/", bytes.NewBuffer(body), &result)

	return result
}

// Druid native 쿼리를 실행하여 결과를 반환합니다.
func SendNativeQuery(druidClient *druid.DruidClient, nativeQuery string) []map[string]interface{} {
	var result []map[string]interface{}
	druidClient.SendRequest("POST", "broker", "/druid/v2/", nil, &result)

	return result
}
