package services

import (
	"fmt"
	"log"
	"wst/libs"
	"wst/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Génère un mot de passe haché avec bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func AddUserSql(user models.UserModel) error {
	db, err := libs.ConnectMysql()
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}
	defer db.Close()

	// Hacher le mot de passe
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("Erreur lors du hachage du mot de passe : %v", err)
	}

	// Préparer la requête SQL
	stmt, err := db.Prepare("INSERT INTO `User`(`UserName`, `Password`) VALUES (?, ?)")
	if err != nil {
		log.Fatal("Erreur lors de la préparation de la requête SQL:", err)
	}
	defer stmt.Close()

	// Exécuter la requête avec le nom d'utilisateur et le mot de passe haché
	_, err = stmt.Exec(user.UserName, hashedPassword)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'insertion de la donnée : %v", err)
	}

	return nil
}
