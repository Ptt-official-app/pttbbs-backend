from fastapi import APIRouter

from .types import (
    ChallengeResponse,
    IssuerCertStatusResponse,
    LinkVerifyRequest,
    LinkVerifyResponse,
    SMTRootStatusResponse,
)

router = APIRouter(
    tags=['account', 'zk']
)


@router.post('/challenge')
def zk_create_challenge() -> ChallengeResponse:
    ...


@router.get('/challenge/{challenge}')
def zk_get_challenge(challenge: str) -> ChallengeResponse:
    ...


@router.get('/smt-root/status')
def zk_smt_root_status() -> SMTRootStatusResponse:
    ...


@router.get('/issuer-cert/status')
def zk_issuer_cert_status() -> IssuerCertStatusResponse:
    ...


@router.get('/link-verify')
def zk_link_verify(params: LinkVerifyRequest) -> LinkVerifyResponse:
    ...
