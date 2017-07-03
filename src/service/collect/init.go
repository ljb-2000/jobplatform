package collect

import (
	model "model/collect"
	"sync"
)

var ThreadCountGet sync.WaitGroup

func RunOneCycle() error {
	var a GainKubernetes = new(KubernetesAllResource)
	gainResource(a)
	ThreadCountGet.Wait()
	return nil
}

type KubernetesAllResource struct {
	nodes    model.Nodes
	pods     model.Pods
	services model.Services
}
type GainKubernetes interface {
	GainPods() error
	GainNodes() error
	GainServices() error
}

func gainResource(a GainKubernetes) {
	ThreadCountGet.Add(3)
	go a.GainPods()
	go a.GainNodes()
	go a.GainServices()

}
