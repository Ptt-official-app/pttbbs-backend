from typing import Literal

from pydantic import BaseModel

from ..types import ID, DbURL, Time3339

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


class CommunityAction(BaseModel):
    notifications: CommunityNotificationsMode
    follow_state: CommunityFollowerState
    ban_expires_at: Time3339
    received_ban_at: Time3339
    became_moderator_at: Time3339
    blocked_at: Time3339
    followed_at: Time3339


class PersonActions(BaseModel):
    downvotes: int = 0
    upvotes: int = 0
    voted_at: Time3339 = ''
    note: str = ''
    noted_at: Time3339 = ''
    blocked_at: Time3339 = ''


class Person(BaseModel):
    comment_count: int
    post_count: int
    instance_id: ID
    bot_account: bool
    matrix_user_id: str = ''
    deleted: bool
    banner: DbURL = ''
    last_refreshed_at: Time3339
    local: bool
    bio: str = ''
    ap_id: DbURL
    updated_at: Time3339 = ''
    published_at: Time3339
    avatar: DbURL = ''
    display_name: str = ''
    name: str
    id: ID


class PersonView(BaseModel):
    community_actions: CommunityAction | None = None
    ban_expires_at: Time3339 = ''
    banned: bool
    person_actions: PersonActions | None = None
    is_admin: bool
    person: Person
