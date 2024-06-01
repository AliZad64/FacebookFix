package handlers

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
