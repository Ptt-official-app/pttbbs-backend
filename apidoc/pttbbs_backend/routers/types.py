from typing import Annotated

from pydantic import BaseModel, Field

type Time8 = int
type Time3339 = Annotated[
    str,
    'in RFC-3339 format (YYYY-MM-DDTHH:MM:SS.ffffffZ)']


type Username = str
type Brdname = str
type ArticleID = str
type CommentID = str

type FavID = str

type ID = int

type DbURL = str


class Perm(BaseModel):
    is_basic: bool = Field(default=False, description="basic")
    is_chat: bool = Field(default=False, description="chat room (聊天室)")
    is_msg: bool = Field(default=False, description="message (message)")
    is_post: bool = Field(default=False, description="post (發表文章)")
    is_reg_verified: bool = Field(
        default=False, description="register verified (註冊認證通過)")
    is_cloak: bool = Field(default=False, description="hide (隱身)")
    is_seecloak: bool = Field(
        default=False, description="can see hidden users (可以看見隱身的 users)")
    is_permanent_account: bool = Field(
        default=False, description="permanent account (帳號永遠保留)"
    )
    is_sysop_hide: bool = Field(
        default=False, description="sysop hidden (站長隱身術)"
    )
    is_bm: bool = Field(default=False, description='moderators (板主)')
    is_account: bool = Field(
        default=False, description='chief account officer (帳號總管)')
    is_chatroom: bool = Field(
        default=False, description='chief chatroom officer (聊天室總管)')
    is_board: bool = Field(default=False, description='(board officer (看板總管)')
    is_sysop: bool = Field(default=False, description='sysop (站長)')
    is_bbsadm: bool = Field(default=False, description='BBS administrator')
    is_notop: bool = Field(
        default=False, description='no leaderboard (不在排行榜裡)')
    is_violate_law: bool = Field(
        default=False, description='violate law (違反法律)')
    is_angel: bool = Field(
        default=False, description='qual to be an angel (有資格當小天使)')
    is_no_reg_code: bool = Field(
        default=False, description='no registration code (不允許認證碼註冊)')
    is_view_sysop: bool = Field(
        default=False, description='chief design officer (視覺站長)')
    is_log_user: bool = Field(
        default=False, description='check user behaviors (可以觀察使用者行蹤)')
    is_no_citizen: bool = Field(
        default=False, description='deprived rights (褫奪公權)')
    is_class_officer: bool = Field(
        default=False, description='chief class officer (群組長)')
    is_acctreg: bool = Field(
        default=False, description='account officer (帳號組)')
    is_prg: bool = Field(
        default=False, description='software engineer (軟體工程師組)')
    is_action: bool = Field(default=False, description='event officer (活動組)')
    is_paint: bool = Field(default=False, description='design officer (美工組)')
    is_police_man: bool = Field(
        default=False, description='chief police officer (警察總管)')
    is_syssubop: bool = Field(default=False, description='class officer (小組長)')
    is_oldsysop: bool = Field(
        default=False, description='sysop emeritus (退休站長)')
    is_police: bool = Field(default=False, description='police officer (警察)')


type ColorMap = int


class Color(BaseModel):
    foreground: ColorMap
    background: ColorMap
    blink: bool
    highlight: bool
    reset: bool


class Rune(BaseModel):
    text: str
    color0: Color
    color1: Color
    is_url: bool
