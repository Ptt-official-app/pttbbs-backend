package types

import "testing"

func TestInitConfig(t *testing.T) {
	setupTest()
	defer teardownTest()

	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitConfig(); (err != nil) != tt.wantErr {
				t.Errorf("InitConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServiceModeDefault(t *testing.T) {
	setupTest()
	defer teardownTest()

	// DEV skips CSRF and client-register ip checks, so it must be opt-in.
	if origServiceMode != PRODUCTION {
		t.Errorf("default SERVICE_MODE = %v, want %v", origServiceMode, PRODUCTION)
	}
}
