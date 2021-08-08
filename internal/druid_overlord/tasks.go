package druid_overlord

import (
	"bytes"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/humit0/druid_manager/internal/druid"
	"github.com/sirupsen/logrus"
)

var (
	taskEntry = baseEntry.WithFields(logrus.Fields{"type": "task"})
)

type TaskQueryParamType struct {
	State               string `url:"state,omitempty"`
	Datasource          string `url:"datasource,omitempty"`
	CreatedTimeInterval string `url:"createdTimeInterval,omitempty"`
	Max                 int    `url:"max,omitempty"`
	TaskType            string `url:"taskType,omitempty"`
}

// GetAllTaskList 함수는 조건에 맞는 테스크 목록을 반환합니다.
func GetAllTaskList(druidClient *druid.DruidClient, queryParam TaskQueryParamType) interface{} {
	var result interface{}

	taskEntry.Debugf("Get all task list (state=%s)", queryParam.State)

	params, _ := query.Values(queryParam)
	paramString := params.Encode()
	taskEntry.Debugf("param: %s", paramString)

	urlBuf := bytes.NewBufferString("/druid/indexer/v1/tasks")
	if paramString != "" {
		urlBuf.WriteString("?")
		urlBuf.WriteString(paramString)
	}

	druidClient.SendRequest("GET", "overlord", urlBuf.String(), nil, &result)
	return result
}

// GetTaskPayload 함수는 전달한 `taskId`에 해당하는 태스크의 payload를 반환합니다.
func GetTaskPayload(druidClient *druid.DruidClient, taskId string) interface{} {
	var result interface{}

	taskEntry.Debugf("Get task payload (taskId=%s)", taskId)

	urlBuf := bytes.NewBufferString("/druid/indexer/v1/task/")
	urlBuf.WriteString(url.PathEscape(taskId))

	druidClient.SendRequest("GET", "overlord", urlBuf.String(), nil, &result)
	return result
}

// GetTaskStatus 함수는 전달한 `taskId`에 해당하는 태스크의 상태를 반환합니다.
func GetTaskStatus(druidClient *druid.DruidClient, taskId string) interface{} {
	var result interface{}

	taskEntry.Debugf("Get task status (taskId=%s)", taskId)

	urlBuf := bytes.NewBufferString("/druid/indexer/v1/task/")
	urlBuf.WriteString(url.PathEscape(taskId))
	urlBuf.WriteString("/status")

	druidClient.SendRequest("GET", "overlord", urlBuf.String(), nil, &result)
	return result
}
