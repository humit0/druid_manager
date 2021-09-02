package druid

func GetDruidClusterInformation(druidClient *DruidClient, serverType string, serverIndex int) map[string]string {
	var result map[string]string
	druidClient.SendRequestWithIndex("GET", serverType, serverIndex, "/status", nil, &result)
	return result
}

func CheckHealthStatus(druidClient *DruidClient, serverType string, serverIndex int) bool {
	var result bool
	druidClient.SendRequestWithIndex("GET", serverType, serverIndex, "/status/health", nil, &result)
	return result
}

func GetDruidConfigurations(druidClient *DruidClient, serverType string, serverIndex int) map[string]string {
	var result map[string]string
	druidClient.SendRequestWithIndex("GET", serverType, serverIndex, "/status/properties", nil, &result)
	return result
}

func CheckHealthy(druidClient *DruidClient, serverType string, serverIndex int) bool {
	var result struct {
		SelfDiscovered bool `json:"selfDiscovered"`
	}
	druidClient.SendRequestWithIndex("GET", serverType, serverIndex, "/status/health", nil, &result)
	return result.SelfDiscovered
}
