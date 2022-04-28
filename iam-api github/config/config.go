package config

import (
	"log"
	"os"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var config *viper.Viper
var logr = logrus.New()

func init() {
	initConfig()
}
func initConfig() {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("config")
	config.AddConfigPath(".")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err)
		os.Exit(0)
	}
	logr.Out = os.Stdout

	file, err := os.OpenFile("iam-api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logr.Out = file
	} else {
		logr.Info("Failed to log to file, using default stderr")
	}

}

func GetConfig() *viper.Viper {
	return config
}

var SkyTracer *go2sky.Tracer

func initSkyWalking() {
	var err error
	r, err := reporter.NewGRPCReporter(GetConfig().GetString("sky_walking.host"))
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	defer r.Close()
	SkyTracer, err = go2sky.NewTracer(GetConfig().GetString("sky_walking.service_name"), go2sky.WithReporter(r))
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
}
