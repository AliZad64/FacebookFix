package engine

import (
	"facebookfix/constants"

	"github.com/gin-gonic/gin"
)

func CustomHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Accept", constants.Accept)
		c.Writer.Header().Set("Accept-Language", constants.AcceptLanguage)
		c.Writer.Header().Set("Cache-Control", constants.CacheControl)
		c.Writer.Header().Set("Pragma", constants.Pragma)
		c.Writer.Header().Set("Sec-Ch-Prefers-Color-Scheme", constants.SecChPrefersColorScheme)
		c.Writer.Header().Set("Sec-Ch-UA", constants.SecChUA)
		c.Writer.Header().Set("Sec-Ch-UA-Full-Version-List", constants.SecChUAFullVersionList)
		c.Writer.Header().Set("Sec-Ch-UA-Mobile", constants.SecChUAMobile)
		c.Writer.Header().Set("Sec-Ch-UA-Platform", constants.SecChUAPlatform)
		c.Writer.Header().Set("Sec-Ch-UA-Platform-Version", constants.SecChUAPlatformVersion)
		c.Writer.Header().Set("Sec-Fetch-Dest", constants.SecFetchDest)
		c.Writer.Header().Set("Sec-Fetch-Mode", constants.SecFetchMode)
		c.Writer.Header().Set("Sec-Fetch-Site", constants.SecFetchSite)
		c.Writer.Header().Set("Sec-Fetch-User", constants.SecFetchUser)
		c.Writer.Header().Set("Upgrade-Insecure-Requests", constants.UpgradeInsecureRequests)
		c.Writer.Header().Set("User-Agent", constants.UserAgent)
	}
}
