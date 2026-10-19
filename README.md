# 📌 Encurtador de URL (Go + Chi + SQLite)

Um encurtador de URL simples, rápido e funcional desenvolvido em **Go**, utilizando:

- **Chi Router**
- **SQLite (modernc.org/sqlite — sem CGO)**
- **Arquitetura limpa (handlers, service, repository)**
- **Página HTML integrada**
- **Rate Limit por IP**
- **Logs estruturados**

Ideal para estudos, portfólio ou como base para projetos maiores.

---

## 🚀 Funcionalidades

- Encurta URLs longas em códigos curtos
- Redireciona automaticamente para a URL original
- Página HTML simples para testes
- Rate limit por IP (proteção contra abuso)
- Banco SQLite criado automaticamente
- Estrutura modular e escalável

---

## 📁 Estrutura do Projeto

```
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── http/
│   │   ├── handlers/
│   │   │   ├── shorten_handler.go
│   │   │   ├── redirect_handler.go
│   │   │   └── static_handler.go
│   │   ├── middleware/
│   │   │   └── rate_limit.go
│   │   └── router.go
│   ├── model/
│   │   └── url.go
│   ├── repository/
│   │   └── url_repository.go
│   └── service/
│       └── url_service.go
├── pkg/
│   └── logger/
│       └── logger.go
├── public/
│   └── index.html
├── go.mod
└── README.md
```

---

## ⚙️ Instalação

### 1. Clone o repositório

```
git clone https://github.com/seu-usuario/encurtador-url-go.git
cd encurtador-url-go
```

### 2. Instale dependências
```
go mod tidy
```

### 3. Execute o servidor
```
go run cmd/api/main.go
```
O servidor iniciará em:

```
http://localhost:8080
```

## 🧪 Testando a API
### 🔹 1. Encurtar uma URL
POST /shorten
```
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://google.com"}'
```

Resposta:

json
```
{
  "short_code": "AbC123"
}
```
### 🔹 2. Redirecionar
Acesse no navegador:

```
http://localhost:8080/AbC123
```

Você será redirecionado para a URL original.

## 🌐 Página HTML
Acesse:

```
http://localhost:8080/
```

A página permite:

Digitar uma URL

Encurtar

Clicar no link gerado

Testar o redirecionamento

## 🛡 Rate Limit
O middleware limita cada IP a 10 requisições por minuto, evitando abuso no endpoint /shorten.

## 🗄 Banco de Dados
O projeto usa:

```
modernc.org/sqlite
```
Driver 100% Go, sem CGO, compatível com Windows, Linux e Mac.

O arquivo do banco é criado automaticamente:

```
data.db
```

## 👩‍💻 Autora
Projeto desenvolvido por Gisele de Melo em caráter de estudo.

## 📄 Licença
MIT — sinta-se livre para usar e modificar.
