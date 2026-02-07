package common

type MethodType string

const (
	GET    MethodType = "GET"
	POST   MethodType = "POST"
	PUT    MethodType = "PUT"
	PATCH  MethodType = "PATCH"
	DELETE MethodType = "DELETE"
)

func (s MethodType) ToString() string {
	return string(s)
}

type ServicePort string

const LOCAL_SERVICE_PORT ServicePort = "8080"
