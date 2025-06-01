# tasks-api

Para executar o projeto será necessário ter o PostegreSQL

O Postegre pode ser obtido de duas fomas: através do docker ou baixando localmente

Para add no docker basta accessar a pasta app e executar o comando
```
docker compose up
```

Com o postegres localmente será necessario criar a tabela no BD

para isso basta rodar o script que se encontra na pasta app/infra/ToDo.sql

Após isso na pasta app execute o seguinte comando para levantar a aplicação

```
go run main.go
```

A collection do insomnia com as chamadas se encontra dentro da pasta app 