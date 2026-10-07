from pydantic import BaseModel

from ...types import ID, DbURL, Time3339


class Person(BaseModel):
    comment_count: int
    post_count: int
    instance_id: ID
    bot_account: bool
    matrix_user_id: str = ''
    deleted: bool
    banner: DbURL = ''
    last_refreshed_at: Time3339
    local: bool
    bio: str = ''
    ap_id: DbURL
    updated_at: Time3339 = ''
    published_at: Time3339
    avatar: DbURL = ''
    display_name: str = ''
    name: str
    id: ID
