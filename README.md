## Telegram task manager bot
Is a simple telegram bot to store your school or university subjects, anything else and the tasks they contain
### Setup dev environment
**Requirements**
* docker
* golang (>=1.18)
* golang-migrate

**Preparations**
```shell
# run Postgres in docker
docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=bot -d postgres

# migrate database schema
make migrate-up
```
**Run**
```shell
# run local bot
make run
```