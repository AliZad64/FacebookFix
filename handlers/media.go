package handlers

type CometFeedStoryDefaultMessageRenderingStrategy struct {
	TypeName string     `json:"__typename"`
	Story    MediaStory `json:"story"`
}

type MediaStory struct {
	Message struct {
		Ranges []struct {
			Entity struct {
				TypeName  string `json:"__typename"`
				Name      string `json:"name"`
				IsNode    string `json:"__isNode"`
				ID        string `json:"id"`
				IsEntity  string `json:"__isEntity"`
				URL       string `json:"url"`
				MobileURL string `json:"mobileUrl"`
			} `json:"entity"`
			EntityIsWeakReference bool `json:"entity_is_weak_reference"`
			Length                int  `json:"length"`
			Offset                int  `json:"offset"`
		} `json:"ranges"`
		Text string `json:"text"`
	} `json:"message"`
	ID string `json:"id"`
}

type CometFeedStoryActorPhotoStrategy struct {
	TypeName string `json:"__typename"`
	Story    Story  `json:"story"`
}

type CometFeedStoryMinimizedTimestampStrategy struct {
	TypeName string `json:"__typename"`
	Story    struct {
		CreationTime int         `json:"creation_time"`
		URL          string      `json:"url"`
		GhlLabel     interface{} `json:"ghl_label"`
		ID           string      `json:"id"`
	} `json:"story"`
}

type CometFeedStoryAudienceStrategy struct {
	TypeName             string `json:"__typename"`
	IsICometStorySection string `json:"__isICometStorySection"`
	IsProdEligible       bool   `json:"is_prod_eligible"`
	Story                struct {
		PrivacyScope struct {
			IconImage struct {
				Name string `json:"name"`
			} `json:"icon_image"`
			Description string `json:"description"`
		} `json:"privacy_scope"`
		ID string `json:"id"`
	} `json:"story"`
	ModuleOperationCometFeedStoryMetadataSectionStory struct {
		DR string `json:"__dr"`
	} `json:"__module_operation_CometFeedStoryMetadataSection_story"`
	ModuleComponentCometFeedStoryMetadataSectionStory struct {
		DR string `json:"__dr"`
	} `json:"__module_component_CometFeedStoryMetadataSection_story"`
}

