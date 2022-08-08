package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/exp/slices"
)

func New(consul string) *Util {
	return &Util{
		consulServer: consul,
	}
}

func (u *Util) GetService(name string) ConsulServiceResponse {
	svcResp := makeRequest(u.buildUrl(fmt.Sprintf("/v1/catalog/service/%s", name)))
	defer svcResp.Body.Close()

	var aSvc ConsulServiceResponse
	json.NewDecoder(svcResp.Body).Decode(&aSvc)

	return aSvc
}

// filters services based on the supplied 'tag'
func (u *Util) FindServicesToDelete(tag string) []string {
	allSvcResp := makeRequest(u.buildUrl("/v1/catalog/services"))

	defer allSvcResp.Body.Close()

	var services allServicesResponse
	json.NewDecoder(allSvcResp.Body).Decode(&services)

	var deleteSvc []string

	for svc, svcTags := range services {
		if slices.Contains(svcTags, tag) {
			deleteSvc = append(deleteSvc, svc)
		}
	}

	return deleteSvc
}

// unregisters a service
func (u *Util) DeleteService(node string, serviceId string) error {
	data := catalogDeregisterRequest{
		Node:      node,
		ServiceID: serviceId,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", u.buildUrl("/v1/catalog/deregister"), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	client := http.Client{}
	respDelete, err := client.Do(req)
	if err != nil {
		return err
	}

	defer respDelete.Body.Close()

	b, _ := ioutil.ReadAll(respDelete.Body)
	fmt.Println(string(b))

	return nil
}

func (u *Util) buildUrl(path string) string {
	return fmt.Sprintf("%s%s", u.consulServer, path)
}

// wraps around GET
func makeRequest(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return resp
}
