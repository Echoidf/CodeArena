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
	HttpPort        int    `yaml:"http_port"`
	Environment     string `yaml:"environment"`
	GrpcPort        int    `yaml:"grpc_port"`
	AuthSecret      string `yaml:"auth_secret"`
	MongoDBUri      string `yaml:"mongodb_uri"`
	XAPIKey         string `yaml:"x_api_key"`
	TranslateServer string `yaml:"translate_server"`
	SMMSToken       string `yaml:"smms_token"`
	Aliyun          Aliyun `yaml:"aliyun"`
	Logger          Logger `yaml:"logger"`
}

type Aliyun struct {
	AccessKeyID     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	Endpoint        string `yaml:"endpoint"`
	BucketName      string `yaml:"bucket_name"`
}

type Logger struct {
	Level      string `yaml:"level"`
	LogPath    string `yaml:"log_path"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}

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
