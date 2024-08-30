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
	log.Println("the market url is ", marketUrl)
	request, err := FacebookRequest(marketUrl)
	if err != nil {
		c.HTML(http.StatusBadRequest, constants.BaseTermplate, nil)
		return
	}

	marketData, err := GetMarket(c, string(request), marketID)
	if err != nil {
		c.HTML(http.StatusBadRequest, constants.BaseTermplate, nil)
		return
	}

	marketData.Url = marketUrl
	c.HTML(http.StatusOK, constants.BaseTermplate, marketData)
}

func GetMarket(c *gin.Context, request string, ID string) (HTMLData, error) {
	market, err := GetMarketContent(string(request))
	if err != nil {
		return HTMLData{}, err
	}
	if len(market.BboxInner.Result.Data.Viewer.MarketplaceProductDetailsPage.Target.ListingPhotos) == 0 {
		market.ListingPhotos, err = GetMarketPlaceListingPhotos(string(request))
		if err != nil {
			return HTMLData{}, err
		}
	} else {
		market.ListingPhotos = market.BboxInner.Result.Data.Viewer.MarketplaceProductDetailsPage.Target.ListingPhotos
	}
	var width, height string
	if len(market.ListingPhotos) > 0 {
		width = strconv.Itoa(market.ListingPhotos[0].Image.Width)
		height = strconv.Itoa(market.ListingPhotos[0].Image.Height)
	}
	// gridImages, err := gridHandler(c, market.ListingPhotos.LisitingPhotos, ID)
	// if err != nil {
	// 	log.Println(err)
	// 	return HTMLData{}, err
	// }
	htmlData := HTMLData{
		Title:       market.BboxInner.Result.Data.Viewer.MarketplaceProductDetailsPage.Target.MarketplaceListingTitle,
		Image:       market.ListingPhotos[0].Image.URI,
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
	//save responseMatch to file

	err := json.Unmarshal([]byte(responseMatch), &market)
	if err != nil {
		log.Println(err)
		return market, err
	}

	return market, nil
}

func GetMarketPlaceListingPhotos(request string) ([]ProductImage, error) {
	var listingPhotos []ProductImage
	re := regexp.MustCompile(constants.MARKET_LISTING_PHOTOS_REGEX)
	match := re.FindStringSubmatch(request)
	if len(match) == 0 {
		return nil, errors.New("marketplace listing not found because no match")
	}
	responseMatch := match[1]
	responseMatch = responseMatch[17:]
	err := json.Unmarshal([]byte(responseMatch), &listingPhotos)
	if err != nil {
		return nil, err
	}
	return listingPhotos, nil

}
