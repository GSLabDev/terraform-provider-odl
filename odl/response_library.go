package odl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// VtnList ... data structure for holding vtn list from odl
type VtnList struct {
	Vtns struct {
		Vtn []struct {
			Name    string `json:"name"`
			Vbridge []struct {
				BridgeStatus struct {
					PathFaults int    `json:"path-faults"`
					State      string `json:"state"`
				} `json:"bridge-status"`
				Name          string `json:"name"`
				VbridgeConfig struct {
					AgeInterval int `json:"age-interval"`
				} `json:"vbridge-config"`
				VInterface []struct {
					Name             string `json:"name"`
					VInterfaceConfig struct {
						Description string `json:"description"`
						Enabled     bool   `json:"enabled"`
					} `json:"vinterface-config"`
					VInterfaceStatus struct {
						EntityState string `json:"entity-state"`
						State       string `json:"state"`
					} `json:"vinterface-status"`
				} `json:"vinterface"`
			} `json:"vbridge"`
			VtenantConfig struct {
				HardTimeout int `json:"hard-timeout"`
				IdleTimeout int `json:"idle-timeout"`
			} `json:"vtenant-config"`
		} `json:"vtn"`
	} `json:"vtns"`
}

// ErrorCase ... data structure of error case from odl
type ErrorCase struct {
	Errors struct {
		Error []struct {
			AppTag  string `json:"error-app-tag"`
			Info    string `json:"error-info"`
			Message string `json:"error-message"`
			Tag     string `json:"error-tag"`
			Type    string `json:"error-type"`
		} `json:"error"`
	} `json:"errors"`
}

// OutputCase ... data structure of output case from odl
type OutputCase struct {
	Output struct {
		Status string `json:"status"`
	} `json:"output"`
}

// Status ... Parses output of odl
func Status(response *http.Response) (bool, *OutputCase, *ErrorCase, error) {
	respString, err := getResponseAsString(response)
	if err != nil {
		return false, nil, nil, err
	}
	if strings.Contains(respString, "output") {
		data := &OutputCase{}
		err := json.Unmarshal([]byte(respString), data)
		if err != nil {
			return false, nil, nil, err
		}
		return true, data, nil, nil
	}
	if strings.Contains(respString, "error-message") {
		data := &ErrorCase{}
		err := json.Unmarshal([]byte(respString), data)
		if err != nil {
			return false, nil, nil, err
		}
		return false, nil, data, nil
	}
	return false, nil, nil, nil
}

//CheckResponseVirtualTenantNetworkExists ... checks response if vtn exists
func CheckResponseVirtualTenantNetworkExists(response *http.Response, name string) (bool, error) {
	respString, err := getResponseAsString(response)
	data := &VtnList{}
	err = json.Unmarshal([]byte(respString), data)
	if err != nil {
		return false, err
	}
	if data.Vtns.Vtn != nil {
		for i := range data.Vtns.Vtn {
			if data.Vtns.Vtn[i].Name == name {
				return true, nil
			}
		}
	}
	return false, nil
}

//CheckResponseVirtualBridgeExists ... checks response if vtn exists
func CheckResponseVirtualBridgeExists(response *http.Response, tenantName, bridgeName string) (bool, error) {
	respString, err := getResponseAsString(response)
	data := &VtnList{}
	err = json.Unmarshal([]byte(respString), data)
	if err != nil {
		return false, err
	}
	if data.Vtns.Vtn != nil {
		for i := range data.Vtns.Vtn {
			if data.Vtns.Vtn[i].Name == tenantName {
				if data.Vtns.Vtn[i].Vbridge != nil {
					for j := range data.Vtns.Vtn[i].Vbridge {
						if data.Vtns.Vtn[i].Vbridge[j].Name == bridgeName {
							return true, nil
						}
					}
				}
				return false, nil
			}
		}
	}
	return false, nil
}

//CheckResponseVirtualInterfaceExists ... checks response if vInterface exists
func CheckResponseVirtualInterfaceExists(response *http.Response, tenantName, bridgeName, interfaceName string) (bool, error) {
	respString, err := getResponseAsString(response)
	data := &VtnList{}
	err = json.Unmarshal([]byte(respString), data)
	if err != nil {
		return false, err
	}
	if data.Vtns.Vtn != nil {
		for i := range data.Vtns.Vtn {
			if data.Vtns.Vtn[i].Name == tenantName {
				if data.Vtns.Vtn[i].Vbridge != nil {
					for j := range data.Vtns.Vtn[i].Vbridge {
						if data.Vtns.Vtn[i].Vbridge[j].Name == bridgeName {
							if data.Vtns.Vtn[i].Vbridge[j].VInterface != nil {
								for k := range data.Vtns.Vtn[i].Vbridge[j].VInterface {
									if data.Vtns.Vtn[i].Vbridge[j].VInterface[k].Name == interfaceName {
										return true, nil
									}
								}
							}
							return false, nil
						}
					}
				}
				return false, nil
			}
		}
	}
	return false, nil
}

func getResponseAsString(response *http.Response) (string, error) {
	if response.Status != "200 OK" {
		return "", fmt.Errorf("[ERROR] %s ", response.Status)
	}
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("[ERROR] Reading body %s", err.Error())
		return "", fmt.Errorf("[ERROR] Reading body %s", err.Error())
	}
	return string(buf), nil
}
