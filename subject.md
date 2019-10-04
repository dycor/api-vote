# API de vote

## Contexte

Après le Brexit l'Europe a pris un coup dans l'aile politiquement et économiquement.
Plusieurs mouvement séparatistes tentent des actions afin de continuer à affaiblir le pouvoir.

Un mouvement populaire non violent dont vous faites parti est en train d'émerger.
Entant que développeur vous avez la tâche de créer des outils permettant de voter des propositions de loi permettant aux citoyens de sortir de cette crise.

## Règles du projet

### Requis

- Votre équipe doit être constituée de 3 à 4 personnes.
- Le projet doit pourvoir builder.
- Les équipes ne doivent pas partager leur code et librairies créé lors du projet avec d'autres équipes.
- Vous avez la possibilité d'utiliser des librairies externe au projet dans la mesure ou elles sont open source.
- La date limite du rendu du projet doit se faire le dernier jour des cours.
- Pour les dépendances externes au projet tel qu'une base de donnée veuillez à n'utiliser que Docker.
- l'API créé doit permettre la persistance dans une base de donnée PostgreSQL.

### Points bonnus

- Veuillez à ajouter un fichier README dans votre projet permettant de prendre en main les commandes créés.
- Le projet doit être versionné avec GIT, une partie de la note se fait sur la propreté de l'historique :)
- Ajouter dans tous les modules nécessaire afin que le projet puisse se builder à toutes les étapes des commits.
- Afin de builder le projet et lancer l'API veuillez utiliser un fichier make à la racine avec une commande run.
- Toutes les fonctions exportés doivent être documentés et ce en Anglais.
- Le projet doit avoir des tests unitaire afin de vérifier des fonctionnement logique des fonctions et de l'API.

### Outils requis pour le bon déroulement du projet

- Go 1.13 installé sur sa machine et être capable de compiler sur sa machine.
- Docker installé avec une image PostgreSQL.
- Visual Studio Code pour l'IDE ou VI pour ceux qui préfèrent.
- Postman d'installé afin de tester les endpoints ou curl.
- Git.
- GORM qui est un ORM en Go permettant de modéliser la base de donnée et gérer les CRUDS
- Gin qui est un framework permettant la réalisation des services et appels de méthodes

## API

L'API permet de réaliser des opérations de CRUD sur la base de donnée, le modèle utilisé est REST; Afin de ne pas dévoiler les IDs interne de la base de donnée ce sont des UUIDs qui sont utilisés dans les endpoints. Les codes de retours à utiliser sont 200  et 401 pour un appel non autorisé. 

### Login

Le login permet à un utilisateur déjà enregistré de se logger dans l'application via un jeton JWT.
Le JWT contient l'UUID et son AccessLevel de l'utilisateur connecté permettant ainsi d'ajouter des règeles métier en fonction des droits d'accès.

```ssh
POST /login
```

- Envoie un Login (Email) et le password de l'utilisateur.
- Le password est haché avec la même méthode qui est utilisé pour la création d'un utilisateur.
- Une vérification est faite dans la base de donnée permettant de s'assurer que sur la base du même login le password correspond; si il correspond envoyer un JWT avec en valeur l'UUID et l'AccessLevel de l'utilisateur loggé
- Aprés trois tentatives infructueuses de login l'utilisateur est bloqué sur la base de son IP

```json
# call service
{
 "login":"admin",
 "password":"pass"
}
# response
{
 "jwt":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```

### User

Il y a deux catégories de Users en fonction de leur AccessLevel, l'administrateur et l'utilisateur classique (le votant).
Le CRUD en REST permet d'ajouter modifier et supprimer des Utilisateurs en fonction des règles de validation des données et des règles métier décrites ci-dessous :

#### Création d'un nouvel User

```ssh
POST /users
```

les champs FirstName, LastName, Email, DateOfBirth et Password à la création sont obligatoire
les champs FirstName, et LastName ne peuvent contenir des nombres ou espace et doivent faire au minimum 2 caractères
le champ Email doit être bien formé selon la RFC en vigeur
le champ DateOfBirth doit être vérifié pour s'assurer que la personne est majeur (>= 18 ans) à la date de création
le password doit être haché avant d'être stocké dans la base de donnée
cette ressource renvoie les informations envoyés en dehors de l'ID interne crée dans la base de donnée et du mot de passe; elle y ajoute l'UUID.
Seul un administrateur peut créer un autre administrateur.

```json
# call service
{
 "first_name":"Karim",
 "last_name":"Benzema",
 "email":"k@benzema.io",
 "pass":"Mostafa87",
 "birth_date":"19-12-1987"
}
# response
{
 "uuid":"210b546a-d61c-11e9-9451-333e2779ce66",
 "first_name":"Karim",
 "last_name":"Benzema",
 "email":"k@benzema.io",
 "birth_date":"19-12-1987"
}
```
#### Modification d'un User

``` PUT /users/[:uuid]```

les champs qui penvent être changés sont uniquement FirstName, LastName, Email, Password
les règles de vérification doivent être les mêmes appliqués lors de la création sur ces derniers champs.
Seul un administrateur peut changer les données d'un autre utilisateur.

