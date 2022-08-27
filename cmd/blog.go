package cmd

import (
	"fmt"

	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/router"
)

func BlogServer() error {
	r := router.New()
	_ = r.SetTrustedProxies(nil)

	if err := r.Run(fmt.Sprintf(":%d", config.Conf.Port)); err != nil {
		panic(err)
	}

	return nil
}
