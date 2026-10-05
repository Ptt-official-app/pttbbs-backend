from typing import Literal

from pydantic import BaseModel, Field

from ..types import ArticleID, Brdname, CommentID, Rune, Time8, Username

type ArticleType = Literal[
    '',

    'draft-post',
    'draft-re-post',
    'draft-fw-post',

    'draft-ascii',
    'draft-re-ascii',
    'draft-fw-ascii',

    'draft-movie',
    'draft-re-movie',
    'draft-fw-movie',

    'draft-rtl',
    'draft-re-rtl',
    'draft-fw-rtl',

    'post',
    're-post',
    'fw-post',

    'ascii',
    're-ascii',
    'fw-ascii',

    'movie',
    're-movie',
    'fw-movie',

    'rtl',
    're-rtl',
    'fw-rtl',
]

type CommentType = Literal[
    '',
    'recommend',
    'boo',
    'comment',
    'forward',

    'reply',
    'edit',
    'deleted',
]


class ArticleSummary(BaseModel):
    brdname: Brdname
    aid: ArticleID
    deleted: bool = False
    create_time: Time8
    modified: Time8
    recommend: int = Field(description='the score of recommend')
    n_comments: int = Field(description='total number of comments')
    owner: Username
    title: str
    money: int
    type: ArticleType = Field(default='', description='article type')
    the_class: str = Field(
        alias='class', description='moderator-defined class')

    url: str
    read: bool

    locked: bool = Field(description='locked and cannot be modified')
    marked: bool = Field(
        description='marked and cannot be deleted')

    idx: str
    tokenuser: Username

    # deprecated
    bid: str = Field(description='(deprecated)')
    subject_type: int = Field(description='(deprecated)')
    rank: int = Field(description='(deprecated)')


class ArticleDetail(ArticleSummary):
    ip: str
    host: str
    bbs: str

    Content: list[list[Rune]]
    ContentPrefix: list[list[Rune]]

    # deprecated


class ArticleBlock(BaseModel):
    Content: list[list[Rune]] | None = None

    deleted: bool = False
    create_time: Time8 = 0
    modified: Time8 = 0
    recommend: int = Field(default=0, description='the score of recommend')
    n_comments: int = Field(default=0, description='total number of comments')
    owner: Username = ''
    title: str = ''
    money: int = 0
    type: ArticleType = Field(default='', description='article type')
    the_class: str = Field(
        alias='class', default='', description='moderator-defined class')

    ip: str = ''
    host: str = ''
    bbs: str = ''

    next_idx: str = ''

    tokenuser: Username = ''

    # deprecated
    subject_type: int = Field(default=0, description='(deprecated)')
    rank: int = Field(default=0, description='(deprecated)')


class Comment(BaseModel):
    brdname: Brdname
    aid: ArticleID
    cid: CommentID
    type: CommentType
    refid: CommentID
    deleted: bool
    create_time: Time8
    sort_time: Time8
    owner: Username

    content: list[list[Rune]]

    ip: str
    host: str

    idx: str
    tokenuser: Username

    # deprecated
    bid: str = Field(description='(deprecated)')


class LoadGeneralArticlesParams(BaseModel):
    title: str = Field(default='', description='title (match)')
    start_idx: str = ''
    limit: int = Field(default=0, description="limit. 0: no limit.")
    desc: bool = Field(default=False, description='descending')


class LoadGeneralArticlesResult(BaseModel):
    list: list[ArticleSummary]
    next_idx: str

    tokenuser: Username = ''


class LoadBottomArticlesResult(BaseModel):
    list: list[ArticleSummary]
    tokenuser: Username = ''


class GetArticleBlocksParams(BaseModel):
    start_idx: str = ''
    limit: int = Field(default=0, description='limit. 0: no limit.')


class LoadArticleCommentsParams(BaseModel):
    start_idx: str = ''
    desc: bool = Field(default=False, description='is descending')
    limit: int = Field(default=0, description='limit. 0: no limit.')


class LoadArticleCommentsResult(BaseModel):
    list: list[Comment]
    next_idx: str

    tokenuser: Username = ''


class CreateArticleParams(BaseModel):
    the_class: str = Field(
        alias='class', description='moderator-defined class')

    title: str
    content: list[list[Rune]]


class DeleteArticleParams(BaseModel):
    reason: str


class DeleteEachArticleParams(BaseModel):
    aid: ArticleID
    reason: str


class DeleteArticlesParams(BaseModel):
    list: list[DeleteEachArticleParams]


class DeleteCommentParams(BaseModel):
    reason: str


class DeleteEachCommentParams(BaseModel):
    cid: CommentID
    reason: str


class DeleteCommentsParams(BaseModel):
    list: list[DeleteEachCommentParams]


class ReplyCommentParams(BaseModel):
    cid: CommentID
    content: list[list[Rune]]


class ReplyCommentsParams(BaseModel):
    list: list[ReplyCommentParams]
