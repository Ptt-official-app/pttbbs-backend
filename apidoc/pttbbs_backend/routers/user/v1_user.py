from fastapi import APIRouter

from ..types import Username
from .v1_types import (
    UserDetail,
)

router = APIRouter(
    tags=['person']
)


@router.get('/user/{username}')
def get_user_info(username: Username) -> UserDetail:
    ...


@router.get('/username')
def get_username():
    ...
