# go-ecommerce-api

# psql -U postgres
###  /l 
#### to show db-table list
### CREATE DATABE [name of db]
### DROP DATABE [name of db]

#### migrate create -ext sql -seq [name of file format .sql]

####  migrate -path pkg/databases/migrations -database 'postgresql://postgres:postgres@localhost:4444/go_ecommerce?sslmode=disable' -verbose up