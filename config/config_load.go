package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"playground/implementation-rabbitMQ-golang/helpers"
)

// Config merepresentasikan credential yang konfigurasi
type Config struct {
	Endpoint  string `yaml:"ENDPOINT"`
	Port      string `yaml:"PORT"`
	AccessKey string `yaml:"ACCESS_KEY"`
	SecretKey string `yaml:"SECRET_KEY"`
}

// Membaca file JSON konfigurasi credential
func ReadConfigFromFile() (*Config, error) {
	var credentials Config

	dir := helpers.GetCurrentDirectory()
	filePath := filepath.Join(dir, "config/config.yml")

	// Membaca file YAML
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("gagal membaca file YAML: %s", err)
	}

	// Parsing data YAML
	err = yaml.Unmarshal(yamlFile, &credentials)
	if err != nil {
		log.Fatalf("gagal membaca data file YAML: %s", err)
	}

	return &credentials, nil
}
