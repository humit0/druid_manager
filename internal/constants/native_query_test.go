package constants_test

import (
	"testing"

	"github.com/humit0/druid_manager/internal/constants"
	"github.com/stretchr/testify/assert"
)

func TestGetTimeBoundaryQuery(t *testing.T) {
	query1 := constants.GetTimeBoundaryQuery("datasource1", "maxTime")
	assert.Equal(t, query1, `{"queryType":"timeBoundary","dataSource":"datasource1","bound":"maxTime"}`)

	query2 := constants.GetTimeBoundaryQuery("datasource2", "")
	assert.Equal(t, query2, `{"queryType":"timeBoundary","dataSource":"datasource2"}`)
}

func TestGetSegmentMetadataQuery(t *testing.T) {
	query1 := constants.GetSegmentMetadataQuery(constants.SegmentMetadataQueryType{Datasource: "datasource1"})
	assert.Equal(t, query1, `{"queryType":"segmentMetadata","dataSource":"datasource1"}`)

	query2 := constants.GetSegmentMetadataQuery(constants.SegmentMetadataQueryType{Datasource: "datasource2", Merge: true})
	assert.Equal(t, query2, `{"queryType":"segmentMetadata","dataSource":"datasource2","merge":true}`)

	queryArg3 := constants.SegmentMetadataQueryType{Datasource: "datasource3"}
	queryArg3.ToInclude = &struct {
		Type    string   "json:\"type\""
		Columns []string "json:\"columns,omitempty\""
	}{Type: "list", Columns: []string{"col1"}}
	query3 := constants.GetSegmentMetadataQuery(queryArg3)
	assert.Equal(t, query3, `{"queryType":"segmentMetadata","dataSource":"datasource3","toInclude":{"type":"list","columns":["col1"]}}`)
}

func TestGetDatasourceMetadataQuery(t *testing.T) {
	query := constants.GetDatasourceMetadataQuery(constants.DatasourceMetadataQueryType{Datasource: "datasource1"})
	assert.Equal(t, query, `{"queryType":"dataSourceMetadata","dataSource":"datasource1"}`)
}
