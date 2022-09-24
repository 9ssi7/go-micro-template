# Go Micro Template

This is a template for a microservice written in Go.

## Overview

You can start your microservice applications that you will develop with Go without any architectural concerns. This template is for illustrative purposes only. Technologies used may vary. This template has been developed in a way that does not hinder most Technologies.

## Project Structure

📦microservice-app
┣ [`📂src`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src)
┃ ┣ [`📂app`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/app)
┃ ┃ ┗ 📜app.go
┃ ┣ [`📂config`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/config)
┃ ┃ ┗ 📜config.go
┃ ┣ [`📂dto`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/dto)
┃ ┃ ┣ 📜some_create.go
┃ ┃ ┣ 📜some_created.go
┃ ┃ ┣ 📜some_find.go
┃ ┃ ┗ 📜some_found.go
┃ ┣ [`📂entity`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/entity)
┃ ┃ ┗ 📜some.go
┃ ┣ [`📂event`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/event)
┃ ┃ ┣ 📜some_created.go
┃ ┃ ┗ 📜some_deleted.go
┃ ┣ [`📂event_handler`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/event_handler)
┃ ┃ ┣ 📜event_handler.go
┃ ┃ ┗ 📜some-feature_handler.go
┃ ┣ [`📂event_publisher`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/event_publisher)
┃ ┃ ┣ 📜event_publisher.go
┃ ┃ ┣ 📜some-created_publisher.go
┃ ┃ ┗ 📜some-deleted_publisher.go
┃ ┣ [`📂internal`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/internal)
┃ ┃ ┣ 📜api.go
┃ ┃ ┣ 📜handler.go
┃ ┃ ┣ 📜repo.go
┃ ┃ ┗ 📜service.go
┃ ┣ [`📂locales`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/locales)
┃ ┃ ┣ 📜en.toml
┃ ┃ ┗ 📜tr.toml
┃ ┣ [`📂mapper`](https://github.com/ssibrahimbas/go-micro-template/tree/main/src/mapper)
┃ ┃ ┣ 📜mapper.go
┃ ┃ ┗ 📜some_mapper.go
┃ ┣ 📜app.env
┃ ┗ 📜main.go
┣ 📜.gitignore
┣ 📜Dockerfile
┣ [`📜Sample.md`](https://github.com/ssibrahimbas/go-micro-template/tree/main/Sample.md)
┣ 📜go.mod
┗ 📜go.sum

## Folder Descriptions

In this repository, there is a `README.md` file under each folder, and in this file, what the folder does is explained in detail. For example, [you may want to visit the `src` folder.](https://github.com/ssibrahimbas/go-micro-template/tree/main/src)