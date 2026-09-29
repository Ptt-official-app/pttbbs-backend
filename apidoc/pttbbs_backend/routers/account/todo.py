from fastapi import APIRouter

from .types import AttemptChangeEmailParams, AttemptSet6238Params, ChangeEmailParams, Set6238Params

router = APIRouter(
    tags=['TODO']
)


@router.post('/account/attemptset6238')
def attempt_set_6238(body: AttemptSet6238Params):
    '''
    (TODO) attempt setting RFC-6238 TOTP authenticator.
    '''
    ...


@router.post('/account/set6238')
def set_6238(body: Set6238Params):
    '''
    (TODO) set RFC-6238 TOTP authenticator.
    '''
    ...


@router.post('/account/attemptchangeemail')
def attempt_change_email(body: AttemptChangeEmailParams):
    '''
    (TODO) attempt changing email

    1. generate verification code and send to original email address.
    2. generate new verification code and send to new email address.
    '''
    ...


@router.post('/account/changeemail')
def change_email(body: ChangeEmailParams):
    '''
    (TODO) change email.
    '''
    ...
