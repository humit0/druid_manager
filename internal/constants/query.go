package constants

import (
	"bytes"
	"text/template"

	"github.com/sirupsen/logrus"
)

var (
	entry = logrus.WithFields(logrus.Fields{"type": "query"})
)

// 일반 쿼리 목록
const (
	ServerListQuery    = "SELECT server, server_type FROM sys.servers"
	TotalSegmentsQuery = `SELECT
  datasource,
  SUM("size") AS total_size,
  CASE WHEN SUM("size") = 0 THEN 0 ELSE SUM("size") / (COUNT(*) FILTER(WHERE "size" > 0)) END AS avg_size,
  CASE WHEN SUM(num_rows) = 0 THEN 0 ELSE SUM("num_rows") / (COUNT(*) FILTER(WHERE num_rows > 0)) END AS avg_num_rows,
  COUNT(*) AS num_segments
FROM sys.segments
GROUP BY 1
ORDER BY 2 DESC`
)

// template으로 써야하는 쿼리 목록 (private)
const (
	getRollupRatioByDatasourceQueryTpl = `SELECT{{ if .time }}
  __time,{{ end }}
  SUM("{{ .countColName }}") * 1.0 / COUNT(*) AS rollup_ratio
FROM "{{ .datasourceName }}"{{ if .time }}
GROUP BY 1{{ end }}`
	getDataCntByDatasourceQueryTpl = `SELECT{{ if .time }}
  __time,{{ end }}
  SUM("{{ .countColName }}") AS data_cnt
FROM "{{ .datasourceName }}"{{ if .time }}
GROUP BY 1{{ end }}`
)

// templateQuery 함수는 템플릿으로 만든 쿼리를 사용할 수 있는 쿼리로 변환해줍니다.
func templateQuery(queryName string, queryTpl string, val interface{}) string {
	tpl := template.New(queryName)
	tpl, err := tpl.Parse(queryTpl)

	if err != nil {
		entry.WithField("queryName", queryName).Fatalf("Cannot parse template %v", err)
	}

	var result bytes.Buffer

	err1 := tpl.Execute(&result, &val)

	if err1 != nil {
		entry.WithField("queryName", queryName).Fatalf("Cannot execute template %v", err)
	}

	return result.String()
}

// GetRollupRatioByDatasourceQuery 함수는 count 컬럼과 데이터소스 명을 입력하여 일자별 rollup 비율을 계산하는 쿼리를 반환합니다.
func GetRollupRatioByDatasourceQuery(countColName string, datasourceName string, groupByTime bool) string {
	var mapping = map[string]string{"countColName": countColName, "datasourceName": datasourceName}
	if groupByTime {
		mapping["time"] = "1"
	}
	return templateQuery("rollupQuery", getRollupRatioByDatasourceQueryTpl, &mapping)
}

// GetDataCntByDatasourceQuery 함수는 count 컬럼과 데이터소스 명을 입력하여 일자별 데이터 건 수 계산하는 쿼리를 반환합니다.
func GetDataCntByDatasourceQuery(countColName string, datasourceName string, groupByTime bool) string {
	var mapping = map[string]string{"countColName": countColName, "datasourceName": datasourceName}
	if groupByTime {
		mapping["time"] = "1"
	}
	return templateQuery("dataCntQuery", getDataCntByDatasourceQueryTpl, &mapping)
}
