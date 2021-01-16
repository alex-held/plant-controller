# Plant-Controller
[![Copyright 2021](https://img.shields.io/badge/copyright-2021-blue.svg?maxAge=60&style=flat-square)](https://github.com/alex-held/plant-controller) ![GitHub](https://img.shields.io/github/license/alex-held/template?style=flat-square)

[![GitHub](https://img.shields.io/badge/github-issues-red.svg?maxAge=60&style=flat-square)](https://github.com/alex-held/plant-controller/issues) [![GitHub Wiki](https://img.shields.io/badge/github-wiki-181717.svg?maxAge=60&style=flat-square)](https://github.com/alex-held/plant-controller)
![GitHub pull requests](https://img.shields.io/github/issues-pr/alex-held/plant-controller?logo=GitHub&style=flat-square)          ![Commit Activity](https://img.shields.io/github/commit-activity/m/alex-held/plant-controller?logo=when-i-work&logoColor=white&style=flat-square)

![frontend docker shield](https://img.shields.io/docker/v/alexheld/plant-controller-frontend?sort=semver&label=frontend&logo=docker&style=flat-square)
![backend docker shield](https://img.shields.io/docker/v/alexheld/plant-controller-backend?sort=semver&label=backend&logo=docker&style=flat-square)




## Environment setup


You need to have [Go](https://golang.org/),                                                      

[Node.js](https://nodejs.org/),
[Docker](https://www.docker.com/), and
[Docker Compose](https://docs.docker.com/compose/)


Verify the tools by running the following commands:

```sh
go version
npm --version
docker --version
docker-compose --version
```

## Start in development mode
                                               
### Start the database
In the project directory run the command (you might
need to prepend it with `sudo` depending on your setup):
```sh
docker-compose -f docker-compose-dev.yml up
```

This starts a local PostgresSQL on `localhost:5432`.
The database will be populated with test records
from the [init-db.sql](data/init-db.sql) file.
                                             
### Start the backend service
Navigate to the `server` folder and start the back end:

```sh
cd server
export LOG_LEVEL=debug
export PG_URL=postgres://postgres:postgres@localhost:5432/plant-controller?sslmode=disable
export HTTP_ADDR=localhost:8080
go run main.go
```

The back end will serve on http://localhost:8080.

Navigate to the `frontend` folder, install dependencies,
and start the front end development server by running:

```sh
cd frontend
npm install
npm start
```
The application will be available on http://localhost:3000.
 
## Start in production mode (docker-compose)

Perform:
```sh
docker-compose up
```

This will build the application and start it together with
its database. Access the frontend on http://localhost:80, the backend on http://localhost:8080
