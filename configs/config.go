package configs

import "github.com/spf13/viper"

var cfg *configs

type configs struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

/*
função que faz parte do cliclo de vida do go

sempre é chamada no start das aplicações
*/
func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5433")

}

/*
função que carrega as configurações
*/
func Load() error {
	viper.SetConfigName("config") // nome do arquivo que ele vai procurar
	viper.SetConfigType("toml")   // tipo do arquivo que ele vai procurar
	viper.AddConfigPath(".")      // local onde ele vai procurar

	//lendo o arqivo
	err := viper.ReadInConfig()
	if err != nil {
		//validando o tipo de erro nao é do tipo ConfigFileNotFoundError
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	//Criando um ponteiro da nossa struct é a mesma coisa que : $configs{}
	cfg = new(configs)

	//Atribuindo os dados lidos na minha struct
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	//Return nil porque nao teve nenhum erro
	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetAPI() string {
	return cfg.API.Port
}
