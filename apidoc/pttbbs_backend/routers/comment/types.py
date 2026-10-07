from pydantic import BaseModel

from ..community.types import Community, CommunityTagsView
from ..post.types import Post
from ..types import ID, DbURL, Report, Time3339
from ..user.types import CommentActions, CommunityActions, PersonActions
from ..user.types.person import Person


class Comment(BaseModel):
    locked: bool
    federation_pending: bool
    unresolved_report_count: int
    report_count: int
    child_count: int
    downvotes: int
    upvotes: int
    score: int
    language_id: ID
    distinguished: bool
    path: str
    local: bool
    ap_id: DbURL
    deleted: bool
    updated_at: Time3339 = ''
    published_at: Time3339
    removed: bool
    content: str
    post_id: ID
    creator_id: ID
    id: ID


class CommentReport(Report):
    violates_instance_rules: bool
    original_comment_text: str
    comment_id: ID
    creator_id: ID


class CommentReportView(BaseModel):
    tags: list[CommunityTagsView]
    creator_community_ban_expires_at: Time3339 = ''
    creator_banned_from_community: bool
    creator_ban_expires_at: Time3339 = ''
    creator_banned: bool
    creator_is_moderator: bool
    creator_is_admin: bool
    community_actions: CommunityActions | None = None
    person_actions: PersonActions | None = None
    resolver: Person | None = None
    comment_actions: CommentActions | None = None
    comment_creator: Person
    creator: Person
    community: Community
    post: Post
    comment: Comment
    comment_report: CommentReport
