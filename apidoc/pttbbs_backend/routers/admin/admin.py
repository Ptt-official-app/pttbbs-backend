from typing import Annotated

from fastapi import APIRouter, Query

from ..types import ListParams, Result
from .types import (
    AddAdminParams,
    AddAdminResult,
    AllowInstanceParams,
    ApproveRegistrationApplicationParams,
    ApproveRegistrationApplicationResult,
    BanPersonFromSiteParams,
    BanPersonFromSiteResult,
    BlockInstanceParams,
    CreateTagLineParams,
    CreateTagLineResult,
    DeleteCommentParams,
    DeleteCommunityParams,
    DeleteMediaParams,
    DeletePostParams,
    DeleteTagLineParams,
    DeleteUserParams,
    EditTagLineParams,
    EditTagLineResult,
    GetRegistrationApplicationParams,
    GetRegistrationApplicationResult,
    HideCommunityParams,
    ListCommentLikeParams,
    ListCommentLikeResult,
    ListMediaParams,
    ListMediaResult,
    ListPostLikeParams,
    ListPostLikeResult,
    ListRegistrationApplicationParams,
    ListRegistrationApplicationResult,
    ListReportParams,
    ListReportResult,
    ListTagLineResult,
    ListUserParams,
    ListUserResult,
    ResolveCommunityReportResult,
    ResolvePrivateMessageReportResult,
    ResolveReportParams,
)

router = APIRouter(
    tags=['admin']
)


@router.post('/admin/leave')
def leave_admin():
    ...


@router.delete('/image', tags=['media'])
def delete_media(params: DeleteMediaParams) -> Result:
    ...


@router.get('/image/list', tags=['media'])
def list_media(params: Annotated[ListMediaParams, Query()]) -> ListMediaResult:
    ...


@router.put('/community/hide', tags=['community'])
def hide_community(params: HideCommunityParams) -> Result:
    ...


@router.put('/community/report/resolve', tags=['community'])
def resolve_community_report(params: ResolveReportParams) -> ResolveCommunityReportResult:
    ...


@router.get('/post/like/list', tags=['post'])
def list_post_like(params: ListPostLikeParams) -> ListPostLikeResult:
    ...


@router.get('/comment/like/list', tags=['comment'])
def list_comment_like(params: ListCommentLikeParams) -> ListCommentLikeResult:
    ...


@router.put('/private_message/report/resolve', tags=['community'])
def resolve_private_message_report(
        params: ResolveReportParams
) -> ResolvePrivateMessageReportResult:
    ...


@router.post('/admin/ban', tags=['person'])
def ban_person_from_site(params: BanPersonFromSiteParams) -> BanPersonFromSiteResult:
    ...


@router.get('/admin/users', tags=['person'])
def list_user(params: Annotated[ListUserParams, Query()]) -> ListUserResult:
    ...


@router.post('/admin/add')
def add_admin(params: AddAdminParams) -> AddAdminResult:
    ...


@router.get('/admin/registration_application/list', tags=['person'])
def list_registration_application(
    params: Annotated[ListRegistrationApplicationParams, Query()]
) -> ListRegistrationApplicationResult:
    ...


@router.put('/admin/registration_application/approve', tags=['person'])
def approve_registration_application(
    params: ApproveRegistrationApplicationParams,
) -> ApproveRegistrationApplicationResult:
    ...


@router.get('/admin/registration_application', tags=['person'])
def get_registration_application(
    params: Annotated[GetRegistrationApplicationParams, Query()]
) -> GetRegistrationApplicationResult:
    ...


@router.post('/admin/purge/person', tags=['person'])
def delete_user(params: DeleteUserParams) -> Result:
    ...


@router.post('/admin/purge/community', tags=['community'])
def delete_community(params: DeleteCommunityParams) -> Result:
    ...


@router.post('/admin/purge/post', tags=['post'])
def delete_post(params: DeletePostParams) -> Result:
    ...


@router.post('/admin/purge/comment', tags=['post'])
def delete_comment(params: DeleteCommentParams) -> Result:
    ...


@router.post('/admin/tagline', tags=['site'])
def create_tagline(params: CreateTagLineParams) -> CreateTagLineResult:
    ...


@router.put('/admin/tagline', tags=['site'])
def edit_tagline(params: EditTagLineParams) -> EditTagLineResult:
    ...


@router.delete('/admin/tagline', tags=['site'])
def delete_tagline(params: DeleteTagLineParams) -> Result:
    ...


@router.get('/admin/tagline/list', tags=['site'])
def list_taglines(params: Annotated[ListParams, Query()]) -> ListTagLineResult:
    ...


@router.get('/report/list', tags=['user'])
def list_reports(params: Annotated[ListReportParams, Query()]) -> ListReportResult:
    ...


@router.post('/admin/instance/block', tags=['federation'])
def block_instance(params: BlockInstanceParams) -> Result:
    ...


@router.post('/admin/instance/allow', tags=['federation'])
def allow_instance(params: AllowInstanceParams) -> Result:
    ...
