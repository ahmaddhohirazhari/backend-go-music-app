<h1 align="center">Backend Go Music App</h1>
<p align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/2560px-Go_Logo_Blue.svg.png" width="400px" alt="Golang.jpg" /></p>
<p align="center">
    <a href="https://golang.org/" target="blank">More about Golang</a>
</p>

## 🔗 Description

This Backend Application is used for vehicle rental systems such as car rental, motorbikes, and bicycles. In the application, users can add, change, delete, and read the data of the vehicle they want to rent. In addition, users can also see the rental history. This application was built using the Golang programming language with the Gorilla / Mux Framework and uses GORM, a Database that is used using PostgreSQL and deployed on the Heroku website.

## 🔗 Installation Gorilla/Mux

- Install Gorilla/Mux

```sh
  go get -u github.com/gorilla/mux
```

## 🔗 Feature

- CRUD Music

## 🔗 Installation Step

- Go to the project directory

```sh
  mkdir BackendGo
  cd BackendGo

  go mod init BackendGo
  # add file main.go
```

- Clone the project

```sh
  git clone https://github.com/ahmaddhohirazhari/backend-go-music-app.git (HTTPS)
  git clone git@github.com:ahmaddhohirazhari/backend-go-music-app.git (SSH)
```

- Add Env

```sh
  APP_PORT= Your Port
  JWT_KEYS= Your Secret Keys

  DB_USER = Your DB User
  DB_HOST = Your DB Host
  DB_NAME = Your DB Name
  DB_PASS = Your DB Password
```

- Install dependencies

```sh
  go get -u ./..
  # or
  go mod tidy
```

- Start the server

```sh
  go run main.go server
```

## 💻 Built with

- [Golang](https://go.dev/): Programming
- [gorilla/mux](https://github.com/gorilla/mux): for handle http request
- [Postgres](https://www.postgresql.org/): for DBMS

## 💻 Deploy

Link Deploy :

## 💻 End Point

Postmant Documentation :

- Postman : [link](https://documenter.getpostman.com/view/19983829/2s8ZDbVfs6)

## 🚀 About Me

- Github : [ahmaddhohirazhari](https://github.com/ahmaddhohirazhari/)
