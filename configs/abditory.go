package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/ztrue/tracerr"
	"log"
	"reflect"
)

type Secrets struct {
	DbUser    string `mapstructure:"dbUser"`
	DbPass    string `mapstructure:"dbPass"`
	SrvJwtKey string `mapstructure:"serverJwtKey"`
}

type Env struct {
	DbMigrationLocation string `env:"FLAGR_MIGRATION_LOCATION"`
	DbSchema            string `env:"FLAGR_DB_SCHEMA"`
	DbHost              string `env:"FLAGR_DB_HOST"`
	DbPort              string `env:"FLAGR_DB_PORT"`
	SrvPort             string `env:"FLAGR_SERVER_PORT"`
	SrvSecretsFile      string `env:"FLAGR_SECRETS_FILE"`
	Secrets
}

func Configure() (*Env, error) {
	vipe := viper.GetViper()

	env := Env{}
	if err := env.RegisterEnvironmentVariables(vipe); err != nil {
		return nil, tracerr.Wrap(err)
	}
	if err := env.LoadSecretsFile(vipe); err != nil {
		return nil, tracerr.Wrap(err)
	}
	return &env, nil
}

func (env *Env) RegisterEnvironmentVariables(vipe *viper.Viper) error {
	vipe.AutomaticEnv()

	//var newEnvConfiguration Env
	t := reflect.TypeOf(*env)

	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		if field.Name == "Secrets" {
			continue
		}

		// Get the field tag value
		tag := field.Tag.Get("env")

		if tag == "" {
			continue
		}
		v := reflect.ValueOf(env).Elem().FieldByName(field.Name)
		if v.IsValid() {
			tagValue := vipe.GetString(tag)
			v.Set(reflect.ValueOf(tagValue))
		}
	}
	return nil
}

func (env *Env) LoadSecretsFile(vipe *viper.Viper) error {
	vipe.SetConfigFile(env.SrvSecretsFile)

	err := vipe.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}
	err = vipe.Unmarshal(&env.Secrets)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	fmt.Printf("%+v", env)

	return nil
}
