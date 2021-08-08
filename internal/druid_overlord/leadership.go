package druid_overlord

import (
	"github.com/humit0/druid_manager/internal/druid"
	"github.com/sirupsen/logrus"
)

var (
	leadershipEntry = baseEntry.WithFields(logrus.Fields{"type": "leadership"})
)

// GetCurrentLeader 함수는 현재 leader에 해당하는 overlord 서버를 반환합니다.
func GetCurrentLeader(druidClient *druid.DruidClient) string {
	var result string

	leadershipEntry.Debug("Get current leader")
	druidClient.SendRequest("GET", "overlord", "/druid/indexer/v1/leader", nil, &result)

	return result
}

// CheckIsLeader 함수는 해당 druid 서버가 leader 서버인지 여부를 반환합니다.
func CheckIsLeader(druidClient *druid.DruidClient, serverIndex int) bool {
	result := struct {
		Leader bool `json:"leader"`
	}{}

	leadershipEntry.Debug("Check current is leader")
	druidClient.SendRequest("GET", "overlord", "/druid/indexer/v1/isLeader", nil, &result)

	return result.Leader
}
