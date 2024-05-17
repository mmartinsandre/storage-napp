# Storage Napp API

Este projeto é uma API para gerenciar produtos em um estoque, utilizando o framework Gin para a criação da API e o Gorm para interações com o banco de dados PostgreSQL. A autenticação é feita usando JWT.

## Pré-requisitos

- Docker e Docker Compose instalados
- Go (Golang) instalado

## Estrutura do Projeto

```
storage-napp/
│
├── docker-compose.yml
├── go.mod
├── main.go
├── middleware/
│   └── middleware.go
└── handlers/
    └── handlers.go
```

## Configuração e Execução

### 1. Subir os Contêineres

Execute o seguinte comando para subir os contêineres definidos no `docker-compose.yml`:

```sh
docker-compose up -d
```

### 2. Configurar Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis de ambiente:

```
DB_HOST=db
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdb
```

### 3. Instalar Dependências

Na raiz do projeto, execute o comando para instalar as dependências do Go:

```sh
go mod tidy
```

### 4. Executar a Aplicação

Para rodar a aplicação, execute:

```sh
go run main.go
```

A aplicação estará rodando na porta `8080`.

## Endpoints

### Autenticação

#### Login

**Endpoint**: `/login`

**Método**: `POST`

**Corpo da Requisição**:

```json
{
    "username": "admin",
    "password": "password"
}
```

**Resposta**:

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Produtos

Todos os endpoints abaixo exigem o cabeçalho `Authorization` com o token JWT obtido no login.

**Cabeçalho**:

```
Authorization: Bearer <seu_token_jwt>
```

#### Criar Produto

**Endpoint**: `/produto`

**Método**: `POST`

**Corpo da Requisição**:

```json
{
    "codigo": "PROD001",
    "nome": "Produto Exemplo",
    "estoque_total": 100,
    "estoque_corte": 10,
    "preco_de": 50.0,
    "preco_por": 45.0
}
```

**Resposta**:

```json
{
    "ID": 1,
    "CreatedAt": "2024-05-17T12:34:56Z",
    "UpdatedAt": "2024-05-17T12:34:56Z",
    "DeletedAt": null,
    "codigo": "PROD001",
    "nome": "Produto Exemplo",
    "estoque_total": 100,
    "estoque_corte": 10,
    "estoque_disponivel": 90,
    "preco_de": 50.0,
    "preco_por": 45.0
}
```

#### Listar Todos os Produtos

**Endpoint**: `/produto`

**Método**: `GET`

**Resposta**:

```json
[
    {
        "ID": 1,
        "CreatedAt": "2024-05-17T12:34:56Z",
        "UpdatedAt": "2024-05-17T12:34:56Z",
        "DeletedAt": null,
        "codigo": "PROD001",
        "nome": "Produto Exemplo",
        "estoque_total": 100,
        "estoque_corte": 10,
        "estoque_disponivel": 90,
        "preco_de": 50.0,
        "preco_por": 45.0
    }
]
```

#### Consultar Produto por ID

**Endpoint**: `/produto/:id`

**Método**: `GET`

**Resposta**:

```json
{
    "ID": 1,
    "CreatedAt": "2024-05-17T12:34:56Z",
    "UpdatedAt": "2024-05-17T12:34:56Z",
    "DeletedAt": null,
    "codigo": "PROD001",
    "nome": "Produto Exemplo",
    "estoque_total": 100,
    "estoque_corte": 10,
    "estoque_disponivel": 90,
    "preco_de": 50.0,
    "preco_por": 45.0
}
```

#### Atualizar Produto

**Endpoint**: `/produto/:id`

**Método**: `PUT`

**Corpo da Requisição**:

```json
{
    "codigo": "PROD002",
    "nome": "Produto Atualizado",
    "estoque_total": 150,
    "estoque_corte": 20,
    "preco_de": 55.0,
    "preco_por": 50.0
}
```

**Resposta**:

```json
{
    "ID": 1,
    "CreatedAt": "2024-05-17T12:34:56Z",
    "UpdatedAt": "2024-05-17T13:45:67Z",
    "DeletedAt": null,
    "codigo": "PROD002",
    "nome": "Produto Atualizado",
    "estoque_total": 150,
    "estoque_corte": 20,
    "estoque_disponivel": 130,
    "preco_de": 55.0,
    "preco_por": 50.0
}
```

#### Deletar Produto

**Endpoint**: `/produto/:id`

**Método**: `DELETE`

**Resposta**:

```json
{}
```

## Observações Finais

- Certifique-se de que todas as dependências do projeto estejam instaladas corretamente antes de executar a aplicação.
- O banco de dados PostgreSQL deve estar em execução e acessível conforme as configurações definidas nas variáveis de ambiente.
- Utilize ferramentas como Postman para testar os endpoints e garantir que a autenticação JWT está funcionando corretamente.