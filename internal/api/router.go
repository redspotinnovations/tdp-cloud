package api

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/internal/midware"

	"tdp-cloud/internal/api/config"
	"tdp-cloud/internal/api/domain"
	"tdp-cloud/internal/api/machine"
	"tdp-cloud/internal/api/qcloud"
	"tdp-cloud/internal/api/socket"
	"tdp-cloud/internal/api/sshkey"
	"tdp-cloud/internal/api/tat_history"
	"tdp-cloud/internal/api/tat_script"
	"tdp-cloud/internal/api/user"
	"tdp-cloud/internal/api/vendor"
	"tdp-cloud/internal/api/workhub"
	"tdp-cloud/internal/api/worktask"
)

func Router(engine *gin.Engine) {

	api := engine.Group("/api")

	api.Use(midware.AbortHandle())

	{
		config.Router(api)
		domain.Router(api)
		machine.Router(api)
		qcloud.Router(api)
		vendor.Router(api)
		sshkey.Router(api)
		tat_history.Router(api)
		tat_script.Router(api)
		user.Router(api)
		workhub.Router(api)
		worktask.Router(api)
	}

	// websocket interface

	wsi := engine.Group("/wsi/:auth")

	wsi.Use(midware.SocketPreset())

	{
		socket.Socket(wsi)
	}

}