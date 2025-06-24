# Documentação do Projeto LibGen (pedro21ribeiro.com)

Este documento fornece uma visão geral do projeto LibGen, incluindo sua arquitetura, componentes, configuração, e instruções de execução.

## Visão Geral

O projeto LibGen é uma aplicação web para gerenciamento de uma biblioteca, permitindo a criação e visualização de livros. A aplicação utiliza o framework Echo para lidar com as requisições HTTP e renderizar as páginas web. O banco de dados PostgreSQL é usado para persistir os dados, e a biblioteca GORM é utilizada como ORM.

## Arquitetura

A aplicação segue uma arquitetura MVC (Model-View-Controller):

*   **Model:** Representa os dados da aplicação. Os modelos estão localizados no diretório `models`. Por exemplo, o arquivo `models/book/book.go` define a estrutura de um livro.
*   **View:** Responsável por apresentar os dados ao usuário. As views são templates HTML localizados no diretório `views`. A biblioteca `html/template` do Go é utilizada para renderizar as views.
*   **Controller:** Lida com as requisições do usuário, interage com os models para buscar ou persistir dados, e escolhe a view apropriada para retornar ao usuário. Os controllers estão localizados no diretório `controllers`.

Adicionalmente, o projeto possui as seguintes camadas:

*   **DTO (Data Transfer Object):** Objetos simples utilizados para transferir dados entre as camadas da aplicação. Eles estão localizados no diretório `dtos`.
*   **Service:** Contém a lógica de negócios da aplicação. Os services estão localizados no diretório `services`.
*   **Repository:** Responsável por interagir com o banco de dados. Os repositórios estão localizados no diretório `repositories`.
*   **SecretLoader:** Responsável por carregar as variáveis de ambiente do arquivo `.env`. Os secrets loaders estão localizados no diretório `secretLoader`.

## Componentes Principais

*   **main.go:** Ponto de entrada da aplicação. Inicializa o servidor Echo, configura os middlewares, conecta ao banco de dados e define as rotas.
*   **controllers/:** Contém os controllers que lidam com as requisições HTTP.
    *   **book/bookController.go:** Controller responsável pelas funcionalidades relacionadas a livros (criar, visualizar).
    *   **user/userController.go:** Controller responsável pelas funcionalidades relacionadas a usuários (login, registro).
    *   **test/testController.go:** Controller para testes
*   **dbConector/dbConector.go:** Contém a lógica para conectar ao banco de dados PostgreSQL usando a biblioteca GORM.
*   **dtos/:** Contém os DTOs usados na aplicação.
    *   **book/bookDto.go:** Define a estrutura de dados para transferência de informações relacionadas a livros.
*   **models/:** Define as estruturas de dados (models) que representam as entidades do domínio (e.g., User, Book).
    *   **book/book.go:** Define a estrutura do modelo de livro.
    *   **user/user.go:** Define a estrutura do modelo de usuário.
*   **repositories/:** Contém os repositórios que abstraem o acesso ao banco de dados.
    *   **book/bookRepository.go:** Responsável por interagir com o banco de dados para operações relacionadas a livros.
    *   **user/userRepository.go:** Responsável por interagir com o banco de dados para operações relacionadas a usuários.
*   **services/:** Contém a lógica de negócios da aplicação.
    *   **book/bookService.go:** Contém a lógica para criação e validação de livros.
*   **views/:** Contém os templates HTML que são renderizados para gerar a interface do usuário.
    *   **book_landing.html:** Template para a página principal de livros.
    *   **book_form.html:** Template para o formulário de livros.
    *   **user.html:** Template para a página de usuário/login.
    *   **login_form.html:** Template para o formulário de login.
    *   **form_error.html:** Template para exibir erros de formulário.
*   **static/:** Contém os arquivos estáticos da aplicação (CSS, JavaScript).
    *   **css/:** Arquivos CSS para estilização da interface do usuário.
        *   **book.css:** Estilos para a página de livros.
        *   **login.css:** Estilos para a página de login.
    *   **js/:** Arquivos JavaScript para adicionar interatividade à interface do usuário.
        *   **login.js:** Scripts para a página de login.
