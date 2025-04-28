package options

import "github.com/nacos-group/nacos-sdk-go/vo"

type SubscribeConfigOptions struct {
	DataId   string
	Group    string
	OnChange func(namespace, group, dataId, data string)
}

func (sco *SubscribeConfigOptions) GetConfigParam() vo.ConfigParam {
	return vo.ConfigParam{
		DataId:   sco.DataId,
		Group:    sco.Group,
		OnChange: sco.OnChange,
	}
}
