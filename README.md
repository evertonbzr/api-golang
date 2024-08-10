# API de Gestão de Biblioteca em Golang

Esta API foi desenvolvida em Golang para gerenciar uma biblioteca, permitindo o cadastro de usuários, a gestão de livros e o controle de empréstimos. O sistema inclui autenticação e autorização, com diferentes níveis de acesso (Administrador e Usuário Regular).

## Componentes do Sistema

### 1. Usuário

- **Atributos**: ID, Nome, E-mail, Data de Registro.
- **Operações**:
  - Registrar novo usuário
  - Listar todos os usuários

### 2. Livro

- **Atributos**: ID, Título, Autor, Status (Disponível/Emprestado).
- **Operações**:
  - Adicionar novo livro
  - Editar dados do livro
  - Listar todos os livros

### 3. Empréstimo

- **Atributos**: ID, ID do Usuário, ID do Livro, Data de Empréstimo, Data de Devolução, Status (Emprestado/Devolvido).
- **Operações**:
  - Registrar novo empréstimo
  - Registrar devolução
  - Listar empréstimos ativos

### 4. Autenticação e Autorização

- Sistema de gerenciamento de autenticação, incluindo diferentes níveis de acesso (Administrador e Usuário Regular).

## Como Usar

1. **Clonar o repositório**
   ```bash
   git clone https://github.com/seu-usuario/nome-do-repositorio.git
   ```
2. **Instalar as dependências**

   ```bash
   go mod download
   ```

3. **Configurar o ambiente**

   - Configure as variáveis de ambiente necessárias no arquivo `.env`.

4. **Executar a aplicação**
   ```bash
   go run main.go
   ```

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests para melhorias e correções.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
