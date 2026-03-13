// Separate show service entrypoint.
package main

import (
	"flag"
	"fmt"
	"net/http"

	"workspace/internal/config"
	"workspace/internal/handler"
	"workspace/internal/svc"
	"workspace/pkg/base62"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "config/shortener-show.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	fmt.Printf("Config: %+v\n", c)

	// base62 module init
	base62.MustInit(c.BaseString)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	server.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/:shortUrl",
			Handler: handler.ShowHandler(ctx),
		},
	})

	fmt.Printf("Starting show server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
