## What is this project about?

This is just me trying to build a project in Go. I chose to build a simple CRUD application using Gin because Gin is probably the most popular HTTP framework for Go out there

## What do I need to have installed in order to run this project?

1. You need to have an installation of Go. I am using Go version 1.22.2
2. You need to have an installation of a PostgreSQL database.

## How do I run this project?

1. First, clone the repo.  
   `git clone git@github.com:arimotearipo/movies.git`

2. Within the cloned directory, create a new file and name it `.env`
3. Add the configurations as follows. You can copy the values below directly and paste it in your `.env` file.

   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=abcdef
   DB_NAME=postgres

   PORT=8080
   ```

   `DB_HOST` shall be the IP address of your PostgreSQL database

   `DB_PORT` shall be the port number of your PostgreSQL database

   `DB_USER` shall be the user id of your PostgreSQL database

   `DB_PASSWORD` shall be the password for your PostgreSQL database

   `DB_NAME` shall be the name of your PostgreSQL database

   `PORT` shall be the port number where you want to host this movies server application

4. i) To run the project in development mode run this command:

   `go run .`

   ii) Or if you want to build the project, run the two commands below:

   `go build`

   `./movies`
