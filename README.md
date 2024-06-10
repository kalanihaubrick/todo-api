# TodoApi

TodoApi é uma aplicação web simples escrita em Go utilizando o framework Gin. Esta aplicação demonstra como criar uma API RESTful com conexão a um banco de dados SQL Server e como organizar o código em diferentes pacotes para melhor manutenção e escalabilidade.

## Estrutura do Projeto

ChatGPT
Claro! Aqui está um exemplo de um README.md bem formatado e detalhado para o seu projeto Go. Esse arquivo deve ser colocado na raiz do seu projeto.

markdown
Copiar código
# MyApp

MyApp é uma aplicação web simples escrita em Go utilizando o framework Gin. Esta aplicação demonstra como criar uma API RESTful com conexão a um banco de dados SQL Server e como organizar o código em diferentes pacotes para melhor manutenção e escalabilidade.

## Estrutura do Projeto
```go
todo-api/
├── api/
│ └── v1/
│ └── todo.go
├── cmd/
│ └── main.go
├── configs/
│ └── config.go
├── internal/
│ ├── controllers/
│ │ └── todo_controller.go
│ ├── models/
│ │ └── todo.go
│ ├── services/
│ │ └── todo_service.go
│ └── database/
│ └── database.go
├── pkg/
│ └── logger/
│ └── logger.go
├── go.mod
├── go.sum
└── README.md
```

## Instalação

1. **Clone o repositório:**

    ```sh
    git clone https://github.com/kalanihaubrick/todo-api.git
    cd todo-api
    ```

2. **Instale as dependências:**

    ```sh
    go mod tidy
    ```

3. **Configure o arquivo `.env`:**

    Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:

    ```env
    DB_USER=yourUserName
    DB_PASSWORD=yourPassword
    DB_HOST=yourHost
    DB_NAME=TodoDB
    PORT=8080
    ```

    Substitua os `yours` pelos seus próprios dados do SQL Server.

4. **Execute a aplicação:**

    ```sh
    go run cmd/myapp/main.go
    ```

## Estrutura do Código

### API

A pasta `api/v1` contém os handlers das rotas da versão 1 da API.

### Configs

A pasta `configs` contém a configuração do DSN para a conexão com o banco de dados.

### Internal

A pasta `internal` contém a lógica de negócio do aplicativo. Esta pasta é dividida em:

- **controllers**: Controladores que gerenciam as requisições HTTP.
- **models**: Definições de modelos de dados e migrações.
- **services**: Lógica de negócios.
- **database**: Configuração e inicialização da conexão com o banco de dados.

### Pkg

A pasta `pkg` contém pacotes reutilizáveis, como o logger configurado com `log/slog`.

## Logger

O logger é configurado para escrever logs tanto no terminal quanto em um arquivo chamado `todo-api.log`. Ele oferece três métodos principais: `Info`, `Error` e `Fatal`.

Exemplo de uso:

```go
logger.Info("This is an info message", "key", "value")
logger.Error("This is an error message", "key", "value")
logger.Fatal("This is a fatal message", "key", "value")
```

## Estrutura da API

*Endpoints*
- **GET /todos**: Lista todos os todos.
- **POST /todos**: Cria um novo todo.
- **GET /todo/:id**: Retorna um todo pelo ID.
- **PUT /todo/:id**: Atualiza um todo pelo ID.
- **DELETE /todo/:id**: Deleta um todo pelo ID.

*Exemplo de Request e Response*
**GET /todos**
Response:
```go
[
    {
        "id": 1,
        "task": "Buy groceries",
        "completed": false
    },
    {
        "id": 2,
        "task": "Read a book",
        "completed": true
    }
]
```

**POST /todos**
Request:
```go
{
    "task": "Walk the dog",
    "completed": false
}
```

Response:
```go
{
    "id": 3,
    "task": "Walk the dog",
    "completed": false
}
```
## Contribuindo

1. **Faça um fork do projeto**
2. **Crie uma branch para sua feature (git checkout -b feature/nova-feature)**
3. **Faça commit das suas alterações (git commit -am 'Adiciona nova feature')**
4. **Faça push para a branch (git push origin feature/nova-feature)**
5. **Crie um novo Pull Request**

## Contato
Criado por *Kalani Haubrick* - Sinta-se à vontade para me contatar!
