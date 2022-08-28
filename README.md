# blog-server
blog-server is a simple blog server. It provides the following functions：

* user registration and login
* create, edit, post, list, get draft
* list&&search, get, like, dislike posts
* comment posts

The blog-server uses JWT for login authentication, MySQL for data storage and Redis for caching.

## Preparation

- Go1.17
  
- MySQL 5.7.19
  
- Redis 6.0.6

## Start Server

Before you start the server, you should make sure that the database and tables exists.

`make run`

## DB Migration

`make migrate`

If you don't have the tables ready, you can use this script to initialize your tables.

`make rollback`

If you want to rollback last migrate, run this command.

## Test

Run `make test`, you can run test cases.

## API docs

blog-server use [swaggo](https://github.com/swaggo/swag) to generate and serve api docs.

When you run blog-server in develop mode, browse to `http://localhost:{your port}/swagger/index.html`, then you can see the api docs.
