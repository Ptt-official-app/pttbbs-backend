from typing import Literal

from pydantic import BaseModel

from ..types import ID, DbURL, Report, Time3339
from ..user.types.person import Person

type CommunityVisibility = Literal[
    'public',
    'unlisted',
    'local_only_public',
    'local_only_private',
    'private']


class Community(BaseModel):
    local_removed: bool
    unresolved_report_count: int
    report_count: int
    subscribers_local: int
    users_active_half_year: int
    users_active_month: int
    users_active_week: int
    users_active_day: int
    comments: int
    posts: int
    subscribers: int
    summary: str = ''
    visibility: CommunityVisibility
    instance_id: ID
    posting_restricted_to_mods: bool
    banner: DbURL = ''
    icon: DbURL = ''
    last_refreshed_at: Time3339
    ap_id: DbURL
    nsfw: bool
    deleted: bool
    updated_at: Time3339 = ''
    published_at: Time3339
    removed: bool
    sidebar: str = ''
    title: str = ''
    name: str
    id: ID


class CommunityReport(Report):
    original_community_banner: str = ''
    original_community_icon: str = ''
    original_community_sidebar: str = ''
    original_community_summary: str = ''
    original_community_title: str = ''
    original_community_name: str = ''
    community_id: ID
    creator_id: ID


class CommunityReportView(BaseModel):
    creator_community_ban_expires_at: Time3339 = ''
    creator_banned_from_community: bool
    creator_ban_expires_at: Time3339 = ''
    creator_banned: bool
    creator_is_moderator: bool
    creator_is_admin: bool
    resolver: Person | None = None
    creator: Person
    community: Community
    community_report: CommunityReport


type Color = Literal[
    'color01', 'color02', 'color03', 'color04', 'color05', 'color06', 'color07', 'color08', 'color09', 'color10']  # noqa


class CommunityTagsView(BaseModel):
    color: Color
    deleted: bool
    updated_at: Time3339 = ''
    published_at: Time3339
    community_id: ID
    summary: str = ''
    display_name: str = ''
    name: str
    ap_id: DbURL
    id: ID
