package druid_coordinator

import (
	"bytes"
	"net/url"

	"github.com/humit0/druid_manager/internal/druid"
	"github.com/sirupsen/logrus"
)

var (
	lookupEntry = baseEntry.WithFields(logrus.Fields{"type": "lookup"})
)

// GetLookupList 함수는 해당 druid 서버에 있는 전체 Lookup 목록을 반환합니다.
func (coordinatorService *CoordinatorServiceImp) GetLookupList(tier string) []string {
	var result []string

	lookupEntry.Debugf("Get lookup list (tier: %s)", tier)

	urlBuff := bytes.NewBufferString("/druid/coodinator/v1/lookups/config/")
	urlBuff.WriteString(url.PathEscape(tier))

	coordinatorService.DruidClient.SendRequest("GET", "coordinator", urlBuff.String(), nil, &result)
	return result
}

// GetLookupsStatus 함수는 해당 druid 서버에 있는 전체 Lookup에 대한 상태를 조회합니다.
func (coordinatorService *CoordinatorServiceImp) GetLookupsStatus(tier string) map[string]bool {
	var respBody map[string]struct {
		Loaded bool `json:"loaded"`
	}

	lookupEntry.Debugf("Get lookup status list (tier: %s)", tier)

	urlBuff := bytes.NewBufferString("/druid/coordinator/v1/lookups/status/")
	urlBuff.WriteString(url.PathEscape(tier))

	coordinatorService.DruidClient.SendRequest("GET", "coordinator", urlBuff.String(), nil, &respBody)

	result := make(map[string]bool)

	for lookup_name, lookup_status := range respBody {
		result[lookup_name] = lookup_status.Loaded
	}

	return result
}

// GetLookupConfig 함수는 해당 druid 서버에 있는 Lookup에 대한 설정을 반환합니다.
func (coordinatorService *CoordinatorServiceImp) GetLookupConfig(tier string, lookupName string) druid.LookupConfigType {
	var result druid.LookupConfigType

	lookupEntry.Debugf("Get lookup config (tier: %s, lookupName: %s)", tier, lookupName)

	urlBuff := bytes.NewBufferString("/druid/coordinator/v1/lookups/config/")
	urlBuff.WriteString(url.PathEscape(tier))
	urlBuff.WriteString("/")
	urlBuff.WriteString(url.PathEscape(lookupName))

	coordinatorService.DruidClient.SendRequest("GET", "coordinator", urlBuff.String(), nil, &result)
	return result
}
