package druid_overlord

import (
	"github.com/sirupsen/logrus"
)

var (
	leadershipEntry = baseEntry.WithFields(logrus.Fields{"type": "leadership"})
)

// GetCurrentLeader 함수는 현재 leader에 해당하는 overlord 서버를 반환합니다.
func (overlordSvc *OverlordServiceImp) GetCurrentLeader() string {
	var result string

	leadershipEntry.Debug("Get current leader")
	overlordSvc.DruidClient.SendRequest("GET", "overlord", "/druid/indexer/v1/leader", nil, &result)

	return result
}

// CheckIsLeader 함수는 해당 druid 서버가 leader 서버인지 여부를 반환합니다.
func (overlordSvc *OverlordServiceImp) CheckIsLeader(serverIndex int) bool {
	var result struct {
		Leader bool `json:"leader"`
	}

	leadershipEntry.Debug("Check current is leader")
	overlordSvc.DruidClient.SendRequest("GET", "overlord", "/druid/indexer/v1/isLeader", nil, &result)

	return result.Leader
}
