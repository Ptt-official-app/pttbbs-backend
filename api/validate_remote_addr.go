package api

func isValidRemoteAddr(remoteAddr string) bool {
	if IsTest {
		return true
	}
	switch remoteAddr {
	case "":
		return false
	default:
		return true
	}
}

// isLocalRemoteAddr reports whether the request comes from the server itself,
// for operator-only apis.
func isLocalRemoteAddr(remoteAddr string) bool {
	switch remoteAddr {
	case "127.0.0.1", "::1", "localhost":
		return true
	default:
		return false
	}
}
