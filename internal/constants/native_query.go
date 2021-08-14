package constants

import "encoding/json"

type TimeBoundaryQueryType struct {
	QueryType  string `json:"queryType"`
	Datasource string `json:"dataSource"`
	Bound      string `json:"bound"`
}

// GetTimeBoundaryQuery 함수는 주어진 `datasource`의 최소 날짜와 최대 날짜를 알아내는 쿼리를 반환합니다.
func GetTimeBoundaryQuery(datasource string, bound string) string {
	query := TimeBoundaryQueryType{
		QueryType:  "timeBoundary",
		Datasource: datasource,
		Bound:      bound,
	}

	res, err := json.Marshal(query)

	if err != nil {
		entry.WithField("queryType", "timeBoundary").Fatalf("Cannot create json object %v", err)
	}

	return string(res)
}