*   **.env:** Arquivo que contém as variáveis de ambiente da aplicação, como as credenciais do banco de dados.  **IMPORTANTE**: Este arquivo não deve ser versionado e deve ser mantido em segredo.
*   **go.mod:** Arquivo que gerencia as dependências do projeto Go.

## Configuração

A aplicação utiliza variáveis de ambiente para configurar a conexão com o banco de dados e outras configurações. As variáveis de ambiente são definidas no arquivo `.env`.

### Variáveis de Ambiente Necessárias

*   `DB_HOST`: Host do banco de dados PostgreSQL.
*   `DB_PORT`: Porta do banco de dados PostgreSQL.
*   `DB_NAME`: Nome do banco de dados PostgreSQL.
*   `DB_USER`: Usuário do banco de dados PostgreSQL.
*   `DB_PASSWORD`: Senha do usuário do banco de dados PostgreSQL.
*   `SALT`: Valor de "sal" para a criptografia.
*   `SECRET`: Segredo usado para assinar os tokens JWT (JSON Web Tokens).

### Exemplo de arquivo .env

```
DB_HOST=postgres.meu.host.com
DB_PORT=5432
DB_NAME=meu_db
DB_USER=meu_usuario
DB_PASSWORD=senha_segura
SALT=salt_das_senhas
SECRET=segredo_de_minimamente_256_bits
```

**IMPORTANTE:** Substitua os valores de exemplo pelas suas credenciais reais do banco de dados.

## Execução

Para executar a aplicação, siga os seguintes passos:

1.  **Pré-requisitos:**
    *   Go instalado (versão 1.20 ou superior).
    *   PostgreSQL instalado e configurado.
    *   GORM instalado e configurado.
2.  **Clonar o repositório:**

    ```
    git clone [URL do seu repositório]
    cd [nome do diretório do repositório]
    ```

3.  **Instalar as dependências:**

    ```
    go mod download
    go mod tidy
    ```

4.  **Configurar o banco de dados:**
    *   Crie um banco de dados PostgreSQL com o nome definido na variável `DB_NAME` do arquivo `.env`.
    *   Crie as tabelas necessárias. Você pode usar as definições de modelo no diretório `models` para criar as tabelas.
    *   A aplicação cria a tabela "books" baseado no model, utilizando a convenção de nomes do GORM. As colunas no banco de dados seguem o camelCase, convertendo "ReleaseDate" para "data_lancamento".
5.  **Definir as variáveis de ambiente:**
    *   Crie um arquivo `.env` na raiz do projeto.
    *   Adicione as variáveis de ambiente necessárias no arquivo `.env`, conforme descrito na seção "Configuração".
6.  **Executar a aplicação:**

    ```
    go run main.go
    ```

    A aplicação será iniciada e estará disponível no endereço `http://localhost:42069`.

## Usando o server.exe

Se o `server.exe` foi previamente compilado, a aplicação pode ser executada diretamente através dele.

1.  **Certifique-se de que o `server.exe` está na raiz do projeto (ou em um diretório acessível).**

2. **Garanta que o arquivo .env esta na mesma pasta do server.exe.**

3.  **Abra um terminal ou prompt de comando.**

4.  **Navegue até o diretório onde o `server.exe` está localizado (se necessário).**

5.  **Execute o `server.exe`:**

    ```
    ./server.exe
    ```

    Ou (no Windows):

    ```
    server.exe
    ```

    A aplicação será iniciada, e a saída do console indicará que o servidor está rodando.

## Rotas da Aplicação

*   **GET /user:** Renderiza a página de login.
*   **POST /user/login:** Autentica o usuário e retorna um token JWT.
*   **GET /book:** Renderiza a página principal de livros (requer autenticação JWT).
*   **POST /book:** Cria um novo livro (requer autenticação JWT).
*   **GET /test:** Rota para testes.

## Segurança

A aplicação utiliza JWT (JSON Web Tokens) para autenticação. O middleware `echojwt` é usado para proteger as rotas que requerem autenticação.

**IMPORTANTE:** Mantenha o valor da variável de ambiente `SECRET` em segredo. Este valor é usado para assinar os tokens JWT.

## Próximos Passos

*   Implementar a funcionalidade de edição e exclusão de livros.
*   Adicionar paginação à listagem de livros.
*   Implementar testes unitários e de integração.
*   Melhorar a interface do usuário.
*   Adicionar mais validações de dados no backend.