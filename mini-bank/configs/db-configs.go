package configs

import (
	"github.com/spf13/viper"
)

var Dbconfigs *dbconfig

type dbconfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
}

func init() {
	loadConfigFile()
	cfg := new(dbconfig)
	cfg.Host = viper.GetString("database.host")
	cfg.Port = viper.GetString("database.port")
	cfg.User = viper.GetString("database.user")
	cfg.Password = viper.GetString("database.password")
	cfg.Db = viper.GetString("database.db")

}

func loadConfigFile() {
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if nok := checkErrorViper(err); nok != nil {
		panic("Erro ao carregar configuracaoes de banco de dados")
	}
}

func checkErrorViper(err error) error {
	if err != nil {
		if nok := checkErrorFileNotFound(err); nok != nil {
			return nok
		}
	}
	return nil
}

func checkErrorFileNotFound(err error) error {
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		return err
	}
	return nil
}
