# golang-crud

1. create datebase in postgress sql for crud.
# postgres

[Nice tutorial](https://www.tutorialspoint.com/postgresql/postgresql_create_database.htm)
[documenation reference](https://www.postgresql.org/docs/9.4/static/app-psql.html)

## install
[postgresql website](https://www.postgresql.org/download/)

##create database

```
CREATE DATABASE golang;
```

##create table
```
CREATE TABLE books (
   isbn INT PRIMARY KEY     NOT NULL,
   title           TEXT    NOT NULL,
   author        CHAR(50),
   price         REAL DEFAULT 25500.00,
);
```
## final step
open diretory and run code in your command line 
```
go run main.go
```
 
