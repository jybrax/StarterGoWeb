package services

import (
	"fmt"
	"log"
	"wst/libs"
	"wst/models"
)

func AddUserSql(user models.UserModel) error {
	db, err := libs.ConnectMysql()
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO `User`(`UserName`, `Password`) VALUES (?, ?)")
	if err != nil {
		log.Fatal("Erreur lors de la préparation de la requête SQL:", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.UserName, user.Password)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'insertion de la donnée : %v", err)
	}

	return nil
}
