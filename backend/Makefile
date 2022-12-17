PKG := neural_storage

# Для сборки с -race нужно иметь glibc не меньше 2.16 -- см. https://github.com/golang/go/issues/37485
ifeq ($(shell readlink /lib64/libc.so.6),libc-2.12.so)
TESTFLAGS:=
else
TESTFLAGS:=-race
endif


# lint =========================================================================

define lint_project
	golangci-lint run --build-tags="unit" ./$(1)/... && golangci-lint run --build-tags="unit" ./cmd/$(1)/...
endef

define lint_folder
	golangci-lint run --build-tags="testing" ./$(1)/...
endef

.PHONY: lint
lint: lint_cube lint_cache lint_config lint_database lint_pkg

.PHONY: lint_cube
lint_cube:
	$(call lint_project,cube)

.PHONY: lint_cache
lint_cache:
	$(call lint_folder,cache)

.PHONY: lint_config
lint_config:
	$(call lint_folder,config)

.PHONY: lint_database
lint_database:
	$(call lint_folder,database)

.PHONY: lint_pkg
lint_pkg:
	$(call lint_folder,pkg)

# lint ^========================================================================

# unit =========================================================================

.PHONY: unit-test
unit-test: OUT:=./out
unit-test:
	mkdir -p $(OUT)/unit
	go test -cover -coverprofile=$(OUT)/unit/coverage.out -covermode=atomic \
	-coverpkg=$(PKG)/... $(TESTFLAGS) --tags="unit" $(PKG)/...

.PHONY: coverage
coverage: OUT:=./out
coverage:
	go tool cover -html=$(OUT)/unit/coverage.out -o $(OUT)/unit/index.html


.PHONY: cobertura
cobertura: OUT:=./out
cobertura:
	go install github.com/t-yuki/gocover-cobertura@latest
	$(GOPATH)/bin/gocover-cobertura < $(OUT)/unit/coverage.out > $(OUT)/unit/coverage.xml

# unit ^========================================================================

# integration ==================================================================

.PHONY: bootstrap-database
bootstrap-database:
	$(MAKE) -C ./database bootstrap-database

.PHONY: bootstrap-backend
bootstrap-backend: ci :=
bootstrap-backend:
	$(MAKE) -C ./utils bootstrap
	python3 ./utils/j2.py ./cube/cube.yml.j2 $(if ci,./cube/vars-ci.yml,./cube/vars.yml) > cube.yml

.PHONY: build
build: OUT=./out
build:
	go build -o $(OUT)/cube.out ./cmd/cube/main.go

.PHONY: integration-test
integration-test:
	$(MAKE) -C ./test integration-bootstrap integration-test

.PHONY: smoke-test
smoke-test: OUT=./out/smoke-coverage.out
smoke-test: HTML=./out/smoke-coverage.html
smoke-test:
	docker compose -f ./docker/docker-compose.smoke.yml up --force-recreate -d || true
	go test $(TESTFLAGS) --tags="smoke" -coverprofile=$(OUT) $(PKG)/... || true
	docker compose -f ./docker/docker-compose.smoke.yml down || true
	go tool cover -html=$(OUT) -o $(HTML) || true
	go tool cover -func=$(OUT) || true

# integration ^=================================================================

# misc =========================================================================

.PHONY: generate
generate:
	go generate ./...
	cd cube && swag init -d ../cmd/cube/,./  -g ../../cube/handlers/fabric.go --parseDependency && cd ..

.PHONY: generate-crypto
generate-crypto: OUT:=./out/crypto
generate-crypto:
	@mkdir -p $(OUT)
	openssl genrsa -out $(OUT)/keypair.pem 2048
	openssl rsa -in $(OUT)/keypair.pem -pubout -out $(OUT)/publickey.crt
	openssl pkcs8 -topk8 -inform PEM -outform PEM -nocrypt -in $(OUT)/keypair.pem -out $(OUT)/pkcs8.key
	@rm $(OUT)/keypair.pem

# misc ^========================================================================
