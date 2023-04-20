package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path"
	"strings"
)

const (
	keyEnvPrefix = "env_prefix" // keyEnvPrefix is the environment variable prefix
	keyEnv       = "env"        // keyEnv is the environment variable key used to specify the application's environment (e.g. dev, test, staging, prod)
	keyConfig    = "config"     // keyConfig key used to specify the path of the configuration file
)

// Config wraps around viper.Viper
type Config struct {
	*viper.Viper
}

func (c *Config) Save(key string, val interface{}) error {
	c.Viper.Set(key, val)
	return c.WriteConfig()
}

func (c *Config) WriteConfig() error {
	err := c.Viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("failed to write changes to config file: %w", err)
	}
	return err
}

var cfg *Config

// Load loads configuration file and return *Config.
// First, it will try load .env file, either from NHM_ENV env var or from default '.env' file.
// Next, it will try to load the config file from NHM_CONFIG env var (if set), or from default "config.json"
// Note that Load() will panic if it can't load the config file.
// We assume that
// - config is loaded at least once during start up
// - if there is a misconfiguration, we panic by default because no point continuing?
func Load() *Config {
	if cfg != nil {
		return cfg
	}
	return load()
}

// Reload forces refresh of configuration
func Reload() *Config {
	return load()
}

// Watch watches for changes in the config file
func Watch() {
	// Watch config file for changes
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
}

func load() *Config {
	// Prepare environment variable configuration
	// Need to set AutomaticEnv() first so that we can get the env prefix properly
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix(strings.ToLower(viper.GetString(keyEnvPrefix)))
	viper.SetTypeByDefaultValue(true)

	// Register key aliases
	viper.RegisterAlias("environment", "env")

	// Set defaults'
	viper.SetDefault(keyEnv, "")
	viper.SetDefault(keyConfig, "")

	// Load config file
	viper.SetConfigFile(GetConfigPath())
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file, ensure that it exists: %w", err))
	}

	cfg = &Config{
		Viper: viper.GetViper(),
	}
	return cfg
}

// GetConfigPath returns the path of the configuration file
// If NHM_CONFIG environment variable is set, the path will be taken from that.
// Otherwise, the path is expected to be "config.{NHM_ENV}.json", default: "config.dev.json"
func GetConfigPath() string {
	rootPath, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("failed to get current working directory: %w", err))
	}
	// setting the config path via environment variable
	p := viper.GetString(keyConfig)
	if p != "" {
		// absolute path
		if strings.HasPrefix(p, "/") {
			return p
		}
		// relative path
		return path.Join(rootPath, p)
	}
	// setting the config path via convention
	env := viper.GetString(keyEnv)
	if env == "" {
		env = "dev"
	}
	p = "./config." + env + ".json"
	return path.Join(rootPath, p)
}
