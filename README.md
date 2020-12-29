
### Go com MongoDB
1. Através do seu terminal, navegue até o diretório ```chapter06/mongodb/example```
2. Execute ```sudo docker-compose up -d```
3. Crie duas variáveis de ambiente ```export MONGODBUSERNAME=root``` e ```export MONGODBPASSWORD=example```
4. Execute ```go run main.go```

O aplicativo deve retornar algo como:
 ```go
State: mongodb.State{Name:"Washington", Population:7062000}
State: mongodb.State{Name:"Oregon", Population:3970000}
 ```
 Para verificar a gravação dos registros no MongoDB, siga os passos abaixo:

1. Através do Terminal, execute ```sudo docker exec -it go-mongodb bash```
2. No bash do container, execute ```mongo -u root -p example```
3. Após o login, execute ```show dbs```

```mongo
admin        0.000GB
config       0.000GB
gocoockbook  0.000GB
local        0.000GB
```
4. Defina o banco de dados, ```use gocoockbook```
5. Liste as coleções disponíveis, ```show collections```

```mongo
example
```
6.Busque todos os dados da coleção **example**, ```db.example.find()```

```mongo
{ "_id" : ObjectId("5feb702d59dbbd90ffd84fd0"), "name" : "Washington", "pop" : 7062000 }
{ "_id" : ObjectId("5feb702d59dbbd90ffd84fd1"), "name" : "Oregon", "pop" : 3970000 }
``` 