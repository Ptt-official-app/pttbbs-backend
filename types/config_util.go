package types

import (
	"github.com/Ptt-official-app/pttbbs-backend/configutil"
)

const configPrefix = "pttbbs-backend:types"

func InitConfig() (err error) {
	config()

	err = postConfig()
	if err != nil {
		return err
	}

	return nil
}

func setStringConfig(idx string, orig string) string {
	return configutil.SetStringConfig(configPrefix, idx, orig)
}

func setListStringConfig(idx string, orig []string) []string {
	return configutil.SetListStringConfig(configPrefix, idx, orig)
}

func setBytesConfig(idx string, orig []byte) []byte {
	return configutil.SetBytesConfig(configPrefix, idx, orig)
}

func setBoolConfig(idx string, orig bool) bool {
	return configutil.SetBoolConfig(configPrefix, idx, orig)
}

func setIntConfig(idx string, orig int) int {
	return configutil.SetIntConfig(configPrefix, idx, orig)
}

func setInt64Config(idx string, orig int64) int64 {
	return configutil.SetInt64Config(configPrefix, idx, orig)
}
