package main

import (
	"k8s_go_client/util/apibyk8s"
	"k8s_go_client/util/apibylocal"
)

// https://blog.csdn.net/weixin_45413603/article/details/107349333

func main() {
	apibylocal.ClientSetGetEndpoints()

	apibyk8s.ClientSetGetEndpoints()
	apibyk8s.RestClientGetPods()
	apibyk8s.ClientSetGetPods()
	apibyk8s.DiscoveryClient()
	apibyk8s.DynamicClientGetPods()
	select {

	}
}