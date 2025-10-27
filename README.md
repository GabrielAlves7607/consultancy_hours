# consultancy_hours
project for technical development in GO


* **`models/schedule.go`** (O "O QUÊ")
    * **Função:** Define a **estrutura de dados**.
    * **Descrição:** É o "molde" ou "blueprint". Ele apenas diz quais campos um agendamento (`Schedule`) tem, como `IDTime` e `CustomerName`. Ele não sabe salvar no banco nem falar com a web.

* **`services/schedule_service.go`** (O "COMO FAZER")
    * **Função:** Contém a **lógica de negócios** e o **acesso ao banco de dados**.
    * **Descrição:** É o "cérebro" da aplicação. É este arquivo que sabe como se conectar no MongoDB, como buscar todos os agendamentos (`GetAllSchedules`) e como inserir um novo (`CreateSchedule`).

* **`controllers/schedule_controller.go`** (O "ATENDENTE")
    * **Função:** Lida com as **requisições HTTP** (web).
    * **Descrição:** É o "garçom" ou "atendente". Ele recebe o pedido do cliente (a requisição HTTP), vai até o `service` ("cérebro") para pedir os dados ou salvar algo, e depois formata a resposta (o JSON) para devolver ao cliente. Ele é a ponte entre a web e a sua lógica.