package handlers

import "encoding/json"

type CometFeedStoryDefaultMessageRenderingStrategy struct {
	TypeName                                              string     `json:"__typename"`
	Story                                                 MediaStory `json:"story"`
	ModuleOperationCometTahoeVideoContextSectionBodyVideo struct {
		DR string `json:"__dr"`
	} `json:"__module_operation_CometTahoeVideoContextSectionBody_video"`
	ModuleComponentCometTahoeVideoContextSectionBodyVideo struct {
		DR string `json:"__dr"`
	} `json:"__module_component_CometTahoeVideoContextSectionBody_video"`
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
		Text              string        `json:"text"`
		DelightRanges     []interface{} `json:"delight_ranges"`
		ImageRanges       []interface{} `json:"image_ranges"`
		InlineStyleRanges []interface{} `json:"inline_style_ranges"`
		AggregatedRanges  []interface{} `json:"aggregated_ranges"`
		ColorRanges       []interface{} `json:"color_ranges"`
	} `json:"message"`
	ID string `json:"id"`
}

type CometFeedStoryActorPhotoStrategy struct {
	TypeName                                            string `json:"__typename"`
	IsICometStorySection                                string `json:"__isICometStorySection"`
	IsProdEligible                                      bool   `json:"is_prod_eligible"`
	Story                                               Story  `json:"story"`
	HasCommerceAttachment                               bool   `json:"has_commerce_attachment"`
	ModuleOperationCometFeedStoryActorPhotoSectionStory struct {
		DR string `json:"__dr"`
	} `json:"__module_operation_CometFeedStoryActorPhotoSection_story"`
	ModuleComponentCometFeedStoryActorPhotoSectionStory struct {
		DR string `json:"__dr"`
	} `json:"__module_component_CometFeedStoryActorPhotoSection_story"`
}

