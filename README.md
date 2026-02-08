# Go Crash Course - Curso Rápido de Go

Curso de aprendizado rápido das principais funcionalidades da linguagem Go, ferramentas, boas práticas para desenvolver APIs prontas para ambiente produtivo.

## 📚 Índice

- [Pré-requisitos](#pré-requisitos)
- [Capítulos](#capítulos)
- [Projeto Prático](#projeto-prático)
- [Como Usar Este Repositório](#como-usar-este-repositório)

---

## 🎯 Pré-requisitos

- Conhecimento básico de programação
- Go 1.20 ou superior instalado
- Editor de código (VS Code, GoLand, ou similar)
- Git instalado

---

## 📖 Capítulos

### **Capítulo 1: Introdução ao Go**
- O que é Go?
- Por que usar Go?
- Instalação e configuração do ambiente
- Primeiro programa: Hello World
- Estrutura de um projeto Go
- Go Modules

### **Capítulo 2: Fundamentos da Linguagem**
- Tipos de dados básicos (int, float, string, bool)
- Variáveis e constantes
- Operadores
- Estruturas de controle (if, switch, for)
- Arrays e Slices
- Maps
- Structs

### **Capítulo 3: Funções e Métodos**
- Declaração de funções
- Múltiplos retornos
- Funções variádicas
- Funções anônimas e closures
- Defer, panic e recover
- Métodos em structs
- Interfaces

### **Capítulo 4: Concorrência**
- Goroutines
- Channels
- Select statement
- Sync package (WaitGroup, Mutex)
- Padrões de concorrência
- Context package

### **Capítulo 5: Tratamento de Erros**
- Filosofia de erros em Go
- Criação de erros customizados
- Error wrapping
- Boas práticas de tratamento de erros

### **Capítulo 6: Pacotes Importantes da Standard Library**
- fmt - Formatação de entrada/saída
- strings - Manipulação de strings
- time - Trabalho com datas e tempo
- encoding/json - Serialização JSON
- net/http - Cliente e servidor HTTP
- io - Operações de I/O
- os - Interação com sistema operacional

### **Capítulo 7: Desenvolvimento de APIs REST**
- Conceitos de REST
- Roteamento HTTP
- Handlers e middleware
- Manipulação de requisições e respostas
- Validação de dados
- Autenticação e autorização (JWT)
- Versionamento de APIs

### **Capítulo 8: Banco de Dados**
- Database/sql package
- Conexão com PostgreSQL
- CRUD operations
- Prepared statements
- Transações
- ORMs (GORM)
- Migrations

### **Capítulo 9: Testes**
- Testing package
- Testes unitários
- Testes de integração
- Table-driven tests
- Mocks e stubs
- Benchmarks
- Coverage

### **Capítulo 10: Boas Práticas e Ferramentas**
- Convenções de código Go
- Go fmt e Go vet
- Linters (golangci-lint)
- Documentação com godoc
- Estrutura de projetos
- Dependency management
- Logging estruturado

### **Capítulo 11: Performance e Otimização**
- Profiling (CPU, Memory)
- Benchmarking
- Otimização de código
- Gerenciamento de memória
- Garbage collector

### **Capítulo 12: Deploy e Produção**
- Build e compilação
- Variáveis de ambiente
- Configuração de aplicações
- Docker e containerização
- CI/CD
- Monitoramento e observabilidade
- Graceful shutdown
- Health checks

---

## 🚀 Projeto Prático

Este repositório inclui uma **API REST completa** como projeto prático, demonstrando:

- ✅ Estrutura de projeto profissional
- ✅ Endpoints CRUD completos
- ✅ Middleware (logging, recovery, CORS)
- ✅ Configuração via variáveis de ambiente
- ✅ Tratamento de erros robusto
- ✅ Validação de dados
- ✅ Documentação de API
- ✅ Boas práticas de código
- ✅ Pronto para produção

### Rodando a API

```bash
# Navegar para o diretório da API
cd api

# Instalar dependências
go mod download

# Rodar a aplicação
go run cmd/api/main.go

# Ou construir e executar
go build -o bin/api cmd/api/main.go
./bin/api
```

A API estará disponível em `http://localhost:8080`

### Endpoints Disponíveis

- `GET /health` - Health check
- `GET /api/v1/users` - Listar usuários
- `GET /api/v1/users/:id` - Buscar usuário
- `POST /api/v1/users` - Criar usuário
- `PUT /api/v1/users/:id` - Atualizar usuário
- `DELETE /api/v1/users/:id` - Deletar usuário

---

## 📝 Como Usar Este Repositório

1. **Clone o repositório**
   ```bash
   git clone https://github.com/yunesnoronha/go-crash-course.git
   cd go-crash-course
   ```

2. **Estude os capítulos sequencialmente**
   - Cada capítulo constrói sobre o anterior
   - Pratique os conceitos antes de avançar

3. **Explore o código da API**
   - Leia e entenda cada arquivo
   - Experimente modificar e adicionar funcionalidades
   - Execute e teste a API

4. **Pratique**
   - Crie seus próprios projetos
   - Implemente novos endpoints
   - Adicione funcionalidades à API base

---

## 📚 Recursos Adicionais

- [Documentação Oficial Go](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Wiki](https://github.com/golang/go/wiki)

---

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.

---

## 📄 Licença

Este projeto está sob a licença MIT.
