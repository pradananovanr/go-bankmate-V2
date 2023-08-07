/*
 * Author : Pradana Novan Rianto (https://github.com/pradananovanr)
 * Created on : Thu Apr 13 2023
 * Copyright : Pradana Novan Rianto © 2023. All rights reserved
 */

package config

import "go-bankmate/util"

type APIConfig struct {
	ApiPort string
}

type DBConfig struct {
	Host, Port, User, Password, Name, SSLMode string
}

type AppConfig struct {
	APIConfig
	DBConfig
}

func (conf *AppConfig) readConfiguration() {
	api := util.DotEnv("SERVER_PORT")
	conf.DBConfig = DBConfig{
		Host:     util.DotEnv("DB_HOST"),
		Port:     util.DotEnv("DB_PORT"),
		User:     util.DotEnv("DB_USER"),
		Password: util.DotEnv("DB_PASSWORD"),
		Name:     util.DotEnv("DB_NAME"),
		SSLMode:  util.DotEnv("SSL_MODE"),
	}

	conf.APIConfig = APIConfig{ApiPort: api}
}

func NewConfiguration() AppConfig {
	config := AppConfig{}
	config.readConfiguration()

	return config
}