type CometFeedStoryMinimizedTimestampStrategy struct {
	TypeName             string      `json:"__typename"`
	IsICometStorySection string      `json:"__isICometStorySection"`
	IsProdEligible       bool        `json:"is_prod_eligible"`
	OverrideURL          interface{} `json:"override_url"`
	VideoOverrideURL     interface{} `json:"video_override_url"`
	Story                struct {
		CreationTime int         `json:"creation_time"`
		URL          string      `json:"url"`
		GhlLabel     interface{} `json:"ghl_label"`
		ID           string      `json:"id"`
	} `json:"story"`
	ModuleOperationCometFeedStoryMetadataSectionStory struct {
		DR string `json:"__dr"`
	} `json:"__module_operation_CometFeedStoryMetadataSection_story"`
	ModuleComponentCometFeedStoryMetadataSectionStory struct {
		DR string `json:"__dr"`
	} `json:"__module_component_CometFeedStoryMetadataSection_story"`
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
	ModuleOperationCometFeedStoryTitleSectionStory struct {
		DR string `json:"__dr"`
	} `json:"__module_operation_CometFeedStoryTitleSection_story"`
	ModuleComponentCometFeedStoryTitleSectionStory struct {
		DR string `json:"__dr"`
	} `json:"__module_component_CometFeedStoryTitleSection_story"`
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
	VideoOwner struct {
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
	} `json:"video_owner"`
	IsPassiveContent       bool `json:"is_passive_content"`
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

type Rc struct {
	Dr string `json:"__dr"`
}

type StoryBucketNode struct {
	ShowCloseFriendBadge bool        `json:"should_show_close_friend_badge"`
	ID                   string      `json:"id"`
	FirstStoryToShow     interface{} `json:"first_story_to_show"`
}

type StoryBucket struct {
	Nodes []StoryBucketNode `json:"nodes"`
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
	Typename                  string         `json:"__typename"`
	IsActor                   string         `json:"__isActor"`
	ID                        string         `json:"id"`
	IsEntity                  string         `json:"__isEntity"`
	URL                       string         `json:"url"`
	WorkForeignEntityInfo     interface{}    `json:"work_foreign_entity_info"`
	WorkInfo                  interface{}    `json:"work_info"`
	StoryBucket               StoryBucket    `json:"story_bucket"`
	LiveVideoForCometLiveRing interface{}    `json:"live_video_for_comet_live_ring"`
	AnswerAgentGroupID        interface{}    `json:"answer_agent_group_id"`
	ProfileURL                string         `json:"profile_url"`
	Name                      string         `json:"name"`
	ProfilePicture            ProfilePicture `json:"profile_picture"`
	IsAdditionalProfilePlus   bool           `json:"is_additional_profile_plus"`
	DelegatePage              DelegatePage   `json:"delegate_page"`
}

type Story struct {
	Actors        []Actor            `json:"actors"`
	CometSections StoryCometSections `json:"comet_sections"`
	SponsoredData interface{}        `json:"sponsored_data"`
	ID            string             `json:"id"`
}

type ActorPhoto struct {
	Typename              string `json:"__typename"`
	IsICometStorySection  string `json:"__isICometStorySection"`
	IsProdEligible        bool   `json:"is_prod_eligible"`
	Story                 Story  `json:"story"`
	HasCommerceAttachment bool   `json:"has_commerce_attachment"`
	ModuleOperation       Rc     `json:"__module_operation_CometFeedStoryActorPhotoSectionMatchRenderer_story"`
	ModuleComponent       Rc     `json:"__module_component_CometFeedStoryActorPhotoSectionMatchRenderer_story"`
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
	OverrideURL          interface{}   `json:"override_url"`
	VideoOverrideURL     interface{}   `json:"video_override_url"`
	Story                MetadataStory `json:"story"`
	ModuleOperation      Rc            `json:"__module_operation_CometFeedStoryMetadataSectionMatchRenderer_story"`
	ModuleComponent      Rc            `json:"__module_component_CometFeedStoryMetadataSectionMatchRenderer_story"`
}

type CometSections struct {
	Message struct {
		CometFeedStoryDefaultMessageRenderingStrategy
	} `json:"message"`
	ActorPhoto ActorPhoto                           `json:"actor_photo"`
	Metadata   []Metadata                           `json:"metadata"`
	Footer     interface{}                          `json:"footer"`
	Title      CometFeedStoryTitleWithActorStrategy `json:"title"`
}

type StoryCometSections struct {
	ActionLink interface{} `json:"action_link"`
}
type MessagePreferredBody struct {
	Typename          string        `json:"__typename"`
	DelightRanges     []interface{} `json:"delight_ranges"`
	ImageRanges       []interface{} `json:"image_ranges"`
	InlineStyleRanges []interface{} `json:"inline_style_ranges"`
	AggregatedRanges  []interface{} `json:"aggregated_ranges"`
	Ranges            []interface{} `json:"ranges"`
	ColorRanges       []interface{} `json:"color_ranges"`
	Text              string        `json:"text"`
	ModuleOperation   Rc            `json:"__module_operation_CometMediaViewerDescriptionSection_media"`
	ModuleComponent   Rc            `json:"__module_component_CometMediaViewerDescriptionSection_media"`
}

type PrivacyScopeRenderer struct {
	Typename                  string             `json:"__typename"`
	IsPrivacySelectorRenderer string             `json:"__isPrivacySelectorRenderer"`
	PrivacyRowInput           PrivacyRowInput    `json:"privacy_row_input"`
	EntryPointRenderer        EntryPointRenderer `json:"entry_point_renderer"`
	ModuleOperation           Rc                 `json:"__module_operation_CometPrivacySelectorForScope_scope"`
	ModuleComponent           Rc                 `json:"__module_component_CometPrivacySelectorForScope_scope"`
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
	Typename        string `json:"__typename"`
	Source          Source `json:"source"`
	ModuleOperation Rc     `json:"__module_operation_CometPrivacySelectorEntryPointMatchContainer_renderer"`
	ModuleComponent Rc     `json:"__module_component_CometPrivacySelectorEntryPointMatchContainer_renderer"`
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
	TargetGroup          interface{}          `json:"target_group"`
	SponsoredData        interface{}          `json:"sponsored_data"`
	IsPlayable           bool                 `json:"is_playable"`
	ID                   string               `json:"id"`
	CreatedTime          int                  `json:"created_time"`
	PrivacyScope         PrivacyScope         `json:"privacy_scope"`
	Owner                Owner                `json:"owner"`
	PhotoProductTags     []interface{}        `json:"photo_product_tags"`
	ProductMatches       []interface{}        `json:"product_matches"`
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
	IfMediaCanShowLink   interface{}      `json:"if_media_can_show_link_to_container_story"`
	AttachedComment      interface{}      `json:"attached_comment"`
	Feedback             PhotoFeedBack    `json:"feedback"`
	CopyrightBannerInfo  interface{}      `json:"copyright_banner_info"`
	IsNode               string           `json:"__isNode"`
	IsPlayable           bool             `json:"is_playable"`
	CixScreen            interface{}      `json:"cix_screen"`
	CanViewerAddTags     bool             `json:"can_viewer_add_tags"`
	Owner                Owner            `json:"owner"`
	CreatedTime          int64            `json:"created_time"`
	Image                Image            `json:"image"`
	AccessibilityCaption string           `json:"accessibility_caption"`
	PhotoProductTags     []interface{}    `json:"photo_product_tags"`
	CometPhotoRenderer   interface{}      `json:"comet_photo_renderer"`
	Tags                 Tags             `json:"tags"`
	CreationStory        CreationStory    `json:"creation_story"`
	DefaultMediaSet      DefaultMediaSet  `json:"default_mediaset"`
}

type ContainerStory struct {
	CanViewerBoost           bool                     `json:"can_viewer_boost"`
	BusinessContentType      string                   `json:"business_content_type"`
	LegacyStoryHideableID    string                   `json:"legacy_story_hideable_id"`
	DelegatePageID           string                   `json:"delegate_page_id"`
	PostPromotionInfo        interface{}              `json:"post_promotion_info"`
	PostAudienceBuildingInfo PostAudienceBuildingInfo `json:"post_audience_building_info"`
	PostHasInvalidTags       bool                     `json:"post_has_invalid_tags"`
	BoostPostButtonData      BoostPostButtonData      `json:"boost_post_button_data"`
	CanShowInlineBoost       bool                     `json:"can_show_inline_boost_for_bizweb_feed_view"`
	LiveBoostingTooltip      interface{}              `json:"live_boosting_tooltip"`
	IsEligibleForLwiFeature  bool                     `json:"is_eligible_for_lwi_feature"`
	ProfilePlusID            string                   `json:"profile_plus_id"`
	ID                       string                   `json:"id"`
	Suffix                   interface{}              `json:"suffix"`
}

type PostAudienceBuildingInfo struct {
	BoostPostUnavailableInfo PhotoBoostPostUnavailableInfo `json:"boost_post_unavailable_info"`
}

type PhotoBoostPostUnavailableInfo struct {
	Typename                               string                                   `json:"__typename"`
	EligibilityRule                        string                                   `json:"eligibility_rule"`
	CodedException                         CodedException                           `json:"coded_exception"`
	TemplatePostSpecOverride               interface{}                              `json:"template_post_spec_override"`
	IsBoostedComponentBoostUnavailableInfo string                                   `json:"__isBoostedComponentBoostUnavailableInfo"`
	ConfirmLabel                           string                                   `json:"confirm_label"`
	ConfirmLink                            string                                   `json:"confirm_link"`
	AdsLwiPost                             AdsLwiPost                               `json:"ads_lwi_post"`
	BoostUnavailableActionRendererComet    PhotoBoostUnavailableActionRendererComet `json:"boost_unavailable_action_renderer_comet"`
}

type CodedException struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
	Code        int    `json:"code"`
}

