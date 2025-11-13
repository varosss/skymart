REGISTRY := vaross/private-projects
BUILD_DATE := $(shell date +%Y_%m_%d_%H_%M_%S)

SERVICES := auth-service ai-service api-gateway payments-service

# ============ COMMON TASKS ============
build.all: $(SERVICES:%=build.%)

build.%:
	@echo "🚀 Building service: $*"
	@docker build \
		-f docker/$*/Dockerfile . \
		-t $(REGISTRY):clirzy-$*-$(BUILD_DATE) \
		-t $(REGISTRY):clirzy-$*-latest \

clean:
	@docker rmi -f $(shell docker images -q $(REGISTRY):clirzy*) || true

# ============ PUSH TASKS ============
push.%:
	@echo "📤 Pushing $(REGISTRY):clirzy-$*-latest"
	@docker push $(REGISTRY):clirzy-$*-latest

push.all: $(SERVICES:%=push.%)
