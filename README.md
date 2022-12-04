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
curl localhost:8080/api/v0/device -d '{"label": "TestDevice", "algorithm": "ECC"}'
curl localhost:8080/api/v0/device/e804c3be-fd82-4ef5-ba41-fc69b2ff8f10/sign -d '{"deviceId": "e804c3be-fd82-4ef5-ba41-fc69b2ff8f10", "data": "lorem ipsum"}'
curl localhost:8080/api/v0/device/e804c3be-fd82-4ef5-ba41-fc69b2ff8f10/verify -d '{"deviceId": "e804c3be-fd82-4ef5-ba41-fc69b2ff8f10", "data": "0_lorem ipsum_MTJhODVmNGUtMTg4ZS00ZjIwLWFkMGMtYjJhMTIyMDIwNTAx", "signature": "MGQCMBEvznpvlCKc5U8+bjAiwCTFniEMnVPfqcfeHgEtnw0srtR3cEfGUTVOgR4cB3dgZQIwQIKqXp7P1T9b1i+SDNGpcRlN4/8ao6Cu0nRkoXU6O4m9mYjEnioq2iQrahUnc4a7"}'
```
