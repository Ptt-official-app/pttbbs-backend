from typing import Annotated

from fastapi import APIRouter, Query

from ..types import Brdname
from .v1_types import (
    BoardDetail,
    BoardSummary,
    LoadAutoCompleteBoardsParams,
    LoadGeneralBoardsParams,
    LoadGeneralBoardsResult,
    LoadPopularBoardsResult,
    RefreshBrdnameBlackListMapResult,
    RefreshBrdnameWhiteListMapResult,
)

router = APIRouter(
    tags=['board']
)


@router.get('/boards/popular')
def load_popular_boards() -> LoadPopularBoardsResult:
    ...


@router.get('/board/{brdname}')
def get_board_detail(brdname: str) -> BoardDetail:
    ...


@router.get('/board/{brdname}/summary')
def get_board_summary(brdname: str) -> BoardSummary:
    ...


@router.get('/boards/refresh-white-list')
def refresh_brdname_white_list_map() -> RefreshBrdnameWhiteListMapResult:
    ...


@router.get('/boards/refresh-black-list')
def refresh_brdname_black_list_map() -> RefreshBrdnameBlackListMapResult:
    ...


@router.get('/boards')
def load_general_boards(
    params: Annotated[LoadGeneralBoardsParams, Query()],
) -> LoadGeneralBoardsResult:
    ...


@router.get('/boards/autocomplete')
def load_auto_complete_boards(
    params: Annotated[LoadAutoCompleteBoardsParams, Query()],
) -> LoadGeneralBoardsResult:
    ...
