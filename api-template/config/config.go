package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceName    string
	GitVersion     string
	BaseURL        string
	Port           string
	RunMode        string
	LogLevel       string
	LogVersion     int
	MediaURL       string
	AllowedOrigins string
	DBHost         string
	DBPort         string
	DBUser         string
	DBName         string
	DBPass         string
	DBSSLMode      string

	IndexFromBlock string
	BscRPCEndPoint string
	ChainId        int

	OperatorWallet     string
	OperatorPrivateKey string
	FactoryContract    string
	ReferralContract   string
	MineContract       string
	SchoolContract     string
	PvEContract        string
	PvEUtilsContract   string
}

// GetCORS in config
func (c *Config) GetCORS() []string {
	cors := strings.Split(c.AllowedOrigins, ";")
	rs := []string{}
	for idx := range cors {
		itm := cors[idx]
		if strings.TrimSpace(itm) != "" {
			rs = append(rs, itm)
		}
	}
	return rs
}

// Loader load config from reader into Viper
type Loader interface {
	Load(viper.Viper) (*viper.Viper, error)
}

// generateConfigFromViper generate config from viper data
func generateConfigFromViper(v *viper.Viper) Config {

	return Config{
		ServiceName: v.GetString("SERVICE_NAME"),
		Port:        v.GetString("PORT"),
		BaseURL:     v.GetString("BASE_URL"),
		RunMode:     v.GetString("RUN_MODE"),
		LogLevel:    v.GetString("LOG_LEVEL"),
		LogVersion:  v.GetInt("LOG_VERSION"),
		GitVersion:  v.GetString("GIT_VERSION"),
		MediaURL:    v.GetString("MEDIA_URL"),

		AllowedOrigins: v.GetString("ALLOWED_ORIGINS"),

		DBHost:    v.GetString("DB_HOST"),
		DBPort:    v.GetString("DB_PORT"),
		DBUser:    v.GetString("DB_USER"),
		DBName:    v.GetString("DB_NAME"),
		DBPass:    v.GetString("DB_PASS"),
		DBSSLMode: v.GetString("DB_SSLMODE"),

		IndexFromBlock:     v.GetString("INDEX_FROM_BLOCK"),
		BscRPCEndPoint:     v.GetString("BSC_RPC_ENDPOINT"),
		ChainId:            v.GetInt("CHAIN_ID"),
		OperatorWallet:     v.GetString("OPERATOR_WALLET"),
		OperatorPrivateKey: v.GetString("OPERATOR_PRIVATE_KEY"),

		FactoryContract:  v.GetString("FACTORY_CONTRACT"),
		ReferralContract: v.GetString("REFERRAL_CONTRACT"),
		MineContract:     v.GetString("MINE_CONTRACT"),
		SchoolContract:   v.GetString("SCHOOL_CONTRACT"),
		PvEContract:      v.GetString("PVE_CONTRACT"),
		PvEUtilsContract: v.GetString("PVE_UTILS_CONTRACT"),
	}
}

// DefaultConfigLoaders is default loader list
func DefaultConfigLoaders() []Loader {
	loaders := []Loader{}
	fileLoader := NewFileLoader(".env", ".")
	loaders = append(loaders, fileLoader)
	loaders = append(loaders, NewENVLoader())

	return loaders
}

// LoadConfig load config from loader list
func LoadConfig(loaders []Loader) Config {
	v := viper.New()
	v.SetDefault("PORT", "8080")
	v.SetDefault("RUN_MODE", "local")

	for idx := range loaders {
		newV, err := loaders[idx].Load(*v)

		if err == nil {
			v = newV
		}
	}
	return generateConfigFromViper(v)
}
