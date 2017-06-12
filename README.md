# HoneyBee
The Honey Bee DB behaves like a beehive. A key-value database built in GOLANG.

## How to run the database using docker:
### Prerequisites:
1. Golang
2. Docker + docker-compose

### Running the database using docker (run from the project root folder):
```
docker-compose up
```

## Testing:
### How to run tests:
```bash
go test github.com/talbor49/HoneyBee/tests
```

## How to use?
Clients at https://github.com/talbor49/HoneyBeeClient

## Things to remember while developing:
1. Compress data - save pointers to data, etc.
2. RAM is the cache, everything is saved to memory eventually
3. Distributing the DB into multiple machines. Split the data, split the tasks, synchronize.
4. Make it stable & durable - have replica, backup data, keep logs, avoid single point of failures.


TODO:
1. Properly test.
2. Properly documentate.
3. Properly log.
