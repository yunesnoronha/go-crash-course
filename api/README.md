# API REST com Go - Documentação

Esta é uma API REST base construída com Go usando apenas a biblioteca padrão, demonstrando boas práticas para ambiente produtivo.

## 📋 Estrutura do Projeto

```
api/
├── cmd/
│   └── api/
│       └── main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── config/
│   │   └── config.go            # Gerenciamento de configuração
│   ├── handlers/
│   │   ├── health_handler.go   # Health check endpoint
│   │   └── user_handler.go     # Handlers de usuários
│   ├── middleware/
│   │   ├── cors.go              # CORS middleware
│   │   ├── logger.go            # Logging middleware
│   │   └── recovery.go          # Recovery middleware (panic)
│   └── models/
│       └── user.go              # Modelo de usuário
├── pkg/
│   └── response/
│       └── response.go          # Funções auxiliares de resposta
├── .env.example                 # Exemplo de variáveis de ambiente
├── go.mod                       # Dependências do módulo
└── README.md                    # Esta documentação
```

## 🚀 Como Executar

### Pré-requisitos

- Go 1.20 ou superior

### Instalação

1. Navegue até o diretório da API:
   ```bash
   cd api
   ```

2. Instale as dependências:
   ```bash
   go mod download
   ```

3. (Opcional) Configure variáveis de ambiente:
   ```bash
   cp .env.example .env
   ```

### Executar

```bash
# Modo desenvolvimento (com hot reload usando air, se instalado)
go run cmd/api/main.go

# Ou construir e executar
go build -o bin/api cmd/api/main.go
./bin/api
```

A API estará disponível em: `http://localhost:8080`

## 📡 Endpoints

### Health Check

```
GET /health
```

**Resposta de Sucesso (200 OK):**
```json
{
  "status": "ok",
  "timestamp": "2024-01-15T10:30:00Z",
  "service": "go-crash-course-api"
}
```

---

### Listar Usuários

```
GET /api/v1/users
```

**Resposta de Sucesso (200 OK):**
```json
[
  {
    "id": "20240115103000",
    "name": "João Silva",
    "email": "joao@example.com",
    "age": 30,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
]
```

---

### Buscar Usuário

```
GET /api/v1/users/:id
```

**Resposta de Sucesso (200 OK):**
```json
{
  "id": "20240115103000",
  "name": "João Silva",
  "email": "joao@example.com",
  "age": 30,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

**Resposta de Erro (404 Not Found):**
```json
{
  "error": "Not Found",
  "message": "usuário não encontrado"
}
```

---

### Criar Usuário

```
POST /api/v1/users
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "João Silva",
  "email": "joao@example.com",
  "age": 30
}
```

**Resposta de Sucesso (201 Created):**
```json
{
  "id": "20240115103000",
  "name": "João Silva",
  "email": "joao@example.com",
  "age": 30,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

**Resposta de Erro (400 Bad Request):**
```json
{
  "error": "Bad Request",
  "message": "nome é obrigatório"
}
```

---

### Atualizar Usuário

```
PUT /api/v1/users/:id
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "João Silva Santos",
  "email": "joao.santos@example.com",
  "age": 31
}
```

**Resposta de Sucesso (200 OK):**
```json
{
  "id": "20240115103000",
  "name": "João Silva Santos",
  "email": "joao.santos@example.com",
  "age": 31,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:35:00Z"
}
```

---

### Deletar Usuário

```
DELETE /api/v1/users/:id
```

**Resposta de Sucesso (200 OK):**
```json
{
  "message": "Usuário deletado com sucesso"
}
```

---

## 🏗️ Arquitetura e Boas Práticas

### Estrutura de Diretórios

- **cmd/**: Contém os pontos de entrada da aplicação
- **internal/**: Código privado da aplicação (não pode ser importado por outros projetos)
- **pkg/**: Código público que pode ser reutilizado

### Middleware

A aplicação usa três middlewares principais:

1. **Logger**: Registra todas as requisições HTTP
2. **Recovery**: Captura panics e retorna erro 500
3. **CORS**: Permite requisições cross-origin

### Configuração

As configurações são gerenciadas através de variáveis de ambiente:

- `PORT`: Porta do servidor (padrão: 8080)
- `ENVIRONMENT`: Ambiente de execução (development/production)

### Tratamento de Erros

Todos os erros são retornados em formato JSON consistente:

```json
{
  "error": "Status Text",
  "message": "Mensagem descritiva do erro"
}
```

### Graceful Shutdown

A aplicação implementa graceful shutdown, aguardando até 10 segundos para:
- Finalizar requisições em andamento
- Fechar conexões
- Liberar recursos

## 🧪 Testando a API

### Usando curl

```bash
# Health check
curl http://localhost:8080/health

# Listar usuários
curl http://localhost:8080/api/v1/users

# Criar usuário
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"João Silva","email":"joao@example.com","age":30}'

# Buscar usuário
curl http://localhost:8080/api/v1/users/20240115103000

# Atualizar usuário
curl -X PUT http://localhost:8080/api/v1/users/20240115103000 \
  -H "Content-Type: application/json" \
  -d '{"name":"João Silva Santos","email":"joao.santos@example.com","age":31}'

# Deletar usuário
curl -X DELETE http://localhost:8080/api/v1/users/20240115103000
```

## 🔒 Segurança

### Implementado

- ✅ Timeouts de conexão configurados
- ✅ Recovery de panics
- ✅ Validação de entrada
- ✅ CORS configurado
- ✅ Graceful shutdown

### Próximos Passos (Produção)

- [ ] Autenticação JWT
- [ ] Rate limiting
- [ ] HTTPS/TLS
- [ ] Validação mais robusta
- [ ] Sanitização de entrada
- [ ] Logging estruturado
- [ ] Métricas e observabilidade

## 📦 Build e Deploy

### Build

```bash
# Build para o sistema atual
go build -o bin/api cmd/api/main.go

# Build otimizado
go build -ldflags="-s -w" -o bin/api cmd/api/main.go

# Cross-compile para Linux
GOOS=linux GOARCH=amd64 go build -o bin/api-linux cmd/api/main.go
```

### Docker

```dockerfile
# Exemplo de Dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o api cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["./api"]
```

## 📚 Recursos

- [Documentação Go](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Standard Library](https://pkg.go.dev/std)

## 🤝 Contribuindo

Contribuições são bem-vindas! Este é um projeto educacional.

## 📝 Notas

- Esta é uma implementação de referência para fins educacionais
- Para produção, considere usar frameworks como Gin, Echo ou Chi
- A persistência de dados está em memória (use banco de dados em produção)
- IDs são gerados com timestamp (use UUID em produção)
