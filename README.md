Event Driven Microservice
===
Abstract:
Event-driven architecture (EDA) is a software architecture paradigm promoting the production, detection, consumption of, and reaction to events.

![](https://github.com/faozimipa/micro/blob/development/_img/image.png)
## Install & Dependence
- Docker
- golang 

## Use
- Build services
  ```
  docker compose up
  ```
- Setup Keyloack
  ```
  go to administration console
  localhost:8080/auth

  create client in master realm. ex: openid-micro
  set Client Protocol to ``openid-connect``
  set Access Type to ``confidential``

  on Credential tabs set Client Authenticator to ``Client id and secret key``
  genrate secret key 
  on Keys tab set Use JWKS URL to on
  set JWKS URL to ``http://localhost:8080/auth/realms/master/protocol/openid-connect/certs``
  
  save configuration 
  ```
- Setup Services Identity
  ```
  set KEY_CLIENT_ID and KEY_CLIENT_SECRET .env on ``services.identity/src/app.env``
  rebuild Identity service
  ```

## Directory Hierarchy
```
|—— .gitignore
|—— _deploy_to_k8s
|    |—— Dockerfile
|    |—— README.MD
|    |—— deployment.yml
|    |—— krakend.json
|    |—— service.yml
|—— _file_server
|    |—— jwk
|        |—— symmetric.json
|—— _grafana
|    |—— Dockerfile
|    |—— dashboards
|        |—— all.yml
|    |—— datasources
|        |—— all.yaml
|    |—— krakend
|        |—— dashboard.json
|—— _img
|    |—— image.png
|—— _pacts
|    |—— order-service-customer-service.json
|—— _postman
|—— _script
|    |—— stop_all_win.bat
|    |—— wait-for-it.sh
|—— api_gateway
|    |—— krakend.json
|—— db_all.sql
|—— docker-compose-infra-only.yml
|—— docker-compose.yml
|—— keyloack
|    |—— Dockerfile
|—— services.customer
|    |—— Dockerfile
|    |—— README.MD
|    |—— go.mod
|    |—— go.sum
|    |—— src
|        |—— app.env
|        |—— entity
|            |—— basket.go
|            |—— basket_item.go
|            |—— customer.go
|            |—— product.go
|        |—— event
|            |—— order_created.go
|        |—— event_handler
|            |—— order_completed_handler.go
|            |—— product_created_handler.go
|            |—— user_created_handler.go
|        |—— internal
|            |—— api.go
|            |—— handler.go
|            |—— repository.go
|            |—— service.go
|            |—— verify_contract_test.go
|        |—— kafka
|            |—— consumer.go
|        |—— main.go
|—— services.identity
|    |—— Dockerfile
|    |—— README.MD
|    |—— go.mod
|    |—— go.sum
|    |—— src
|        |—— app.env
|        |—— entity
|            |—— user.go
|        |—— event
|            |—— user_created.go
|        |—— internal
|            |—— api.go
|            |—— handler.go
|            |—— repository.go
|            |—— service.go
|        |—— jwt
|            |—— jwt.go
|        |—— main.go
|—— services.notification
|    |—— Dockerfile
|    |—— README.MD
|    |—— go.mod
|    |—— go.sum
|    |—— src
|        |—— app.env
|        |—— kafka
|            |—— consumer.go
|        |—— main.go
|—— services.order
|    |—— Dockerfile
|    |—— README.MD
|    |—— go.mod
|    |—— go.sum
|    |—— src
|        |—— app.env
|        |—— dto
|            |—— basket_item_dto.go
|        |—— entity
|            |—— order.go
|            |—— order_item.go
|        |—— event
|            |—— order_created.go
|        |—— event_handler
|            |—— product_reserve_failed_handler.go
|            |—— product_reserved_handler.go
|        |—— http_client
|            |—— customer_contract_test.go
|            |—— customer_http_client.go
|        |—— internal
|            |—— api.go
|            |—— handler.go
|            |—— repository.go
|            |—— service.go
|        |—— kafka
|            |—— consumer.go
|        |—— main.go
|—— services.product
|    |—— Dockerfile
|    |—— README.MD
|    |—— go.mod
|    |—— go.sum
|    |—— src
|        |—— app.env
|        |—— entity
|            |—— product.go
|        |—— event
|            |—— order_created.go
|            |—— product_created.go
|        |—— event_handler
|            |—— order_created_handler.go
|        |—— internal
|            |—— api.go
|            |—— handler.go
|            |—— repository.go
|            |—— service.go
|        |—— kafka
|            |—— consumer.go
|        |—— main.go
|—— shared
|    |—— config
|        |—— config.go
|    |—— go.mod
|    |—— go.sum
|    |—— kafka
|        |—— kafka.go
|    |—— keyloack
|        |—— keyloack.go
|    |—— server
|        |—— server.go
```
## Code Details
## References
- [Golang](https://go.dev)
- [Gorm](https://github.com/go-gorm/gorm)
- [Gin](https://github.com/gin-gonic/gin)
- [Krakend](https://github.com/devopsfaith/krakend)
- [Keyloack](https://github.com/keycloak/keycloak)
- [Kafka - Zookeeper]()
- [Grafana](https://grafana.com)
- [Influxdb]()
- [Jaeger](https://www.jaegertracing.io)
  
## License

