package config

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/spf13/afero"
)

func TestGetConfig(t *testing.T) {
	fs := getConfigsTestFs()
	tests := []struct {
		name       string
		configPath string
		want       Config
		wantErr    bool
	}{
		{"MissingConfig", "./doesntExist.yaml", Config{}, true},
		{"EmptyConfig", "./emptyConfig.yaml", Config{}, true},
		{"InvalidFieldValue", "./invalidValue.yaml", Config{}, true},
		{"BadCharacterInKey", "./invalidCharacterKey.yaml", Config{}, true},
		{"BadCharacterInValue", "./invalidCharacterValue.yaml", Config{}, true},
		{"Valid config", "./validConfig.yaml", Config{
			DebugLog:         true,
			WssRpcUrl:        "wss://rpc.com",
			WebsocketTimeout: time.Second,
			ContractAddress:  "0x000232323",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("VIPER_CONFIG", tt.configPath)
			got, err := GetConfig(fs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateConfig(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Empty config", args{Config{}}, true},
		{"WebsocketTimeout is less than a second, unable to run", args{Config{
			WebsocketTimeout: time.Nanosecond,
		}}, true},
		{"wss_rpc_url required, unable to run", args{Config{
			WebsocketTimeout: time.Minute,
		}}, true},
		{"wss_rpc_url should be valid url", args{Config{
			WebsocketTimeout: time.Minute,
			WssRpcUrl:        "ss://rpc.com",
		}}, true},
		{"subgraph_url required, unable to run", args{Config{
			WebsocketTimeout: time.Minute,
			WssRpcUrl:        "wss://rpc.com",
		}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateConfig(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func getConfigsTestFs() afero.Fs {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "emptyConfig.yaml", []byte(""), 0644)
	afero.WriteFile(fs, "invalidValue.yaml", []byte(`debug_log: "test"`), 0644)
	afero.WriteFile(fs, "invalidCharacterKey.yaml", []byte(`%infura_url: "https://---.infura.io/v3/---"`), 0644)
	afero.WriteFile(fs, "invalidCharacterValue.yaml", []byte(`infura_url: %"https://---.infura.io/v3/---"`), 0644)
	afero.WriteFile(fs, "validConfig.yaml", []byte(`
debug_log: true
websocket_timeout: 1s
wss_rpc_url: "wss://rpc.com"
`), 0644)
	return fs
}
