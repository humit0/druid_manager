package druid_broker

import (
	"bytes"
	"encoding/json"

	"github.com/humit0/druid_manager/internal/druid"
)

var (
	queryEntry = baseEntry.WithField("type", "query")
)

// Druid SQL 쿼리를 실행하여 결과를 반환합니다.
func SendSQLQuery(druidClient *druid.DruidClient, sqlQuery string, queryResult interface{}) {
	jsonBody := make(map[string]interface{})
	jsonBody["query"] = sqlQuery
	jsonBody["resultFormat"] = "object"

	body, err := json.Marshal(jsonBody)
	if err != nil {
		queryEntry.Fatalf("Cannot serialize json (%v)", jsonBody)
	}

	queryEntry.Debugf("Sending sql query (%s)", body)

	druidClient.SendRequest("POST", "broker", "/druid/v2/sql/", bytes.NewBuffer(body), &queryResult)
}

// Druid native 쿼리를 실행하여 결과를 반환합니다.
func SendNativeQuery(druidClient *druid.DruidClient, nativeQuery string, queryResult interface{}) {
	queryEntry.Debugf("Sending native query (%s)", nativeQuery)

	druidClient.SendRequest("POST", "broker", "/druid/v2/", bytes.NewBuffer([]byte(nativeQuery)), &queryResult)
}
