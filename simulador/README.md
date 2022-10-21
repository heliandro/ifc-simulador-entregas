# Simulador de Entrega

Componente responsável por simular a criação de uma entrega.

## Build

Exemplo:

```bash
doocker build -t heliandro/ifc-simulador:v0.1.0 .
```

## Run

```bash
docker run --rm -e PORT=8001 -p 8001:8001 heliandro/ifc-simulador:<tag>
```

## Run Dev Mode

Rode o comando abaixo para builda e executar a api do simulador em modo de desenvolvimento.

```bash
docker-compose up --build -d
docker-compose exec simulador bash
go run main.go
```

## Endpoints

- GET /health

### GET /health

check api health

```bash
curl -i -X GET http://localhost:8001/health
```