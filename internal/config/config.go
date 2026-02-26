package config

import (
	"github.com/spf13/viper"
)

// Config 应用配置结构
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Player   PlayerConfig   `mapstructure:"player"`
	UI       UIConfig       `mapstructure:"ui"`
	Cache    CacheConfig    `mapstructure:"cache"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	URL      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// PlayerConfig 播放器配置
type PlayerConfig struct {
	Volume    int  `mapstructure:"volume"`
	Crossfade bool `mapstructure:"crossfade"`
}

// UIConfig UI配置
type UIConfig struct {
	Theme   string `mapstructure:"theme"`
	Artwork bool   `mapstructure:"artwork"`
}

// CacheConfig 缓存配置
type CacheConfig struct {
	Enabled bool `mapstructure:"enabled"`
	TTL     int  `mapstructure:"ttl"`
}

// NewConfig 创建默认配置
func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			URL:      "http://localhost:4040",
			Username: "admin",
			Password: "",
		},
		Player: PlayerConfig{
			Volume:    80,
			Crossfade: false,
		},
		UI: UIConfig{
			Theme:   "default",
			Artwork: true,
		},
		Cache: CacheConfig{
			Enabled: true,
			TTL:     3600,
		},
	}
}

// Load 加载配置
func Load() (*Config, error) {
	v := viper.New()
	
	// 设置默认值
	v.SetDefault("server.url", "http://localhost:4040")
	v.SetDefault("server.username", "admin")
	v.SetDefault("server.password", "")
	v.SetDefault("player.volume", 80)
	v.SetDefault("player.crossfade", false)
	v.SetDefault("ui.theme", "default")
	v.SetDefault("ui.artwork", true)
	v.SetDefault("cache.enabled", true)
	v.SetDefault("cache.ttl", 3600)
	
	// 设置配置文件路径
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("$HOME/.config/laetibeat")
	
	// 读取环境变量
	v.AutomaticEnv()
	
	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
		// 配置文件不存在，使用默认值
	}
	
	// 解析配置
	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	
	return &config, nil
}

// Save 保存配置
func (c *Config) Save() error {
	v := viper.New()
	
	// 设置配置文件路径
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("$HOME/.config/laetibeat")
	
	// 设置配置值
	v.Set("server.url", c.Server.URL)
	v.Set("server.username", c.Server.Username)
	v.Set("server.password", c.Server.Password)
	v.Set("player.volume", c.Player.Volume)
	v.Set("player.crossfade", c.Player.Crossfade)
	v.Set("ui.theme", c.UI.Theme)
	v.Set("ui.artwork", c.UI.Artwork)
	v.Set("cache.enabled", c.Cache.Enabled)
	v.Set("cache.ttl", c.Cache.TTL)
	
	// 写入配置文件
	err := v.WriteConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件不存在，创建新文件
			err = v.SafeWriteConfig()
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	
	return nil
}
