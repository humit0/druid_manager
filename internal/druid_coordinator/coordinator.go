// Package druid_coordinator는 Druid Coordinator 클러스터에서 제공하는 API를
// 구현한 패키지입니다.

package druid_coordinator

import (
	"github.com/humit0/druid_manager/internal/druid"
	"github.com/sirupsen/logrus"
)

var (
	baseEntry = logrus.WithFields(logrus.Fields{"server": "coordinator"})
)

type CoordinatorServiceImp struct {
	DruidClient druid.DruidClient
}
