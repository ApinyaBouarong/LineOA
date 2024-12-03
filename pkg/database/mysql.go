package database

import (
	"LineOA/internal/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Agent struct {
		Authtoken string `yaml:"authtoken"`
	} `yaml:"agent"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("ไม่สามารถอ่านไฟล์ config.yaml: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("ไม่สามารถแปลงไฟล์ YAML: %v", err)
	}

	return &config, nil
}

func ConnectToDB() (*sql.DB, error) {
	var config models.Config
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Error reading config file:", err)
		return nil, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Error parsing config file:", err)
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Name,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
