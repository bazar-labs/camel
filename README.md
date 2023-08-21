## Dev setup

1. Create an `.env` file
2. Run `docker compose up` (it will start the db and run the [migrations](https://github.com/golang-migrate/migrate#cli-usage))
3. Run `go run .`

## Resetting local database

- Run `docker compose down --volumes` to reset the database state, then run `docker compose up` (you should be able to see all migrations running again)

## Creating new migration files

0. Install migrate cli from `https://github.com/golang-migrate/migrate/tree/master/cmd/migrate`
1. Run `migrate create -ext sql -dir db/migrations <short_name_describing_the_migration>`
2. Restart the docker container

## Testing

- Some packages require the docker compose file to run (like the `store` package)
- Some packages make use of mocks. Install [`moq`](https://github.com/matryer/moq) to generate mocks for interfaces

## Note

- ciaoooooooo
