# StarterGoWeb

StarterGoWeb est un kit de démarrage léger pour framework web, basé sur Go, conçu pour construire des applications web évolutives. Ce projet offre une architecture modulaire, avec une organisation claire du backend en différentes couches telles que les contrôleurs, modèles, services, et middlewares.

## Fonctionnalités

- Structure de code modulaire
- Configuration SQL incluse
- Intégration HTML/CSS/JS pour le frontend
- Routage RESTful

## Installation

1. Clonez le dépôt :
   ```bash
   git clone https://github.com/jybrax/StarterGoWeb.git

2. Accédez au répertoire du projet :
    ```bash
    cd StarterGoWeb

3. Installez les dépendances :
    ```bash
    go mod tidy

## Utilisation
1. Lancez le serveur :
    ```bash
    go run server.go

## Structure du Projet

- controllers/ : Gère les requêtes et réponses HTTP.
- models/ : Définit les structures de données et interactions avec la base de données.
- routers/ : Définit les routes de l'API.
- middlewares/ : Gère la logique intermédiaire comme l'authentification ou les logs.
- public/ : Contient les fichiers statiques (HTML, CSS, JS).