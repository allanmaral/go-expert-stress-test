# Desafio Go - Stress Testing

Bem-vindo ao Desafio de Stress Testing da Pós-Graduação Go Expert! Este projeto consiste na implementação de uma CLI que realiza requisições em uma API de forma paralela e emite um relatório de requisições por segundo, tempo mínimo e máximo de respostas.

## Pré-requisitos

Antes de começar, certifique-se de ter instalado os seguintes requisitos:

- [Go SDK](https://golang.org/dl/): Linguagem de programação Go.
- [Docker](https://docs.docker.com/get-docker/): Plataforma de conteinerização.
- [Make](https://www.gnu.org/software/make/): Utilizado para automatização de tarefas.

## Executando o Projeto

1. Clone este repositório em sua máquina local:

```bash
git clone https://github.com/allanmaral/go-expert-stress-test.git
```

1. Navegue até o diretório do projeto:

```bash
cd go-expert-stress-test
```

1. Instale as dependências do projeto:

```bash
go mod tidy
```

1. Finalmente, rode a CLI com:

```bash
go run main.go --url https://google.com --requests 100 --concurrency 10
```

## Rodando a CLI usando Docker

Além de rodar a CLI usando a SDK do Go, também é possível roda usando somente o Docker, para isso, execute o comando:

```bash

docker run --rm allanmaral/stress:latest --url https://google.com --requests 100 --concurrency 10

```

Você deverá ver algo como:

```txt
Unable to find image 'allanmaral/stress:latest' locally
latest: Pulling from allanmaral/stress
6e415c5c724b: Pull complete
Digest: sha256:22782faba07511107e122aeff36e3489c459f0b7a1fa20e96a4248b8eb563c63
Status: Downloaded newer image for allanmaral/stress:latest
Stress testing https://google.com with 100 request(s) using 10 concurrent requests

Summary:
    Elapsed:             4.484128628s
    Count:               100
        200:             87
        302:             13
    Errors:
        Get "https://www.google.com/sorry/index?continue=https://google.com/&q=EgSxSnVsGISfhrIGIjDlzRMQzgci08DnDKk7EDGW8HE_Vs-DBJzKqX0_hgqtwt9aQViPgCkEan5C5HsulVYyBj5qY25kcloBQw": stopped after 10 redirects: 2
        Get "https://www.google.com/sorry/index?continue=https://google.com/&q=EgSxSnVsGIWfhrIGIjADiC9MJOKOCeQ_Rp9NKewqplQ41YBzYVHjLsfbX6knnJhmwlYtxWsO1lxK1Gy5VUAyBj5qY25kcloBQw": stopped after 10 redirects: 3
        Get "https://www.google.com/": stopped after 10 redirects: 8
    Requests per Second: 22.30

Statistics:
    Min:  249.095792ms
    Max:  1.213293668s
    Mean: 392.530075ms +/- 232.476005ms
```

## Estrutura do projeto

A maior parte da implementação do projeto foi feito no pacote `stresstest` na pasta `internal/stresstest`. O stresstest foi quebrado em alguns componentes:

- **Tester**: Objeto que cria uma go routine para cada para cada `concurrency` e executa as requisições.
- **Reporter**: Objeto que fica escutando o canal de `Result`s do `Tester` e agrega o resultado em um `Resport`.
- **Report**: Objeto com as estatísticas do teste e responsável por formatar o resultado.
