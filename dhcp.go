package freebox

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	endPointDhcpDynamicLease = "dhcp/dynamic_lease"
)

type ErrorResponse struct {
	Success   bool         `json:"success"`
	ErrorCode string       `json:"error_code"`
	Message   string       `json:"msg"`
	Result    *ErrorResult `json:"result"`
}

type ErrorResult struct {
	PasswordSalt   string `json:"password_salt,omitempty"`
	Challenge      string `json:"challenge,omitempty"`
}

// ConnectionStatusResponse :
type DhcpDynamicLeaseResponse struct {
	Success   bool                `json:"success"`
	ErrorCode string              `json:"error_code,omitempty"`
	Message   string              `json:"msg,omitempty"`
	Result    []*DhcpDynamicLease `json:"result"`
}

type DhcpDynamicLease struct {
	Mac            string  `json:"mac,omitempty"`
	Hostname       string  `json:"hostname,omitempty"`
	IP             string  `json:"ip,omitempty"`
	LeaseRemaining int     `json:"lease_remaining,omitempty"`
	AssignTime     int     `json:"assign_time,omitempty"`
	RefreshTime    int     `json:"refresh_time,omitempty"`
	IsStatic       bool    `json:"is_static,omitempty"`
	Host           LanHost `json:"host,omitempty"`
}

type LanHost struct {
	ID                string                  `json:"id"`
	DefaultName       string                  `json:"default_name"`
	PrimaryName       string                  `json:"primary_name"`
	HostType          string                  `json:"host_type"`
	PrimaryNameManual bool                  `json:"primary_name_manual"`
	L2Ident           LanHostL2Ident          `json:"l2ident"`
	VendorName        string                  `json:"vendor_name"`
	Persistent        bool                    `json:"persistent"`
	Reachable         bool                    `json:"reachable"`
	LastTimeReachable int                     `json:"last_time_reachable"`
	Active            bool                    `json:"active"`
	FirstActivity     int                     `json:"first_activity"`
	LastActivity      int                     `json:"last_activity"`
	Names             []LanHostName           `json:"names"`
	L3Connectivities  []LanHostL3Connectivity `json:"l3connectivity"`
	Interface         string                  `json:"interface"`
	AccessPoint       LanHostAccessPoint      `json:"access_point"`
}

type LanHostName struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

type LanHostL2Ident struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type LanHostL3Connectivity struct {
	Addr              string `json:"addr"`
	Active            bool   `json:"active"`
	Reachable         bool   `json:"reachable"`
	LastActivity      int    `json:"last_activity"`
	AF                string `json:"af"`
	LastTimeReachable int    `json:"last_time_reachable"`
}

type LanHostAccessPoint struct {
	UID                 string `json:"uid"`
	Mac                 string `json:"mac"`
	RxBytes             int `json:"rx_bytes"`
	TxBytes             int `json:"tx_bytes"`
	RxRate             int `json:"rx_rate"`
	TxRate             int `json:"tx_rate"`
	Type                string `json:"type"`
	ConnectivityType    string `json:"connectivity_type"`
	WifiInformation struct {
		Band string `json:"band"`
		PhyRxRate int `json:"phy_rx_rate"`
		PhyTxRate int `json:"phy_tx_rate"`
		SSID string `json:"ssid"`
		Standard string `json:"standard"`
		Signal int `json:"signal"`
	} `json:"wifi_information"`
	EthernetInformation struct {
		Duplex string `json:"duplex"`
		Speed  string `json:"speed"`
		Link   string `json:"link"`
	} `json:"ethernet_information"`
}

func (device Device) DynamicLease(sessionToken string) (*DhcpDynamicLeaseResponse, *ErrorResponse, error) {
//func (device Device) DynamicLease(sessionToken string) (*string, error) {
	api := fmt.Sprintf("%s%s/", device.Url(""), endPointDhcpDynamicLease)

	client := &http.Client{}
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("X-Fbx-App-Auth", sessionToken)

	httpResponse, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, nil, err
	}

	if httpResponse.StatusCode != 200 {
		response := new(ErrorResponse)
		if err := json.Unmarshal(body, response); err != nil {
			return nil, nil, err
		}

		return nil, response, nil
	}

	response := new(DhcpDynamicLeaseResponse)
	if err := json.Unmarshal(body, response); err != nil {
		return nil, nil, err
	}

	return response, nil, nil
}
