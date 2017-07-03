package collect

import (
	"dao"
	"common"
	"strconv"
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	model "model/collect"
	"log"
)

var PodList model.PodList
var NodeList model.NodeList
var ServiceList model.ServiceList_k
var all_containers int
var KuberMasterIp string
var KuberMasterStatus bool

func init() {

	KuberMasterIp = "http://10.110.18.107:8080"
	_, err := http.Get(KuberMasterIp + "/version")
	if err != nil {
		common.LogErr(err)
		KuberMasterStatus = false
	} else {
		KuberMasterStatus = true
	}
	log.Printf("%s\t%s\t%s\t", "KuberMasterStatus status is ", strconv.FormatBool(KuberMasterStatus), time.Now())

}
func get_time() string {
	get_time := time.Now().UTC().Add(8 * time.Hour).Format("2006-01-02 15:04:05")
	return get_time
}
func GainResourceFromK8s(resource interface{}, urls string) error {
	resp, err := http.Get(KuberMasterIp + urls)
	common.LogErr(err)
	body, err := ioutil.ReadAll(resp.Body)
	common.LogErr(err)
	err = json.Unmarshal(body, &resource)
	common.LogErr(err)
	return nil
}

func (resource KubernetesAllResource) GainPods() error {
	GainResourceFromK8s(&PodList, "/api/v1/pods")
	i := 0
	n_containers := 0
	tagTemp := common.Gen_id(5)
	for _, v := range PodList.Items {
		var x = resource.pods
		i = i + 1
		x.All_pod_numbers = strconv.Itoa(len(PodList.Items))
//		x.Pod_No = strconv.Itoa(i)
		x.Pod_hostIP = v.Status.HostIP
		x.Pod_name = v.Name
		x.Create_time = v.CreationTimestamp.Format("2006-01-02 15:04:05")
		x.Record_time = get_time()
		x.Containers_numbers = strconv.Itoa(len(v.Status.Conditions))
		tem_containers := len(v.Status.Conditions)
		n_containers = tem_containers + n_containers
		if i == len(PodList.Items) {
			all_containers = n_containers
		}
		x.All_container_numbers = strconv.Itoa(all_containers)
		x.Tag = tagTemp
		dao.Db_insert(&x)

	}
	common.DebugPrint("pods is insert")
	ThreadCountGet.Done()
	return nil
}

func (resource KubernetesAllResource) GainNodes() error {
	GainResourceFromK8s(&NodeList, "/api/v1/nodes")

	tagTemp := common.Gen_id(5)
	for _, v := range NodeList.Items {
		var nodes = resource.nodes
		nodes.Node_name = v.Name
		nodes.Create_time = v.CreationTimestamp.Format("2006-01-02 15:04:05")
		fmt.Println(len(v.Status.Images))
		for k, v := range v.Status.Capacity {
			switch k {
			case "cpu":
				nodes.Numbers_cpu_core = v.String()
			case "memory":
				nodes.Memory_size = v.String()
			case "alpha.kubernetes.io/nvidia-gpu":
				nodes.Numbers_gpu_core = v.String()
			case "pods":
				nodes.Pod_limit = v.String()
			}
		}
		nodes.Record_time = get_time()
		nodes.Tag = tagTemp
		dao.Db_insert(&nodes)

	}
	common.DebugPrint("nodes is insert")
	ThreadCountGet.Done()
	return nil

}

func (resource KubernetesAllResource) GainServices() error {
	GainResourceFromK8s(&ServiceList, "/api/v1/services")
	tagTemp := common.Gen_id(5)
	for _, v := range ServiceList.Items {
		var service = resource.services
		service.Create_time = v.CreationTimestamp.Format("2006-01-02 15:04:05")
		service.Service_name = v.Name
		service.Service_numbers = strconv.Itoa(len(ServiceList.Items))
		service.Record_time = get_time()
		for k,v:=range v.Labels {
			common.DebugPrint(k,v)
		}

		service.Tag = tagTemp
		dao.Db_insert(&service)
	}
	common.DebugPrint("service is insert")
	ThreadCountGet.Done()
	return nil
}
