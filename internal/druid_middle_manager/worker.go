package druid_middle_manager

import "github.com/sirupsen/logrus"

var (
	workerEntry = baseEntry.WithFields(logrus.Fields{"type": "worker"})
)

func (middleManagerSvc *MiddleManagerServiceImp) GetWorkerStatus(serverIndex int) map[string]string {
	var result map[string]string

	workerEntry.Debugf("Get worker status (serverIndex: %d)", serverIndex)

	middleManagerSvc.DruidClient.SendRequestWithIndex("GET", "middleManager", serverIndex, "/druid/worker/v1/enabled", nil, &result)

	return result
}
