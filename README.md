<h1 align="center">
  Tasks - Api
</h1>

## Description
Esse projeto foi desenvolvido com finalidade de criar uma API para gerenciamente de tarefas

O banco de dados escolhido foi o PostgreSQL, por ser capaz de lidar eficientemente com grandes volumes de dados e acessos simultâneos. Além disso, optei por utilizá-lo como um desafio pessoal, já que é a primeira vez que o utilizo em um projeto.

## Técnologias necessárias

- [Go](https://go.dev/)
- [PostreSQL](https://www.postgresql.org/)

## Testando localmente


Para rodar o projeto localmente será necessários ter instalado a linguagem [Golang](https://go.dev/dl/)

O projeto pode ser clonado utiizando o seguite comando:
```
https://github.com/JoyceEllenNeryTeles/tasks-api.git
cd tasks-api
```

Para o banco de dados será necerrário ter o [PostgreSQL](https://www.postgresql.org/download/) que pode ser obtido de duas fomas: através do docker ou baixando localmente

Caso esteja usando o docker, basta accessar a pasta app e executar o comando
```
docker compose up
```

Com o Postegres localmente será necessario criar a tabela no BD, para isso pode-se utilizar o [DBeaver](https://dbeaver.io/download/)

Rode o script que se encontra na pasta app/infra/ToDo.sql

Após isso, dentro da pasta app execute o seguinte comando para levantar a aplicação

```
go run main.go
```

A collection do insomnia com as chamadas se encontra dentro da pasta app 


