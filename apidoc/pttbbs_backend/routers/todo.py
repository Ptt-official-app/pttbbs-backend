from fastapi import APIRouter

from .account import todo as account_todo
from .article import todo as article_todo
from .board import todo as board_todo
from .user import todo as user_todo

router = APIRouter()

router.include_router(account_todo.router)
router.include_router(board_todo.router)
router.include_router(article_todo.router)
router.include_router(user_todo.router)
