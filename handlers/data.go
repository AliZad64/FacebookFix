package handlers

import (
	"facebookfix/engine"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetVideoDataHandler(ctx *gin.Context) {
	param := ctx.Param("id")
	if param == "" {
		ctx.Error(fmt.Errorf("no video id found"))
		return
	}
	log.Info().Msgf("test endpoint and video id: %s", param)

	key := fmt.Sprintf("video:%s", param)
	videoURL, err := engine.GetRedisContent(ctx, key)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Redirect(http.StatusFound, videoURL)
}
