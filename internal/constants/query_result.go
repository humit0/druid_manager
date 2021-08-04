package constants

// 일반 쿼리 결과
type ServerListItem struct {
	Server     string `json:"server"`
	ServerType string `json:"server_type"`
}

type TotalSegmentsItem struct {
	Datasource  string  `json:"datasource"`
	TotalSize   int64   `json:"total_size"`
	AvgSize     float64 `json:"avg_size"`
	AvgNumRows  float64 `json:"avg_num_rows"`
	NumSegments int64   `json:"num_segments"`
}

// 템플릿 쿼리 결과
type RollupRatioByDatasourceItem struct {
	Time           string  `json:"__time,omitempty"`
	RollupRatio    float64 `json:"rollup_ratio"`
	DatasourceName string  `json:"datasource_name"`
}

type DataCntByDatasourceItem struct {
	Time           string `json:"__time,omitempty"`
	DataCnt        int64  `json:"data_cnt"`
	DatasourceName string `json:"datasource_name"`
}
