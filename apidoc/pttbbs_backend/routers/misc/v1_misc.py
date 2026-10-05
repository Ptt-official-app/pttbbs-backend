from typing import Annotated

from fastapi import APIRouter, Query

from .v1_types import GetVersionResult, IndexParams, IndexResult

router = APIRouter(
    tags=['misc']
)


@router.get('/')
def index(params: Annotated[IndexParams, Query()]) -> IndexResult:
    ...


@router.get('/version')
def get_version() -> GetVersionResult:
    ...
