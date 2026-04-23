# Wallet Experience - EDA Microservices

Este projeto é um ecossistema de microserviços focado em processamento de transações financeiras e atualização de saldos em tempo real, utilizando uma arquitetura orientada a eventos (**Event Driven Architecture**).

## 🏗️ Arquitetura do Sistema

O projeto é dividido em dois serviços principais que se comunicam de forma assíncrona:

1. **Wallet Core (Go)**: Gerencia clientes, contas e o fluxo de transações. Após validar e persistir uma transação, ele dispara eventos para o Apache Kafka.
2. **Balance Core (Go)**: Consome eventos de atualização de saldo do Kafka e mantém um banco de dados de leitura rápida para consulta de saldos atualizados.

## 🛠️ Stack Tecnológica

- **Linguagem:** Go (Golang)
- **Mensageria:** Apache Kafka & Confluent Control Center
- **Banco de Dados:** MySQL 5.7 (Instâncias separadas para cada microserviço)
- **Containerização:** Docker & Docker Compose
- **Padrões de Projeto:** Unit of Work (Uow), Repository Pattern, Clean Architecture, DTOs

## 🚀 Como Executar

### Pré-requisitos

- Docker e Docker Compose instalados
- Go 1.26+ (caso queira rodar os testes localmente).

### Passo a Passo

1. **Subir os containers:**

    ```bash
    docker compose up -d --build
    ```

2. **Popular o banco inicial (Wallet):**

    Para testar as transações, é necessário criar clientes e contas no `db-wallet`:

    ```bash
    docker exec -it db-wallet mysql -uroot -proot wallet -e "INSERT INTO clients (id, name, email, created_at) VALUES ('1', 'User 1', 'user1@test.com', NOW()), ('2', 'User 2', 'user2@test.com', NOW()); INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('123', '1', 1000, NOW()), ('456', '2', 1000, NOW());"
    ```

3. **Monitorar o Kafka:**

    Acesse o **Confluent Control Center** em: `http://localhost:9021`.

### Comandos Adicionais

1. **Reiniciar os containers:**

    ```bash
    docker compose down -v
    docker compose up -d --build
    ```

2. **Conferir tabelas criadas:**

    ```bash
    docker exec -it db-wallet mysql -uroot -proot wallet -e "SHOW TABLES;"
    docker exec -it db-balance mysql -uroot -proot balances -e "SHOW TABLES;"
    ```

3. **Logs:**

    ```bash
    docker logs wallet-core
    docker logs balance-core
    ```

## 🔌 Endpoints Principais

### Wallet-Core (Porta 8080)

```http
### criar cliente
POST http://localhost:8080/clients HTTP/1.1
Content-Type: application/json
{
    "name": "John Doe",
    "email": "j@j.com"
}

### criar conta
POST http://localhost:8080/accounts HTTP/1.1
Content-Type: application/json
{
    "client_id": "[UUID]"
}

### criar transação
POST http://localhost:8080/transactions HTTP/1.1
Content-Type: application/json
{
    "account_id_from": "[UUID]",
    "account_id_to": "[UUID]",
    "amount": 5
}
```

### Balance-Core (Porta 8081)

```http
### buscar saldo consolidado
GET http://localhost:8081/accounts/{{account_id}}
```

## 🧪 Testes

Para rodar os testes unitários do UseCase e garantir que a lógica de negócio e os mocks estão íntegros:

```bash
go test ./...
```

## 📝 Notas de Implementação

- **Unit of Work:** Implementado para garantir a atomicidade entre a gravação da transação e o update de saldo no banco do Wallet.
- **Idempotência:** O Balance-Core foi desenhado para processar os eventos garantindo que o estado final do saldo seja consistente mesmo em caso de reprocessamento de mensagens.

## Considerações Finais

Este projeto faz parte do desafio de Microserviços da Full Cycle, demonstrando proficiência em sistemas distribuídos, consistência eventual e alta disponibilidade.