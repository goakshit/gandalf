## Overview
Billing services for car parking slot using microservices

#### Services:
    - Database
    - Consumer
    - Billing
    - Zookeeper
    - Broker(Kafka)

#### Build (Docker)
```
git clone git@github.com:goakshit/gandalf.git
cd gandalf
docker-compose -f ./build/docker/docker-compose.yaml up -d --build
```
Docker-compose starts up the services mentioned above.
1. `Database(Postgres)`: Docker container boots up and executes the sql file in ./build/scripts/db which creates the table where data is persisted from kafka.
2. `Zookeeper`: Runs at port 2181.
3. `Broker(Kafka)`: Exposes 29092(Internal network) & 9092 (Host Network)
4. `Consumer`: Listens to kafka messages and persists the data in db.
5. `Billing`: Exposes two api(s) at port 80 -> **/api/duration/{id}** & **/api/cost/{id}**. Eg: curl -XGET http://localhost:80/api/duration/5
6. `client`: Pushes the vehicle information to Kafka. To exec: **go run ./cmd/client** 

#### Swagger Docs
To generate swagger json: 
Need to have the swagger CLI then `swagger generate spec -o ./swagger.json --scan-models`
or exec `make`

#### Tests
Mocks are added for persistence(DB layer). Tests are added for billing service. To exec: `go test ./internal/pkg/billing`
