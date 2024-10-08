package constants

const (
	VIDEO_MEDIA_REGEX           = `{"comet_sections":{"message":{"__typename":"CometFeedStoryDefaultMessageRenderingStrategy".*?"collaborators":`
	VIDEO_HD_REGEX              = `"browser_native_hd_url":"(.*?)"`
	VIDEO_SD_REGEX              = `"browser_sd_url":"(.*?)"`
	REEL_MEDIA_REGEX            = `{"short_form_video_context":{"self_view_boost":.*?},"extensions"`
	PHOTO_MEDIA_REGEX           = `({"__bbox":{"complete".+?{"__typename":"CometFeedStoryActorPhotoStrategy".+?"sequence_number":)`
	PHOTO_REGEX                 = `({"data":{"currMedia".+?"is_final".+?}})`
	MARKET_REGEX                = `("adp_MarketplacePDPContainerQueryRelayPreloader_[0-9a-f]{23}".+?{"__bbox":{.+?"sequence_number":)`
	MARKET_LISTING_PHOTOS_REGEX = `("listing_photos".+?])`
)
