build:
	go build -o /bin/device

run:
	./bin/device

clean:
	rm -rf /bin/device

test:
	go test -v ./...

docker-build:
	docker build -t device .

docker-run:
	docker run -p 8080:8080 device

test-get:
	curl -X GET http://localhost:8080/v1/devices

test-post:
	curl -X POST http://localhost:8080/v1/devices -d '{"name": "device1", "brand": "type1"}' -H "Content-Type: application/json"

test-put:
	curl -X PUT http://localhost:8080/v1/devices/3572bed8-b9a7-421a-93e6-3f206e55555f -d '{"name": "device1", "brand": "type2"}' -H "Content-Type: application/json"

test-delete:
	curl -X DELETE http://localhost:8080/v1/devices/1
