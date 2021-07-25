package druid

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type DruidClient struct {
	// Master server
	CoordinatorUrls []string
	OverlordUrls    []string
	// Data server
	HistoricalUrls    []string
	MiddleManagerUrls []string
	// Query server
	BrokerUrls []string
	RouterUrls []string

	// Authentication
	Username string
	Password string
}

type SimpleServerEntity struct {
	Host     string
	Tier     string
	Type     string
	Priority string
	CurrSize string
	MaxSize  string
}

func (druidClient *DruidClient) InitClient(brokerUrl string) {
	if druidClient.Username == "" || druidClient.Password == "" {
		log.Fatalf("You should specify username and password")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/druid/coordinator/v1/servers", brokerUrl), nil)

	if err != nil {
		log.Fatalf("Failed to create request (%v)", err)
	}

	req.SetBasicAuth(druidClient.Username, druidClient.Password)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to get response (%v)", err)
	}
	defer resp.Body.Close()

	result := []SimpleServerEntity{}

	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Fatalf("Failed to parse response (%v)", err)
	}

	for _, server := range result {
		switch server.Type {
		case "coordinator":
			druidClient.CoordinatorUrls = append(druidClient.CoordinatorUrls, server.Host)
		case "overlord":
			druidClient.OverlordUrls = append(druidClient.OverlordUrls, server.Host)
		case "historical":
			druidClient.HistoricalUrls = append(druidClient.HistoricalUrls, server.Host)
		case "middleManager":
			druidClient.MiddleManagerUrls = append(druidClient.MiddleManagerUrls, server.Host)
		case "broker":
			druidClient.BrokerUrls = append(druidClient.BrokerUrls, server.Host)
		case "router":
			druidClient.RouterUrls = append(druidClient.RouterUrls, server.Host)
		default:
			log.Fatalf("Unsupported server type (%s)", server.Type)
		}
	}
}
