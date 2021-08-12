package druid_middle_manager

import (
	"github.com/humit0/druid_manager/internal/druid"
	"github.com/sirupsen/logrus"
)

var (
	baseEntry = logrus.WithFields(logrus.Fields{"server": "middleManager"})
)

type MiddleManagerServiceImp struct {
	DruidClient *druid.DruidClient
}
