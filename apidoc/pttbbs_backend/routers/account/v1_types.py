from pydantic import BaseModel, Field


class AttemptRegisterUserParams(BaseModel):
    email: str


class AttemptLoginParams(BaseModel):
    input: str = Field(description='can be username or email')


class LoginParams(BaseModel):
    client_id: str = ''
    client_secret: str = ''

    input: str

    verify_code: str

    oidc_id: str = Field(
        default='',
        description='id of the OIDC auth request to complete, empty for a normal web login')


class LoginResult(BaseModel):
    username: str
    access_token: str
    token_type: str
    refresh_token: str
    access_expire: int = Field(
        description='access expire in unix-timestamp (second)')
    refresh_expire: int = Field(
        description='refresh expire in unix-timestamp (second)')


class RefreshParams(BaseModel):
    client_id: str
    client_secret: str

    refresh_token: str


class AttemptSet6238Params(BaseModel):
    pass


class Set6238Params(BaseModel):
    pass


class AttemptChangeEmailParams(BaseModel):
    email: str
    oidc_id: str


class ChangeEmailParams(BaseModel):
    pass
