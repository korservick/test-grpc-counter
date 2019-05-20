TARGET=grpc_counter

all: build docker-up

clean: docker-down
#	rm -rf $(TARGET)
	docker rmi korservick/$(TARGET):build
	docker rmi korservick/$(TARGET):compile
	docker rmi korservick/$(TARGET):latest


build: docker-down
#	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(TARGET) ./main.go
#	docker build . -t counter
#	docker build --target counter_build -t korservick/counter-build:latest .
#	docker build --target counter_compile -t korservick/counter-compile:latest .
#
#	docker build --cache-from counter-build --target counter-compile -t counter-compile .
#	docker build --target counter-compile --target counter -t counter .
#	docker build --target counter-compile -t counter .
#	docker build --cache-from counter-compile --target counter -t counter .
#	docker build --no-cache --cache-from counter-build -t counter .
	docker build -t korservick/$(TARGET):build -f env/Dockerfile.build .
	docker rmi -f korservick/$(TARGET):compile
	docker rmi -f korservick/$(TARGET):latest
	docker build -t korservick/$(TARGET):compile -f env/Dockerfile.compile .
	docker build -t korservick/$(TARGET):latest -f env/Dockerfile .

proto:
	protoc --go_out=plugins=grpc:. ./pb/*.proto

docker-up:
	docker-compose -f env/docker-compose.yml -p $(TARGET) up -d

docker-down:
	docker-compose -f env/docker-compose.yml -p $(TARGET) down