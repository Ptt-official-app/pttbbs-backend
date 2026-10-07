from pydantic import BaseModel, Field

from ..community.types import CommunityReportView
from ..media.types import LocalImageView
from ..message.types import PrivateMessageReportView
from ..post.types import VoteView
from ..report.types import ReportCombinedView
from ..site.types import TagLine
from ..types import ID, ListParams, ListResult, ReportSortType, ReportType, Time3339
from ..user.types import (
    LocalUserSortType,
    LocalUserView,
    PersonView,
    RegistrationApplicationView,
)


class DeleteMediaParams(BaseModel):
    filename: str


class ListMediaParams(ListParams):
    pass


class ListMediaResult(ListResult[LocalImageView]):
    pass


class HideCommunityParams(BaseModel):
    reason: str
    hidden: bool
    community_id: ID


class ResolveReportParams(BaseModel):
    conclusion: str
    resolved: bool
    report_id: ID


class ResolveCommunityReportResult(BaseModel):
    community_report_view: CommunityReportView


class ListPostLikeParams(ListParams):
    post_id: ID


class ListPostLikeResult(ListResult[VoteView]):
    pass


class ListCommentLikeParams(ListParams):
    post_id: ID


class ListCommentLikeResult(ListResult[VoteView]):
    pass


class ResolvePrivateMessageReportResult(BaseModel):
    private_message_report_view: PrivateMessageReportView


class BanPersonFromSiteParams(BaseModel):
    expires_at: Time3339 = ''
    reason: str
    remove_or_restore_data: bool = False
    ban: bool
    person_id: ID


class BanPersonFromSiteResult(BaseModel):
    person_view: PersonView


class ListUserParams(ListParams):
    sort: LocalUserSortType = 'new'
    banned_only: bool = False


class ListUserResult(ListResult[LocalUserView]):
    pass


class AddAdminParams(BaseModel):
    added: bool
    person_id: ID


class AddAdminResult(BaseModel):
    admins: list[PersonView]


class ListRegistrationApplicationParams(ListParams):
    unread_only: bool = False


class ListRegistrationApplicationResult(ListResult[RegistrationApplicationView]):
    pass


class ApproveRegistrationApplicationParams(BaseModel):
    approve: bool
    id: ID


class ApproveRegistrationApplicationResult(BaseModel):
    registration_application: RegistrationApplicationView


class GetRegistrationApplicationParams(BaseModel):
    person_id: ID


class GetRegistrationApplicationResult(BaseModel):
    registration_application: RegistrationApplicationView


class DeleteUserParams(BaseModel):
    reason: str
    person_id: ID


class DeleteCommunityParams(BaseModel):
    reason: str
    community_id: ID


class DeletePostParams(BaseModel):
    reason: str
    post_id: ID


class DeleteCommentParams(BaseModel):
    reason: str
    comment_id: ID


class CreateTagLineParams(BaseModel):
    content: str


class CreateTagLineResult(BaseModel):
    tagline: TagLine


class EditTagLineParams(BaseModel):
    content: str
    id: ID


class EditTagLineResult(BaseModel):
    tagline: TagLine


class DeleteTagLineParams(BaseModel):
    id: ID


class ListTagLineResult(ListResult[TagLine]):
    ...


class ListReportParams(ListParams):
    my_reports_only: bool = False
    show_community_rule_violations: bool = Field(
        default=False,
        description='Only for admins: also show reports with violates_instance_rules=false')
    sort: ReportSortType = 'old'
    community_id: ID = 0
    post_id: ID = 0
    type_: ReportType = 'all'
    unresolved_only: bool = False


class ListReportResult(ListResult[ReportCombinedView]):
    pass


class BlockInstanceParams(BaseModel):
    expires_at: int = Field(
        default=0,
        description='An i64 unix timestamp is used for a simpler API client implementation.'
    )
    reason: str
    block: bool
    instance: str


class AllowInstanceParams(BaseModel):
    reason: str
    allow: bool
    instance: str
