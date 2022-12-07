# Coding Challenges

This Repository is the base for our challenges in the development application process. 
Make sure to check out our [open positions](https://fiskaly.com/jobs)

Challenges are intended for the following "skill-levels" and domains of expertise: 
- crm-challenge-ts: Junior / Mid Fullstack developer
- hex-editor-challenge-react-ts: Web developer all levels. Implementations are expected to vary for each skill-level
- signing-service-challenge-go: Mid / Senior Backend developer
- signing-service-challenge-ts: Mid / Senior Backend developer


## Curl Examples
```
curl localhost:8080/api/v0/device -d '{"label": "TestDevice", "algorithm": "RSA"}'
curl localhost:8080/api/v0/device/90683de8-1259-42e3-9adb-0d40d2d78108/sign -d '{"deviceId": "90683de8-1259-42e3-9adb-0d40d2d78108", "data": "lorem ipsum"}'
curl localhost:8080/api/v0/device/a0a58abe-acdf-459a-968b-dd442314731f/verify -d '{"deviceId": "a0a58abe-acdf-459a-968b-dd442314731f", "data": "0_lorem ipsum_YjZlNDc2OWQtYThlOS00ZDIyLWI4MDEtYTA2NTQxZDc1OGZi", "signature": "JPWBXscSJcY+l2CwDej/OWbTjClEt6qpN290L/PfnYjyDDjjlJqII52Ed/bUN6RS3ewOMhosv6u3IZYcplXcLg=="}'
```
