// package druid는 드루이드 API를 호출할 때 사용하는 클라이언트 내용을 구현합니다.
package druid

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/humit0/druid_manager/internal/constants"
	"github.com/sirupsen/logrus"
)

// DruidClient 구조체는 드루이드 클라이언트 관련 정보를 포함합니다.
type DruidClient struct {
	// Master server
	CoordinatorURLs []string
	OverlordURLs    []string
	// Data server
	HistoricalURLs    []string
	MiddleManagerURLs []string
	// Query server
	BrokerURLs []string
	RouterURLs []string

	// Authentication
	Username string
	Password string

	// Druid Service
	BrokerSvc        BrokerService
	CoordinatorSvc   CoordinatorService
	OverlordSvc      OverlordService
	HistoricalSvc    HistoricalService
	MiddleManagerSvc MiddleManagerService
}

// SimpleServerEntity 구조체는 서버 목록에 대한 응답을 받을 때 사용하는 JSON 타입입니다.
type SimpleServerEntity struct {
	Host     string
	Tier     string
	Type     string
	Priority string
	CurrSize string
	MaxSize  string
}

var (
	entry = logrus.StandardLogger()
)

// sortServer 함수는 서버 주소를 정렬합니다.
func (druidClient *DruidClient) sortServer() {
	sort.Strings(druidClient.CoordinatorURLs)
	sort.Strings(druidClient.OverlordURLs)
	sort.Strings(druidClient.HistoricalURLs)
	sort.Strings(druidClient.MiddleManagerURLs)
	sort.Strings(druidClient.BrokerURLs)
	sort.Strings(druidClient.RouterURLs)
}

// InitClient 함수는 브로커 주소를 받아서 클라이언트 관련 정보를 초기화합니다.
func (druidClient *DruidClient) InitClient(brokerURL string) {
	brokerURL = strings.TrimRight(brokerURL, "/")
	// broker 쪽에 주소를 추가했다가
	druidClient.BrokerURLs = append(druidClient.BrokerURLs, brokerURL)

	var result []constants.ServerListItem
	druidClient.BrokerSvc.SendSQLQuery(constants.ServerListQuery, &result)

	// 다시 제거합니다.
	druidClient.BrokerURLs = druidClient.BrokerURLs[:0]

	// 서버 정보를 순회하면서 추가합니다.
	for _, server := range result {
		switch server.ServerType {
		case "coordinator":
			druidClient.CoordinatorURLs = append(druidClient.CoordinatorURLs, server.Server)
		case "overlord":
			druidClient.OverlordURLs = append(druidClient.OverlordURLs, server.Server)
		case "historical":
			druidClient.HistoricalURLs = append(druidClient.HistoricalURLs, server.Server)
		case "middleManager":
			druidClient.MiddleManagerURLs = append(druidClient.MiddleManagerURLs, server.Server)
		case "broker":
			druidClient.BrokerURLs = append(druidClient.BrokerURLs, server.Server)
		case "router":
			druidClient.RouterURLs = append(druidClient.RouterURLs, server.Server)
		default:
			entry.Fatalf("Unsupported server type (%s)", server.ServerType)
		}
	}
	druidClient.sortServer()
}

// GetServerList 함수는 druid 클러스터 종류를 입력받아서 해당 서버 목록을 반환합니다.
func (druidClient *DruidClient) GetServerList(serverType string) []string {
	switch serverType {
	case "coordinator":
		return druidClient.CoordinatorURLs
	case "overlord":
		return druidClient.OverlordURLs
	case "historical":
		return druidClient.HistoricalURLs
	case "middleManager":
		return druidClient.MiddleManagerURLs
	case "broker":
		return druidClient.BrokerURLs
	case "router":
		return druidClient.RouterURLs
	default:
		entry.Fatalf("Invalid server type (%s)", serverType)
	}
	return nil
}

