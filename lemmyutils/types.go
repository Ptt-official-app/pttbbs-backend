package lemmyutils

import "github.com/Ptt-official-app/pttbbs-backend/types"

type CommunityView struct {
	Tags             []CommunityTagsView `json:"tags"`
	CanModify        bool                `json:"can_mod"`
	CommunityActions CommunityActions    `json:"community_actions,omitempty"`
	Community        Community           `json:"community"`
}

type CommunityTagsView struct {
	Color       string         `json:"color"`
	Deleted     bool           `json:"deleted"`
	UpdatedAt   types.Time8861 `json:"updated_at,omitempty"`
	PublishedAt types.Time8861 `json:"published_at"`
	CommunityID int            `json:"community_id"`
	Summary     string         `json:"summary,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Name        string         `json:"name"`
}

type CommunityNotificationsMode string

type CommunityFollowerState string

type CommunityActions struct {
	Notifications     CommunityNotificationsMode `json:"notifications,omitempty"`
	FollowState       CommunityFollowerState     `json:"follow_state,omitempty"`
	BanExpiresAt      types.Time8861             `json:"ban_expires_at,omitempty"`
	ReceivedBanAt     types.Time8861             `json:"received_ban_at,omitempty"`
	BecomeModeratorAt types.Time8861             `json:"became_moderator_at,omitempty"`
	BlockedAt         types.Time8861             `json:"blocked_at,omitempty"`
	FollowedAt        types.Time8861             `json:"followed_at,omitempty"`
}

type VoteSettings string

type CommunityVisibility string

type DbURL string

type CommunityId int

type Community struct {
	CommentDownvotes        VoteSettings        `json:"comment_downvotes"`
	CommentUpvotes          VoteSettings        `json:"comment_upvotes"`
	PostDownvotes           VoteSettings        `json:"post_downvotes"`
	PostUpvotes             VoteSettings        `json:"post_upvotes"`
	LocalRemoved            bool                `json:"local_removed"`
	UnresolvedReportCount   int                 `json:"unresolved_report_count"`
	ReportCount             int                 `json:"report_count"`
	SubscribersLocal        int                 `json:"subscribers_local"`
	UsersActiveHalfYear     int                 `json:"users_active_half_year"`
	UsersActiveMonth        int                 `json:"users_active_month"`
	UsersActiveWeek         int                 `json:"users_active_week"`
	UsersActiveDay          int                 `json:"users_active_day"`
	Comments                int                 `json:"comments"`
	Posts                   int                 `json:"posts"`
	Subscribers             int                 `json:"subscribers"`
	Summary                 string              `json:"summary,omitempty"`
	Visibility              CommunityVisibility `json:"visibility"`
	InstanceID              int                 `json:"instance_id"`
	PostingRestrictedToMods bool                `json:"posting_restricted_to_mods"`
	Banner                  DbURL               `json:"banner,omitempty"`
	Icon                    DbURL               `json:"icon,omitempty"`
	LastRefreshedAt         types.Time8861      `json:"last_refreshed_at"`
	Local                   bool                `json:"local"`
	APID                    DbURL               `json:"ap_id"`
	NSFW                    bool                `json:"nsfw"`
	Deleted                 bool                `json:"deleted"`
	UpdatedAt               types.Time8861      `json:"updated_at,omitempty"`
	PublishedAt             types.Time8861      `json:"published_at"`
	Removed                 bool                `json:"removed"`
	Sidebar                 string              `json:"sidebar,omitempty"`
	Title                   string              `json:"title,omitempty"`
	Name                    string              `json:"name"`
	ID                      CommunityId         `json:"id"`
}
