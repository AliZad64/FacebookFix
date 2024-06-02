package handlers

import (
	"encoding/json"
	"facebookfix/constants"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type PhotoParams struct {
	Fbid string `json:"fbid" form:"fbid"`
}

func GetPhotoHandler(c *gin.Context) {
	var photoParams PhotoParams
	if err := c.ShouldBindQuery(&photoParams); err != nil {
		c.HTML(http.StatusBadRequest, constants.BaseTermplate, nil)
		return
	}
	photoUrl := constants.PhotoURL + photoParams.Fbid
	request, err := FacebookRequest(photoUrl)
	if err != nil {
		c.HTML(http.StatusBadRequest, constants.BaseTermplate, nil)
		return
	}

	photo, err := GetPhoto(string(request))
	if err != nil {
		c.HTML(http.StatusBadRequest, constants.BaseTermplate, nil)
		return
	}

	photo.Url = photoUrl
	c.HTML(http.StatusOK, constants.BaseTermplate, photo)
}

func GetPhoto(request string) (HTMLData, error) {

	photo, err := GetPhotoContent(string(request))
	if err != nil {
		return HTMLData{}, err
	}

	media, err := GetPhotoMediaContent(string(request))
	if err != nil {
		return HTMLData{}, err
	}
	htmlData := HTMLData{
		Title:       media.Bbox.Result.Data.CreationStory.CometSections.ActorPhoto.Story.Actors[0].Name,
		Image:       photo.Data.CurrMedia.Image.URI,
		Video:       "",
		Card:        "summary_large_image",
		Description: media.Bbox.Result.Data.MessagePreferredBody.Text,
	}

	return htmlData, nil
}

func GetPhotoContent(request string) (PhotoStruct, error) {
	re := regexp.MustCompile(constants.PHOTO_REGEX)
	match := re.FindStringSubmatch(request)
	if len(match) == 0 {
		return PhotoStruct{}, nil
	}
	urlResponse := match[1]
	var photo PhotoStruct
	if err := json.Unmarshal([]byte(urlResponse), &photo); err != nil {
		return PhotoStruct{}, err
	}
	return photo, nil

}

func GetPhotoMediaContent(request string) (PhotoMedia, error) {
	re := regexp.MustCompile(constants.PHOTO_MEDIA_REGEX)
	match := re.FindStringSubmatch(request)
	if len(match) == 0 {
		return PhotoMedia{}, nil
	}
	urlResponse := match[1]
	urlResponse += "0}}"

	var photo PhotoMedia
	if err := json.Unmarshal([]byte(urlResponse), &photo); err != nil {
		return PhotoMedia{}, err
	}
	return photo, nil
}