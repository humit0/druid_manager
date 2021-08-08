package druid_broker

import (
	"bytes"
	"encoding/json"
)

var (
	queryEntry = baseEntry.WithField("type", "query")
)

// SendSQLQuery 함수는 Druid SQL 쿼리를 실행하여 결과를 반환합니다.
func (brokerService BrokerServiceImp) SendSQLQuery(sqlQuery string, queryResult interface{}) {
	jsonBody := make(map[string]interface{})
	jsonBody["query"] = sqlQuery
	jsonBody["resultFormat"] = "object"

	body, err := json.Marshal(jsonBody)
	if err != nil {
		queryEntry.Fatalf("Cannot serialize json (%v)", jsonBody)
	}

	queryEntry.Debugf("Sending sql query (%s)", body)

	brokerService.DruidClient.SendRequest("POST", "broker", "/druid/v2/sql/", bytes.NewBuffer(body), &queryResult)
}

// SendNativeQuery 함수는 Druid native 쿼리를 실행하여 결과를 반환합니다.
func (brokerService BrokerServiceImp) SendNativeQuery(nativeQuery string, queryResult interface{}) {
	queryEntry.Debugf("Sending native query (%s)", nativeQuery)

	brokerService.DruidClient.SendRequest("POST", "broker", "/druid/v2/", bytes.NewBuffer([]byte(nativeQuery)), &queryResult)
}
