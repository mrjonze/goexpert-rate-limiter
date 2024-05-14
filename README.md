# goexpert-rate-limiter
Resposta ao desafio técnico Rate Limiter da pós Go Expert.

Para rodar o banco de dados Redis, execute o comando abaixo:
```shell
docker-compose up -d
```

Para rodar o servidor da aplicação, execute o comando abaixo dentro da pasta ```server```:
```shell
go run main.go
```

Para rodar os testes unitários da aplicação, execute o comando abaixo dentro da pasta ```server```:
```shell
go test -v
```

O arquivo ```.env``` se encontra dentro da pasta ```server``` e contém as variáveis de ambiente necessárias para a aplicação.

```
TOKEN_NAME=Nome do token a ser passado no cabeçalho da requisição
REQUEST_LIMIT_TOKEN=Limite de requisições com o token
REQUEST_LIMIT_IP=Limite de requisições por IP
BLOCK_TIME_TOKEN=Tempo de bloqueio do token
BLOCK_TIME_IP=Tempo de bloqueio do IP
DATABASE_URL=Endereço do servidor Redis disponibilizado pelo docker-compose
```

Uma requisição simples pode ser realizada da seguinte maneira: 
```shell
curl http://localhost:8080
```

Uma requisição com token válido (em que o valor de ```API_KEY``` é igual ao da propriedade ```TOKEN_NAME``` do arquivo ```.env```) pode ser realizada da seguinte maneira:
```shell
curl -H "API_KEY:abc123" http://localhost:8080/
```

De maneira semelhante, uma requisição com token inválido (em que o valor de ```API_KEY``` é diferente do da propriedade ```TOKEN_NAME``` do arquivo ```.env```) pode ser realizada da seguinte maneira:
```shell
curl -H "API_KEY:invalidHeader" http://localhost:8080/
```
