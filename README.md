## Projeto

Projeto de simulação de entregas. 

O objetivo é simular uma corrida/entrega (sem a feature de produto) e mostra-las em tempo real no mapa.

Nota:
A base desse projeto é o simulador de entregas, projeto este que foi desenvolvido na Imersão Full Cycle v10 do canal [Full Cycle](https://www.youtube.com/c/FullCycle).

Esse projeto tem modificações significativas se comparado ao original, como por exemplo, a infraestrutura, bibliotecas diferentes, versoes atualizadas das linguagens e frameworks, etc. Resumindo, é uma recriação do projeto original.

## Arquitetura

todo

## Tecnologias

- NodeJS
- Typescript
- Go
- React
- Kafka
- Docker
- Google Maps

## Rodando o projeto

```bash
docker-compose up -d
```

para buildar as imagens do docker e subir

```bash
docker-compose up --build
```

## Endpoints

- Painel do Kafka: http://localhost:9021
- Aplicacao Frontend: http://localhost:3000
- Aplicacao Backend: http://localhost:8000
- Aplicacao Simulador: http://localhost:8001

### API Simulador

base url: http://localhost:8001

- GET /health

```bash
curl -i -X GET http://localhost:8001/health
```

### API Backend

base url: http://localhost:8000

- GET /health

```bash
curl -i -X GET http://localhost:8000/health
```

## Topicos

- delivery

### Topic delivery

Para criar o tópico delivery, execute o comando abaixo:

```bash
docker-compose exec kafka kafka-topics \
    --create \
    --topic delivery \
    --bootstrap-server localhost:9092
```

Para publicar no topico

```bash
docker-compose exec kafka bash -c "echo '{ \"id\": \"1\", \"name\": \"Heliandro\", \"lat\": \"-23.5505\", \"lng\": \"-46.6333\" }' | kafka-console-producer --topic delivery --broker-list localhost:9092"
```

Para consumir o topico

```bash
docker-compose exec kafka kafka-console-consumer \
    --bootstrap-server localhost:9092 --topic delivery \
    --from-beginning
```

Curiosidade: para logar no Kafka por console
    
```bash
docker-compose exec kafka bash
```
