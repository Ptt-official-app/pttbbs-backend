from fastapi import APIRouter

from ..types import ArticleID, Brdname
from .types import (
    ArticleSummary,
    CreateArticleParams,
    DeleteArticleParams,
    DeleteArticlesParams,
    DeleteCommentsParams,
    ReplyCommentsParams,
)

router = APIRouter(
    tags=['TODO']
)


@router.post('/board/{brdname}/article/draft')
def create_article_draft(brdname: Brdname, body: CreateArticleParams) -> ArticleSummary:
    ...


@router.put('/board/{brdname}/article/{aid}/draft')
def update_article_draft(
    brdname: Brdname, aid: ArticleID, body: CreateArticleParams
) -> ArticleSummary:
    ...


@router.post('/board/{brdname}/article/{aid}')
def create_article(brdname: Brdname, aid: ArticleID) -> ArticleSummary:
    ...


@router.put('/board/{brdname}/article/{aid}')
def edit_article(brdname: Brdname, aid: ArticleID, body: CreateArticleParams) -> ArticleSummary:
    ...


@router.post('/board/{brdname}/article{aid}/comments/reply')
def reply_comments(brdname: Brdname, aid: ArticleID, body: ReplyCommentsParams):
    ...


@router.delete('/board/{brdname}/article/{aid}/comments')
def delete_comments(brdname: Brdname, aid: ArticleID, body: DeleteCommentsParams):
    ...


@router.delete('/board/{brdname}/article/{aid}')
def delete_article(brdname: Brdname, aid: ArticleID, body: DeleteArticleParams):
    ...


@router.delete('/board/{brdname}/articles')
def delete_articles(brdname: Brdname, body: DeleteArticlesParams):
    ...
