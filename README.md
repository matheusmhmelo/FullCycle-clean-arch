# Desafio Clean Arch pós graduação Full Cycle

## Iniciando dependências do sistema
1. Na pasta raíz do projeto execute o comando para preparar o MySQL e RabbitMQ: \
```make up```
2. Após finalizada a inicialização das dependências, execute o comando para rodar as migrations e iniciar o sistema: \
```make init```

## Acessando o serviço
O serviço pode ser acessado da seguint forma:

- **HTTP**: Através da porta :8000

  - POST http://localhost:8000/order
  - GET http://localhost:8000/order

- **GraphQL**: Através da porta :8080
  - query listOrder
  - mutation creteOrder
- **GRPC**: Através da porta :50051
  - rpc CreateOrder
  - rpc ListOrder
