## Sistema de Gestão de Estoque

Um sistema web de gestão de estoque desenvolvido em **Golang**, utilizando o framework **Gin**, para ajudar a gerenciar categorias, produtos, status e controlar a quantidade de itens no estoque.

## Funcionalidades

- **Categorias**: Classifique os produtos em categorias específicas.
- **Produtos**: Gerencie informações detalhadas de cada produto.
- **Controle**: Acompanhe a quantidade disponível de cada produto no estoque.
- **Status**: Identifique se os produtos estão disponíveis ou indisponíveis.

## Estrutura do Projeto

- **Categoria**: Representa a categoria a que o produto pertence (ex.: Eletrônicos, Alimentos, Roupas).
- **Produto**: Representa o produto, contendo informações como nome, descrição, preço, etc.
- **Controle**: Gerencia a quantidade de produtos disponíveis no estoque.
- **Status**: Define o status de disponibilidade do produto (disponível ou não).

## Tecnologias Utilizadas

- **Golang**: Linguagem principal para desenvolvimento do backend.
- **Gin**: Framework para criação de APIs REST em Golang.
- **Docker** e **Docker Compose**: Para containerizar a aplicação.
- **SQL Server**: Banco de dados utilizado para armazenar informações.
- **Postman**: utilizado para testar as solicitação HTTP

## Como Rodar o Projeto

### Pré-requisitos
Certifique-se de que você tem as seguintes ferramentas instaladas:
- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)
- Um banco de dados SQL Server configurado.
- Ferramenta para testar as rotas (recomendo o [Postman](https://www.postman.com/downloads/))

### Passos

1. Clone este repositório:
   ```bash
   git clone https://github.com/DaviFernandes034/SGE--gestao-de-estoque.git
   cd SGE--gestao-de-estoque

2. Configure seu arquivo .env com as configurações do seu banco de dados
   ```bash
      DB_HOST=localhost
      DB_PORT=1433
      DB_USER=seu_usuario
      DB_PASSWORD=sua_senha
      DB_NAME=nome_do_banco

3. Configure o arquivo docker-compose.yml
    - Certifique-se de que o arquivo já está configurado para orquestrar:
        - Um container para o banco de dados SQL Server.   
        - Um container para a aplicação em Golang.
        - 
4. Construa e inicialize os Conteiners:
      ```bash
       docker-compose up --build
5. Acesse a aplicação:
      ```bash
      http://localhost:8080
6. Verifique os logs:
      ```bash
         docker-compose logs -f
