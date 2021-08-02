package constants_test

import (
	"testing"

	"github.com/humit0/druid_manager/internal/constants"
)

func TestGetRollupRatioByDatasourceQuery(t *testing.T) {
	query1 := constants.GetRollupRatioByDatasourceQuery("count", "datasource_1", false)
	if query1 != `SELECT
  SUM("count") * 1.0 / COUNT(*) AS rollup_ratio
FROM "datasource_1"` {
		t.Error("Wrong!")
	}
	query2 := constants.GetRollupRatioByDatasourceQuery("count", "datasource_1", true)
	if query2 != `SELECT
  __time,
  SUM("count") * 1.0 / COUNT(*) AS rollup_ratio
FROM "datasource_1"
GROUP BY 1` {
		t.Error("Wrong!")
	}
}

func TestGetDataCntByDatasourceQuery(t *testing.T) {
	query1 := constants.GetDataCntByDatasourceQuery("count", "datasource_2", false)
	if query1 != `SELECT
  SUM("count") AS data_cnt
FROM "datasource_2"` {
		t.Error("Wrong!")
	}
	query2 := constants.GetDataCntByDatasourceQuery("count", "datasource_2", true)
	if query2 != `SELECT
  __time,
  SUM("count") AS data_cnt
FROM "datasource_2"
GROUP BY 1` {
		t.Error("Wrong!")
	}
}
