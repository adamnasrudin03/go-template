# go-template
 This basic REST API service uses the main features of Golang: user, log, and auth with bearer tokens.

## Technology Used
- Versioning using Git (See [Git Installation](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git))
- Programming Language using Go 1.22 or later. (See [Golang Installation](https://golang.org/doc/install))
- DB using PostgreSQL 14 or later. (See [PostgreSQL Installation](https://www.postgresql.org/download/))
- Cache Using Redis. (See [Redis Installation](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/) )
- Routing Using Gin (See [Gin Framework Doc](https://gin-gonic.com/docs/quickstart/))
- ORM Library using GORM. (See [GORM Guides](https://gorm.io/docs/index.html) )
- Auth using Golang-JWT. (See [Golang-JWT Guides](https://github.com/golang-jwt/jwt))
- Message Broker using RabbitMQ. (See [RabbitMQ Installation](https://www.rabbitmq.com/docs/download/) )
- Deploy using Docker. (See [Docker Installation](https://docs.docker.com/desktop/) )


## Development Guide

### Documentations
- not done

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