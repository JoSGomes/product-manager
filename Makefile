db-up:
	docker compose up -d

db-down:
	docker compose down

run-localhost:
	export SERVER_CONTEXT=product-manager
	export SERVER_HOST=localhost
	export SERVER_PORT=8080
	export DATABASE_USERNAME=products_admin
	export DATABASE_PASSWORD=123qwe
	export DATABASE_HOST=localhost
	export DATABASE_PORT=5432
	export DATABASE_NAME=products_api

	go run ./cmd/products-manager-server