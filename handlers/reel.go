package handlers

import (
	"encoding/json"
	"facebookfix/constants"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func GetReelHandler(c *gin.Context) {
	id := c.Param("id")
	reelUrl := constants.ReelURL + id

	request, err := FacebookRequest(reelUrl)
	if err != nil {
		log.Println("error1 ", err)
		c.Redirect(http.StatusFound, "/")
		return
	}
	log.Println("the request is ", string(request))
	data, err := GetReel(string(request))
	if err != nil {
		log.Println("error2 ", err)
		c.Redirect(http.StatusFound, "/")
		return
	}
	videoUrl := data.Video
	c.Redirect(http.StatusFound, videoUrl)
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
