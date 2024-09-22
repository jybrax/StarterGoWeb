package libs

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Config struct {
	Username     string `json:"Username"`
	Password     string `json:"Password"`
	Host         string `json:"Host"`
	DataBaseName string `json:"DataBaseName"`
	SslMode      string `json:"SslMode"`
}

func LoadConfig() Config {
	file, err := os.Open("sqlConfig.json")
	if err != nil {
		return Config{}
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}
	}

	return config
}

// Méthode connexion à la base MySQL
func ConnectMysql() (*sql.DB, error) {
	config := LoadConfig()
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Username, config.Password, config.Host, config.DataBaseName)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping MySQL: %v", err)
	}

	return db, nil
}

// Méthode connexion à la base PostgreSQL
func ConnectPostgreSql() (*sql.DB, error) {
	config := LoadConfig()
	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s", config.Username, config.Password, config.Host, config.DataBaseName, config.SslMode)

	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		DB.Close()
		return nil, fmt.Errorf("failed to ping PostgreSQL: %v", err)
	}

	return DB, nil
}
