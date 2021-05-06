package options

import "github.com/nacos-group/nacos-sdk-go/vo"

type DeregisterServiceOptions struct {
	Ip          string
	Port        uint64
	ClusterName string
	ServiceName string
	GroupName   string
}

func (drso *DeregisterServiceOptions) GetDeregisterInstanceParam() *vo.DeregisterInstanceParam {
	return &vo.DeregisterInstanceParam{
		Ip:          drso.Ip,
		Port:        drso.Port,
		Cluster:     drso.ClusterName,
		ServiceName: drso.ServiceName,
		GroupName:   drso.GroupName,
	}
}
