package options

import "github.com/nacos-group/nacos-sdk-go/vo"

// SelectServiceInstanceOptions The options for select the service instance
type SelectServiceInstanceOptions struct {
	Clusters    []string // The cluster name of the service instances
	ServiceName string   // The name of the service
	GroupName   string   // The group name of the service instances
}

func (ssio *SelectServiceInstanceOptions) SelectOneHealthInstanceParam() vo.SelectOneHealthInstanceParam {
	return vo.SelectOneHealthInstanceParam{
		Clusters:    ssio.Clusters,
		ServiceName: ssio.ServiceName,
		GroupName:   ssio.GroupName,
	}
}

func (ssio *SelectServiceInstanceOptions) SelectInstancesParam() vo.SelectInstancesParam {
	return vo.SelectInstancesParam{
		Clusters:    ssio.Clusters,
		ServiceName: ssio.ServiceName,
		GroupName:   ssio.GroupName,
		HealthyOnly: true,
	}
}

func (ssio *SelectServiceInstanceOptions) GetServiceParam() vo.GetServiceParam {
	return vo.GetServiceParam{
		Clusters:    ssio.Clusters,
		ServiceName: ssio.ServiceName,
		GroupName:   ssio.GroupName,
	}
}
