
from typing import Literal

from pydantic import BaseModel, Field

from ..types import Perm, Time8

type LogCommentType = Literal['', 'ip', 'country']

type BoardType = Literal['board', 'class', 'folder']

type SortBy = Literal['name', 'class']


class BrdAttr(BaseModel):
    is_no_stats: bool = Field(
        default=False, description='not in stats (不列入統計)')
    is_group_board: bool = Field(
        default=False, description='class board (class 板)')
    is_hide: bool = Field(default=False, description='is hidden (隱板)')
    is_post_mask: bool = Field(
        default=False, description='restricting post/read (限制發文/閱讀)')
    is_anony: bool = Field(
        default=False, description='is anonymous (post-users) (匿名板)')
    is_default_anony: bool = Field(
        default=False, description='is default anonymous (預設匿名)')
    is_no_credit: bool = Field(
        default=False, description='no rewarding for posts (發文無獎勵)')
    is_vote_board: bool = Field(
        default=False, description='is vote board (連署板)')
    is_warn_eol: bool = Field(
        default=False, description='is in eol mode (即將廢除)')
    is_no_comment: bool = Field(
        default=False, description='no comments (不可推文)')
    is_angel_anony: bool = Field(
        default=False, description='angels can be anonymized (小天使可以匿名)')
    is_no_boo: bool = Field(default=False, description='no boo (不可噓)')
    is_board_member_only_post: bool = Field(
        default=False, description='board-member-only post (板友才可發文)')
    is_geust_post: bool = Field(
        default=False, description='guests can post (guest 可以發文)')
    is_cooldown: bool = Field(default=False, description='cooldown (靜)')
    is_no_fast_comment: bool = Field(
        default=False, description='no fast comments (禁止快速推文)')
    log_comment_type: LogCommentType = Field(
        default='', description="type of logging in comment ('', 'ip', 'country') (推文紀錄 (無) 或 IP 或 國家)")  # noqa
    is_no_reply: bool = Field(default=False, description='no reply (不可回文)')
    is_aligned_comment: bool = Field(
        default=False, description='aligned comments (推文對齊)')
    is_no_self_del_post: bool = Field(
        default=False, description='post-users cannot delete articles (不可自己刪文)')
    is_bm_mask_content: bool = Field(
        default=False, description='moderators can delete/forbid some words (板主可以刪除/禁止特定文字)')  # noqa


class BoardSummaryCore(BaseModel):
    brdname: str
    title: str

    the_type: BoardType = Field(alias='type')
    the_class: str = Field(
        alias='class', description='moderator-defined classes')
    nuser: int
    moderators: list[str]
    reason: str
    read: bool
    fav: bool
    total: int

    last_post_time: Time8

    level_idx: str = ''

    url: str = ''

    n_vote_question: int = Field(
        description='number of voting questions (正在進行的投票案)')
    next_vote_expire_time: Time8 = Field(
        description='next vote expire time (下次投票結束時間)')

    n_gamble: int = Field(description='number of gambles (正在進行的賭盤)')
    next_gamble_expire_time: Time8 = Field(
        description='next gamble expire time (下次賭盤結束時間)')

    clsname: str

    is_popular: bool
    is_over_18: bool

    idx: str = ''
    tokenuser: str = ''

    # to deprecate
    bid: str = Field(description='(deprecated)')
    flag: int = Field(description='(deprecated)')
    stat_attr: int = Field(description='(deprecated)')
    gid: int = Field(description='(deprecated)')
    pttbid: int = Field(description='(deprecated)')
    is_sym_link: bool = Field(default=False, description='(deprecated)')
    is_ip_log_comment: bool = Field(default=False, description='(deprecated)')


class BoardSummary(BrdAttr, BoardSummaryCore):
    pass


class BoardDetail(Perm, BoardSummary):
    update_time: Time8

    vote_limit_logins: int
    post_limit_logins: int
    vote_limit_bad_post: int
    post_limit_bad_post: int

    post_type: list[str] = Field(description='moderator-defined post types')
    post_tmpl: list[bool] = Field(description='with templates')

    fast_recommend_pause: Time8

    # to deprecate
    vote: int = Field(description='(deprecated)')
    link_pttbid: int = Field(description='(deprecated)')
    chesscountry: str = Field(default='', description='(deprecated)')


class LoadPopularBoardsResult(BaseModel):
    list: list[BoardSummary]
    tokenuser: str


class RefreshBrdnameBlackListMapResult(BaseModel):
    total: int


class RefreshBrdnameWhiteListMapResult(BaseModel):
    total: int


class LoadGeneralBoardsParams(BaseModel):
    title: str = ''
    keyword: str = ''
    start_idx: str = ''
    asc: bool = Field(default=False, description='ascending')
    limit: int = Field(
        default=0, description='limit of the boards. 0: no limit.')


class LoadGeneralBoardsResult(BaseModel):
    list: list[BoardSummary]
    next_idx: str = ''

    tokenuser: str = ''


class LoadAutoCompleteBoardsParams(BaseModel):
    brdname: str = Field(default='', description='brdname (starting with)')
    start_idx: str = ''
    asc: bool = Field(default=False, description='ascending')
    limit: int = Field(
        default=0, description='limit of the boards. 0: no limit.')


class CreateBoardParamsCore(BaseModel):
    brdname: str
    the_class: str = Field(alias='class')
    title: str
    bms: list[str] = Field(description='moderators')

    is_group: bool = Field(default=False, description='is group')


class CreateBoardParams(Perm, BrdAttr, CreateBoardParamsCore):
    pass


class LoadClassBoardsParams(BaseModel):
    start_idx: str = ''
    sortby: SortBy = 'class'
    asc: bool = Field(default=False, description='is ascending')
    limit: int = Field(
        default=0, description='limit of the boards. 0: no limit.')
