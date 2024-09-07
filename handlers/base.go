package handlers

import (
	"facebookfix/constants"
	"facebookfix/engine"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTMLData struct {
	Title       string
	Image       string
	Video       string
	Width       string
	Height      string
	Card        string
	Url         string
	Description string
}

func FacebookRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	engine.SetHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	request, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func GetEmbedHandler(c *gin.Context) {
	url := c.Request.URL.String()
	param := c.Query("v")
	videoTitle := fmt.Sprintf("/watch/%s", param)
	if param == "" {
		param = c.Param("id")
		videoTitle = fmt.Sprintf("/videos/%s", param)
	}

	htmlData := HTMLData{
		Title:       "Facebook",
		Image:       "",
		Video:       videoTitle,
		Card:        "player",
		Description: "",
		Width:       "0",
		Height:      "0",
		Url:         url,
	}
	c.HTML(http.StatusOK, constants.BaseTermplate, htmlData)
}
