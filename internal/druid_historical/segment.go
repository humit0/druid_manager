package druid_historical

import "github.com/sirupsen/logrus"

var (
	segmentEntry = baseEntry.WithFields(logrus.Fields{"type": "segment"})
)

func (historicalSvc HistoricalServiceImp) GetLoadedStatus(serverIndex int) bool {
	var result struct {
		CacheInitialized bool `json:"cacheInitialized"`
	}

	segmentEntry.Debugf("Get segment loaded status (serverIndex: %d)", serverIndex)

	historicalSvc.DruidClient.SendRequestWithIndex("GET", "historical", serverIndex, "/druid/historical/v1/loadstatus", nil, &result)

	return result.CacheInitialized
}
