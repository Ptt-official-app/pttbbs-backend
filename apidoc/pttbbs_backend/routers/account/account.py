from fastapi import APIRouter

from .types import (
    AttemptLoginParams,
    AttemptRegisterUserParams,
    LoginParams,
    LoginResult,
    RefreshParams,
)

router = APIRouter(
    tags=['account']
)


@router.post('/account/attemptregister')
def attempt_register_user(body: AttemptRegisterUserParams) -> None:
    '''
    attempt registering user with email.

    sending register url to the email address.
    '''


@router.get('/account/register')
def register_user(token: str):
    '''
    register user based on token info.
    '''


@router.post('/account/attemptlogin')
def attempt_login(body: AttemptLoginParams):
    '''
    attempt login

    generate verification code and send to email address.
    '''


@router.post('/account/login')
def login(body: LoginParams) -> LoginResult:
    ...


@router.post('/account/logout')
def logout():
    ...


@router.post('/account/refresh')
def refresh(body: RefreshParams) -> LoginResult:
    '''
    token refresh
    '''
    ...
