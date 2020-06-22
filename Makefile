
BINARY=go-mysql-api
test:
	go test -v -cover -covermode=atomic ./...

build:
	go build -o ${BINARY} main.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t go-mysql-api .

run:
	docker-compose -f ./docker/docker-compose.yml up --build -d

stop:
	docker-compose -f ./docker/docker-compose.yml down

.PHONY: test docker run stop build make