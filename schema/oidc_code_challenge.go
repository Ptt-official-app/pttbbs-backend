package schema

type OIDCCodeChallenge struct {
	Challenge string `json:"challenge"`
	Method    string `json:"method"`
}
