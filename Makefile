# misc =========================================================================

.PHONY: generate-crypto
generate-crypto:
	$(MAKE) -C ./backend generate-crypto
	$(MAKE) -C ./nginx generate-crypto

# misc ^========================================================================

# frontend =====================================================================

.PHONY: bootstrap-frontend
bootstrap-frontend:
	$(MAKE) -C ./frontend bootstrap

.PHONY: lint-frontend
lint-frontend:
	$(MAKE) -C ./frontend lint

# frontend ^====================================================================

# backend ======================================================================

.PHONY: lint-backend
lint-backend:
	$(MAKE) -C ./backend lint

.PHONY: unit-test
unit-test:
	$(MAKE) -C ./backend unit-test

# backend ^=====================================================================

# integration ==================================================================

.PHONY: build
build:
	$(MAKE) -C ./backend build

.PHONY: build
bootstrap-ci: OUT := ./out
bootstrap-ci:
	$(MAKE) -C ./backend bootstrap-database

	$(MAKE) -C ./backend generate-crypto OUT=./$(OUT)
	PRIVATE_KEY=$(shell pwd)/backend/$(OUT)/pkcs8.key \
	PUBLIC_KEY=$(shell pwd)/backend/$(OUT)/publickey.crt \
		$(MAKE) -C ./backend bootstrap-backend ci=1
	PORT=10001 CONFIG_PATH=$(shell pwd)/backend/cube.yml build/backend/cube.out 2>&1 >/dev/null &

.PHONY: integration-test
integration-test:
	$(MAKE) -C ./backend integration-test

.PHONY: smoke-test
smoke-test:
	$(MAKE) -C ./backend smoke-test

.PHONY: load-test
load-test:
	$(MAKE) -C ./backend/test load-test

.PHONY: coverage
coverage:
	$(MAKE) -C ./backend coverage

.PHONY: cobertura
cobertura:
	$(MAKE) -C ./backend cobertura

# integration ^=================================================================

# dev ==========================================================================

.PHONY: dev-start
dev-start:
	docker compose -f ./docker/docker-compose.dev.yml up --build

# dev ^=========================================================================

# prod =========================================================================

.PHONY: prod-start
prod-start:
	$(MAKE) -C ./frontend build
	docker compose -f ./docker/docker-compose.prod.yml up --build -d

.PHONY: prod-stop
prod-stop:
	docker compose -f ./docker/docker-compose.prod.yml down

# prod ^========================================================================
