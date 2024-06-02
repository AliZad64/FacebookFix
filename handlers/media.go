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
