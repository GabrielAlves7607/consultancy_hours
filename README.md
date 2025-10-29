# Projeto de Agendamento de Consultoria em Go

Este é um projeto de backend em Go (Golang) para gerenciar o agendamento de horários de consultoria.

O projeto utiliza um banco de dados MongoDB para persistir os dados e é organizado em uma **arquitetura em camadas** para separar as responsabilidades:

* **`controllers`**: Lida com as requisições HTTP e respostas.
* **`services`**: Contém a lógica de negócios e regras de validação.
* **`models`**: Define as estruturas de dados (structs).
* **`constants`**: Armazena mensagens de erro e valores constantes.

## 🚀 Funcionalidades Principais

* **Listar Horários Disponíveis**: Consulta o banco e retorna os slots de horário que ainda não foram agendados.
* **Criar Agendamento**: Permite que um cliente agende um horário vago.
* **Validação de Regras**:
    * Impede o agendamento de horários que já foram reservados.
    * Impede o agendamento de horários fora do formato válido (ex: "12:30" ou "99:00").

## 🛠️ Tecnologias Utilizadas

* **Go (Golang)**
* **MongoDB** (com o driver `go.mongodb.org/mongo-driver`)
* **`net/http`** (biblioteca padrão do Go para o servidor)
* **`godotenv`** (para carregar variáveis de ambiente)

## 📂 Estrutura do Projeto

```
consultancy_hours/
│
├── constants/
│   └── constants.go        # Mensagens de erro e valores constantes
│
├── controllers/
│   └── schedule_controller.go # Camada de Apresentação (HTTP, JSON)
│
├── models/
│   └── schedule.go         # Camada de Dados (Structs)
│
├── services/
│   └── schedule_service.go # Camada de Negócios (Lógica e interação com DB)
│
├── .env                    # Arquivo de configuração (Exemplo)
├── .gitignore
├── go.mod
├── go.sum
├── main.go                 # Ponto de entrada: Conexão com DB, Injeção de Dependência
└── README.md
```

## ⚙️ Configuração e Instalação

### Pré-requisitos

* Go (Versão 1.18 ou superior)
* MongoDB (uma instância local ou um cluster no Atlas)

### Passos para Executar

1.  **Clone o repositório:**

    ```bash
    git clone [URL_DO_SEU_REPOSITORIO]
    cd consultancy_hours
    ```

2.  **Instale as dependências:**

    ```bash
    go mod tidy
    ```

3.  **Crie o arquivo de ambiente:**
    Crie um arquivo chamado `.env` na raiz do projeto e adicione sua string de conexão do MongoDB:

    ```env
    # .env
    MONGO_URI="mongodb://seu_usuario:sua_senha@localhost:27017/agenda?authSource=admin"
    ```

    *Substitua pela sua string de conexão correta. O banco de dados usado no código é `agenda`.*

4.  **Execute o projeto:**

    ```bash
    go run main.go
    ```

    O servidor será iniciado na porta `8080`.

    ```
    Connect to the MongoDB!
    Iniciando servidor na porta 8080...
    ```

## 📖 Como Usar (Endpoints HTTP)

Após iniciar o projeto, você pode interagir com ele através das seguintes rotas:

### 1\. Consultar Horários Disponíveis

Retorna uma lista de todos os slots de horário que **não** possuem agendamento.

* **Rota**: `GET /horarios/disponiveis`
* **Exemplo (curl + jq):**
  ```bash
  curl http://localhost:8080/hours/available | jq .
  ```
* **Exemplo de Resposta (`200 OK`):**
  ```json
  [
    {
      "id_horario": "00:00",
      "status": "disponivel"
    },
    {
      "id_horario": "01:00",
      "status": "disponivel"
    },
    // ... (etc) ...
  ]
  ```

### 2\. Agendar um Horário

Cria um novo agendamento para um cliente.

* **Rota**: `POST /agendar`
* **Corpo da Requisição (JSON):**
  ```json
  {
    "id_horario": "HH:00",
    "nome_cliente": "Nome do Cliente"
  }
  ```
* **Exemplo (curl):**
  ```bash
  curl -X POST http://localhost:8080/toSchedule \
       -H "Content-Type: application/json" \
       -d '{"id_horario": "14:00", "nome_cliente": "Lais"}'
  ```

#### Respostas Possíveis

* **Sucesso (`201 Created`):**

  ```json
  {
    "message": "Shedule completed successfully!",
    "id_inserido": "6724f7e6e5a4f216a9a067a3"
  }
  ```

* **Erro - Horário Ocupado (`409 Conflict`):**

  ```json
  {
    "erro": "This time is already booked"
  }
  ```

* **Erro - Horário Inválido (`400 Bad Request`):**

  ```json
  {
    "erro": "The id_horario sent is not a valid slot"
  }
  ```