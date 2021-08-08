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