type CometFeedStoryTitleWithActorStrategy struct {
	TypeName             string `json:"__typename"`
	IsICometStorySection string `json:"__isICometStorySection"`
	IsProdEligible       bool   `json:"is_prod_eligible"`
	Story                struct {
		ID     string `json:"id"`
		Actors []struct {
			TypeName              string      `json:"__typename"`
			Name                  string      `json:"name"`
			ID                    string      `json:"id"`
			IsActor               string      `json:"__isActor"`
			IsEntity              string      `json:"__isEntity"`
			URL                   string      `json:"url"`
			WorkForeignEntityInfo interface{} `json:"work_foreign_entity_info"`
			WorkInfo              interface{} `json:"work_info"`
		} `json:"actors"`
		Title         interface{} `json:"title"`
		CometSections struct {
			ActionLink interface{} `json:"action_link"`
		} `json:"comet_sections"`
		EncryptedTracking string `json:"encrypted_tracking"`
	} `json:"story"`
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

type PrivacyScope struct {
	Description  string      `json:"description"`
	Label        string      `json:"label"`
	DisplayLabel interface{} `json:"display_label"`
	IconImage    struct {
		Height int    `json:"height"`
		Scale  int    `json:"scale"`
		Uri    string `json:"uri"`
		Width  int    `json:"width"`
	} `json:"icon_image"`
	PrivacyScopeRenderer PrivacyScopeRenderer `json:"privacy_scope_renderer"`
}

type Feedback struct {
	AssociatedGroup interface{} `json:"associated_group"`
	ID              string      `json:"id"`
}

type Hashtag struct {
	Offset int `json:"offset"`
	Length int `json:"length"`
	Entity struct {
		Typename  string `json:"__typename"`
		IsEntity  string `json:"__isEntity"`
		MobileUrl string `json:"mobileUrl"`
		Url       string `json:"url"`
		ID        string `json:"id"`
		IsNode    string `json:"__isNode"`
	} `json:"entity"`
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
type Message struct {
	Text              string        `json:"text"`
	Ranges            []Hashtag     `json:"ranges"`
	DelightRanges     []interface{} `json:"delight_ranges"`
	ImageRanges       []interface{} `json:"image_ranges"`
	InlineStyleRanges []interface{} `json:"inline_style_ranges"`
	AggregatedRanges  []interface{} `json:"aggregated_ranges"`
	ColorRanges       []interface{} `json:"color_ranges"`
}

type ReelMediaStruct struct {
	ShortFormVideoContext  ShortFormVideoContext `json:"short_form_video_context"`
	PrivacyScope           PrivacyScope          `json:"privacy_scope"`
	PostID                 string                `json:"post_id"`
	Tracking               string                `json:"tracking"`
	CreationTime           int64                 `json:"creation_time"`
	Feedback               Feedback              `json:"feedback"`
	BrandedContentPostInfo interface{}           `json:"branded_content_post_info"`
	Message                Message               `json:"message"`
}

type ProfilePicture struct {
	URI    string `json:"uri"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Scale  int    `json:"scale"`
}

type DelegatePage struct {
	IsBusinessPageActive bool   `json:"is_business_page_active"`
	ID                   string `json:"id"`
}

type Actor struct {
	Typename                string         `json:"__typename"`
	IsActor                 string         `json:"__isActor"`
	ID                      string         `json:"id"`
	IsEntity                string         `json:"__isEntity"`
	URL                     string         `json:"url"`
	ProfileURL              string         `json:"profile_url"`
	Name                    string         `json:"name"`
	ProfilePicture          ProfilePicture `json:"profile_picture"`
	IsAdditionalProfilePlus bool           `json:"is_additional_profile_plus"`
	DelegatePage            DelegatePage   `json:"delegate_page"`
}

type Story struct {
	Actors        []Actor            `json:"actors"`
	CometSections StoryCometSections `json:"comet_sections"`
	ID            string             `json:"id"`
}

type ActorPhoto struct {
	Typename              string `json:"__typename"`
	IsICometStorySection  string `json:"__isICometStorySection"`
	IsProdEligible        bool   `json:"is_prod_eligible"`
	Story                 Story  `json:"story"`
	HasCommerceAttachment bool   `json:"has_commerce_attachment"`
}

type MetadataStory struct {
	CreationTime int         `json:"creation_time"`
	URL          string      `json:"url"`
	GhlLabel     interface{} `json:"ghl_label"`
	ID           string      `json:"id"`
}

type Metadata struct {
	Typename             string        `json:"__typename"`
	IsICometStorySection string        `json:"__isICometStorySection"`
	IsProdEligible       bool          `json:"is_prod_eligible"`
	Story                MetadataStory `json:"story"`
}

type CometSections struct {
	Message struct {
		CometFeedStoryDefaultMessageRenderingStrategy
	} `json:"message"`
	ActorPhoto ActorPhoto                           `json:"actor_photo"`
	Metadata   []Metadata                           `json:"metadata"`
	Title      CometFeedStoryTitleWithActorStrategy `json:"title"`
}

type StoryCometSections struct {
	ActionLink interface{} `json:"action_link"`
}
type MessagePreferredBody struct {
	Typename string `json:"__typename"`
	Text     string `json:"text"`
}

type PrivacyScopeRenderer struct {
	Typename                  string             `json:"__typename"`
	IsPrivacySelectorRenderer string             `json:"__isPrivacySelectorRenderer"`
	PrivacyRowInput           PrivacyRowInput    `json:"privacy_row_input"`
	EntryPointRenderer        EntryPointRenderer `json:"entry_point_renderer"`
	ID                        string             `json:"id"`
}

type PrivacyRowInput struct {
	Allow             []interface{} `json:"allow"`
	BaseState         string        `json:"base_state"`
	Deny              []interface{} `json:"deny"`
	PrivacyTargeting  interface{}   `json:"privacy_targeting"`
	TagExpansionState string        `json:"tag_expansion_state"`
}

type EntryPointRenderer struct {
	Typename string `json:"__typename"`
	Source   Source `json:"source"`
}

type Source struct {
	Typename string `json:"__typename"`
	Scope    Scope  `json:"scope"`
	ID       string `json:"id"`
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
	IsMedia              string               `json:"__isMedia"`
	AttachedComment      interface{}          `json:"attached_comment"`
	Message              Message              `json:"message"`
	MessagePreferredBody MessagePreferredBody `json:"message_preferred_body"`
	ContainerStory       ContainerStory       `json:"container_story"`
	CreationStory        CreationStory        `json:"creation_story"`
	EncryptedTracking    string               `json:"encrypted_tracking"`
	Actors               []Actor              `json:"actors"`
	IsPlayable           bool                 `json:"is_playable"`
	ID                   string               `json:"id"`
	CreatedTime          int                  `json:"created_time"`
	PrivacyScope         PrivacyScope         `json:"privacy_scope"`
	Owner                Owner                `json:"owner"`
}

type CurrMedia struct {
	Typename             string           `json:"__typename"`
	IsMedia              string           `json:"__isMedia"`
	CanViewerEdit        bool             `json:"can_viewer_edit"`
	ContainerStory       ContainerStory   `json:"container_story"`
	ID                   string           `json:"id"`
	PostID               string           `json:"post_id"`
	EncryptedTracking    string           `json:"encrypted_tracking"`
	ViewabilityConfig    []int            `json:"viewability_config"`
	ClientViewConfig     ClientViewConfig `json:"client_view_config"`
	Feedback             PhotoFeedBack    `json:"feedback"`
	IsNode               string           `json:"__isNode"`
	IsPlayable           bool             `json:"is_playable"`
	CanViewerAddTags     bool             `json:"can_viewer_add_tags"`
	Owner                Owner            `json:"owner"`
	CreatedTime          int64            `json:"created_time"`
	Image                Image            `json:"image"`
	AccessibilityCaption string           `json:"accessibility_caption"`
	Tags                 Tags             `json:"tags"`
	CreationStory        CreationStory    `json:"creation_story"`
	DefaultMediaSet      DefaultMediaSet  `json:"default_mediaset"`
}

type ContainerStory struct {
	CanViewerBoost          bool                `json:"can_viewer_boost"`
	BusinessContentType     string              `json:"business_content_type"`
	LegacyStoryHideableID   string              `json:"legacy_story_hideable_id"`
	DelegatePageID          string              `json:"delegate_page_id"`
	PostHasInvalidTags      bool                `json:"post_has_invalid_tags"`
	BoostPostButtonData     BoostPostButtonData `json:"boost_post_button_data"`
	CanShowInlineBoost      bool                `json:"can_show_inline_boost_for_bizweb_feed_view"`
	IsEligibleForLwiFeature bool                `json:"is_eligible_for_lwi_feature"`
	ProfilePlusID           string              `json:"profile_plus_id"`
	ID                      string              `json:"id"`
}

type CodedException struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
	Code        int    `json:"code"`
}

type AdsLwiPost struct {
	Story Story `json:"story"`
}

type BoostPostButtonData struct {
	ButtonLabel     string `json:"button_label"`
	ButtonActionURI string `json:"button_action_uri"`
	TooltipLabel    string `json:"tooltip_label"`
	ButtonLabelType string `json:"button_label_type"`
}

type ClientViewConfig struct {
	CanDelayLogImpression bool `json:"can_delay_log_impression"`
	UseBanzaiSignalImp    bool `json:"use_banzai_signal_imp"`
	UseBanzaiVitalImp     bool `json:"use_banzai_vital_imp"`
}

type PhotoFeedBack struct {
	ID                         string          `json:"id"`
	CanViewerReact             bool            `json:"can_viewer_react"`
	ViewerActor                interface{}     `json:"viewer_actor"`
	ViewerFeedbackReactionInfo interface{}     `json:"viewer_feedback_reaction_info"`
	AssociatedVideo            interface{}     `json:"associated_video"`
	TopReactions               TopReactions    `json:"top_reactions"`
	SupportedReactionInfos     []ReactionInfo  `json:"supported_reaction_infos"`
	UnifiedReactors            UnifiedReactors `json:"unified_reactors"`
	Reactors                   Reactors        `json:"reactors"`
}

type TopReactions struct {
	Edges []ReactionEdge `json:"edges"`
}

type ReactionEdge struct {
	ReactionCount int          `json:"reaction_count"`
	Node          ReactionNode `json:"node"`
}

type ReactionNode struct {
	ID string `json:"id"`
}

type ReactionInfo struct {
	ID string `json:"id"`
}

type UnifiedReactors struct {
	Count int `json:"count"`
}

type Reactors struct {
	Count   int  `json:"count"`
	IsEmpty bool `json:"is_empty"`
}

type Owner struct {
	Typename                             string         `json:"__typename"`
	ID                                   string         `json:"id"`
	UserID                               string         `json:"user_id"`
	AdditionalProfileHasTaggableProducts bool           `json:"additional_profile_has_taggable_products"`
	Name                                 string         `json:"name"`
	ProfilePicture                       ProfilePicture `json:"profile_picture"`
	IsProfile                            string         `json:"__isProfile"`
}

type Image struct {
	URI    string `json:"uri"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Tags struct {
	Edges []interface{} `json:"edges"`
}

type CreationStory struct {
	URL           string        `json:"url"`
	Attachments   []Attachment  `json:"attachments"`
	Tracking      string        `json:"tracking"`
	CometSections CometSections `json:"comet_sections"`
}

type Attachment struct {
	ActionLinks []interface{} `json:"action_links"`
}

type DefaultMediaSet struct {
	Typename   string     `json:"__typename"`
	IsMediaSet string     `json:"__isMediaSet"`
	PrevMedia  MediaEdges `json:"prevMedia"`
	NextMedia  MediaEdges `json:"nextMedia"`
	ID         string     `json:"id"`
}

type MediaEdges struct {
	Edges []MediaEdge `json:"edges"`
}

type MediaEdge struct {
	Node MediaNode `json:"node"`
}

type MediaNode struct {
	Typename   string `json:"__typename"`
	ID         string `json:"id"`
	IsMedia    string `json:"__isMedia"`
	IsPlayable bool   `json:"is_playable"`
	Owner      Owner  `json:"owner"`
	IsNode     string `json:"__isNode"`
}

type Extensions struct {
	PrefetchUrisV2 []PrefetchURI `json:"prefetch_uris_v2"`
	IsFinal        bool          `json:"is_final"`
}

type PrefetchURI struct {
	URI   string      `json:"uri"`
	Label interface{} `json:"label"`
}

type PhotoStruct struct {
	Data       Data       `json:"data"`
	Extensions Extensions `json:"extensions"`
}

type Result struct {
	Label          string     `json:"label"`
	Path           []string   `json:"path"`
	Data           Data       `json:"data"`
	Extensions     Extensions `json:"extensions"`
	SequenceNumber int        `json:"sequence_number"`
}

type Bbox struct {
	Complete bool   `json:"complete"`
	Result   Result `json:"result"`
}

type PhotoMedia struct {
	Bbox Bbox `json:"__bbox"`
}

type LatitudeLongitude struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type MarketLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type MarketplaceListingRenderableTarget struct {
	SeoVirtualCategory                      interface{}    `json:"seo_virtual_category"`
	Location                                MarketLocation `json:"location"`
	IsShippingOffered                       bool           `json:"is_shipping_offered"`
	IsMarketplaceListingWithSweepstake      string         `json:"__isMarketplaceListingWithSweepstake"`
	SweepstakeEnabled                       bool           `json:"sweepstake_enabled"`
	Typename                                string         `json:"__typename"`
	ID                                      string         `json:"id"`
	IsMarketplaceListingRenderable          string         `json:"__isMarketplaceListingRenderable"`
	IsMarketplaceListingWithPersonalization string         `json:"__isMarketplaceListingWithPersonalization"`
	PersonalizationInfo                     interface{}    `json:"personalization_info"`
}

type TranslatabilityForViewer struct {
	SourceDialectName string `json:"source_dialect_name"`
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

type SaveInfo struct {
	ViewerSaveState string `json:"viewer_save_state"`
}

type MarketStory struct {
	TranslationAvailableForViewer bool                     `json:"translation_available_for_viewer"`
	TranslatabilityForViewer      TranslatabilityForViewer `json:"translatability_for_viewer"`
	TranslatedMessageForViewer    interface{}              `json:"translated_message_for_viewer"`
	ID                            string                   `json:"id"`
	Actors                        []MarketActor            `json:"actors"`
	IsEntity                      string                   `json:"__isEntity"`
	URL                           string                   `json:"url"`
	IsNode                        string                   `json:"__isNode"`
	Tracking                      interface{}              `json:"tracking"`
	PostID                        string                   `json:"post_id"`
	SaveInfo                      SaveInfo                 `json:"save_info"`
}

type RedactedDescription struct {
	Text string `json:"text"`
}

type LocationText struct {
	Text string `json:"text"`
}

type FormattedPrice struct {
	Text string `json:"text"`
}

type ListingPrice struct {
	AmountWithOffset string `json:"amount_with_offset"`
	Currency         string `json:"currency"`
	Amount           string `json:"amount"`
}

type MarketplaceLeadGenForm struct {
	ViewerPurchaseLimit       int         `json:"viewer_purchase_limit"`
	BoostedMarketplaceListing interface{} `json:"boosted_marketplace_listing"`
	PromotedListing           interface{} `json:"promoted_listing"`
}

type MarketplaceUserProfile struct {
	C2COrdersShipped interface{} `json:"c2c_orders_shipped"`
	ID               string      `json:"id"`
}

type MarketplaceRatingsStatsByRole struct {
	SellerStats struct {
		FiveStarRatingsAverage         int `json:"five_star_ratings_average"`
		FiveStarTotalRatingCountByRole int `json:"five_star_total_rating_count_by_role"`
	} `json:"seller_stats"`
	SellerRatingsArePrivate bool `json:"seller_ratings_are_private"`
}

type MarketplaceListingSeller struct {
	Typename                               string                        `json:"__typename"`
	ID                                     string                        `json:"id"`
	UserID                                 string                        `json:"user_id"`
	IsProfile                              string                        `json:"__isProfile"`
	MarketplaceUserProfile                 MarketplaceUserProfile        `json:"marketplace_user_profile"`
	IsActor                                string                        `json:"__isActor"`
	Name                                   string                        `json:"name"`
	ProfilePicture                         ProfilePicture                `json:"profile_picture"`
	ProfilePicture160                      ProfilePicture                `json:"profile_picture_160"`
	ProfilePicture112                      ProfilePicture                `json:"profile_picture_112"`
	ProfilePicture64                       ProfilePicture                `json:"profile_picture_64"`
	ProfilePicture60                       ProfilePicture                `json:"profile_picture_60"`
	ProfilePicture50                       ProfilePicture                `json:"profile_picture_50"`
	IsActorWithCustomizableCommerceProfile string                        `json:"__isActorWithCustomizableCommerceProfile"`
	CommerceProfilePictureWithFallback160  ProfilePicture                `json:"commerce_profile_picture_with_fallback_160"`
	CommerceProfilePictureWithFallback112  ProfilePicture                `json:"commerce_profile_picture_with_fallback_112"`
	CommerceProfilePictureWithFallback64   ProfilePicture                `json:"commerce_profile_picture_with_fallback_64"`
	CommerceProfilePictureWithFallback60   ProfilePicture                `json:"commerce_profile_picture_with_fallback_60"`
	CommerceProfilePictureWithFallback50   ProfilePicture                `json:"commerce_profile_picture_with_fallback_50"`
	JoinTime                               int                           `json:"join_time"`
	MarketplaceRatingsStatsByRole          MarketplaceRatingsStatsByRole `json:"marketplace_ratings_stats_by_role"`
	MarketplaceShouldDisplayVerifiedBadge  bool                          `json:"marketplace_should_display_verified_badge"`
}

type ProductItemAttributeData struct {
	Typename      string `json:"__typename"`
	AttributeName string `json:"attribute_name"`
	Value         string `json:"value"`
	Label         string `json:"label"`
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

type OriginTarget struct {
	Typename string `json:"__typename"`
	ID       string `json:"id"`
}

type CommerceBadgesInfo struct {
	SourceSummary interface{}   `json:"source_summary"`
	Badges        []interface{} `json:"badges"`
}

type VehicleOdometerData struct {
	Unit  interface{} `json:"unit"`
	Value interface{} `json:"value"`
}

type PrimaryMpEnt struct {
	Typename string `json:"__typename"`
	ID       string `json:"id"`
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
	Typename                           string                             `json:"__typename"`
	ProductDetailsType                 string                             `json:"product_details_type"`
	IsMarketplaceProductDetailsPage    string                             `json:"__isMarketplaceProductDetailsPage"`
	MarketplaceListingRenderableTarget MarketplaceListingRenderableTarget `json:"marketplace_listing_renderable_target"`
	Target                             Target                             `json:"target"`
	Viewer                             struct {
		Typename                            string `json:"__typename"`
		MarketplaceActorWithIntegrityStatus struct {
			MarketplaceUserInDma bool `json:"marketplace_user_in_dma"`
		} `json:"marketplace_actor_with_integrity_status"`
		MarketplaceSettings struct {
			LoanPaymentOptions interface{} `json:"loan_payment_options"`
		} `json:"marketplace_settings"`
	} `json:"viewer"`
	ModuleOperationMarketplacePDPPage struct {
		Dr string `json:"__dr"`
	} `json:"__module_operation_MarketplacePDP_page"`
	ModuleComponentMarketplacePDPPage struct {
		Dr string `json:"__dr"`
	} `json:"__module_component_MarketplacePDP_page"`
	ID string `json:"id"`
}

type Target struct {
	Typename                                string              `json:"__typename"`
	MarketplaceListingTitle                 string              `json:"marketplace_listing_title"`
	ID                                      string              `json:"id"`
	IsMarketplaceListingRenderable          string              `json:"__isMarketplaceListingRenderable"`
	Story                                   MarketStory         `json:"story"`
	RedactedDescription                     RedactedDescription `json:"redacted_description"`
	CreationTime                            int                 `json:"creation_time"`
	LocationText                            LocationText        `json:"location_text"`
	LocationVanityOrID                      string              `json:"location_vanity_or_id"`
	IsViewerSeller                          bool                `json:"is_viewer_seller"`
	ListingInventoryType                    string              `json:"listing_inventory_type"`
	FormattedPrice                          FormattedPrice      `json:"formatted_price"`
	ListingPrice                            ListingPrice        `json:"listing_price"`
	IsMarketplaceVehicleListing             string              `json:"__isMarketplaceVehicleListing"`
	PaymentTimePeriod                       interface{}         `json:"payment_time_period"`
	Condition                               string              `json:"condition"`
	CustomTitle                             interface{}         `json:"custom_title"`
	IsLive                                  bool                `json:"is_live"`
	IsPending                               bool                `json:"is_pending"`
	IsSold                                  bool                `json:"is_sold"`
	LoggingID                               string              `json:"logging_id"`
	IsMarketplaceListingWithLeadGen         string              `json:"__isMarketplaceListingWithLeadGen"`
	IsMarketplaceRealEstateListing          string              `json:"__isMarketplaceRealEstateListing"`
	IsMarketplaceMessageable                string              `json:"__isMarketplaceMessageable"`
	MessagingEnabled                        bool                `json:"messaging_enabled"`
	IsMarketplaceListingWithIntegrityStatus string              `json:"__isMarketplaceListingWithIntegrityStatus"`
	ListingIsRejected                       bool                `json:"listing_is_rejected"`
	ProductItem                             struct {
		ID                        string      `json:"id"`
		ViewerPurchaseLimit       int         `json:"viewer_purchase_limit"`
		BoostedMarketplaceListing interface{} `json:"boosted_marketplace_listing"`
		PromotedListing           interface{} `json:"promoted_listing"`
	} `json:"product_item"`
	SellerMessageThread                     interface{}              `json:"seller_message_thread"`
	IsCheckoutEnabled                       bool                     `json:"is_checkout_enabled"`
	IsDraft                                 bool                     `json:"is_draft"`
	CanSellerEdit                           bool                     `json:"can_seller_edit"`
	OriginGroup                             interface{}              `json:"origin_group"`
	IsMarketplaceListingWithVariants        string                   `json:"__isMarketplaceListingWithVariants"`
	DefaultVariantListing                   interface{}              `json:"default_variant_listing"`
	PrimaryMpEnt                            PrimaryMpEnt             `json:"primary_mp_ent"`
	IsHidden                                bool                     `json:"is_hidden"`
	MarketplaceListingSeller                MarketplaceListingSeller `json:"marketplace_listing_seller"`
	ReportableEntID                         string                   `json:"reportable_ent_id"`
	IsMarketplaceListingHideableFromFriends string                   `json:"__isMarketplaceListingHideableFromFriends"`
	HiddenFromFriends                       string                   `json:"hidden_from_friends"`
	CanShare                                bool                     `json:"can_share"`
	CanSellerChangeAvailability             bool                     `json:"can_seller_change_availability"`
	IsOnMarketplace                         bool                     `json:"is_on_marketplace"`
	IsMarketplaceListingWithChildListings   string                   `json:"__isMarketplaceListingWithChildListings"`
	HasChildren                             bool                     `json:"has_children"`
	CrossPostInfo                           struct {
		AllListings []struct {
			Typename        string `json:"__typename"`
			IsOnMarketplace bool   `json:"is_on_marketplace"`
			ID              string `json:"id"`
		} `json:"all_listings"`
	} `json:"cross_post_info"`
	IsMarketplaceListingWithDeliveryOptions    string                     `json:"__isMarketplaceListingWithDeliveryOptions"`
	PrimaryListingPhoto                        ProductImage               `json:"primary_listing_photo"`
	IsMarketplaceListingWithEmailCommunication string                     `json:"__isMarketplaceListingWithEmailCommunication"`
	IsEmailCommunicationEnabled                bool                       `json:"is_email_communication_enabled"`
	VehicleOdometerData                        VehicleOdometerData        `json:"vehicle_odometer_data"`
	AttributeData                              []ProductItemAttributeData `json:"attribute_data"`
	Location                                   Location                   `json:"location"`
	OriginTarget                               OriginTarget               `json:"origin_target"`
	ShippingOffered                            bool                       `json:"shipping_offered"`
	LegalDisclosureImpressumURL                interface{}                `json:"legal_disclosure_impressum_url"`
	IsMarketplaceListingWithBadges             string                     `json:"__isMarketplaceListingWithBadges"`
	CommerceBadgesInfo                         CommerceBadgesInfo         `json:"commerce_badges_info"`
	Seller                                     struct {
		Typename               string `json:"__typename"`
		ID                     string `json:"id"`
		MarketplaceUserProfile struct {
			BusinessInformation interface{} `json:"business_information"`
			ID                  string      `json:"id"`
		} `json:"marketplace_user_profile"`
		IsActor string `json:"__isActor"`
	} `json:"seller"`
	EnergyEfficiencyClassEu interface{}    `json:"energy_efficiency_class_eu"`
	ListingPhotos           []ProductImage `json:"listing_photos"`
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
	Complete   bool         `json:"complete"`
	Result     MarketResult `json:"result"`
	Extensions Extensions   `json:"extensions"`
}

type MarketSchema struct {
	BboxInner     BboxInner      `json:"__bbox"`
	LisitngPhotos []ProductImage `json:"listing_photos"`
}
