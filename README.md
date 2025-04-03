# goAPI
# API de Produtos em Go

## Visão Geral
Esta é uma API RESTful desenvolvida em Go utilizando o framework Fiber e o ORM GORM para manipulação do banco de dados PostgreSQL. A API permite a criação, leitura, atualização e exclusão (CRUD) de produtos.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação utilizada para desenvolvimento da API.
- **Fiber**: Framework web inspirado no Express.js, otimizado para alto desempenho.
- **GORM**: ORM (Object-Relational Mapping) para interação com o banco de dados.
- **PostgreSQL**: Banco de dados relacional utilizado para armazenar os produtos.
- **Swagger**: Documentação interativa da API gerada automaticamente.

## Endpoints da API

| Método  | Rota             | Descrição                        |
|----------|----------------|--------------------------------|
| GET      | /products      | Lista todos os produtos       |
| GET      | /products/{id} | Obtém um produto por ID       |
| POST     | /products      | Cria um novo produto          |
| PUT      | /products/{id} | Atualiza um produto existente |
| DELETE   | /products/{id} | Remove um produto            |
| GET      | /swagger/*     | Acessa a documentação Swagger |

## Como Rodar o Projeto

1. Clone este repositório:
   ```sh
   git clone https://github.com/LuizVanzo/goAPI.git
   ```

2. Entre no diretório do projeto:
   ```sh
   cd Trab
   ```

3. Instale as dependências:
   ```sh
   go mod tidy
   ```

4. Configure o banco de dados no arquivo `main.go`:
   ```go
   dsn := "host=localhost user=postgres password=postgres dbname=products port=5432 sslmode=disable"
   ```

5. Execute a aplicação:
   ```sh
   go run main.go
   ```

6. Acesse a API em `http://localhost:3000`

## Comparativo com Outras Tecnologias e Tendências

### Conteúdos Estáticos e Renderização no Servidor
A API é puramente backend e não lida com renderização de conteúdo estático.

### Microsserviços
A arquitetura atual é monolítica

### Middleware de Aplicações Descentralizadas (Blockchain)
A API atualmente usa um banco de dados relacional centralizado. 

