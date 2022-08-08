package utils

type Util struct {
	consulServer string
}

// key is the service, the list contains tags
type allServicesResponse map[string][]string

// consul service
type ConsulServiceResponse []map[string]interface{}

type catalogDeregisterRequest struct {
	Node      string `json:"Node"`
	ServiceID string `json:"ServiceID"`
}
