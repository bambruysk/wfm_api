package repo

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// GetDBConfig read database configuration from flags, environment or config file.  If no one not found return error
func GetDBConfig (configPath string) string {

	viper.AutomaticEnv()
	viper.SetEnvPrefix("DB")
	//	viper.AddConfigPath("./../config")
	if configPath != "" {
		log.Println("Parsing config from file: ", configPath)
		viper.SetConfigFile(configPath)
		err := viper.ReadInConfig()
		if err !=  nil {
			log.Fatalf("Unable read conffig file: %s %v", configPath, err)
		}
	} else {
		log.Println("Config file not specified")
	}

	config := viper.Sub("database")
	if config ==  nil {
		log.Fatalln("Database config not found")
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		config.GetString("host"),
		config.GetString("user"),
		config.GetString("password"),
		config.GetString("dbname"),
		config.GetInt("port"),
	)
}
