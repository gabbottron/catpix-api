# catpix-api
A challenge project to implement a cat pictures API

## Local dev setup
#### Dependencies
1. Docker
2. Docker Compose
3. Copy env_example file to .env

**IMPORTANT:** Before running this project locally, you MUST have a .env file in the project root (see steps above!)

### Container Commands
```
docker-compose up   : Bring up the API and dependencies (building them if necessary)
docker-compose down : Bring down the API and dependencies
```

**DEFAULT URI:** http://localhost:8080

**POSTMAN COLLECTION:** https://www.getpostman.com/collections/fa46ec11f62bb24ad333
The above link will load all of the requests and macros for authentication into Postman for you.

## Open Endpoints

Open endpoints require no Authentication.
* [HealthCheck](doc/healthcheck.md)                     : `GET  /check`
* [Login](doc/login.md)                                 : `POST /login`
* [New User](doc/new_user.md)                           : `POST /user`
* [Get Picture by ID](doc/get_picture.md)               : `GET /picture/:picture_id`
* [Get All Pictures](doc/get_pictures.md)               : `GET /pictures`
* [Get All Pictures for User](doc/get_user_pictures.md) : `GET /user/:user_id/pictures`

## Endpoints that require Authentication

Closed endpoints require a valid Token to be included in the header of the
request. A Token can be acquired from the Login view above.
* [Create Picture](doc/new_picture.md)    : `POST  /auth/picture`
* [Update Picture](doc/update_picture.md) : `PUT  /auth/picture/:picture_id`
* [Delete Picture](doc/delete_picture.md) : `DELETE  /auth/picture/:picture_id`
* [New User](doc/new_user.md)             : `POST /user`