```json
# call service
{
 "email":"k@benzema.com",
 "pass":"AKvhsu&29$fm",
}
# response
{
 "uuid":"210b546a-d61c-11e9-9451-333e2779ce66",
 "first_name":"Karim",
 "last_name":"Benzema",
 "email":"k@benzema.com",
 "birth_date":"19-12-1987"
}
```

#### Destruction d'un User

```ssh
DELETE /users/[:uuid]
```

L'utilisateur doit continuer à exister dans la base de donnée (soft delete)
Seul un administrateur est en mesure d'appeler ce service 

### Vote

Afin de voter un utilisateur doit avoir un jeton JWT
toutes actions de vote se fait sur la base de son UUID dans le JWT.

#### Création d'une proposition d'un Vote

``` POST /votes```

Seul les administrateurs ont les droits pour créer une proposition de vote
la proposition doit avoir les champs Titre et description de rempli.


```json
# call service
{
 "title":"Propreté des trottoirs",
 "desc":"Dans le budget qui sera soumis au vote des conseillers de Paris lundi, 32 M€ seront consacrés à l’achat de nouvelles machines, plus silencieuses, plus propres et plus performantes. Il n’y aura pas d’embauche d’agents supplémentaires."
}
# response
{
 "uuid":"c9e10e68-d61c-11e9-9bdf-070102731d2b",
 "title":"Propreté des trottoirs",
 "desc":"Dans le budget qui sera soumis au vote des conseillers de Paris lundi, 32 M€ seront consacrés à l’achat de nouvelles machines, plus silencieuses, plus propres et plus performantes. Il n’y aura pas d’embauche d’agents supplémentaires."
}
```

#### Récupération d'un Vote

```ssh
GET /votes/[:uuid]
```

Ce cervice permet de récupérer les détails sur un vote avec les votes associés

```json
# response
{
 "uuid":"c9e10e68-d61c-11e9-9bdf-070102731d2b",
 "title":"Propreté des trottoirs",
 "desc":"Dans le budget qui sera soumis au vote des conseillers de Paris lundi, 32 M€ seront consacrés à l’achat de nouvelles machines, plus silencieuses, plus propres et plus performantes. Il n’y aura pas d’embauche d’agents supplémentaires.",
 "uuid_votes":["a2926a2c-d61d-11e9-83fa-bb435499d1ac","a3951cbc-d61d-11e9-8d4e-2bcfbd18769d","a4da3f08-d61d-11e9-839a-43822a8be55e"]
}
```

#### Édition d'un Vote

```ssh
PUT /votes/[:uuid]
```

Seul un administrateur peut changer le titre et la description.
L'administrateur n'a pas les droits pour changer la collection d'UUIDVote.
Si un votant appel ce service c'est son UUID contenu dans le JWT qui est utilisé pour l'ajouter à la collection de UUIDVote.


```json
# call service
{
 "start_date":"10-09-2019",
 "end_date":"20-09-2019"
}
# response
{
 "uuid":"c9e10e68-d61c-11e9-9bdf-070102731d2b",
 "title":"Propreté des trottoirs",
 "desc":"Dans le budget qui sera soumis au vote des conseillers de Paris lundi, 32 M€ seront consacrés à l’achat de nouvelles machines, plus silencieuses, plus propres et plus performantes. Il n’y aura pas d’embauche d’agents supplémentaires.",
 "uuid_votes":null,
 "start_date":"10-09-2019",
 "end_date":"20-09-2019"
}
```

## Entités

**User** : ID (int), UUID (string), AccessLevel (int), FirstName (string), LastName (string), Email (string), Password (string), DateOfBirth (time.Time), CreatedAt (time.Time), UpdatedAt (time.Time), DeletedAt (*time.Time).
l'ID est créé par la base de donnée en incrémental.
le champs UUID est créé coté serveur à la réception du payload après validation des données.
le password est haché coté serveur est stocké ensuite dans la base de donnée.

**Vote** : ID (int), UUID (string), Title (string), Description (text), UUIDVote (collection), StartDate (time.Time), EndDate (time.Time), CreatedAt (time.Time), UpdatedAt (time.Time), DeletedAt (*time.Time).
L'UUID vote est une collection d'UUID d'utilisateurs ayant voté.


## Liens

- [gorm](https://github.com/jinzhu/gorm)
- [Makefile](https://www.gnu.org/software/make/manual/make.html)
- [Golang specs](https://golang.org/ref/spec)
- [SHA256](https://golang.org/pkg/crypto/sha256/)
- [UUID](https://github.com/satori/go.uuid)
- [JWT](https://jwt.io/)
- [JWT middleware with Gin](https://github.com/appleboy/gin-jwt)
- [Gin](https://github.com/gin-gonic/gin)
- [status code](https://fr.wikipedia.org/wiki/Liste_des_codes_HTTP)
- [Postman](https://www.getpostman.com/)
- [Docker](https://www.docker.com/)
- [Docker image PostgreSQL](https://docs.docker.com/samples/library/postgres/)
- [Docker image Alpine](https://docs.docker.com/samples/library/alpine/)
