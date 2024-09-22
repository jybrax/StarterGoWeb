package services

import (
	"database/sql"
	"errors"
	"log"
	"wst/libs"
	"wst/models"

	"golang.org/x/crypto/bcrypt"
)

func VerifUserSql(user models.UserModel) error {
	// Connexion à la base de données
	db, err := libs.ConnectMysql()
	if err != nil {
		log.Println("Erreur de connexion à la base de données:", err)
		return err
	}
	defer db.Close()

	// Requête pour récupérer l'utilisateur par son nom
	query := "SELECT Password FROM `User` WHERE UserName = ?"

	var storedPassword string

	// Exécuter la requête avec le nom d'utilisateur fourni
	err = db.QueryRow(query, user.UserName).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			// Si aucun utilisateur n'a été trouvé
			return errors.New("Utilisateur non trouvé")
		}
		// Autre erreur lors de l'exécution de la requête
		return err
	}

	// Utiliser bcrypt pour comparer les mots de passe
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		return errors.New("Mot de passe incorrect")
	}

	// Si tout est bon, renvoyer nil (pas d'erreur)
	return nil
}
