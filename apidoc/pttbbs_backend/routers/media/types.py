from typing import Literal

from pydantic import BaseModel

from ..post.types import Post
from ..types import ID, Time3339
from ..user.types import Person

type ImageMode = Literal['none', 'store_link_previews', 'proxy_all_images']


class LocalImage(BaseModel):
    thumbnail_for_post_id: ID
    person_id: ID
    published_at: Time3339
    pictrs_alias: str


class LocalImageView(BaseModel):
    post: Post
    person: Person
    local_image: LocalImage
