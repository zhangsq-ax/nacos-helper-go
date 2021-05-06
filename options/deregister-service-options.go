package options

import "github.com/nacos-group/nacos-sdk-go/vo"

// DeregisterServiceOptions The options of deregister a service instance
type DeregisterServiceOptions struct {
	Ip          string // the IP of the service instance
	Port        uint64 // the port of the service instance listened
	ClusterName string // optional, default is DEFAULT
	ServiceName string // the name of the service
	GroupName   string // optional, default is DEFAULT_GROUP
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
