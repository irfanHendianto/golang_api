# golang_api

Before Running Program
- create db first and set file .env base on your computer

Running Program 
- go run server.go

Runing Unit Test
- set data in function testing
- to running unit test using command go test -v - run (NameFunctionTesting)

Runing Docker Compose : docker compose up

for running command for unit test 
  - running command docker exec -it containerName /bin/sh
  - after inside docker container running go test -v -vet=off -run funcNameTesting
 


Format data yang di harus isi
DB_USER=
DB_PASS=
DB_HOST=
DB_NAME=
