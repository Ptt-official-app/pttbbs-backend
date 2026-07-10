package zk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/sirupsen/logrus"
)

// XXX The reason why we need to have separated definition of types
//     is because go-zkid-verifier requires additional lib to compile.

// from: https://github.com/privacy-ethereum/go-zkid-verifier/blob/main/linkverify/linkverify.go
type PublicSignals = VerifierPublicSignals

// SmtRootOutcome is the outcome of the proof-vs-trusted-root comparison.
type SmtRootOutcome struct {
	Issuer      SmtRootIssuerID `json:"-"`
	IssuerName  string          `json:"issuer"`
	Match       bool            `json:"match"`
	Expected    string          `json:"expected"`
	Observed    string          `json:"observed"`
	TrustSource string          `json:"trust_source,omitempty"`
	TrustedAt   time.Time       `json:"trusted_at,omitempty"`
}

type IssuerModulusOutcome struct {
	Issuer         SmtRootIssuerID `json:"-"`
	IssuerName     string          `json:"issuer"`
	Match          bool            `json:"match"`
	ExpectedSHA256 string          `json:"expected_sha256"`
	TrustSource    string          `json:"trust_source,omitempty"`
	TrustedAt      time.Time       `json:"trusted_at,omitempty"`
}

type AppIDOutcome struct {
	Match    bool   `json:"match"`
	Expected string `json:"expected"`
	Observed string `json:"observed"`
}

// ChallengeOutcome reports whether the per-session challenge in the device-sig
// proof matches the value the verifier issued for this challenge.
type ChallengeOutcome struct {
	Match    bool   `json:"match"`
	Expected string `json:"expected"`
	Observed string `json:"observed"`
}

// from: https://github.com/privacy-ethereum/go-zkid-verifier/blob/main/verifier/verifier.go
type VerifierPublicSignals struct {
	CertChain []string `json:"cert_chain"`
	UserSig   []string `json:"user_sig"`
}

// from: https://github.com/privacy-ethereum/go-zkid-verifier/blob/main/verifier/public_inputs.go
type VerifierParsedInputs struct {
	PkCommit         string   `json:"pk_commit"`
	Nullifier        string   `json:"nullifier"`
	AppID            string   `json:"app_id"`
	AppIDPacked      string   `json:"app_id_packed"`
	Challenge        string   `json:"challenge"`
	IssuerRSAModulus []string `json:"issuer_rsa_modulus"`
	SmtRoot          string   `json:"smt_root"`
}

// from: https://github.com/privacy-ethereum/go-zkid-verifier/blob/main/httpapi/dto.go
type VerifyResponse struct {
	// VerifySuccessResponse
	Verified      bool                  `json:"verified"`
	Nullifier     string                `json:"nullifier,omitempty"`
	IDVerified    bool                  `json:"id_verified,omitempty"`
	Persisted     bool                  `json:"persisted,omitempty"`
	PublicSignals *PublicSignals        `json:"public_signals,omitempty"`
	ParsedInputs  *VerifierParsedInputs `json:"parsed_inputs,omitempty"`
	SmtRoot       *SmtRootOutcome       `json:"smt_root,omitempty"`
	IssuerModulus *IssuerModulusOutcome `json:"issuer_modulus,omitempty"`
	AppID         *AppIDOutcome         `json:"app_id,omitempty"`
	Challenge     *ChallengeOutcome     `json:"challenge,omitempty"`

	// VerifyFailResponse
	Reason string `json:"reason,omitempty"`
}

// from: https://github.com/privacy-ethereum/go-zkid-verifier/blob/main/smtroot/smtroot.go
type SmtRootIssuerID int

func NewZKProxy() *httputil.ReverseProxy {
	theURL, _ := url.Parse(types.ZK_PREFIX)
	logrus.Infof("NewZKProxy: ZK_PREFIX: %v theURL: %v", types.ZK_PREFIX, theURL)

	proxy := httputil.NewSingleHostReverseProxy(theURL)

	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Del("Access-Control-Allow-Origin")
		resp.Header.Del("Access-Control-Allow-Methods")
		resp.Header.Del("Access-Control-Allow-Headers")
		resp.Header.Del("Access-Control-Allow-Credentials")
		return nil
	}

	return proxy
}

func NewZKLinkVerifyProxy() *httputil.ReverseProxy {
	url, _ := url.Parse(types.ZK_PREFIX)

	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Del("Access-Control-Allow-Origin")
		resp.Header.Del("Access-Control-Allow-Methods")
		resp.Header.Del("Access-Control-Allow-Headers")
		resp.Header.Del("Access-Control-Allow-Credentials")

		// Only log successful or specific status codes if desired
		if resp.StatusCode != http.StatusOK {
			return nil
		}

		// 1. Read the body bytes
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			logrus.Errorf("ZKLinkVerifyProxy: Failed to read ZK verifier response body: %v", err)
			return err
		}
		resp.Body.Close() // Close the original body
		defer func() {
			// 4. CRITICAL: Restore the resp.Body so Gin can send it to the client
			resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}()

		// 2. Record the result (Log it, save to DB, or pass to a channel)
		var verifyResponse *VerifyResponse
		err = json.Unmarshal(bodyBytes, &verifyResponse)
		if err != nil {
			logrus.Errorf("ZKLinkVerifyProxy: unable to unmarshal: bodyBytes: %v e: %v", bodyBytes, err)
			return err
		}
		if !verifyResponse.Verified {
			return nil
		}

		// 3. schema update UserIsGovernmentID
		userID, ok := resp.Request.Context().Value(types.ZK_USER_ID_KEY).(bbs.UUserID)
		if !ok {
			return nil
		}

		nowNS := types.NowNanoTS()
		err = schema.UpdateUserIsGovernmentID(userID, verifyResponse.Verified, nowNS)
		if err != nil {
			return err
		}

		return nil
	}
	return proxy
}
