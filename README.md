# blog-server
blog-server is a simple blog server. It provides the following functions：

* user registration and login
* create/edit/post/list/get draft
* list&&search/get/comment/like/dislike posts
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

`make run arg=migrate`

If you don't make the tables ready, you can use this script to initialize your tables.
