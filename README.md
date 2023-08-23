# Go Books Store

This is an online book store web app.

## Technologies

- GO 1.20
- MySql

## Functionality

- See all the books available
- See single book by id
- Create new book
- Update existing book by id
- Delete existing book by id

## Project Start Command

```bash
go install
swag init
go run .
```

or for development purposes use below nodemon command for live refresh as you make changes

```bash
go install
swag init
nodemon -e go --signal SIGTERM --exec 'go' run .
```

Defaultly server will start running on port 3000 which can be modified from main.go file

## Authors

- [@neelthakkar](https://www.github.com/Neel278)
