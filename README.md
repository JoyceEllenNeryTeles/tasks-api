<h1 align="center">
  Tasks - Api
</h1>

<br />
<p align="center">
  <img src="https://img.shields.io/static/v1?label=license&message=MIT" alt="License">
  <img src="https://img.shields.io/github/repo-size/Lissone/point-control-api" alt="Repo size" />
  <img src="https://img.shields.io/github/languages/top/Lissone/point-control-api" alt="Top lang" />
  <img src="https://img.shields.io/github/stars/Lissone/point-control-api" alt="Stars repo" />
  <img src="https://img.shields.io/github/forks/Lissone/point-control-api" alt="Forks repo" />
  <img src="https://img.shields.io/github/issues-pr/Lissone/point-control-api" alt="Pull requests" >
  <img src="https://img.shields.io/github/last-commit/Lissone/point-control-api" alt="Last commit" />
</p>

<p align="center">
  <a href="https://github.com/Lissone/point-control-api/issues">Report bug</a>
  ·
  <a href="https://github.com/Lissone/point-control-api/issues">Request feature</a>
</p>

<br />

## Técnologias necessárias

- [Go](https://go.dev/)
- [PostreSQL](https://www.postgresql.org/)


## Testando localmente


Para será necessários ter instalado a linguagem [Golang](https://go.dev/dl/)

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


