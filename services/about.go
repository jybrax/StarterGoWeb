package services

import (
	"fmt"
	"log"
	"os"
	"wst/libs"
	"wst/models"

	"github.com/joho/godotenv"
)

func AddUserSql(user models.UserModel) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	sqlData := libs.SqlData{
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		Host:         os.Getenv("DB_HOST"),
		DataBaseName: os.Getenv("DB_NAME"),
	}

	db, err := sqlData.ConnectMysql()
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO `User`(`Name`, `UserName`) VALUES (?, ?)")
	if err != nil {
		log.Fatal("Erreur lors de la préparation de la requête SQL:", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.UserName)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'insertion de la donnée : %v", err)
	}

	return nil
}
