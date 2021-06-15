package nacos_helper

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/zhangsq-ax/nacos-helper-go/options"
)

var nacosClient *naming_client.INamingClient

// GetNamingClient Get Nacos naming client
func GetNamingClient(opts *options.NacosOptions) (*naming_client.INamingClient, error) {
	if nacosClient == nil {
		if opts == nil {
			var err error
			opts, err = options.GetNacosOptionsByEnv()
			if err != nil {
				return nil, err
			}
		}

		serverConfig := []constant.ServerConfig{
			{
				Scheme:      opts.Server.Scheme,
				IpAddr:      opts.Server.IpAddr,
				Port:        opts.Server.Port,
				ContextPath: opts.Server.ContextPath,
			},
		}

		clientConfig := constant.ClientConfig{
			NamespaceId:         opts.Client.NamespaceId,
			NotLoadCacheAtStart: true,
			LogLevel:            "debug",
			Username:            opts.Client.Username,
			Password:            opts.Client.Password,
			AppName:             opts.Client.AppName,
		}

		client, err := clients.NewNamingClient(vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfig,
		})

		if err != nil {
			return nil, err
		}

		nacosClient = &client
	}

	return nacosClient, nil
}

// RegisterServiceInstance Register a service instance
func RegisterServiceInstance(client *naming_client.INamingClient, opts *options.RegisterServiceOptions) error {
	var err error
	if client == nil {
		client, err = GetNamingClient(nil)
		if err != nil {
			return err
		}
	}

	_, err = (*client).RegisterInstance(*opts.GetRegisterInstanceParam())

	return err
}

// DeregisterServiceInstance Deregister a specified service instance
func DeregisterServiceInstance(client *naming_client.INamingClient, opts *options.DeregisterServiceOptions) error {
	var err error
	if client == nil {
		client, err = GetNamingClient(nil)
		if err != nil {
			return err
		}
	}

	_, err = (*client).DeregisterInstance(*opts.GetDeregisterInstanceParam())

	return err
}

// SelectServiceInstance Select an available service instance
func SelectServiceInstance(client *naming_client.INamingClient, opts *options.SelectServiceInstanceOptions) (*model.Instance, error) {
	var err error
	if client == nil {
		client, err = GetNamingClient(nil)
		if err != nil {
			return nil, err
		}
	}

	return (*client).SelectOneHealthyInstance(opts.SelectOneHealthInstanceParam())
}

// SubscribeServiceInstance Subscribe services change and select an available service instance
func SubscribeServiceInstance(client *naming_client.INamingClient, opts *options.SubscribeServiceInstanceOptions) error {
	var err error
	if client == nil {
		client, err = GetNamingClient(nil)
		if err != nil {
			return err
		}
	}

	err = (*client).Subscribe(opts.GetSubscribeParam(func(services []model.SubscribeService, err error) {
		if err != nil {
			opts.SubscribeCallback(nil, err)
		}
		var instance *model.Instance
		instance, err = SelectServiceInstance(client, opts.GetSelectServiceInstanceOptions())
		opts.SubscribeCallback(instance, err)
	}))

	return err
}
