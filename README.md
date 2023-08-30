# Build a REST API with Golang and MongoDB

Este é um exemplo de projeto que demonstra como construir uma API RESTful usando a linguagem de programação Go e o banco de dados MongoDB.

## Pré-requisitos

* [Go](https://golang.org/dl/)
* [MongoDB](https://www.mongodb.com/download-center/community)
* [Postman](https://www.getpostman.com/downloads/)

## Dependências

* [Gin](github.com/gin-gonic/gin)
* [MongoDB Go Driver](go.mongodb.org/mongo-driver/mongo)
* [Godotenv](github.com/joho/godotenv)
* [Validator](https://github.com/go-playground/validator)

## Endpoints

| Método HTTP | Endpoint | Descrição |
| --- | --- | --- |
| GET | /items | Retorna uma lista de itens |
| POST | /items/add | Adiciona um novo item |
