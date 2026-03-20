package configutil

import (
	"strings"
	"time"

	"github.com/go-viper/encoding/ini"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	pttbbsconfigutil "github.com/Ptt-official-app/go-pttbbs/configutil"
)

var myViper *viper.Viper

func InitViper(filename string) (err error) {
	err = pttbbsconfigutil.InitViper(filename)
	if err != nil {
		return err
	}

	codecRegistry := viper.NewCodecRegistry()
	_ = codecRegistry.RegisterCodec("ini", ini.Codec{})

	myViper = viper.NewWithOptions(
		viper.WithCodecRegistry(codecRegistry),
	)

	filenameList := strings.Split(filename, ".")
	if len(filenameList) == 1 {
		return ErrInvalidIni
	}

	filenamePrefix := strings.Join(filenameList[:len(filenameList)-1], ".")
	filenamePostfix := filenameList[len(filenameList)-1]

	myViper.SetConfigName(filenamePrefix)
	myViper.SetConfigType(filenamePostfix)
	myViper.AddConfigPath("/etc/go-pttbbs")
	myViper.AddConfigPath(".")
	err = myViper.ReadInConfig()
	if err != nil {
		return err
	}

	logrus.Infof("viper keys: %v", myViper.AllKeys())

	return nil
}

func SetStringConfig(configPrefix string, idx string, orig string) string {
	if myViper == nil {
		myViper = viper.GetViper()
	}

	idx = configPrefix + "." + strings.ToLower(idx)
	if !myViper.IsSet(idx) {
		return orig
	}

	return myViper.GetString(idx)
}

func SetListStringConfig(configPrefix string, idx string, orig []string) []string {
	if myViper == nil {
		myViper = viper.GetViper()
	}

	idx = configPrefix + "." + strings.ToLower(idx)
	if !myViper.IsSet(idx) {
		return orig
	}

	return strings.Split(myViper.GetString(idx), ",")
}

func SetBytesConfig(configPrefix string, idx string, orig []byte) []byte {
	if myViper == nil {
		myViper = viper.GetViper()
	}

	idx = configPrefix + "." + strings.ToLower(idx)
	if !myViper.IsSet(idx) {
		return orig
	}

	return []byte(myViper.GetString(idx))
}

func SetBoolConfig(configPrefix string, idx string, orig bool) bool {
	if myViper == nil {
		myViper = viper.GetViper()
	}

	idx = configPrefix + "." + strings.ToLower(idx)
	if !myViper.IsSet(idx) {
		return orig
	}

	return myViper.GetBool(idx)
}

func SetDurationConfig(configPrefix string, idx string, orig time.Duration) time.Duration {
	if myViper == nil {
		myViper = viper.GetViper()
	}

	idx = configPrefix + "." + strings.ToLower(idx)
	if !myViper.IsSet(idx) {
		return orig
	}
	return time.Duration(myViper.GetInt(idx))
}

func SetIntConfig(configPrefix string, idx string, orig int) int {
	if myViper == nil {
		myViper = viper.GetViper()
	}

	idx = configPrefix + "." + strings.ToLower(idx)
	if !myViper.IsSet(idx) {
		return orig
	}
	return myViper.GetInt(idx)
}

func SetInt64Config(configPrefix string, idx string, orig int64) int64 {
	if myViper == nil {
		myViper = viper.GetViper()
	}

	idx = configPrefix + "." + strings.ToLower(idx)
	if !myViper.IsSet(idx) {
		return orig
	}
	return myViper.GetInt64(idx)
}

func SetDoubleConfig(configPrefix string, idx string, orig float64) float64 {
	idx = configPrefix + "." + strings.ToLower(idx)
	if !myViper.IsSet(idx) {
		return orig
	}

	return myViper.GetFloat64(idx)
}
