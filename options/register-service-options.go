package options

import "github.com/nacos-group/nacos-sdk-go/vo"

type RegisterServiceOptions struct {
	Ip          string
	Port        uint64
	Weight      float64
	Metadata    map[string]string
	ClusterName string
	ServiceName string
	GroupName   string
}

func (rso *RegisterServiceOptions) GetRegisterInstanceParam() *vo.RegisterInstanceParam {
	weight := rso.Weight
	if weight <= 0 {
		weight = 10
	}
	return &vo.RegisterInstanceParam{
		Ip:          rso.Ip,
		Port:        rso.Port,
		Weight:      weight,
		Enable:      true,
		Healthy:     true,
		Metadata:    rso.Metadata,
		ClusterName: rso.ClusterName,
		ServiceName: rso.ServiceName,
		GroupName:   rso.GroupName,
		Ephemeral:   true,
	}
}
