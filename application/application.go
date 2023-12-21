package application

import (
	"github.com/incubus8/go/pkg/gin"
	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
	"ms-briapi-log-stream/config"
	"ms-briapi-log-stream/router"
)

func init() {
	gotenv.Load()
}

func StartApp() {
	addr := config.Config.ServiceHost + ":" + config.Config.ServicePort
	conf := gin.Config{
		ListenAddr: addr,
		Handler:    router.Router(),
		OnStarting: func() {
			log.Info().Msg("Your service is up and running at " + addr)
		},
	}

	gin.Run(conf)
}
