# Go REST API
A RESTful API with Go

It is build with **golang**, **gofiber** and **mongodb**

## Installation & Run
```bash
# Download this project
go get github.com/Devansh-365/gofiber
```

```bash
# Build and Run
cd gofiber
go build
./gofiber

# API Endpoint : http://127.0.0.1:3000
```

## API

#### /users
* `GET` : Get all users

#### /users/:userId
* `GET` : Get a user

#### /user
* `POST` : Create user

#### /user/:userId
* `PUT` : Update a user
* `DELETE` : Delete a user


## Todo

- [x] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 
