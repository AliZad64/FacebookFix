package handlers

type VideoStruct struct {
	VideoID string `json:"videoID"`
}
type CometFeedStoryDefaultMessageRenderingStrategy struct {
	Story MediaStory `json:"story"`
}

type MediaStory struct {
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
}

type MediaStruct struct {
	CometSections CometSections `json:"comet_sections"`
}

type VideoCreationStory struct {
	PostID string `json:"post_id"`
	ID     string `json:"id"`
}

type DisplayPicture struct {
	Uri string `json:"uri"`
}

type Feedback struct {
	AssociatedGroup interface{} `json:"associated_group"`
	ID              string      `json:"id"`
}

type ShortFormVideoContext struct {
	SelfViewBoost interface{} `json:"self_view_boost"`
	Video         struct {
		CreationStory     VideoCreationStory `json:"creation_story"`
		ID                string             `json:"id"`
		VideoCollaborator interface{}        `json:"video_collaborator"`
	} `json:"video"`
	VideoOwner             VideoOwner `json:"video_owner"`
	IsPassiveContent       bool       `json:"is_passive_content"`
	FbShortsReshareContext struct {
		IsReshare      bool `json:"is_reshare"`
		ReshareCreator struct {
			Typename               string `json:"__typename"`
			IsActor                string `json:"__isActor"`
			ID                     string `json:"id"`
			Name                   string `json:"name"`
			EnableReelsTabDeeplink bool   `json:"enable_reels_tab_deeplink"`
			IsVerified             bool   `json:"is_verified"`
			Url                    string `json:"url"`
		} `json:"reshare_creator"`
	} `json:"fb_shorts_reshare_context"`
	VideoOwnerType string `json:"video_owner_type"`
}

type VideoOwner struct {
	Typename               string         `json:"__typename"`
	DelegatePageID         string         `json:"delegate_page_id"`
	ID                     string         `json:"id"`
	IsActor                string         `json:"__isActor"`
	Name                   string         `json:"name"`
	EnableReelsTabDeeplink bool           `json:"enable_reels_tab_deeplink"`
	IsVerified             bool           `json:"is_verified"`
	Url                    string         `json:"url"`
	DisplayPicture         DisplayPicture `json:"displayPicture"`
	SubscribeStatus        string         `json:"subscribe_status"`
	DelegatePage           DelegatePage   `json:"delegate_page"`
}

type ReelMediaStruct struct {
	ShortFormVideoContext  ShortFormVideoContext `json:"short_form_video_context"`
	PostID                 string                `json:"post_id"`
	Tracking               string                `json:"tracking"`
	CreationTime           int64                 `json:"creation_time"`
	Feedback               Feedback              `json:"feedback"`
	BrandedContentPostInfo interface{}           `json:"branded_content_post_info"`
}

type DelegatePage struct {
	IsBusinessPageActive bool   `json:"is_business_page_active"`
	ID                   string `json:"id"`
}

type Actor struct {
	Name string `json:"name"`
}

type Story struct {
	Actors []Actor `json:"actors"`
}

type ActorPhoto struct {
	Story Story `json:"story"`
}

type CometSections struct {
	Message struct {
		CometFeedStoryDefaultMessageRenderingStrategy
	} `json:"message"`
	ActorPhoto ActorPhoto `json:"actor_photo"`
}

type MessagePreferredBody struct {
	Text string `json:"text"`
}

type Scope struct {
	Label               string    `json:"label"`
	IconImage           IconImage `json:"icon_image"`
	CanViewerEdit       bool      `json:"can_viewer_edit"`
	Description         string    `json:"description"`
	ExtendedDescription string    `json:"extended_description"`
}

type IconImage struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Scale  int    `json:"scale"`
	URI    string `json:"uri"`
	Width  int    `json:"width"`
}

type PrefetchUri struct {
	URI   string      `json:"uri"`
	Label interface{} `json:"label"`
}

type Data struct {
	CurrMedia            CurrMedia            `json:"currMedia"`
	MessagePreferredBody MessagePreferredBody `json:"message_preferred_body"`
	CreationStory        CreationStory        `json:"creation_story"`
}

type CurrMedia struct {
	Image         Image         `json:"image"`
	CreationStory CreationStory `json:"creation_story"`
}

type CodedException struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
	Code        int    `json:"code"`
}

type AdsLwiPost struct {
	Story Story `json:"story"`
}

type Image struct {
	URI    string `json:"uri"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type CreationStory struct {
	CometSections CometSections `json:"comet_sections"`
}

type PhotoStruct struct {
	Data Data `json:"data"`
}

type Result struct {
	Data Data `json:"data"`
}

type Bbox struct {
	Result Result `json:"result"`
}

type PhotoMedia struct {
	Bbox Bbox `json:"__bbox"`
}

type LatitudeLongitude struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type MarketActor struct {
	Typename string `json:"__typename"`
	Name     string `json:"name"`
	ID       string `json:"id"`
}

type Tracking struct {
	StoryLocation        int    `json:"story_location"`
	StoryAttachmentStyle string `json:"story_attachment_style"`
	EntAttachementType   string `json:"ent_attachement_type"`
	Actrs                string `json:"actrs"`
}

type RedactedDescription struct {
	Text string `json:"text"`
}

type LocationText struct {
	Text string `json:"text"`
}

type MarketplaceLeadGenForm struct {
	ViewerPurchaseLimit       int         `json:"viewer_purchase_limit"`
	BoostedMarketplaceListing interface{} `json:"boosted_marketplace_listing"`
	PromotedListing           interface{} `json:"promoted_listing"`
}

type ReverseGeocodeDetailed struct {
	CountryAlphaTwo   string `json:"country_alpha_two"`
	PostalCodeTrimmed string `json:"postal_code_trimmed"`
}

type Location struct {
	Latitude               float64                `json:"latitude"`
	Longitude              float64                `json:"longitude"`
	ReverseGeocodeDetailed ReverseGeocodeDetailed `json:"reverse_geocode_detailed"`
}

type ProductImage struct {
	Typename             string `json:"__typename"`
	AccessibilityCaption string `json:"accessibility_caption"`
	Image                struct {
		Height int    `json:"height"`
		Width  int    `json:"width"`
		URI    string `json:"uri"`
	} `json:"image"`
	ID string `json:"id"`
}

type MarketplaceProductDetailsPage struct {
	Target Target `json:"target"`
}

type Target struct {
	MarketplaceListingTitle string              `json:"marketplace_listing_title"`
	RedactedDescription     RedactedDescription `json:"redacted_description"`
	LocationText            LocationText        `json:"location_text"`         // might be used later
	PrimaryListingPhoto     ProductImage        `json:"primary_listing_photo"` // might be used later
	Location                Location            `json:"location"`              // might be used later
	ListingPhotos           []ProductImage      `json:"listing_photos"`
}
type Viewer struct {
	MarketplaceProductDetailsPage MarketplaceProductDetailsPage `json:"marketplace_product_details_page"`
}

type MarketData struct {
	Viewer Viewer `json:"viewer"`
}

type MarketResult struct {
	Data MarketData `json:"data"`
}

type BboxInner struct {
	Result MarketResult `json:"result"`
}

type MarketSchema struct {
	BboxInner     BboxInner      `json:"__bbox"`
	ListingPhotos []ProductImage `json:"listing_photos"`
}
