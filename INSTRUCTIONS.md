# 🚀 DESAFIO EDA - Microsserviço de Balances (Consumidor Kafka)

## 🎯 Objetivo

Consolidar meu conhecimento sobre **Arquitetura Orientada a Eventos (EDA)**. Devo desenvolver um novo microsserviço (em qualquer linguagem de programação) que atue como **consumidor**.

Ele deve:
- 📥 Ler os eventos de transações gerados pelo microsserviço **"Wallet Core"** (via Kafka)
- 💰 Atualizar o saldo das contas
- 🌐 Disponibilizar uma consulta via API

## 🛠️ Tecnologias e Ferramentas

- 💻 **Linguagem:** Livre (escolha a que preferir) → Go.
- 📨 **Mensageria:** Apache Kafka.
- 🐳 **Infraestrutura:** Docker e Docker Compose.

## ⚙️ Requisitos Técnicos

O Microsserviço **"Balances"** deve:
- 🔌 Se conectar ao Kafka e escutar os tópicos onde o **"Wallet Core"** publica as transações.
- 🗄️ Persistir o saldo atualizado de cada conta em um banco de dados próprio.

## 🌐 API de Consulta

- [X] Criar um endpoint `GET /balances/{account_id}`
- [X] Este endpoint deve retornar o saldo atual da conta solicitada 💰
- [X] A aplicação deve rodar na **porta 3003** 🚪

## 🐳 Requisitos de Dockerização (Automação Total)

Para a correção, preciso que todo o ecossistema suba com um único comando.

O arquivo `docker-compose.yaml` deve orquestrar:

1. 🌍 **O Ecossistema Completo:**
    - Deve subir:
        - 📨 Kafka
        - 🧠 Microsserviço **Wallet Core** (produtor)
        - ⚖️ Microsserviço **Balances** (consumidor)
        - 🗄️ Os respectivos bancos de dados

2. ⚡ **Inicialização Automática:**

    Ao rodar `docker compose up -d`. Deve acontecer automaticamente:
    - 🧱 Criação das tabelas (migrations)
    - 🌱 Inserção de dados fictícios (seeds)
    - ✅ Ambiente pronto para testes imediatos
    
    **🚫 Não deve ser necessário rodar scripts manuais após o container subir.**

## 📄 Arquivos Auxiliares

- 📬 **api.http:** Incluir um arquivo `.http` (ou coleção do Postman exportada) com as chamadas prontas para testar o microsserviço.

## 📦 Entregável

1. 🔗 **Link do Repositório:** O link para o repositório do GitHub contendo todo o projeto (Wallet Core + Balances Service configurados no compose).

2. 📘 **README:** Instruções de execução.

## 📏 Regras de Entrega

1. 📁 **Repositório Único:** O projeto completo deve estar em um único repositório.
2. 🌿 **Branch Principal:** Todo o código deve estar na branch `main`.
3. 🎯 **Foco no Desafio:** 
    - Arquitetura (fluxo de eventos funcionando) 🔄
    - Automação com Docker 🐳
    - (Não é foco principal: Clean Code / Patterns)