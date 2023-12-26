# Calculadora em Go - Projeto em GoLang

Este é um simples projeto de calculadora em Go que suporta quatro operações matemáticas básicas: soma, subtração, multiplicação e divisão. O projeto inclui um servidor HTTP que responde a solicitações para a realização destas operações.

## Funcionalidades

- **Soma:** `/calc/soma/{a}/{b}`
- **Subtração:** `/calc/sub/{a}/{b}`
- **Multiplicação:** `/calc/mul/{a}/{b}`
- **Divisão:** `/calc/div/{a}/{b}`
- **Histórico de Operações:** `/calc/historic`

## Pré-requisitos

Certifique-se de ter o seguinte instalado antes de executar o projeto:

- [Go](https://golang.org/) - Linguagem de programação Go.
- [GoLand](https://www.jetbrains.com/go/) - IDE oficial da JetBrains para Go ou alguma outra IDE de sua preferência.

## Configuração e Execução

1. **Clone o Repositório:**
   ```bash
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
   ```


## Como Testar a Aplicação:

Abra um navegador web de sua preferência e acesse os seguintes URLs para realizar operações matemáticas:

* Soma: http://localhost:8080/calc/soma/2/3
* Subtração: http://localhost:8080/calc/sub/5/2
* Multiplicação: http://localhost:8080/calc/mul/4/6
* Divisão: http://localhost:8080/calc/div/8/2

* Para visualizar o histórico de operações, acesse: http://localhost:8080/calc/history