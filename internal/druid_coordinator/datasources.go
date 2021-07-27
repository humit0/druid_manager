package druid_coordinator

import (
	"github.com/humit0/druid_manager/internal/druid"
)

// 해당 druid 서버에 있는 전체 데이터 소스 목록을 반환합니다.
func GetDatasources(druidClient *druid.DruidClient) []string {
	result := []string{}

	druidClient.SendRequest("GET", "coordinator", "/druid/coordinator/v1/datasources", nil, &result)
	return result
}
