package handlers

import (
	"encoding/json"
	"facebookfix/constants"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type VideoQueryParams struct {
	V string `form:"v"`
}

type VideoStruct struct {
	VideoID string `json:"videoID"`
}

func GetVideoHandler(c *gin.Context) {
	id := c.Param("id")
	var queryParam VideoQueryParams
	if err := c.ShouldBindQuery(&queryParam); err != nil {
		c.HTML(http.StatusBadRequest, constants.BaseTermplate, nil)
		return
	}
	var videoUrl string
	if len(queryParam.V) > 0 {
		videoUrl = constants.VideoURL + queryParam.V
	}
	if len(id) > 0 {
		videoUrl = constants.VideoURL + id
	}

	if len(queryParam.V) == 0 {
		videoUrl = constants.VideoURL + id
	}
	request, err := FacebookRequest(videoUrl)
	if err != nil {
		c.HTML(http.StatusBadRequest, constants.BaseTermplate, nil)
		return
	}

	data, err := GetVideo(string(request))
	if err != nil {
		c.HTML(http.StatusBadRequest, constants.BaseTermplate, nil)
		return
	}
	data.Url = videoUrl
	c.HTML(http.StatusOK, "base.tmpl", data)
}

func GetVideo(request string) (HTMLData, error) {
	videoStruct, err := GetVideoContent(request)
	if err != nil {
		return HTMLData{}, err
	}

	mediaStruct, err := GetMediaContent(request)
	if err != nil {
		return HTMLData{}, err
	}

	data := HTMLData{
		Title:       mediaStruct.CometSections.ActorPhoto.Story.Actors[0].Name,
		Card:        "player",
		Description: mediaStruct.CometSections.Message.Story.Message.Text,
		Video:       videoStruct.VideoID,
		Width:       "720",
		Height:      "1280",
	}

	return data, nil
}

func GetVideoContent(request string) (VideoStruct, error) {
	re := regexp.MustCompile(constants.VIDEO_HD_REGEX)
	matches := re.FindStringSubmatch(request)
	if len(matches) == 0 {
		re = regexp.MustCompile(constants.VIDEO_SD_REGEX)
		matches = re.FindStringSubmatch(request)
	}

	var urlResponse string
	if len(matches) == 2 {
		urlResponse = matches[1]
	} else {
		return VideoStruct{}, fmt.Errorf("no match found")
	}

	jsonString := fmt.Sprintf("{\"videoID\":\"%s\"}", urlResponse)
	var videoStruct VideoStruct
	if err := json.Unmarshal([]byte(jsonString), &videoStruct); err != nil {
		return VideoStruct{}, err
	}
	return videoStruct, nil
}

func GetMediaContent(request string) (MediaStruct, error) {
	watchMetadataDataRegex := regexp.MustCompile(constants.VIDEO_MEDIA_REGEX)

	mediaMatches := watchMetadataDataRegex.FindStringSubmatch(request)
	var mediaString string
	if len(mediaMatches) > 0 {
		mediaString = mediaMatches[0] + "[]}"
	} else {
		return MediaStruct{}, nil
	}
	var mediaStruct MediaStruct
	if len(mediaString) > 0 {
		err := json.Unmarshal([]byte(mediaString), &mediaStruct)
		if err != nil {
			return MediaStruct{}, err
		}
	}
	if len(mediaStruct.CometSections.ActorPhoto.Story.Actors) == 0 {
		mediaStruct.CometSections.ActorPhoto.Story.Actors = append(mediaStruct.CometSections.ActorPhoto.Story.Actors, Actor{
			Name: "Facebook",
		})
	}

	return mediaStruct, nil
}
