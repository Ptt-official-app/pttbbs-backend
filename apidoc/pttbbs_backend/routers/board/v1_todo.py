from fastapi import APIRouter

from ..types import Brdname
from .v1_types import (
    BoardSummary,
    CreateBoardParams,
    LoadGeneralBoardsResult,
    RefreshBrdnameBlackListMapResult,
    RefreshBrdnameWhiteListMapResult,
)

router = APIRouter(
    tags=['TODO']
)


@router.get('/boards/white-list/refresh')
def refresh_brdname_white_list_map() -> RefreshBrdnameWhiteListMapResult:
    ...


@router.get('/boards/black-list/refresh')
def refresh_brdname_black_list_map() -> RefreshBrdnameBlackListMapResult:
    ...


@router.post('/cls/{clsname}/board')
def create_board(clsname: Brdname, body: CreateBoardParams) -> BoardSummary:
    ...


@router.delete('/board/{brdname}')
def delete_board(brdname: Brdname):
    ...


@router.get('/cls/{clsname}/boards')
def load_class_boards(clsname: Brdname) -> LoadGeneralBoardsResult:
    ...


@router.get('/cls/{clsname}')
def get_class(clsname: Brdname) -> BoardSummary:
    ...


@router.post('/cls')
def create_class() -> BoardSummary:
    ...


@router.delete('/cls/{clsname}')
def delete_class(clsname: Brdname):
    ...
