package options

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type NacosServerOptions struct {
	Scheme      string
	ContextPath string
	IpAddr      string
	Port        uint64
}

type NacosClientOptions struct {
	NamespaceId string
	AppName     string
	Username    string
	Password    string
}

type NacosOptions struct {
	Server NacosServerOptions
	Client NacosClientOptions
}

func getEnvString(key string) string {
	return os.Getenv(key)
}

func getEnvUint64(key string) (uint64, error) {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return 0, err
	}
	return uint64(val), nil
}

func GetNacosOptionsByEnv() (*NacosOptions, error) {
	port, err := getEnvUint64("NACOS_PORT")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get Nacos server port: %v", err))
	}

	return &NacosOptions{
		Server: NacosServerOptions{
			IpAddr:      getEnvString("NACOS_HOST"),
			Port:        port,
			Scheme:      getEnvString("NACOS_SCHEME"),
			ContextPath: getEnvString("NACOS_CONTEXT_PATH"),
		},
		Client: NacosClientOptions{
			NamespaceId: getEnvString("NACOS_NAMESPACE_ID"),
			AppName:     getEnvString("NACOS_APP_NAME"),
			Username:    getEnvString("NACOS_USERNAME"),
			Password:    getEnvString("NACOS_PASSWORD"),
		},
	}, nil
}
