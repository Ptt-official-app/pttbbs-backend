from pydantic import BaseModel, Field

from ..community.v1_types import BoardSummary
from ..types import Perm, Time8, Username


class UserDetailCore(BaseModel):
    username: Username
    realname: str
    nickname: str

    is_government_id: bool = Field(
        description='is verified with government-id (zk)')

    login_days: int
    posts: int
    first_login: Time8
    last_login: Time8
    last_ip: str
    last_host: str

    money: int

    # deprecated
    pttemail: str = Field(default='', description='(deprecated)')
    justify: str = Field(default='', description='(deprecated)')


class UserDetail(UserDetailCore, Perm):
    pass


class GetUsernameResult(BaseModel):
    username: Username

    tokenuser: Username = ''


class LoadFavoriteBoardsParams(BaseModel):
    level_idx: str = ''
    start_idx: str = ''
    asc: bool = Field(default=False, description='ascending')
    limit: int = Field(default=0, description='limit. 0: no limit.')


class LoadFavoriteBoardsResult(BaseModel):
    list: list[BoardSummary]
    next_idx: str

    tokenuser: Username = ''


class AddFavoriteBoardParams(BaseModel):
    level_idx: str = ''
    brdname: str


class AddFavoriteFolderParams(BaseModel):
    level_idx: str = ''
    title: str = ''


class AddFavoriteLineParams(BaseModel):
    level_idx: str = ''


class DeleteFavoriteParams(BaseModel):
    level_idx: str = ''
