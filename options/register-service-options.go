package options

import "github.com/nacos-group/nacos-sdk-go/vo"

// RegisterServiceOptions The options to register a service instance
type RegisterServiceOptions struct {
	Ip          string            `json:"ip" yaml:"ip"`                                       // The IP address or host name that the service instance is bound to
	Port        uint64            `json:"port" yaml:"port"`                                   // The port that the service instance listened
	Weight      float64           `json:"weight,omitempty" yaml:"weight,omitempty"`           // Weight of the call, default is 10
	Metadata    map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`       // The metadata of the service instance
	ClusterName string            `json:"clusterName,omitempty" yaml:"clusterName,omitempty"` // The cluster name of the service instance, default is DEFAULT
	ServiceName string            `json:"serviceName" yaml:"serviceName"`                     // The name of the service
	GroupName   string            `json:"groupName,omitempty" yaml:"groupName,omitempty"`     // The group name of the service instance, default is DEFAULT_GROUP
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
