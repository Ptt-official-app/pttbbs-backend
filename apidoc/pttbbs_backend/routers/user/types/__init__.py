from typing import Literal

from pydantic import BaseModel

from ...types import (
    ID,
    CommentSortType,
    ListingType,
    PostListingMode,
    PostNotificationMode,
    PostSortType,
    Time3339,
)
from .person import Person

type CommunityNotificationsMode = Literal[
    'all_posts_and_comments',
    'all_posts',
    'replies_and_mentions',
    'mute'
]

type CommunityFollowerState = Literal[
    'accepted',
    'pending',
    'approval_required',
    'denied',
]


class CommunityActions(BaseModel):
    notifications: CommunityNotificationsMode
    follow_state: CommunityFollowerState
    ban_expires_at: Time3339
    received_ban_at: Time3339
    became_moderator_at: Time3339
    blocked_at: Time3339
    followed_at: Time3339


type LocalUserSortType = Literal['new', 'old']


class PostActions(BaseModel):
    notifications: PostNotificationMode = 'all_comments'
    vote_is_upvote: bool = False
    read_comments_amount: int = 0
    hidden_at: Time3339 = ''
    voted_at: Time3339 = ''
    saved_at: Time3339 = ''
    read_comments_at: Time3339 = ''
    read_at: Time3339 = ''


class PersonActions(BaseModel):
    downvotes: int = 0
    upvotes: int = 0
    voted_at: Time3339 = ''
    note: str = ''
    noted_at: Time3339 = ''
    blocked_at: Time3339 = ''


class CommentActions(BaseModel):
    vote_is_upvote: bool = False
    saved_at: Time3339 = ''
    voted_at: Time3339 = ''


class PersonView(BaseModel):
    community_actions: CommunityActions | None = None
    ban_expires_at: Time3339 = ''
    banned: bool
    person_actions: PersonActions | None = None
    is_admin: bool
    person: Person


type VoteShow = Literal['show', 'show_for_others', 'hide']


class LocalUser(BaseModel):
    show_media: bool
    invited_by_local_user_id: ID = 0
    default_items_per_page: int
    show_person_votes: bool
    show_upvote_percentage: bool
    show_downvotes: VoteShow
    show_upvotes: bool
    show_score: bool
    default_post_time_range_seconds: float = 0
    hide_posts_with_media: bool
    auto_mark_fetched_posts_as_read: bool
    default_comment_sort_type: CommentSortType
    private_messages_enabled: bool
    last_donation_notification_at: Time3339
    collapse_bot_comments: bool
    animated_images_enabled: bool
    totp_2fa_enabled: bool
    post_listing_mode: PostListingMode
    admin: bool
    infinite_scroll_enabled: bool
    blur_nsfw: bool
    open_links_in_new_tab: bool
    accepted_application: bool
    email_verified: bool
    show_read_posts: bool
    show_bot_accounts: bool
    send_notifications_to_email: bool
    show_avatars: bool
    interface_language: str
    default_listing_type: ListingType
    default_post_sort_type: PostSortType
    theme: str
    show_nsfw: bool
    email: str
    person_id: ID
    id: ID


class LocalUserView(BaseModel):
    ban_expires_at: Time3339 = ''
    banned: bool
    person: Person
    local_user: LocalUser


class RegistrationApplication(BaseModel):
    updated_at: Time3339 = ''
    published_at: Time3339
    deny_reason: str = ''
    admin_id: ID = 0
    answer: str
    local_user_id: ID
    id: ID


class RegistrationApplicationView(BaseModel):
    admin: Person | None = None
    creator: Person
    creator_local_user: LocalUser
    registration_application: RegistrationApplication
