package nacos_helper

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"nacos_helper/options"
)

var nacosClient *naming_client.INamingClient

func GetClient(opts *options.NacosOptions) (*naming_client.INamingClient, error) {
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

func RegisterServiceInstance(client *naming_client.INamingClient, opts *options.RegisterServiceOptions) error {
	var err error
	if client == nil {
		client, err = GetClient(nil)
		if err != nil {
			return err
		}
	}

	_, err = (*client).RegisterInstance(*opts.GetRegisterInstanceParam())

	return err
}

func DeregisterServiceInstance(client *naming_client.INamingClient, opts *options.DeregisterServiceOptions) error {
	var err error
	if client == nil {
		client, err = GetClient(nil)
		if err != nil {
			return err
		}
	}

	_, err = (*client).DeregisterInstance(*opts.GetDeregisterInstanceParam())

	return err
}

func SelectServiceInstance(client *naming_client.INamingClient, opts *options.SelectServiceInstanceOptions) (*model.Instance, error) {
	var err error
	if client == nil {
		client, err = GetClient(nil)
		if err != nil {
			return nil, err
		}
	}

	return (*client).SelectOneHealthyInstance(opts.SelectOneHealthInstanceParam())
}
