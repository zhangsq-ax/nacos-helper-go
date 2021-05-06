package options

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// NacosServerOptions The options of the Nacos server
type NacosServerOptions struct {
	Scheme      string `json:"scheme,omitempty" yaml:"scheme,omitempty"`           // optional
	ContextPath string `json:"contextPath,omitempty" yaml:"contextPath,omitempty"` // optional
	IpAddr      string `json:"ipAddr" yaml:"ipAddr"`                               // require, the IP or hostname of the Nacos server
	Port        uint64 `json:"port" yaml:"port"`                                   // require, the port of the Nacos server listened
}

// NacosClientOptions The options of the Nacos client
type NacosClientOptions struct {
	NamespaceId string `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"` // optional
	AppName     string `json:"appName" yaml:"appName"`                             // require, the application name of the client
	Username    string `json:"username,omitempty" yaml:"username,omitempty"`       // optional
	Password    string `json:"password,omitempty" yaml:"password,omitempty"`       // optional
}

// NacosOptions The options to connect Nacos server
type NacosOptions struct {
	Server NacosServerOptions `json:"server" yaml:"server"`
	Client NacosClientOptions `json:"client" yaml:"client"`
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

// GetNacosOptionsByEnv Get the Nacos server connection options through the environment variable
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
