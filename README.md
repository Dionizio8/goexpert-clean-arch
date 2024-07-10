# Desafio Clean Arch

## Descrição

Criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

## Como executar
*Será necessário criar a fila **orders** no RabbitMQ para realizar os testes local.*

### Configurando 

1. Inicializando **Mysql** & **RabbitMQ** com *docker-compose*
```bash
docker-compose up -d 
```

3. No [RabbitMQ](http://localhost:15672/]) criar em **Queues and Streams** a fila com nome *orders* e o Binding *amq.direct*

### Start aplicação
```bash
cd cmd/ordersystem
```
```bash
go run main.go wire_gen.go
```
```
Starting web server on port :8000
Starting gRPC server on port 50051
Starting GraphQL server on port 8080
```
