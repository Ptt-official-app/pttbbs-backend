from typing import Literal

from pydantic import BaseModel

from ..community.types import Community, CommunityTagsView
from ..types import ID, DbURL, Report, Time3339
from ..user.types import CommunityActions, Person, PersonActions, PostActions


class Post(BaseModel):
    embed_video_height: int = 0
    embed_video_width: int = 0
    federation_pending: bool
    unresolved_report_count: int
    report_count: int
    downvotes: int
    upvotes: int
    score: int
    comments: int
    newest_comment_time_at: Time3339 = ''
    scheduled_publish_time_at: Time3339 | Literal['None'] = 'None'
    alt_text: str = ''
    url_content_type: str = ''
    featured_local: bool
    featured_community: bool
    language_id: ID
    embed_video_url: DbURL = ''
    local: bool
    ap_id: DbURL
    thumbnail_url: DbURL = ''
    embed_description: str = ''
    embed_title: str = ''
    nsfw: bool
    deleted: bool
    updated_at: Time3339 = ''
    published_at: Time3339
    locked: bool
    removed: bool
    community_id: ID
    creator_id: ID
    body: str = ''
    url: DbURL = ''
    name: str
    id: ID


class VoteView(BaseModel):
    is_upvote: bool
    creator_banned_from_community: bool
    creator_banned: bool
    creator: Person


class PostReport(Report):
    violates_instance_rules: bool
    original_post_body: str
    original_post_url: DbURL
    original_post_name: str
    post_id: ID
    creator_id: ID


class PostReportView(BaseModel):
    tags: list[CommunityTagsView]
    creator_community_ban_expires_at: Time3339 = ''
    creator_banned_from_community: bool
    creator_ban_expires_at: Time3339 = ''
    creator_banned: bool
    creator_is_moderator: bool
    creator_is_admin: bool
    resolver: Person | None = None
    person_actions: PersonActions | None = None
    post_actions: PostActions | None = None
    community_actions: CommunityActions | None = None
    creator: Person
    community: Community
    post: Post
    post_report: PostReport
