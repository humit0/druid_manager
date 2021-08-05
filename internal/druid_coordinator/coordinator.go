// Package druid_coordinator는 Druid Coordinator 클러스터에서 제공하는 API를
// 구현한 패키지입니다.

package druid_coordinator

import "github.com/sirupsen/logrus"

var (
	baseEntry = logrus.WithFields(logrus.Fields{"server": "coordinator"})
)
