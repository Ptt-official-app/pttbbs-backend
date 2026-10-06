import importlib.metadata

from fastapi import FastAPI
from pydantic import BaseModel

from .routers import account, article, board, misc, site, todo, user, zk

_V1_PREFIX = '/api/v1'

_V4_PREFIX = '/api/v4'


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
    site.router,
    prefix=_V4_PREFIX
)


app.include_router(
    misc.v1_router,
    prefix=_V1_PREFIX
)

app.include_router(
    account.v1_router,
    prefix=_V1_PREFIX
)

app.include_router(
    zk.router,
    prefix=_V1_PREFIX
)

app.include_router(
    user.v1_router,
    prefix=_V1_PREFIX
)

app.include_router(
    board.v1_router,
    prefix=_V1_PREFIX
)

app.include_router(
    article.v1_router,
    prefix=_V1_PREFIX
)


app.include_router(
    todo.v1_router,
    prefix=_V1_PREFIX
)
