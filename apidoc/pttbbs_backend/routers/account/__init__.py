
from fastapi import APIRouter

from . import account

router = APIRouter()

router.include_router(
    account.router,
)
