package remote

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const nodeInfoURL = "http://127.0.0.1:9000/api/v1/nodeinfo"

type NodeInfo struct {
	DhcpServer string `json:"dhcpserver"`
	Key        string `json:"key"`
	Vpn        string `json:"vpn"`
	WorkId     string `json:"workid"`
}

func GetNodeInfo() (nodeInfo NodeInfo, err error) {
	resp, err := http.Get(nodeInfoURL)
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &nodeInfo)
	return
}
