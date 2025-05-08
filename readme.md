## Gerar mocks com mockgen

```sh
mockgen -destination=application/mocks/application.go -source=application/product.go application
```

# SQLITE3
## Entrar no sqlite3
```sh
sqlite3 .files/db/sqlite.db 
```

## Criar a tabela products
```sh
create table products(id string, name string, priva float, status string);
```

## Listar tabelas no sqlite
```sh
.tables
```