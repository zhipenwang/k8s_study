package apibyk8s

import (
	"context"
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	clientcmdV1 "k8s.io/client-go/tools/clientcmd/api/v1"
)

// https://blog.csdn.net/weixin_45413603/article/details/107349333

func DiscoveryClient()  {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	client, err := discovery.NewDiscoveryClientForConfig(config)
	fmt.Println(err)
	_, apiSource, err := client.ServerGroupsAndResources()
	fmt.Println(err)
	for _, v := range apiSource {
		gv, err := schema.ParseGroupVersion(v.GroupVersion)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, resource := range v.APIResources {
			fmt.Println(resource.Name, resource.Group, resource.Namespaced, gv.Group, gv.Version)
		}
	}
}

func DynamicClientGetPods()  {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	client, err := dynamic.NewForConfig(config)
	fmt.Println(err)
	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	gvrObj, err := client.Resource(gvr).Namespace("kubernetes-dashboard").List(context.Background(), metaV1.ListOptions{Limit: 500})
	fmt.Println(err)
	podList := &coreV1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(gvrObj.UnstructuredContent(), podList)
	fmt.Println(err)
	for _, v := range podList.Items {
		fmt.Println(v.Namespace, v.Name, v.Status.PodIPs, v.Status.Phase)
	}
}

func ClientSetGetPods()  {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	client, err := kubernetes.NewForConfig(config)
	fmt.Println(err)
	pods, err := client.CoreV1().Pods("kubernetes-dashboard").List(context.Background(), metaV1.ListOptions{})
	fmt.Println(err)
	for _, v := range pods.Items {
		fmt.Println(v.Namespace, v.Name, v.Status.PodIP)
	}
}

func ClientSetGetEndpoints()  {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	client, err := kubernetes.NewForConfig(config)
	fmt.Println(err)
	//cs.CoreV1().Endpoints(svc.Namespace).Get(context.Background(), svc.Name, v1.GetOptions{})
	endpoints, err := client.CoreV1().Endpoints("kubernetes-dashboard").Get(context.Background(), "nginx", metaV1.GetOptions{})
	for _, v := range endpoints.Subsets {
		//fmt.Println(v.Addresses, v.Ports)
		for _, vv := range v.Addresses {
			fmt.Println(vv.Hostname, vv.IP, *vv.NodeName)
		}
	}
	//endpointList, err := client.CoreV1().Endpoints("kubernetes-dashboard").List(context.Background(), metaV1.ListOptions{})
	//fmt.Println(err)
	//for _, v := range endpointList.Items {
	//	fmt.Println(v.Namespace, v.Name, v.Subsets)
	//}
}

func RestClientGetPods() {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	config.Host = "https://kubernetes.docker.internal:6443"
	config.APIPath = "api"
	config.GroupVersion = &clientcmdV1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	restClient, err := rest.RESTClientFor(config)
	fmt.Println(err)
	result := &coreV1.PodList{}
	err = restClient.Get().
		Namespace("kubernetes-dashboard").
		Resource("pods").
		VersionedParams(&metaV1.ListOptions{Limit: 500}, scheme.ParameterCodec).
		Do(context.Background()).Into(result)
	fmt.Println(err)
	//fmt.Println(result)
	for _, v := range result.Items {
		//fmt.Printf("%+#v\n", v)
		fmt.Println(v.Namespace, v.Name, v.Status.PodIPs, v.ClusterName, v.Labels, v.Kind, v.UID)
	}
	return
}

