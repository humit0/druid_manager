package constants_test

import (
	"encoding/json"
	"testing"

	"github.com/humit0/druid_manager/internal/constants"
)

func TestServerListItemParse(t *testing.T) {
	var result constants.ServerListItem
	valid := `{"server": "druid_host:80", "server_type": "broker"}`

	err := json.Unmarshal([]byte(valid), &result)
	if err != nil {
		t.Error("Failed to parse json")
	}

	if result.Server != "druid_host:80" || result.ServerType != "broker" {
		t.Error("Failed to parse data")
	}

	no_server_type := `{"server": "druid_host:8888"}`

	result = constants.ServerListItem{}

	err = json.Unmarshal([]byte(no_server_type), &result)

	if err != nil {
		t.Error("Failed to parse json")
	}

	if result.Server != "druid_host:8888" || result.ServerType != "" {
		t.Errorf("Failed to parse data")
	}
}

func TestTotalSegmentsItemParse(t *testing.T) {
	var result constants.TotalSegmentsItem
	valid := `{"datasource": "d1", "total_size": 4, "avg_size": 0.5, "avg_num_rows": 0.8, "num_segments": 20}`

	err := json.Unmarshal([]byte(valid), &result)
	if err != nil {
		t.Error("Failed to parse json")
	}
	if result.Datasource != "d1" || result.TotalSize != 4 || result.AvgSize != 0.5 || result.AvgNumRows != 0.8 || result.NumSegments != 20 {
		t.Error("Failed to parse data")
	}

	result = constants.TotalSegmentsItem{}
	with_additional := `{"datasource": "d1", "total_size": 4, "avg_size": 0.5, "avg_num_rows": 0.8, "num_segments": 20, "add1": "add_val"}`
	err = json.Unmarshal([]byte(with_additional), &result)
	if err != nil {
		t.Error("Failed to parse json")
	}
	if result.Datasource != "d1" || result.TotalSize != 4 || result.AvgSize != 0.5 || result.AvgNumRows != 0.8 || result.NumSegments != 20 {
		t.Error("Failed to parse data")
	}

	result = constants.TotalSegmentsItem{}
	without_datasource := `{"total_size": 4, "avg_size": 0.5, "avg_num_rows": 0.8, "num_segments": 20}`
	err = json.Unmarshal([]byte(without_datasource), &result)
	if err != nil {
		t.Error("Failed to parse json")
	}
	if result.Datasource != "" || result.TotalSize != 4 || result.AvgSize != 0.5 || result.AvgNumRows != 0.8 || result.NumSegments != 20 {
		t.Error("Failed to parse data")
	}
}

func TestRollupRatioByDatasourceItem(t *testing.T) {
	var result constants.RollupRatioByDatasourceItem
	valid := `{"__time": "2021-01-01T00:00:00Z", "rollup_ratio": 1.23456}`

	err := json.Unmarshal([]byte(valid), &result)
	if err != nil {
		t.Error("Failed to parse json")
	}
	if result.Time != "2021-01-01T00:00:00Z" || result.RollupRatio != 1.23456 {
		t.Error("Failed to parse data")
	}

	result = constants.RollupRatioByDatasourceItem{}
	with_additional := `{"__time": "2021-01-01T00:00:00Z", "rollup_ratio": 1.23456, "add1": "add_val"}`
	err = json.Unmarshal([]byte(with_additional), &result)
	if err != nil {
		t.Error("Failed to parse json")
	}
	if result.Time != "2021-01-01T00:00:00Z" || result.RollupRatio != 1.23456 {
		t.Error("Failed to parse data")
	}

	result = constants.RollupRatioByDatasourceItem{}
	without_time := `{"rollup_ratio": 1.23456}`
	err = json.Unmarshal([]byte(without_time), &result)
	if err != nil {
		t.Error("Failed to parse json")
	}
	if result.Time != "" || result.RollupRatio != 1.23456 {
		t.Error("Failed to parse data")
	}
}

func TestDataCntByDatasourceItem(t *testing.T) {
	var result constants.DataCntByDatasourceItem
	valid := `{"__time": "2021-01-01T00:00:00Z", "data_cnt": 123456}`

	err := json.Unmarshal([]byte(valid), &result)
	if err != nil {
		t.Error("Failed to parse json")
	}
	if result.Time != "2021-01-01T00:00:00Z" || result.DataCnt != 123456 {
		t.Error("Failed to parse data")
	}

	result = constants.DataCntByDatasourceItem{}
	with_additional := `{"__time": "2021-01-01T00:00:00Z", "data_cnt": 123456, "add1": "add_val"}`
	err = json.Unmarshal([]byte(with_additional), &result)
	if err != nil {
		t.Error("Failed to parse json")
	}
	if result.Time != "2021-01-01T00:00:00Z" || result.DataCnt != 123456 {
		t.Error("Failed to parse data")
	}

	result = constants.DataCntByDatasourceItem{}
	without_time := `{"data_cnt": 123456}`
	err = json.Unmarshal([]byte(without_time), &result)
	if err != nil {
		t.Error("Failed to parse json")
	}
	if result.Time != "" || result.DataCnt != 123456 {
		t.Error("Failed to parse data")
	}
}
