package config

import (
	"net/url"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

type Config struct {
	DebugLog         bool          `mapstructure:"debug_log"`
	WebsocketTimeout time.Duration `mapstructure:"websocket_timeout,required"`
	WssRpcUrl        string        `mapstructure:"wss_rpc_url,required"`
	ContractAddress  string        `mapstructure:"contract_address,required"`
}

func GetConfig(fs afero.Fs) (Config, error) {
	config := Config{}
	viper.SetEnvPrefix("VIPER")
	viper.MustBindEnv("CONFIG")
	configPath := viper.GetString("CONFIG")
	viper.SetConfigFile(configPath)

	viper.SetFs(fs)

	err := viper.ReadInConfig()
	if err != nil {
		return config, errors.Wrap(err, "failed to read config")
	}
	err = viper.Unmarshal(&config, viper.DecodeHook(
		mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			AddressHookFunc(),
		),
	))
	if err != nil {
		return config, errors.Wrap(err, "failed to unmarshal config")
	}
	err = validateConfig(config)
	if err != nil {
		return config, errors.Wrap(err, "invalid config")
	}
	return config, err
}

func validateConfig(config Config) error {
	if config.WebsocketTimeout < time.Second {
		return errors.New("WebsocketTimeout is less than a second, unable to run")
	}
	if config.WssRpcUrl == "" {
		return errors.New("wss_rpc_url required, unable to run")
	}
	ok := IsUrl(config.WssRpcUrl, []string{"wss", "ws"})
	if !ok {
		return errors.New("wss_rpc_url should be valid wss url")
	}

	if config.ContractAddress == "" {
		return errors.New("contract_address required, unable to run")
	}

	return nil
}

func IsUrl(str string, schemas []string) bool {
	u, err := url.Parse(str)
	foundSchema := ""
	for _, schema := range schemas {
		if schema == u.Scheme {
			foundSchema = schema
			break
		}
	}
	ok := (err == nil && u.Scheme != "" && u.Host != "" && u.Scheme == foundSchema)

	return ok
}

// https://sagikazarmark.hu/blog/decoding-custom-formats-with-viper/
func AddressHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(common.Address{}) {
			return data, nil
		}
		result := common.HexToAddress(data.(string))

		return result, nil
	}
}
