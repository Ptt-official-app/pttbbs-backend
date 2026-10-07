from typing import Annotated

from fastapi import APIRouter, Query

from .v1_types import (
    ArticleDetail,
    GetArticleBlocksParams,
    LoadArticleCommentsParams,
    LoadArticleCommentsResult,
    LoadBottomArticlesResult,
    LoadGeneralArticlesParams,
    LoadGeneralArticlesResult,
)

router = APIRouter(
    tags=['post']
)


@router.get('/board/{brdname}/articles')
def load_general_articles(
        brdname: str,
        params: Annotated[LoadGeneralArticlesParams, Query()]
) -> LoadGeneralArticlesResult:
    ...


@router.get('/board/{brdname}/articles/bottom')
def load_bottom_articles(brdname: str) -> LoadBottomArticlesResult:
    ...


@router.get('/board/{brdname}/article/{aid}')
def get_article(brdname: str, aid: str) -> ArticleDetail:
    ...


@router.get('/board/{brdname}/article/{aid}/blocks')
def get_article_blocks(brdname: str, aid: str, params: Annotated[GetArticleBlocksParams, Query()]):
    ...


@router.get('/board/{brdname}/article/{aid}/comments', tags=['comment'])
def load_article_comments(
    params: Annotated[LoadArticleCommentsParams, Query()],
) -> LoadArticleCommentsResult:
    ...
