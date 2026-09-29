from typing import Any

from pydantic import BaseModel, Field


class IndexParams(BaseModel):
    the_in: int = Field(default=0, alias='in')


class IndexResult(BaseModel):
    data: Any


class GetVersionResult(BaseModel):
    version: str
    commit: str

    pttversion: str
    pttcommit: str
