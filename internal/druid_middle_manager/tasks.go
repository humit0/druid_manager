package druid_middle_manager

import (
	"bytes"
	"net/url"

	"github.com/sirupsen/logrus"
)

var (
	taskEntry = baseEntry.WithFields(logrus.Fields{"type": "task"})
)

// GetAllActiveTaskList 함수는 조건에 맞는 테스크 목록을 반환합니다.
func (middleManagerSvc MiddleManagerServiceImp) GetAllActiveTaskList() []string {
	var result []string

	taskEntry.Debugf("Get all active task list")

	urlBuf := bytes.NewBufferString("/druid/indexer/v1/tasks")

	middleManagerSvc.DruidClient.SendRequest("GET", "middleManager", urlBuf.String(), nil, &result)
	return result
}

// GetTaskLog 함수는 전달한 `taskId`에 해당하는 태스크의 log를 반환합니다.
func (middleManagerSvc MiddleManagerServiceImp) GetTaskLog(taskId string) string {
	var result string

	taskEntry.Debugf("Get task log (taskId=%s)", taskId)

	urlBuf := bytes.NewBufferString("/druid/worker/v1/task/")
	urlBuf.WriteString(url.PathEscape(taskId))
	urlBuf.WriteString("/log")

	middleManagerSvc.DruidClient.SendRequest("GET", "middleManager", urlBuf.String(), nil, &result)
	return result
}