// CreateRequestWithIndex 함수는 HTTP 메서드, druid 클러스터 종류, path, body를 입력받아서 해당 요청을 보내는 `http.Request` 객체를 생성합니다.
func (druidClient *DruidClient) CreateRequestWithIndex(method string, serverType string, serverIndex int, path string, requestBody io.Reader) *http.Request {
	serverURLs := druidClient.GetServerList(serverType)

	if len(serverURLs) == 0 || serverIndex >= len(serverURLs) || serverIndex < 0 {
		entry.Fatalf("Cannot get server url (server type: %s) [index: %d]", serverType, serverIndex)
	}
	requestURLBuff := bytes.NewBufferString(serverURLs[serverIndex])
	requestURLBuff.WriteString(path)
	requestURL := requestURLBuff.String()

	entryWithReq := entry.WithField("method", method).WithField("url", requestURL)
	if druidClient.Username == "" || druidClient.Password == "" {
		entryWithReq.Fatal("You should specify username and password")
	}

	entryWithReq.Info("Create new request")
	req, err := http.NewRequest(method, requestURL, requestBody)

	if err != nil {
		entryWithReq.Fatalf("Cannot create request object (%v)", err)
	}
	req.SetBasicAuth(druidClient.Username, druidClient.Password)
	req.Header.Add("Content-Type", "application/json")

	return req
}

// GetResponse 함수는 `http.Request` 객체의 응답을 받아서 JSON으로 파싱합니다.
func GetResponse(req *http.Request, result interface{}) {
	client := &http.Client{}
	resp, err := client.Do(req)

	urlBuff := bytes.NewBufferString(req.URL.Scheme)
	urlBuff.WriteString("://")
	urlBuff.WriteString(req.Host)
	urlBuff.WriteString(req.URL.Path)

	entryWithReq := entry.WithField("method", req.Method).WithField("url", urlBuff.String())

	if err != nil {
		entryWithReq.Fatalf("Failed to get response (%v)", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	entryWithReq.Debugf("respBody: %s", body)

	err = json.Unmarshal(body, &result)

	if err != nil {
		entryWithReq.Fatalf("Failed to parse response (%v)\nBody: %s", err, body)
	}
}

// SendRequestWithIndex 함수는 HTTP 메서드, druid 클러스터 종류, path, body, 서버 index를 입력받아서 해당 요청을 보내고 해당 응답을 JSON으로 파싱합니다.
func (druidClient *DruidClient) SendRequestWithIndex(method string, serverType string, serverIndex int, path string, requestBody io.Reader, result interface{}) {
	req := druidClient.CreateRequestWithIndex(method, serverType, serverIndex, path, requestBody)
	GetResponse(req, &result)
}

// SendRequest 함수는 HTTP 메서드, druid 클러스터 종류, path, body를 입력받아서 해당 요청을 보내고 해당 응답을 JSON으로 파싱합니다.
func (druidClient *DruidClient) SendRequest(method string, serverType string, path string, requestBody io.Reader, result interface{}) {
	druidClient.SendRequestWithIndex(method, serverType, 0, path, requestBody, &result)
}

// ShowServers 함수는 서버 목록을 druid 클러스터 별로 출력합니다.
func (druidClient *DruidClient) ShowServers() {
	entry.Info("Druid cluster list")
	for _, serverURL := range druidClient.CoordinatorURLs {
		entry.WithField("server", "coordinator").Infof("url: %s", serverURL)
	}
	for _, serverURL := range druidClient.OverlordURLs {
		entry.WithField("server", "overlord").Infof("url: %s", serverURL)
	}
	for _, serverURL := range druidClient.HistoricalURLs {
		entry.WithField("server", "historical").Infof("url: %s", serverURL)
	}
	for _, serverURL := range druidClient.MiddleManagerURLs {
		entry.WithField("server", "middleManager").Infof("url: %s", serverURL)
	}
	for _, serverURL := range druidClient.BrokerURLs {
		entry.WithField("server", "broker").Infof("url: %s", serverURL)
	}

	for _, serverURL := range druidClient.RouterURLs {
		entry.WithField("server", "router").Infof("url: %s", serverURL)
	}
}
