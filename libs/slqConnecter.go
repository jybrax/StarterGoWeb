package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type SqlData struct {
	Username     string
	Password     string
	Host         string
	DataBaseName string
	SslMode      string
}

// Interface pour la connexion
type SqlConnecter interface {
	ConnectMysql() (*sql.DB, error)
	ConnectPostgreSql() (*sql.DB, error)
	ExecuteQuery(query string, db *sql.DB) (*sql.Rows, error)
}

// Méthode pour se connecter à MySQL
func (s SqlData) ConnectMysql() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", s.Username, s.Password, s.Host, s.DataBaseName)

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

// Méthode pour se connecter à PostgreSQL
func (s SqlData) ConnectPostgreSql() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s", s.Username, s.Password, s.Host, s.DataBaseName, s.SslMode)

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

func (s SqlData) ExecuteQuery(query string, db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	return rows, nil
}
