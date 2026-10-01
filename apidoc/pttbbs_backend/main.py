import importlib.metadata

from fastapi import FastAPI
from pydantic import BaseModel

from .routers import account, article, board, misc, todo, user, zk

_PREFIX = '/api/v1'


class Resp500(BaseModel):
    '''
    errResult

    [https://github.com/Ptt-official-app/pttbbs-backend/blob/main/api/types.go#L23](https://github.com/Ptt-official-app/pttbbs-backend/blob/main/api/types.go#L23)
    '''
    Msg: str
    tokenuser: str


app = FastAPI(
    version=importlib.metadata.version("pttbbs-backend"),
    title='pttbbs-backend',
    responses={
        500: {"model": Resp500}
    }
)

app.include_router(
    misc.router,
    prefix=_PREFIX
)

app.include_router(
    account.router,
    prefix=_PREFIX
)

app.include_router(
    zk.router,
    prefix=_PREFIX
)

app.include_router(
    user.router,
    prefix=_PREFIX
)

app.include_router(
    board.router,
    prefix=_PREFIX
)

app.include_router(
    article.router,
    prefix=_PREFIX
)


app.include_router(
    todo.router,
    prefix=_PREFIX
)
