package druid_coordinator

import "github.com/sirupsen/logrus"

var (
	baseEntry = logrus.WithFields(logrus.Fields{"server": "coordinator"})
)
