# go-template
 This basic REST API service uses the main features of Golang: user, log, [translate message](https://github.com/adamnasrudin03/go-template/blob/main/pkg/helpers/languange.go#L45) and auth with bearer tokens.

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


## Development Guide

### Documentations
- not done

### Collection Using Insomnia
- ./go-template-{date_export}.json
  
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
- Start service API
    ```sh
        go run main.go
    ```

## Build project by docker
- change data environment in file ./docker-compose.yml
- build with docker compose

    ```sh
        docker-compose -f "docker-compose.yml" up -d --build 
    ```


###

<br clear="both">
<h3 align="left">Connect with me:</h3>
<div align="left">
  <a href="https://www.linkedin.com/in/adam-nasrudin/" target="_blank">
    <img src="https://img.shields.io/static/v1?message=LinkedIn&logo=linkedin&label=&color=0077B5&logoColor=white&labelColor=&style=for-the-badge" height="35" alt="linkedin logo"  />
  </a>
</div>

###