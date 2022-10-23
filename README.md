# Go DDD CQRS Backend

proyecto de practica para aplicar DDD y CQRS

### Comandos
start

    docker-compose up -d --build

stop

    docker-compose down

### REST

`http://localhost:4000`

Sign In con nick `root` y contra `1234`

```shell
curl --location --request POST 'http://localhost:4000/auth/signIn' \
--header 'Content-Type: application/json' \
--data-raw '{
    "nick": "root",
    "password": "1234"
}'
```

Sign Out con refresh token

```shell
curl --location --request POST 'http://localhost:4000/auth/signOut' \
--header 'Content-Type: application/json' \
--data-raw '{
  "refreshToken": ${REFRESH TOKEN}
}'
```

Access con refresh token y conseguir access token

```shell
curl --location --request POST 'http://localhost:4000/auth/access' \
--header 'Content-Type: application/json' \
--data-raw '{
  "refreshToken": ${REFRESH TOKEN}
}'
```

### WEBSOCKET

`http://localhost:4001`

Enviar el siguiente mensaje a websocket para subscribirse a los cambios
```json
{
    "action": "subscribe",
    "event": EVENT_TYPE
}
```
y para desubscribir
```json
{
    "action": "unsubscribe",
    "event": EVENT_TYPE
}
```
evento recibido
```json
{
    "action": "event",
    "event": EVENT_TYPE,
    "data": EVENT_DATA
}
```

**Tipos de eventos**:

- event.auth.signIn
- event.auth.signOut
- event.auth.access