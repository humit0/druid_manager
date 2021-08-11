package druid

type BrokerService interface {
	SendSQLQuery(sqlQuery string, queryResult interface{})
	SendNativeQuery(nativeQuery string, queryResult interface{})

	GetDatasourceDimensions(datasourceName string) []string
	GetDatasourceMetrics(datasourceName string) []string
}

func (druidClient *DruidClient) SetBrokerSvc(brokerSvc BrokerService) {
	druidClient.BrokerSvc = brokerSvc
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

type CoordinatorService interface {
	GetAllCompactionConfiguration() interface{}
	GetCompactionConfigurationByDatasource(datasourceName string) interface{}

	GetDatasourceList() []string

	GetLookupList(tier string) []string
	GetLookupsStatus(tier string) map[string]bool
	GetLookupConfig(tier string, lookupName string) LookupConfigType
}

func (druidClient *DruidClient) SetCoodinatorSvc(coordinatorSvc CoordinatorService) {
	druidClient.CoordinatorSvc = coordinatorSvc
}

type TaskQueryParamType struct {
	State               string `url:"state,omitempty"`
	Datasource          string `url:"datasource,omitempty"`
	CreatedTimeInterval string `url:"createdTimeInterval,omitempty"`
	Max                 int    `url:"max,omitempty"`
	TaskType            string `url:"taskType,omitempty"`
}

type OverlordService interface {
	GetCurrentLeader() string
	CheckIsLeader(serverIndex int) bool

	GetAllTaskList(queryParam TaskQueryParamType) interface{}
	GetTaskPayload(taskId string) interface{}
	GetTaskStatus(taskId string) interface{}
}

func (druidClient *DruidClient) SetOverlordSvc(overlordSvc OverlordService) {
	druidClient.OverlordSvc = overlordSvc
}

type HistoricalService interface {
	GetLoadedStatus(serverIndex int) bool
}

func (druidClient *DruidClient) SetHistoricalSvc(historicalSvc HistoricalService) {
	druidClient.HistoricalSvc = historicalSvc
}

type MiddleManagerService interface {
	GetAllActiveTaskList() []string
	GetTaskLog(taskId string) string
	GetWorkerStatus(serverIndex int) map[string]string
}

func (druidClient *DruidClient) SetMiddleManagerSvc(middleManagerSvc MiddleManagerService) {
	druidClient.MiddleManagerSvc = middleManagerSvc
}
