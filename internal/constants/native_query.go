package constants

import "encoding/json"

type TimeBoundaryQueryType struct {
	QueryType  string `json:"queryType"`
	Datasource string `json:"dataSource"`
	Bound      string `json:"bound,omitempty"`
}

type SegmentMetadataQueryType struct {
	QueryType  string   `json:"queryType"`
	Datasource string   `json:"dataSource"`
	Intervals  []string `json:"intervals,omitempty"`
	ToInclude  *struct {
		Type    string   `json:"type"`
		Columns []string `json:"columns,omitempty"`
	} `json:"toInclude,omitempty"`
	Merge                  bool     `json:"merge,omitempty"`
	AnalysisTypes          []string `json:"analysisTypes,omitempty"`
	LenientAggregatorMerge bool     `json:"lenientAggregatorMerge,omitempty"`
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

// GetSegmentMetadata 함수는 주어진 `datasource`의 세그먼트 관련 정보를 알아내는 쿼리를 반환합니다.
func GetSegmentMetadataQuery(query SegmentMetadataQueryType) string {
	query.QueryType = "segmentMetadata"

	res, err := json.Marshal(query)

	if err != nil {
		entry.WithField("queryType", "segmentMetadata").Fatalf("Cannot create json object %v", err)
	}

	return string(res)
}
