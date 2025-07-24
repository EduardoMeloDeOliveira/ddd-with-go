# 🧪 Meu Primeiro Projeto em Go — API RESTful com Gin + GORM

Este é **meu primeiro projeto desenvolvido em Go**, e como alguém vindo do **mundo Java**, confesso que enfrentei um certo **estranhamento inicial** 😅.

Logo de cara, percebi que a estrutura da linguagem e os conceitos me remetiam bastante ao **JavaScript/TypeScript com Express**, principalmente na forma de organizar rotas, serviços e middlewares. Mas com o tempo fui entendendo melhor a filosofia do Go: **simples, direto ao ponto e com foco em performance**.

No fim das contas, **deu tudo certo!** 🎉  
Aprendi bastante sobre como estruturar um projeto limpo em Go, adotando boas práticas como:

- 🧱 **Separação por camadas** (domain, application, infrastructure, interfaces)
- 🎫 **Uso de DTOs**, **Mappers**, **Services**
- 📜 **Aplicação do padrão Repository**
- 🐘 **Integração com banco de dados usando GORM**
- 🌐 **Criação de uma API RESTful com Gin**

---


## 🚀 Tecnologias Usadas

- 🐹 **Go** — [golang.org](https://golang.org/)
- 🕸 **Gin** — [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- 🎩 **GORM** — [gorm.io](https://gorm.io/)
- 🐘 **PostgreSQL** — [postgresql.org](https://www.postgresql.org/)
- 🐳 **Docker & Docker Compose**

---

## 📦 Funcionalidades

- ➕ **Criar usuário**
- 📋 **Listar usuários**
- ✏️ **Atualizar email de um usuário**
- 🔒 **Validações de entrada com Gin**
- 🐘 **Persistência em PostgreSQL com GORM**
- 🧱 **Entidades com UUID**
- 🔄 **Conversão automática entre entidades e DTOs**



## 🏁 Como Rodar

### 1. **Subir os containers com Docker Compose**

Na raiz do projeto, rode:

```bash
docker compose up -d
```

Esse comando vai iniciar o banco de dados PostgreSQL e demais serviços necessários.

---

### 2. **Executar a aplicação Go**

Com os containers rodando, execute o comando:

```bash
go run ./project/cmd/api
```

A API estará disponível em **`http://localhost:8080`**.

---



## 📁 Estrutura de Pastas

```text
.
├── docker-compose.yml        # 🐳 Configuração do Docker Compose
├── go.mod                    # 📦 Dependências Go (Módulos)
├── go.sum                    # 📦 Hashes de dependências
├── project
│   ├── cmd
│   │   └── api
│   │       └── main.go       # ⚙️ Inicializador da aplicação
│   ├── internal
│   │   ├── application
│   │   │   ├── dto
│   │   │   │   └── user_dto.go
│   │   │   └── mapper
│   │   │       └── user_mapper.go
│   │   ├── domain
│   │   │   ├── entity
│   │   │   │   └── user.go
│   │   │   ├── repository
│   │   │   │   └── user_repository.go
│   │   │   └── service
│   │   │       └── user_service.go
│   │   ├── infrastructure
│   │   │   └── repository_impl
│   │   │       └── user_repository_impl.go
│   │   └── interface
│   │       └── http
│   │           ├── heathcheck_controller.go
│   │           └── user_controller.go
│   └── tmp
│       ├── build-errors.log
│       └── main
└── tmp
    └── build-errors.log
```

---


