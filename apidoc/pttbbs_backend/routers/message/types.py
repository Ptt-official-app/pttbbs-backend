from pydantic import BaseModel

from ..types import ID, DbURL, Report, Time3339
from ..user.types import Person


class PrivateMessage(BaseModel):
    deleted_by_recipient: bool
    removed: bool
    local: bool
    ap_id: DbURL
    updated_at: Time3339 = ''
    published_at: Time3339
    deleted: bool
    content: str
    recipient_id: ID
    creator_id: ID
    id: ID


class PrivateMessageReport(Report):
    original_pm_text: str
    private_message_id: ID
    creator_id: ID


class PrivateMessageReportView(BaseModel):
    creator_ban_expires_at: Time3339 = ''
    creator_banned: bool
    creator_is_admin: bool
    resolver: Person | None = None
    private_message_creator: Person
    creator: Person
    private_message: PrivateMessage
    private_message_report: PrivateMessageReport
