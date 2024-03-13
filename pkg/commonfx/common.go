package commonfx

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Common struct {
	Viper  *viper.Viper
	Logger zerolog.Logger
}

var Module = fx.Module(
	"commonfx",
	fx.Provide(
		fx.Annotate(
			newCommonfx,
			fx.ResultTags(`name:"commonfx"`),
		),
	),
)

func newCommonfx() *Common {
	return &Common{
		Viper:  newViper(),
		Logger: newZeroLog(),
	}
}

func newViper() *viper.Viper {
	fileConfigName := getFileConfigName()

	viper.AddConfigPath(".")
	viper.SetConfigName(fileConfigName)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil
	}

	return viper.GetViper()
}

func getFileConfigName() string {
	var appEnv = flag.String("env", "", "specify environment for running")
	flag.Parse()
	fileConfigName := "config"
	if len(*appEnv) > 0 {
		fileConfigName = fileConfigName + "-" + *appEnv
	}
	return fileConfigName
}

func GetViper() *viper.Viper {
	return viper.GetViper()
}

func newZeroLog() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	return log.Logger
}
