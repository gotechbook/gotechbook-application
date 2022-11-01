package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	logs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v2/logger/interfaces"
	ll "github.com/topfreegames/pitaya/v2/logger/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"os"
	"time"
)

var (
	GOTECHBOOK_GATE      GateConfig
	GOTECHBOOK_CONFIGURE ConfigureConfig
	GOTECHBOOK_AUTH      AuthConfig
	GOTECHBOOK_REDIS     *redis.Client
	GOTECHBOOK_MONGO     *mongo.Client
)

func Viper(path string, m interface{}) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(m); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(m); err != nil {
		fmt.Println(err)
	}
	return v
}

func LoadConfig(path string, m interface{}) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(m); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(m); err != nil {
		fmt.Println(err)
	}
}

func SetLogger(path string, level string, source string) interfaces.Logger {
	writer, _ := logs.New(
		path+".%Y-%m-%d-%H-%M",
		logs.WithLinkName(path),
		logs.WithMaxAge(time.Duration(60)*time.Second*60*24*30),
		logs.WithRotationTime(time.Duration(60)*time.Second*60),
	)
	gLog := logrus.New()
	gLog.Formatter = new(logrus.TextFormatter)
	lv, _ := logrus.ParseLevel(level)
	gLog.Level = lv
	log := gLog.WithFields(logrus.Fields{
		"source": source,
	})
	writers := []io.Writer{
		writer,
		os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	gLog.SetOutput(fileAndStdoutWriter)
	return ll.NewWithFieldLogger(log)
}
