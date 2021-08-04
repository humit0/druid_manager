package druid_broker

import "github.com/sirupsen/logrus"

var (
	baseEntry = logrus.WithFields(logrus.Fields{"server": "broker"})
)
