# go-template
 This project is an example of a REST API project using the Go language and the implementation of other tools, such as Auth with JWT, Logger API, Cache, ORM SQL, Message Broker, Export Excel File, OTP Mechanism, Deploy with Docker, Clean Code (smell code checker by CodeScene), Unit Test, and so on.


## Technology Used
- Versioning using Git (See <a href="https://git-scm.com/book/en/v2/Getting-Started-Installing-Git" target="_blank">Git Installation</a>)
- Programming Language using Go 1.22 or later. (See <a href="https://golang.org/doc/install" target="_blank">Golang Installation</a>)
- DB using PostgreSQL 14 or later. (See <a href="https://www.postgresql.org/download/" target="_blank">PostgreSQL Installation</a>)
- Migration using ORM GORM. (See <a href="https://gorm.io/docs/migration.html" target="_blank">Documentation</a>)
- Cache Using Redis. (See <a href="https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/" target="_blank">Redis Installation</a>)
- Routing Using Gin (See <a href="https://gin-gonic.com/docs/quickstart/" target="_blank">Gin Framework Doc</a>)
- ORM Library using GORM. (See <a href="https://gorm.io/docs/index.html" target="_blank">GORM Guides</a>)
- Auth using Golang-JWT. (See <a href="https://github.com/golang-jwt/jwt" target="_blank">Golang-JWT Guides</a>)
- Message Broker using RabbitMQ. (See <a href="https://www.rabbitmq.com/docs/download/" target="_blank">RabbitMQ Installation</a>)
- Deploy using Docker. (See <a href="https://docs.docker.com/desktop/" target="_blank">Docker Installation</a>)

## Feature
| Method | Auth | Endpoint                                   | Dec                                                                   |
| ------ | ---- | -------------------------------------------| --------------------------------------------------------------------- |
| GET    | no   | /                                          | Welcome API                                                           |
| POST   | no   | /api/v1/auth/sign-up                       | Register user with role USER                                          |
| POST   | yes  | /api/v1/root/auth/sign-up                  | Register user with role USER or ADMIN, auth with super admin (root)   |
| POST   | no   | /api/v1/auth/sign-in                       | Login                                                                 |
| PATCH  | yes  | /api/v1/users/:id                          | Update User                                                           |
| GET    | yes  | /api/v1/users/:id                          | Detail User                                                           |
| GET    | yes  | /api/v1/users                              | List User, auth only admin or super admin (root)                      |
| PATCH  | yes  | /api/v1/users/change-password/:id          | Change Password                                                       |
| GET    | yes  | /api/v1/users/send-email-verify            | Send OTP Email verified                                               |
| POST   | yes  | /api/v1/users/verified-email               | Verified email with otp                                               |
| GET    | no   | /api/v1/users/request-reset-password/:id   | Send OTP Email forgot password                                        |
| PATCH  | no   | /api/v1/users/validate-reset-password/:id  | Verified Reset password                                               |
| GET    | yes  | /api/v1/logs                               | List log activity history                                             |
| GET    | yes  | /api/v1/logs/download                      | Download xlx List log activity history                                |
| GET    | no   | /api/v1/message/translate/id               | Translate text to language id (indonesia)                             |
| GET    | no   | /api/v1/message/consumer                   | Trigger manual consume queue rabbitMQ                                 |

### Role
- ROOT  (role super admin) = create a user the first time the project is run, <a href="https://github.com/adamnasrudin03/go-template/blob/main/pkg/seeders/user.go#L14" target="_blank"> check here </a> 
- ADMIN (role admin)
- USER (role user)
  

## Development Guide

### Documentations
  <a href="https://documenter.getpostman.com/view/10619265/2sA3Qzaooy" target="_blank"> Postman API Documentation </a>

### Collection Using Postman
- ./go-template.postman_collection.json
  
### Installation
- Clone this repo

    ```sh
        git clone https://github.com/adamnasrudin03/go-template.git
    ```

- Copy `.env.example` to `.env`

    ```sh
        cp .env.example .env
    ```
- Setup local database
- If you using RabbitMQ, Please check or create queue for <a href="https://github.com/adamnasrudin03/go-template/blob/main/app/models/queue.go#L8" target="_blank"> ./app/models/queue.go </a>
- Start service API
    ```sh
        go run main.go
    ```

## Build project by docker
- check ip address in terminal
    ```sh
        ipconfig
    ```
- change data environment in file ./docker-compose.yml
- build with docker compose

    ```sh
        docker-compose -f "docker-compose.yml" up -d --build 
    ```
    - with make file
    ```sh
        make docker
    ```

## Coverage Unit Test
  - with make file
  ```sh
      make cover
  ```



### Connect with me
  <a href="https://www.linkedin.com/in/adam-nasrudin/" target="_blank">
    <img 
        src="https://img.shields.io/static/v1?message=LinkedIn&logo=linkedin&label=&color=0077B5&logoColor=white&labelColor=&style=for-the-badge" 
        height="35" alt="linkedin logo"  />
  </a>
