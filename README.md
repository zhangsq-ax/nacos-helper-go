# nacos-helper-go
Nacos helper for Go language

## Install

```bash
$ go get github.com/zhangsq-ax/nacos-helper-go
```

## The environment variables of Nacos

Name | Required | Description | Example
---- | ---- | ---- | ----
`NACOS_HOST` | true | The IP address of host name of the Nacos server | `"127.0.0.1"`
`NACOS_PORT` | true | The port that the Nacos server listened | `"8848"`
`NACOS_SCHEME` | false | The URL scheme of the Nacos server | `"https"`
`NACOS_CONTEXT_PATH` | false | The URL context path of the Nacos server | `"/nacos"`
`NACOS_NAMESPACE_ID` | false | The namespace ID of the client in the Nacos server | -
`NACOS_APP_NAME` | false | The name of the application to which the client belongs | -
`NACOS_USERNAME` | false | The username that connect to the Nacos server | -
`NACOS_PASSWORD` | false | The password that connect to the Nacos server | -

## Usage

### Register a service instance

```go
package main

import (
	nacos_helper "github.com/zhangsq-ax/nacos-helper-go"
	"log"
)

func main() {
	...
	// use environment variables
	err := nacos_helper.RegisterServiceInstance(nil, &nacos_helper.options.RegisterServiceOptions{
		Ip:"127.0.0.1",
		Port: 8080,
		ServiceName: "foo",
	})
	if err != nil {
		log.Fatalln(err)
	}
	...
}
```

### Select an available service instance

```go
package main

import (
	nacos_helper "github.com/zhangsq-ax/nacos-helper-go"
	"log"
	"fmt"
)

func main() {
	// use environment variables
	instance, err := nacos_helper.SelectServiceInstance(nil, &nacos_helper.options.SelectServiceInstanceOptions{
	    ServiceName: "foo",
	})
	if err != nil {
		log.Fatalln(err)
	}
	
	fmt.Println(instance.Ip)
	fmt.Println(instance.Port)
	
	...
}
```