package nacos_helper

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/zhangsq-ax/nacos-helper-go/options"
)

var namingClient *naming_client.INamingClient
var configClient *config_client.IConfigClient

// GetConfigClient Get Nacos config client
func GetConfigClient(opts *options.NacosOptions) (*config_client.IConfigClient, error) {
	if configClient == nil {
		if opts == nil {
			var err error
			opts, err = options.GetNacosOptionsByEnv()
			if err != nil {
				return nil, err
			}
		}

		client, err := clients.NewConfigClient(*opts.GetNacosClientParam())
		if err != nil {
			return nil, err
		}

		configClient = &client
	}

	return configClient, nil
}

// GetNamingClient Get Nacos naming client
func GetNamingClient(opts *options.NacosOptions) (*naming_client.INamingClient, error) {
	if namingClient == nil {
		if opts == nil {
			var err error
			opts, err = options.GetNacosOptionsByEnv()
			if err != nil {
				return nil, err
			}
		}

		client, err := clients.NewNamingClient(*opts.GetNacosClientParam())
		if err != nil {
			return nil, err
		}

		namingClient = &client
	}

	return namingClient, nil
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
