ALL_PACKAGES=$(shell go list ./... | grep -v "vendor" )

assign-vars = $(if $(1),$(1),$(shell grep '$(2):' application.yml | tail -n 1| cut -d ':' -f 2 | sed 's/^\s*//'))
DB_HOST:=$(call assign-vars,$(DB_HOST),DB_HOST)
DB_NAME:=$(call assign-vars,$(DB_NAME),DB_NAME)
DB_USER:=$(call assign-vars,$(DB_USER),DB_USER)
DB_PASSWORD:=$(call assign-vars,$(DB_USER),DB_PASSWORD)
export PGPASSWORD=$(DB_PASSWORD)

setup:
	go get -u golang.org/x/lint/golint

vet:
	GO111MODULE=on go vet $(ALL_PACKAGES)

clean:
	GO111MODULE=on go clean

copy-config:
	cp -n application.yml.sample application.yml | true

db.create:
	createdb -h $(DB_HOST) -U $(DB_USER) -O $(DB_USER) -Eutf8 $(DB_NAME)

db.drop: copy-config
	dropdb -h $(DB_HOST) -U $(DB_USER) $(DB_NAME)

db.reset: db.drop db.create

test:
	go test $(ALL_PACKAGES) -p=1

dep:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor

lint:
	@for p in $(ALL_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

fmt:
	GO111MODULE=on go fmt $(ALL_PACKAGES)

ci: setup copy-config fmt lint vet db.create test