package config

import (
	"github.com/spf13/viper"
	"log"
)

var c *Conf

func Instance() *Conf {
	if c == nil {
		InitConfig("./config.toml")
	}
	return c
}

type Conf struct {
	DB     DBConf
	Redis  RedisConf
	App    AppConf
	Store  StoreConf
	ZapLog ZapLogConf
	Jwt    JwtConf
	Crypt  CryptConf
}

type DBConf struct {
	DBHost string
	DBPort string
	DBUser string
	DBPwd  string
	DBName string
}

type RedisConf struct {
	RedisAddr string
	RedisPWD  string
	RedisDB   int
}

type AppConf struct {
	HttpPort    string
	RunMode     string
	CacheMode   string
	QueueType   string
	CaptchaMode string
	PageSize    int
}

type StoreConf struct {
	StoreType    string
	EndPoint     string
	AccessKey    string
	AccessSecret string
	BucketName   string
	ShowPrefix   string
}

type ZapLogConf struct {
	Director string
	SaveMode string
}

type JwtConf struct {
	Secret string
	Ttl    int
}

type CryptConf struct {
	CryptMode  int
	ChipMode   int
	PrivateKey string
	PublicKey  string
	SM4Key     string
	SM4Iv      string
}

func InitConfig(tomlPath ...string) {
	if len(tomlPath) > 1 {
		log.Fatal("配置路径数量不正确")
	}

	v := viper.New()
	v.SetConfigFile(tomlPath[0])
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("配置文件读取失败: ", err.Error())
	}
	err = v.Unmarshal(&c)
	if err != nil {
		log.Fatal("配置解析失败:", err.Error())
	}
}
