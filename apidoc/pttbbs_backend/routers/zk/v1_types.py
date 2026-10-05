from pydantic import BaseModel

from ..types import Time3339


class LinkVerifyRequest(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/httpapi/dto.go#L30](https://github.com/ethereum/go-zkid-verifier/blob/main/httpapi/dto.go#L30)
    '''
    cert_chain_type: str
    cert_chain_proof: bytes
    user_sig_proof: bytes


class PublicSignals(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/verifier/verifier.go#L24](https://github.com/ethereum/go-zkid-verifier/blob/main/verifier/verifier.go#L24)
    '''
    cert_chain: list[str]
    user_sig: list[str]


class ParsedInputs(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/verifier/public_inputs.go#L134](https://github.com/ethereum/go-zkid-verifier/blob/main/verifier/public_inputs.go#L134)
    '''
    pk_commit: str
    nullifier: str
    app_id: str
    app_id_packed: str
    challenge: str
    issuer_rsa_modulus: list[str]
    smt_root: str


class SmtRootOutcome(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/linkverify/verifier.go#L34](https://github.com/ethereum/go-zkid-verifier/blob/main/linkverify/verifier.go#L34)
    '''
    issuer: str
    match: bool
    expected: str
    observed: str
    trust_source: str = ''
    trusted_at: Time3339 = ''


class IssuerModulusOutcome(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/linkverify/verifier.go#L44](https://github.com/ethereum/go-zkid-verifier/blob/main/linkverify/verifier.go#L44)
    '''
    issuer: str
    match: bool
    expected_sha256: str
    trust_source: str = ''
    trusted_at: Time3339 = ''


class AppIDOutcome(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/linkverify/verifier.go#L53](https://github.com/ethereum/go-zkid-verifier/blob/main/linkverify/verifier.go#L53)
    '''
    match: bool
    expected: str
    observed: str


class ChallengeOutcome(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/linkverify/verifier.go#L61](https://github.com/ethereum/go-zkid-verifier/blob/main/linkverify/verifier.go#L61)
    '''
    match: bool
    expected: str
    observed: str


class LinkVerifySuccessResponse(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/httpapi/dto.go#L8](https://github.com/ethereum/go-zkid-verifier/blob/main/httpapi/dto.go#L8)
    '''
    verified: bool
    nullifier: str
    id_verified: bool = False
    persisted: bool = False
    public_signals: PublicSignals | None = None
    parsed_inputs: ParsedInputs | None = None
    smt_root: SmtRootOutcome | None = None
    issuer_modulus: IssuerModulusOutcome | None = None
    app_id: AppIDOutcome | None = None
    challenge: ChallengeOutcome | None = None


class LinkVerifyFailResponse(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/httpapi/dto.go#L21](https://github.com/ethereum/go-zkid-verifier/blob/main/httpapi/dto.go#L21)
    '''
    verified: bool
    reason: str = ''
    smt_root: SmtRootOutcome | None = None
    issuer_modulus: IssuerModulusOutcome | None = None
    app_id: AppIDOutcome | None = None
    challenge: ChallengeOutcome | None = None


type LinkVerifyResponse = LinkVerifySuccessResponse | LinkVerifyFailResponse


class CertStatus(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/issuercert/issuercert.go#L49](https://github.com/ethereum/go-zkid-verifier/blob/main/issuercert/issuercert.go#L49)
    '''
    issuer: str
    sha256: str
    subject: str
    not_before: Time3339
    not_after: Time3339
    source: str
    fetched_at: Time3339
    modulus_bits: int


class AttemptStat(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/smtroot/smtroot.go#L98](https://github.com/ethereum/go-zkid-verifier/blob/main/smtroot/smtroot.go#L98)
    '''
    count: int
    last_latency_ms: int
    last_err: str = ''


class IssuerCertStatus(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/issuercert/issuercert.go#L60](https://github.com/ethereum/go-zkid-verifier/blob/main/issuercert/issuercert.go#L60)
    '''
    certs: dict[str, CertStatus]
    updated_at: Time3339
    last_attempt_at: Time3339
    last_error: str = ''
    consecutive_fail: int
    cache_age_seconds: int
    attempts: dict[str, AttemptStat]


class IssuerCertStatusResponse(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/httpapi/issuercert.go#L10](https://github.com/ethereum/go-zkid-verifier/blob/main/httpapi/issuercert.go#L10)
    '''
    enforced: bool
    status: IssuerCertStatus


class SMTRootStatus(BaseModel):
    '''
    https://github.com/ethereum/go-zkid-verifier/blob/main/smtroot/smtroot.go#L105
    '''
    source_used: str
    roots: dict[str, str]
    updated_at: Time3339
    last_attempt_at: Time3339
    last_error: str = ''
    consecutive_fail: int
    cache_age_seconds: int
    attempts: dict[str, AttemptStat]


class SMTRootStatusResponse(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/httpapi/smtroot.go#L10](https://github.com/ethereum/go-zkid-verifier/blob/main/httpapi/smtroot.go#L10)
    '''
    enforced: bool
    status: SMTRootStatus


class ChallengeResponse(BaseModel):
    '''
    [https://github.com/ethereum/go-zkid-verifier/blob/main/httpapi/challenge.go#L12](https://github.com/ethereum/go-zkid-verifier/blob/main/httpapi/challenge.go#L12)
    '''
    challenge: str
    app_id: str
    expires_at: Time3339
