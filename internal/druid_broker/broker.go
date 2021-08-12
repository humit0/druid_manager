package druid_broker

import (
	"github.com/humit0/druid_manager/internal/druid"
	"github.com/sirupsen/logrus"
)

var (
	baseEntry = logrus.WithFields(logrus.Fields{"server": "broker"})
)

type BrokerServiceImp struct {
	DruidClient *druid.DruidClient
}
