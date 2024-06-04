package handlers

import (
	"encoding/json"
	"errors"
	"facebookfix/constants"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetMarketHandler(c *gin.Context) {
	marketID := c.Param("id")
	marketUrl := constants.MarketUrl + marketID
	request, err := FacebookRequest(marketUrl)
	if err != nil {
		c.HTML(http.StatusBadRequest, constants.BaseTermplate, nil)
		return
	}

	marketData, err := GetMarket(c, string(request))
	if err != nil {
		c.HTML(http.StatusBadRequest, constants.BaseTermplate, nil)
		return
	}

	marketData.Url = marketUrl
	c.HTML(http.StatusOK, constants.BaseTermplate, marketData)
}

func GetMarket(c *gin.Context, request string) (HTMLData, error) {
	market, err := GetMarketContent(string(request))
	if err != nil {
		return HTMLData{}, err
	}

	width := strconv.Itoa(market.BboxInner.Result.Data.Viewer.MarketplaceProductDetailsPage.Target.ListingPhotos[0].Image.Width)
	height := strconv.Itoa(market.BboxInner.Result.Data.Viewer.MarketplaceProductDetailsPage.Target.ListingPhotos[0].Image.Height)
	log.Println(market)
	gridImages, err := gridHandler(c, market.BboxInner.Result.Data.Viewer.MarketplaceProductDetailsPage.Target.ListingPhotos, "test")
	log.Println(gridImages)
	htmlData := HTMLData{
		Title:       market.BboxInner.Result.Data.Viewer.MarketplaceProductDetailsPage.Target.MarketplaceListingTitle,
		Image:       gridImages,
		Video:       "",
		Width:       width,
		Height:      height,
		Card:        "summary_large_image",
		Url:         "",
		Description: market.BboxInner.Result.Data.Viewer.MarketplaceProductDetailsPage.Target.RedactedDescription.Text,
	}
	return htmlData, nil
}

func GetMarketContent(request string) (MarketSchema, error) {
	var market MarketSchema
	re := regexp.MustCompile(constants.MARKET_REGEX)
	match := re.FindStringSubmatch(request)
	if len(match) == 0 {
		return market, errors.New("marketplace listing not found because no match")
	}
	responseMatch := match[1]
	startKeyword := `{"__bbox":`
	startIndex := strings.Index(responseMatch, startKeyword)
	if startIndex == -1 {
		return market, errors.New("marketplace listing not found because no startIndex")
	}
	responseMatch = responseMatch[startIndex:] + "0}}"
	err := json.Unmarshal([]byte(responseMatch), &market)
	if err != nil {
		return market, err
	}

	return market, nil
}
