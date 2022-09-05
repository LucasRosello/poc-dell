# Todo
- [ ] MUX API
- [ ] RabbitMQ
- [ ] Postgress
- [ ] Docker Compose
- [ ] Kubernetes Kind 

# Docs Used
- [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)
- [jackc pgx](https://github.com/jackc/pgx)
- [Getting started with pgx](https://github.com/jackc/pgx/wiki/Getting-started-with-pgx-through-database-sql)
- [Work with Go & PostgreSQL using pgx](https://medium.com/geekculture/work-with-go-postgresql-using-pgx-caee4573672)


 docker run --name testing-postgres -e POSTGRES_PASSWORD=secret -d -p 5432:5432 postgres

 docker exec -it ID bash
 psql -U postgres

CREATE TABLE product(product_code  VARCHAR (50) PRIMARY KEY, description VARCHAR(50), avaliable BOOL);
INSERT INTO product (product_code, description, avaliable) values ('E35', 'Secaplatos', true);
select * from product;