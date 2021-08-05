package druid_coordinator

import (
	"fmt"
	"net/url"

	"github.com/humit0/druid_manager/internal/druid"
	"github.com/sirupsen/logrus"
)

var (
	lookupEntry = baseEntry.WithFields(logrus.Fields{"type": "lookup"})
)

// GetLookupList 함수는 해당 druid 서버에 있는 전체 Lookup 목록을 반환합니다.
func GetLookupList(druidClient *druid.DruidClient, tier string) []string {
	var result []string

	lookupEntry.Debug("Get lookup list (tier: %s)", tier)
	druidClient.SendRequest("GET", "coordinator", fmt.Sprintf("/druid/coordinator/v1/lookups/config/%s", url.PathEscape(tier)), nil, &result)
	return result
}

// GetLookupsStatus 함수는 해당 druid 서버에 있는 전체 Lookup에 대한 상태를 조회합니다.
func GetLookupsStatus(druidClient *druid.DruidClient, tier string) map[string]bool {
	var respBody map[string]struct {
		Loaded bool `json:"loaded"`
	}

	lookupEntry.Debugf("Get lookup status list (tier: %s)", tier)
	druidClient.SendRequest("GET", "coordinator", fmt.Sprintf("/druid/coordinator/v1/lookups/status/%s", url.PathEscape(tier)), nil, &respBody)

	result := make(map[string]bool)

	for lookup_name, lookup_status := range respBody {
		result[lookup_name] = lookup_status.Loaded
	}

	return result
}

type LookupConfigType struct {
	Version                string `json:"version"`
	LookupExtractorFactory struct {
		Type string `json:"type"`
		// map type
		Mapping map[string]string `json:"map,omitempty"`

		// cachedNamespace type
		ExtractionNamespace struct {
			Type       string `json:"type"`
			PollPeriod string `json:"pollPeriod,omitempty"`

			// URI Lookup
			NamespaceParseSpec struct {
				// csv, tsv format
				Format         string   `json:"format"`
				Columns        []string `json:"columns,omitempty"`
				KeyColumn      string   `json:"keyColumn,omitempty"`
				ValueColumn    string   `json:"valueColumn,omitempty"`
				HasHeaderRow   bool     `json:"hasHeaderRow,omitempty"`
				SkipHeaderRows int32    `json:"skipHeaderRows,omitempty"`
				// tsv format
				Delimiter     string `json:"delimiter,omitempty"`
				ListDelimiter string `json:"listDelimiter,omitempty"`

				// customJson format
				KeyFieldName   string `json:"keyFieldName,omitempty"`
				ValueFieldName string `json:"valueFieldName,omitempty"`
			} `json:"namespaceParseSpec,omitempty"`
			Uri       string `json:"uri,omitempty"`
			UriPrefix string `json:"uriPrefix,omitempty"`
			FileRegex string `json:"fileRegex,omitempty"`

			// JDBC Lookup
			ConnectorConfig struct {
				ConnectURI string `json:"connectURI"`
				User       string `json:"user"`
				Password   string `json:"password"`
			} `json:"connectorConfig,omitempty"`
			Table       string `json:"table,omitempty"`
			KeyColumn   string `json:"keyColumn,omitempty"`
			ValueColumn string `json:"valueColumn,omitempty"`
			Filter      string `json:"filter,omitempty"`
			TsColumn    string `json:"tsColumn,omitempty"`
		} `json:"extractionNamespace,omitempty"`

		FirstCacheTimeout int32 `json:"firstCacheTimeout,omitempty"`
		Injective         bool  `json:"injective,omitempty"`
	} `json:"lookupExtractorFactory"`
}

// GetLookupConfig 함수는 해당 druid 서버에 있는 Lookup에 대한 설정을 반환합니다.
func GetLookupConfig(druidClient *druid.DruidClient, tier string, lookup_name string) LookupConfigType {
	var result LookupConfigType

	lookupEntry.Debugf("Get lookup config (tier: %s, lookup_name: %s)", tier, lookup_name)
	druidClient.SendRequest("GET", "coordinator", fmt.Sprintf("/druid/coordinator/v1/lookups/config/%s/%s", tier, lookup_name), nil, &result)
	return result
}
