package types

type ServiceMode string

const (
	// in DEV mode, client_secret is always test_client_secret, log set to INFO
	DEV ServiceMode = "DEV"

	// in PRODUCTION mode, log set to WARN
	STAGING ServiceMode = "STAGING"

	// in PRODUCTION mode, log set to WARN
	PRODUCTION ServiceMode = "PRODUCTION"

	// in INFO mode, log set to INFO
	INFO ServiceMode = "INFO"

	// in DEBUG mode, log set to DEBUG
	DEBUG ServiceMode = "DEBUG"

	// in test mode
	TEST ServiceMode = "TEST"
)
