package options

import "github.com/nacos-group/nacos-sdk-go/vo"

type SelectServiceInstanceOptions struct {
	Clusters    []string
	ServiceName string
	GroupName   string
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
