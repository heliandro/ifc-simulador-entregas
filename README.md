## Projeto

Este é um projeto de simulação de entregas. O objetivo é simular uma corrida/entrega (sem a feature de produto) e mostrar as entregas em tempo real no mapa.

Este projeto foi desenvolvido com base nos videos do canal do youtube Full Cycle.

Fiz melhorias no projeto original, como por exemplo, configurar o docker-compose para subir o broker kafka, zookeeper e o painel do kafka. Além de atualizar as tecnologias para as ultimas versões.

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
