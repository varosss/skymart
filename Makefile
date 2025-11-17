REGISTRY := vaross/private-projects
BUILD_DATE := $(shell date +%Y_%m_%d_%H_%M_%S)

SERVICES := auth user ai api-gateway payment

# ------------------------------------------------
# Help command
# ------------------------------------------------
help:
	@echo ""
	@echo "Available make commands:"
	@echo ""
	@echo "  build.all                 Build all services"
	@echo "  build.<service>           Build a specific service"
	@echo "  clean                     Remove all built images"
	@echo "  push.all                  Push all services to registry"
	@echo "  push.<service>            Push a specific service to registry"
	@echo "  docker.migrate.<service>.up    Run migrations UP for a service"
	@echo "  docker.migrate.<service>.down  Run migrations DOWN for a service"
	@echo "  migrate.<service>.create name=<migration_name>   Create new migration for a service"
	@echo ""
	@echo "Example usage:"
	@echo "  make build.auth-service"
	@echo "  make push.ai-service"
	@echo "  make docker.migrate.auth-service.up"
	@echo "  make migrate.auth-service.create name=create_users_table"
	@echo ""


# ------------------------------------------------
# Common commands
# ------------------------------------------------
build.all: $(SERVICES:%=build.%)

build.%:
	@echo "🚀 Building service: $*"
	@docker build \
		-f ./services/$*/Dockerfile . \
		-t $(REGISTRY):clirzy-$*-$(BUILD_DATE) \
		-t $(REGISTRY):clirzy-$*-latest \

clean:
	@docker rmi -f $(shell docker images -q $(REGISTRY):clirzy*) || true


# ------------------------------------------------
# Push commands
# ------------------------------------------------
push.%:
	@echo "📤 Pushing $(REGISTRY):clirzy-$*-latest"
	@docker push $(REGISTRY):clirzy-$*-latest

push.all: $(SERVICES:%=push.%)


# ------------------------------------------------
# Migration commands
# ------------------------------------------------
docker.migrate.%.up:
	@echo "⬆️ Running migrations UP for service: $*"
	MSYS_NO_PATHCONV=1 docker compose exec $* sh -c 'migrate -path "$$MIGRATIONS_PATH" -database "$$POSTGRES_URL" up'

docker.migrate.%.down:
	@echo "⬇️ Running migrations DOWN for service: $*"
	MSYS_NO_PATHCONV=1 docker compose exec $* sh -c 'migrate -path "$$MIGRATIONS_PATH" -database "$$POSTGRES_URL" down'

migrate.%.create:
	@echo "🛠 Creating migration for $* with name: $(name)"
	MSYS_NO_PATHCONV=1 migrate create -ext sql -dir "./$*/migrations" $(name)
