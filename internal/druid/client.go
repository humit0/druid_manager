// Druid client package
// 드루이드 API를 호출할 때 사용하는 클라이언트 내용을 구현합니다.
package druid

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

// 드루이드 클라이언트
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
}

// 서버 목록에 대한 응답을 받을 때 사용하는 JSON 타입입니다.
type SimpleServerEntity struct {
	Host     string
	Tier     string
	Type     string
	Priority string
	CurrSize string
	MaxSize  string
}

// 서버 주소를 정렬합니다.
func (druidClient *DruidClient) sortServer() {
	sort.Strings(druidClient.CoordinatorURLs)
	sort.Strings(druidClient.OverlordURLs)
	sort.Strings(druidClient.HistoricalURLs)
	sort.Strings(druidClient.MiddleManagerURLs)
	sort.Strings(druidClient.BrokerURLs)
	sort.Strings(druidClient.RouterURLs)
}

// 클라이언트 관련 정보를 초기화합니다.
func (druidClient *DruidClient) InitClient(routerURL string) {
	// router 쪽에 주소를 추가했다가
	druidClient.RouterURLs = append(druidClient.RouterURLs, routerURL)
	result := []SimpleServerEntity{}
	druidClient.SendRequest("GET", "router", "/druid/coordinator/v1/servers", nil, &result)

	// 다시 제거합니다.
	druidClient.RouterURLs = druidClient.RouterURLs[:0]

	// 서버 정보를 순회하면서 추가합니다.
	for _, server := range result {
		switch server.Type {
		case "coordinator":
			druidClient.CoordinatorURLs = append(druidClient.CoordinatorURLs, server.Host)
		case "overlord":
			druidClient.OverlordURLs = append(druidClient.OverlordURLs, server.Host)
		case "historical":
			druidClient.HistoricalURLs = append(druidClient.HistoricalURLs, server.Host)
		case "middleManager":
			druidClient.MiddleManagerURLs = append(druidClient.MiddleManagerURLs, server.Host)
		case "broker":
			druidClient.BrokerURLs = append(druidClient.BrokerURLs, server.Host)
		case "router":
			druidClient.RouterURLs = append(druidClient.RouterURLs, server.Host)
		default:
			log.Fatalf("Unsupported server type (%s)", server.Type)
		}
	}
	druidClient.sortServer()
}

// druid 클러스터 종류를 입력받아서 해당 서버 목록을 반환합니다.
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
		log.Fatalf("Invalid server type (%s)", serverType)
	}
	return nil
}

// HTTP 메서드, druid 클러스터 종류, path, body를 입력받아서 해당 요청을 보내는 `http.Request` 객체를 생성합니다.
func (druidClient *DruidClient) CreateRequest(method string, serverType string, path string, requestBody io.Reader) *http.Request {
	if druidClient.Username == "" || druidClient.Password == "" {
		log.Fatal("You should specify username and password")
	}
	serverURLs := druidClient.GetServerList(serverType)

	if len(serverURLs) == 0 {
		log.Fatalf("Cannot get server url (server type: %s)", serverType)
	}

	requestURL := fmt.Sprintf("%s/%s", serverURLs[0], path)

	req, err := http.NewRequest(method, requestURL, requestBody)

	if err != nil {
		log.Fatalf("Cannot create request object (%v)", err)
	}
	req.SetBasicAuth(druidClient.Username, druidClient.Password)
	return req
}

// `http.Request` 객체의 응답을 받아서 JSON으로 파싱합니다.
func GetResponse(req *http.Request, result interface{}) {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Failed to get response (%v)", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Fatalf("Failed to parse response (%v)", err)
	}
}

// HTTP 메서드, druid 클러스터 종류, path, body를 입력받아서 해당 요청을 보내고 해당 응답을 JSON으로 파싱합니다.
func (druidClient *DruidClient) SendRequest(method string, serverType string, path string, requestBody io.Reader, result interface{}) {
	req := druidClient.CreateRequest(method, serverType, path, requestBody)
	GetResponse(req, &result)
}

// 서버 목록을 druid 클러스터 별로 출력합니다.
func (druidClient *DruidClient) ShowServers() {
	fmt.Printf("== Coordinate server ==\n")
	for _, serverURL := range druidClient.CoordinatorURLs {
		fmt.Printf("* %s\n", serverURL)
	}
	fmt.Printf("\n== Overlord server ==\n")
	for _, serverURL := range druidClient.OverlordURLs {
		fmt.Printf("* %s\n", serverURL)
	}
	fmt.Printf("\n== Historical server ==\n")
	for _, serverURL := range druidClient.HistoricalURLs {
		fmt.Printf("* %s\n", serverURL)
	}
	fmt.Printf("\n== MiddleManager server ==\n")
	for _, serverURL := range druidClient.MiddleManagerURLs {
		fmt.Printf("* %s\n", serverURL)
	}
	fmt.Printf("\n== Broker server ==\n")
	for _, serverURL := range druidClient.BrokerURLs {
		fmt.Printf("* %s\n", serverURL)
	}

	fmt.Printf("\n== Router server ==\n")
	for _, serverURL := range druidClient.RouterURLs {
		fmt.Printf("* %s\n", serverURL)
	}
}
