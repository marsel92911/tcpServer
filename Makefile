TCP_SERVER_TARGET_BITS?=18
TCP_SERVER_PORT?=12345

# to set TCP_SERVER_PORT or TCP_SERVER_TARGET_BITS values use `make run-server -e TCP_SERVER_PORT="12345"`
.PHONY: run-server
run-server: build-server
	docker run  \
		-e TCP_SERVER_TARGET_BITS=$(TCP_SERVER_TARGET_BITS) \
		-p $(TCP_SERVER_PORT):$(TCP_SERVER_PORT) \
		--rm server

.PHONY: run-client
run-client: build-client
	docker run  \
		--network host \
		--rm client

.PHONY: build-client
build-client:
	cd ./client && \
	docker build -t client .

.PHONY: build-server
build-server:
	cd ./server && \
	docker build -t server .

.PHONY: delete-client-image
delete-client-image:
	docker rmi client

.PHONY: delete-server-image
delete-server-image:
	docker rmi server
