package options

import (
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type SubscribeServiceInstanceOptions struct {
	Clusters          []string // The cluster name of the service instances
	ServiceName       string   // The name of the service
	GroupName         string   // The group name of the service instances
	SubscribeCallback func(*model.Instance, error)
}

func (ssio *SubscribeServiceInstanceOptions) GetSelectServiceInstanceOptions() *SelectServiceInstanceOptions {
	return &SelectServiceInstanceOptions{
		Clusters:    ssio.Clusters,
		ServiceName: ssio.ServiceName,
		GroupName:   ssio.GroupName,
	}
}

func (ssio *SubscribeServiceInstanceOptions) GetSubscribeParam(callback func([]model.SubscribeService, error)) *vo.SubscribeParam {
	return &vo.SubscribeParam{
		Clusters:          ssio.Clusters,
		ServiceName:       ssio.ServiceName,
		GroupName:         ssio.GroupName,
		SubscribeCallback: callback,
	}
}