type AdsLwiPost struct {
	Story Story `json:"story"`
}

type PhotoBoostUnavailableActionRendererComet struct {
	Typename                                                                           string          `json:"__typename"`
	BoostUnavailableInfo                                                               json.RawMessage `json:"boost_unavailable_info"`
	ModuleOperationCometFeedStoryProfilePlusBoostUnavailableButtonAudienceBuildingInfo ModuleOperation `json:"__module_operation_CometFeedStoryProfilePlusBoostUnavailableButton_audienceBuildingInfo"`
	ModuleComponentCometFeedStoryProfilePlusBoostUnavailableButtonAudienceBuildingInfo ModuleComponent `json:"__module_component_CometFeedStoryProfilePlusBoostUnavailableButton_audienceBuildingInfo"`
}

type ModuleOperation struct {
	Dr string `json:"__dr"`
}

type ModuleComponent struct {
	Dr              string `json:"__dr"`
	ModuleOperation Rc     `json:"__module_operation_CometMediaViewerDescriptionSection_media"`
	Component       Rc     `json:"__module_component_CometMediaViewerDescriptionSection_media"`
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
	CanViewerBoost                    bool                     `json:"can_viewer_boost"`
	CanViewerEdit                     bool                     `json:"can_viewer_edit"`
	ID                                string                   `json:"id"`
	LegacyStoryHideableID             string                   `json:"legacy_story_hideable_id"`
	PostPromotionInfo                 interface{}              `json:"post_promotion_info"`
	WhatsappAdContext                 interface{}              `json:"whatsapp_ad_context"`
	BusinessContentType               string                   `json:"business_content_type"`
	DelegatePageID                    string                   `json:"delegate_page_id"`
	PostAudienceBuildingInfo          PostAudienceBuildingInfo `json:"post_audience_building_info"`
	URL                               string                   `json:"url"`
	Shareable                         interface{}              `json:"shareable"`
	SponsoredData                     interface{}              `json:"sponsored_data"`
	InformTreatmentForMessaging       interface{}              `json:"inform_treatment_for_messaging"`
	Attachments                       []Attachment             `json:"attachments"`
	Tracking                          string                   `json:"tracking"`
	TargetGroup                       interface{}              `json:"target_group"`
	ShareableFromPerspectiveOfFeedUfi interface{}              `json:"shareable_from_perspective_of_feed_ufi"`
	VoteAttachments                   []interface{}            `json:"vote_attachments"`
	Suffix                            interface{}              `json:"suffix"`
	SerializedFrtpIdentifiers         interface{}              `json:"serialized_frtp_identifiers"`
	DebugInfo                         interface{}              `json:"debug_info"`
	CometSections                     CometSections            `json:"comet_sections"`
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
	Target                             struct {
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
		MarketplaceLeadGenForm                  interface{}         `json:"marketplace_lead_gen_form"`
		IsMarketplaceRealEstateListing          string              `json:"__isMarketplaceRealEstateListing"`
		ListedBy                                interface{}         `json:"listed_by"`
		RealEstateListingAgent                  interface{}         `json:"real_estate_listing_agent"`
		InventoryCount                          interface{}         `json:"inventory_count"`
		LegalReportingCtaType                   interface{}         `json:"legal_reporting_cta_type"`
		LegalReportingUri                       interface{}         `json:"legal_reporting_uri"`
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
		IsMarketplaceListingWithDeliveryOptions string        `json:"__isMarketplaceListingWithDeliveryOptions"`
		ActiveOrder                             interface{}   `json:"active_order"`
		RebuyOrderReceipt                       interface{}   `json:"rebuy_order_receipt"`
		C2CShippingEligible                     bool          `json:"c2c_shipping_eligible"`
		IsShippingOffered                       bool          `json:"is_shipping_offered"`
		MostRecentActiveOrderAsBuyer            interface{}   `json:"most_recent_active_order_as_buyer"`
		OrderSummaries                          []interface{} `json:"order_summaries"`
		DeliveryTypes                           []string      `json:"delivery_types"`
		ShouldHidePdpShippingContent            bool          `json:"should_hide_pdp_shipping_content"`
		ShippingProfile                         interface{}   `json:"shipping_profile"`
		PrimaryListingPhoto                     ProductImage  `json:"primary_listing_photo"`
		IsSellerBusinessOnboarded               bool          `json:"is_seller_business_onboarded"`
		IsBuyNowEnabled                         bool          `json:"is_buy_now_enabled"`
		IsPurchaseProtected                     bool          `json:"is_purchase_protected"`
		IsMarketplaceListingWithCheckoutOffers  string        `json:"__isMarketplaceListingWithCheckoutOffers"`
		CanBuyerMakeCheckoutOffer               bool          `json:"can_buyer_make_checkout_offer"`
		IsCartEnabled                           bool          `json:"is_cart_enabled"`
		SellerCms                               interface{}   `json:"seller_cms"`
		IsMarketplaceListingWithVacationMode    string        `json:"__isMarketplaceListingWithVacationMode"`
		VacationMode                            interface{}   `json:"vacation_mode"`
		IsGroupCommerceProductItemIsDeprecated  string        `json:"__isGroupCommerceProductItemIsDeprecated"`
		AllMessageThreads                       struct {
			IsEmpty bool `json:"is_empty"`
		} `json:"all_message_threads"`
		MarketplaceListingCategoryID       string      `json:"marketplace_listing_category_id"`
		ShouldShowTxnSurveyOnMas           bool        `json:"should_show_txn_survey_on_mas"`
		IncentiveCampaignForFreeShipping   interface{} `json:"incentive_campaign_for_free_shipping"`
		FirstMessageSuggestedValue         interface{} `json:"first_message_suggested_value"`
		IsMarketplaceListingWithSweepstake string      `json:"__isMarketplaceListingWithSweepstake"`
		SweepstakeEnabled                  bool        `json:"sweepstake_enabled"`
		SweepstakeStatus                   string      `json:"sweepstake_status"`
		SweepstakeContent                  interface{} `json:"sweepstake_content"`
		MarketplaceBumpInfo                interface{} `json:"marketplace_bump_info"`
		IsMarketplaceListingSellerNotices  string      `json:"__isMarketplaceListingSellerNotices"`
		FiEnhancedAppealData               struct {
			UseFiAppeals bool `json:"use_fi_appeals"`
		} `json:"fi_enhanced_appeal_data"`
		ListingSellerNotices                       []interface{}              `json:"listing_seller_notices"`
		IsMarketplaceListingWithEmailCommunication string                     `json:"__isMarketplaceListingWithEmailCommunication"`
		IsEmailCommunicationEnabled                bool                       `json:"is_email_communication_enabled"`
		FairMarketValueData                        interface{}                `json:"fair_market_value_data"`
		VehicleCondition                           interface{}                `json:"vehicle_condition"`
		VehicleExteriorColor                       interface{}                `json:"vehicle_exterior_color"`
		VehicleFeatures                            []interface{}              `json:"vehicle_features"`
		VehicleFuelType                            interface{}                `json:"vehicle_fuel_type"`
		VehicleIdentificationNumber                interface{}                `json:"vehicle_identification_number"`
		VehicleInteriorColor                       interface{}                `json:"vehicle_interior_color"`
		VehicleIsPaidOff                           interface{}                `json:"vehicle_is_paid_off"`
		VehicleMakeDisplayName                     interface{}                `json:"vehicle_make_display_name"`
		VehicleModelDisplayName                    interface{}                `json:"vehicle_model_display_name"`
		VehicleNumberOfOwners                      interface{}                `json:"vehicle_number_of_owners"`
		VehicleOdometerData                        VehicleOdometerData        `json:"vehicle_odometer_data"`
		VehicleRegistrationPlateInformation        interface{}                `json:"vehicle_registration_plate_information"`
		VehicleSellerType                          interface{}                `json:"vehicle_seller_type"`
		VehicleSpecifications                      interface{}                `json:"vehicle_specifications"`
		VehicleTitleStatus                         interface{}                `json:"vehicle_title_status"`
		VehicleTransmissionType                    interface{}                `json:"vehicle_transmission_type"`
		VehicleTrimDisplayName                     interface{}                `json:"vehicle_trim_display_name"`
		VehicleCarfaxReport                        interface{}                `json:"vehicle_carfax_report"`
		AttributeData                              []ProductItemAttributeData `json:"attribute_data"`
		Location                                   Location                   `json:"location"`
		OriginTarget                               OriginTarget               `json:"origin_target"`
		ShippingOffered                            bool                       `json:"shipping_offered"`
		LegalDisclosureImpressumURL                interface{}                `json:"legal_disclosure_impressum_url"`
		IsMarketplaceListingWithBadges             string                     `json:"__isMarketplaceListingWithBadges"`
		CommerceBadgesInfo                         CommerceBadgesInfo         `json:"commerce_badges_info"`
		ListingAddress                             interface{}                `json:"listing_address"`
		SellerPhoneNumber                          interface{}                `json:"seller_phone_number"`
		VehicleWebsiteLink                         interface{}                `json:"vehicle_website_link"`
		DealershipName                             interface{}                `json:"dealership_name"`
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
		PreRecordedVideos       []interface{}  `json:"pre_recorded_videos"`
	} `json:"target"`
	Viewer struct {
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
	BboxInner BboxInner `json:"__bbox"`
}
