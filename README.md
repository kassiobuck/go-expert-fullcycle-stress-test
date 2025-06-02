# Stress Test

Este projeto realiza testes de estresse em URLs utilizando Go e Docker.

## Pré-requisitos

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)

## Build da imagem Docker

```bash
docker build -t stress-test .
```

## Utilização

Execute o container Docker com os parâmetros desejados:

```bash
docker run stress-test --url=http://google.com --requests=1000 --concurrency=10
```
*Substitua os valores conforme necessário para sua aplicação.*



#### Retorno esperado (similar):
```bash ===== Relatório de Teste de Carga =====
Tempo total gasto: 50.394619923s
Total de requests realizados: 1000
Requests com status 200: 1000
Distribuição dos códigos de status HTTP:
Status 200: 1000 
```
