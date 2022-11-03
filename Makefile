run:
	go run main.go

test:
	go test -cover -race ./...

compose-up:
	docker-compose up -d

docker-exec:
	docker exec -it devicemanager /bin/bash

mockary:
	~/go/bin/mockery --all

create-volume:
	docker volume create --name=mysql_devicemanager_data

remove-volume:
	docker volume rm mysql_devicemanager_data
