from typing import Annotated

from fastapi import APIRouter, File, UploadFile

from ..types import Result
from .types import (
    CreateSiteParams,
    NodeInfo,
    SiteInfo,
    SiteView,
    UpdateSiteParams,
    UploadSiteBannerResult,
    UploadSiteIconResult,
)

router = APIRouter(
    tags=['site']
)


@router.get('/site')
def get_site_info() -> SiteInfo:
    ...


@router.post('/site')
def create_site(params: CreateSiteParams) -> SiteView:
    ...


@router.put('/site')
def update_site(params: UpdateSiteParams) -> SiteView:
    ...


@router.post('/site/icon')
def upload_site_icon(image: Annotated[UploadFile, File()]) -> UploadSiteIconResult:
    ...


@router.delete('/site/icon')
def delete_site_icon() -> Result:
    ...


@router.post('/site/banner')
def upload_site_banner(image: Annotated[UploadFile, File()]) -> UploadSiteBannerResult:
    ...


@router.delete('/site/banner')
def delete_site_banner() -> Result:
    ...


@router.get('/nodeinfo/2.1')
def get_node_info() -> NodeInfo:
    ...
