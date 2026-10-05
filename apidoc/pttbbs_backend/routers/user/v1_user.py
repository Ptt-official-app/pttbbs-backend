from typing import Annotated

from fastapi import APIRouter, Query

from ..board.v1_types import BoardSummary
from ..types import FavID, Username
from .v1_types import (
    UserDetail,
)

router = APIRouter(
    tags=['user']
)


@router.get('/user/{username}')
def get_user_info(username: Username) -> UserDetail:
    ...


@router.get('/username')
def get_username():
    ...
