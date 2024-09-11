package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	ProductionEnv  = "production"
	DevelopmentEnv = "development"

	DatabaseTimeout    = 5 * time.Second
	ProductCachingTime = 1 * time.Minute
)

var AuthIgnoreMethods = []string{
	"/user.UserService/Login",
	"/user.UserService/Register",
}

type Schema struct {
	HttpPort    int          `yaml:"http_port"`
	Environment string       `yaml:"environment"`
	GrpcPort    int          `yaml:"grpc_port"`
	AuthSecret  string       `yaml:"auth_secret"`
	MongoDBUri  string       `yaml:"mongodb_uri"`
	Aliyun      Aliyun       `yaml:"aliyun"`
	Logger      Logger       `yaml:"logger"`
	Wechat      WechatConfig `yaml:"wechat"`
	Mysql       MysqlConfig  `yaml:"mysql"`
}

type (
	Aliyun struct {
		AccessKeyID     string `yaml:"access_key_id"`
		AccessKeySecret string `yaml:"access_key_secret"`
		Endpoint        string `yaml:"endpoint"`
		BucketName      string `yaml:"bucket_name"`
	}

	Logger struct {
		Level      string `yaml:"level"`
		LogPath    string `yaml:"log_path"`
		MaxSize    int    `yaml:"max_size"`
		MaxAge     int    `yaml:"max_age"`
		MaxBackups int    `yaml:"max_backups"`
	}

	MysqlConfig struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}

	WechatConfig struct {
		AppID     string `yaml:"app_id"`
		AppSecret string `yaml:"app_secret"`
	}
)

var (
	cfg  Schema
	once sync.Once
)

func LoadConfig() *Schema {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	yamlFile, err := os.ReadFile(filepath.Join(currentDir, "config.yaml"))
	if err != nil {
		log.Printf("Error on reading configuration file, error: %v", err)
	}

	var cfg Schema
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("Error on parsing configuration file, error: %v", err)
	}

	return &cfg
}

func GetConfig() *Schema {
	once.Do(func() {
		cfg = *LoadConfig()
	})
	return &cfg
}
