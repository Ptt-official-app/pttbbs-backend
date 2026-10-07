from fastapi import APIRouter

from .account import v1_todo as account_todo
from .community import v1_todo as board_todo
from .post import v1_todo as article_todo
from .user import v1_todo as user_todo

v1_router = APIRouter()

v1_router.include_router(account_todo.router)
v1_router.include_router(board_todo.router)
v1_router.include_router(article_todo.router)
v1_router.include_router(user_todo.router)
