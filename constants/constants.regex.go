package constants

const (
	VIDEO_MEDIA_REGEX = `{"comet_sections":{"message":{"__typename":"CometFeedStoryDefaultMessageRenderingStrategy".*?"collaborators":`
	VIDEO_HD_REGEX    = `"browser_native_hd_url":"(.*?)"`
	VIDEO_SD_REGEX    = `"browser_sd_url":"(.*?)"`
	REEL_MEDIA_REGEX  = `{"short_form_video_context":{"self_view_boost":.*?},"extensions"`
	PHOTO_MEDIA_REGEX = `({"__bbox":{"complete".+?{"__typename":"CometFeedStoryActorPhotoStrategy".+?"sequence_number":)`
	PHOTO_REGEX       = `({"data":{"currMedia".+?"is_final".+?}})`
)
