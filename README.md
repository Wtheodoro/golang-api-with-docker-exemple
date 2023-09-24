# Go API Project with Docker and Docker Compose

This is a base project for creating a Go (Golang) API using Docker and Docker Compose. It includes a Dockerfile for the API and another for a MySQL database.

## Requirements

Make sure you have Docker and Docker Compose installed on your machine before proceeding.

- Docker: [Install Docker](https://docs.docker.com/get-docker/)
- Docker Compose: [Install Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

1. Clone this repository to your machine:

   ```bash
   git clone https://github.com/Wtheodoro/golang-api-with-docker-exemple
   cd golang-api-with-docker-exemple

   ```

2. To start the development environment, run the following command:
   docker-compose up --build

3. Access the API at http://localhost:8080 (or the address you configured) and begin developing your Go application.

4. To stop the containers, you can run the following command:
   docker-compose down
