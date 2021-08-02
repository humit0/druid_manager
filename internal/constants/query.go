package constants

import (
	"bytes"
	"text/template"

	"github.com/sirupsen/logrus"
)

// 일반 쿼리 목록
const (
	ServerListSqlQuery = "SELECT server, server_type FROM sys.servers"
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
	getRollupRatioQueryTpl = "SELECT SUM(`{{ .countColName }}`) * 1.0 / COUNT(*) as rollup_ratio FROM `{{ .datasourceName }}`"
)

// count 컬럼과 데이터소스 명을 입력하여 rollup 비율을 계산하는 쿼리를 반환합니다.
func GetRollupRatioQuery(countColName string, datasourceName string) string {
	tpl := template.New("rollupQuery")

	tpl, err := tpl.Parse(getRollupRatioQueryTpl)

	if err != nil {
		logrus.Fatalf("Cannot parse template %v", err)
	}

	var result bytes.Buffer

	err1 := tpl.Execute(&result, map[string]string{"countColName": countColName, "datasourceName": datasourceName})

	if err1 != nil {
		logrus.Fatalf("Cannot execute %v", err)
	}

	return result.String()
}
