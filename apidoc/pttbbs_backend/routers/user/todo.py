from typing import Annotated

from fastapi import APIRouter, Query

from ..board.types import BoardSummary
from ..types import FavID, Username
from .types import (
    AddFavoriteBoardParams,
    AddFavoriteFolderParams,
    AddFavoriteLineParams,
    LoadFavoriteBoardsParams,
    LoadFavoriteBoardsResult,
)

router = APIRouter(
    tags=['TODO']
)


@router.get('/user/{username}/favorites')
def load_favorite_boards(
    username: Username, params: Annotated[LoadFavoriteBoardsParams, Query()],
) -> LoadFavoriteBoardsResult:
    ...


@router.post('/user/{username}/favorites/addboard')
def add_favorite_board(username: Username, body: AddFavoriteBoardParams) -> BoardSummary:
    ...


@router.post('/user/{username}/favorites/addfolder')
def add_favorite_folder(username: Username, body: AddFavoriteFolderParams) -> BoardSummary:
    ...


@router.post('/user/{username}/favorites/addline')
def add_favorite_line(username: Username, body: AddFavoriteLineParams) -> BoardSummary:
    ...


@router.delete('/user/{username}/favorite/{fav_id}')
def delete_favorite(username: Username, fav_id: FavID):
    ...
