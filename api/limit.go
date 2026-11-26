package api

// MAX_LIMIT is the upper bound of the `limit` query parameter of list apis.
// It equals DEFAULT_MAX_LIST (the largest page the apis return by default) so that
// one request can never ask for more than a default page, and a client-controlled
// value can neither be used as an allocation size nor turn into an unbounded query.
const MAX_LIMIT = DEFAULT_MAX_LIST

// isValidLimit checks that the client-supplied limit is in [1, MAX_LIMIT].
// Non-positive values would panic on slice indexing (limit < 0) or be passed to
// mongo as "no limit" (limit <= 0).
func isValidLimit(limit int) bool {
	return limit >= 1 && limit <= MAX_LIMIT
}
