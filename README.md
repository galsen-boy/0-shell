# 0-SHELL

Un shell simple implémenté en Go.
L'objectif de ce projet est de créer une shell simple.

Par l'intermédiaire de la 0-shell vous arriverez au cœur d'Unix et explorer une partie importante de l'API de ce système qui est la création et la synchronisation des processus. L'exécution d'une commande à l'intérieur d'un shell implique la création d'un nouveau processus, dont l'exécution et l'état final seront surveillés par les processus de ses parents. Cet ensemble de fonctions sera la clé du succès du projet.

Pour ce projet, vous n'aurez qu'à créer un simple Unix shelloù vous pouvez exécuter certaines des commandes les plus connues. Pour cette partie du projet, aucune fonction avancée, tuyaux ou redirection ne sera demandé, mais vous pouvez les ajouter si vous voulez.

## Description

Ce projet est une implémentation basique d'un shell en Go. Il offre les fonctionnalités suivantes :

- Exécution de commandes shell basiques
- Gestion de la commande `cd` pour changer de répertoire

## Structure du projet

```
.
├── go.mod
├── main.go
├── README.md
├── student-script.sh
└── utils
    ├── command.go
    └── utils.go
```

- `main.go` : Point d'entrée du programme
- `utils/command.go` : Gestion des commandes
- `utils/utils.go` : Fonctions utilitaires

## Installation

1. Assurez-vous d'avoir Go installé sur votre système.
2. Clonez ce dépôt :
   ```
   git clone https://github.com/votre-username/0-shell.git
   ```
3. Naviguez dans le répertoire du projet :
   ```
   cd 0-shell-go
   ```

## Utilisation

Pour lancer le shell, exécutez :

```
go run .
```

Une fois le shell lancé, vous pouvez entrer des commandes comme vous le feriez dans un terminal normal.

### Exemples de commandes :


- **echo**
- **cd**
- **ls** avec les flags -l, -a et -F
- **pwd**
- **cat**
- **cp**
- **rm** avec le  flag -r
- **mv**
- **mkdir**
- **exit**


## Fonctionnalités

- **Exécution de commandes** : Le shell peut exécuter la plupart des commandes Unix standard.
- **Changement de répertoire** : Utilisation de `cd` pour naviguer dans le système de fichiers.
- **Variables d'environnement temporaires** : Possibilité de définir des variables d'environnement pour une seule commande.
- **Prompt personnalisé** : Affichage du répertoire courant en vert dans le prompt.

## Contribution

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir une issue ou à soumettre une pull request.

## Contributers
- [daiba](https://learn.zone01dakar.sn/git/daiba) (Daibou BA)
- [ialimoud](https://learn.zone01dakar.sn/git/ialimoud) 
- [ndiba](https://learn.zone01dakar.sn/git/ndiba) (Ndiaga BA)