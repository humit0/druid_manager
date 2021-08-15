package constants

// ServerListItem 구조체는 일반 쿼리 결과를 나타냅니다.
type ServerListItem struct {
	Server     string `json:"server"`
	ServerType string `json:"server_type"`
}

// TotalSegmentsItem 구조체는 세그먼트 관련 쿼리 결과를 나타냅니다.
type TotalSegmentsItem struct {
	Datasource  string  `json:"datasource"`
	TotalSize   int64   `json:"total_size"`
	AvgSize     float64 `json:"avg_size"`
	AvgNumRows  float64 `json:"avg_num_rows"`
	NumSegments int64   `json:"num_segments"`
}

// 템플릿 쿼리 결과

// RollupRatioByDatasourceItem 구조체는 롤업 비율 쿼리 결과를 나타냅니다.
type RollupRatioByDatasourceItem struct {
	Time           string  `json:"__time,omitempty"`
	RollupRatio    float64 `json:"rollup_ratio"`
	DatasourceName string  `json:"datasource_name"`
}

// DataCntByDatasourceItem 구조체는 총 데이터 수 결과를 나타냅니다.
type DataCntByDatasourceItem struct {
	Time           string `json:"__time,omitempty"`
	DataCnt        int64  `json:"data_cnt"`
	DatasourceName string `json:"datasource_name"`
}

// Native 쿼리 결과

// TimeBoundaryItem 구조체는 TimeBoundary 쿼리 결과를 나타냅니다.
type TimeBoundaryItem struct {
	Timestamp string `json:"timestamp"`
	Result    struct {
		MinTime string `json:"minTime,omitempty"`
		MaxTime string `json:"maxTime,omitempty"`
	} `json:"result"`
}
