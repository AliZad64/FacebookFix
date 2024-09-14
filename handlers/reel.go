package handlers

import (
	"encoding/json"
	"facebookfix/constants"
	"facebookfix/engine"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func GetReelHandler(c *gin.Context) {
	id := c.Param("id")
	reelUrl := constants.ReelURL + id

	videoTitle := fmt.Sprintf("/vid/%s", id)

	request, err := FacebookRequest(reelUrl)
	if err != nil {
		log.Println("error1 ", err)
		c.Redirect(http.StatusFound, "/")
		return
	}

	data, err := GetReel(string(request))
	if err != nil {
		log.Println("error2 ", err)
		c.Redirect(http.StatusFound, "/")
		return
	}
	key := fmt.Sprintf("video:%s", id)
	err = engine.SetRedisContent(c, key, data.Video)
	if err != nil {
		c.Error(err)
		return
	}

	data.Video = videoTitle
	c.HTML(http.StatusOK, constants.BaseTermplate, data)
}

func GetReel(html string) (HTMLData, error) {
	reelStruct, err := GetVideoContent(html)
	if err != nil {
		return HTMLData{}, err
	}

	mediaStruct, err := GetReelMediaContent(html)
	if err != nil {
		return HTMLData{}, err
	}
	data := HTMLData{
		Title:       mediaStruct.ShortFormVideoContext.VideoOwner.Name,
		Card:        "player",
		Description: "Reel",
		Video:       reelStruct.VideoID,
		Width:       "0",
		Height:      "0",
	}

	return data, nil
}

func GetReelMediaContent(request string) (ReelMediaStruct, error) {
	watchMetadataDataRegex := regexp.MustCompile(constants.REEL_MEDIA_REGEX)

	mediaMatches := watchMetadataDataRegex.FindStringSubmatch(request)
	var mediaString string
	if len(mediaMatches) > 0 {
		mediaString = mediaMatches[0]
		mediaString = mediaString[:len(mediaString)-13]
	} else {
		return ReelMediaStruct{}, nil
	}
	var mediaStruct ReelMediaStruct
	if len(mediaString) > 0 {
		err := json.Unmarshal([]byte(mediaString), &mediaStruct)
		if err != nil {
			return ReelMediaStruct{}, err
		}
	}
	return mediaStruct, nil
}
