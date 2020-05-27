package sdk

/*
 * Author: imuge
 * Date: 2020/5/27 下午5:28
 */

type ServiceEndpoint struct {
	Host    string
	Port    int32
	BaseUrl string
}

func NewServiceEndpoint(host string, port int32) ServiceEndpoint {
	return ServiceEndpoint{
		Host:    host,
		Port:    port,
		BaseUrl: "http://" + host + ":" + string(port),
	}
}
