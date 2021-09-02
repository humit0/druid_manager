package monitor

import (
	"github.com/humit0/druid_manager/internal/druid"
	"github.com/sirupsen/logrus"
)

var (
	baseEntry = logrus.WithFields(logrus.Fields{"type": "monitor"})
)

type MonitorImp struct {
	DruidClient *druid.DruidClient
}
