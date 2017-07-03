package collect

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Nodes), new(Pods), new(Services))
}
type Nodes struct {
	Id               int64    `json:"id" orm:"pk;auto"`
	Node_name        string `json:"pod_name" orm:"column(Node_name)"`
	Numbers_cpu_core string  `json:"pod_name" orm:"column(Numbers_cpu_core)"`
	Numbers_gpu_core string  `json:"pod_name" orm:"column(Numbers_gpu_core)"`
	Memory_size      string  `json:"pod_name" orm:"column(Memory_size)"`
	Pod_limit        string  `json:"pod_name" orm:"column(Pod_limit)"`
	Create_time      string `json:"Creat_time" orm:"column(Create_time)"`
	Record_time      string `json:"Record_time" orm:"column(Record_time)"`
	Tag              string `json:"tag" orm:"column(tag)"`
}

type Pods struct {
	Id                   int64    `json:"id" orm:"pk;auto"`
	Pod_name              string `json:"pod_name" orm:"column(pod_name)"`
	Pod_hostIP            string`json:"pod_name" orm:"column(pod_hostIP)"`
	Containers_numbers    string`json:"pod_name" orm:"column(containers_numbers)"`
	Create_time           string `json:"Creat_time" orm:"column(create_time)"`
	Record_time           string `json:"Record_time" orm:"column(Record_time)"`
	All_pod_numbers       string `json:"All_pod_numbers" orm:"column(All_pod_numbers)"`
	All_container_numbers string `json:"All_container_numbers" orm:"column(All_container_numbers)"`
	Tag                   string `json:"tag" orm:"column(tag)"`
}

type Services struct {
	Id             int64    `json:"id" orm:"pk;auto"`
	Service_name    string `json:"pod_name" orm:"column(Service_name)"`
	Service_numbers string `json:"pod_name" orm:"column(Service_numbers)"`
	Create_time     string `json:"Creat_time" orm:"column(Creat_time)"`
	Record_time     string `json:"Record_time" orm:"column(Record_time)"`
	Tag             string `json:"tag" orm:"column(tag)"`
}


