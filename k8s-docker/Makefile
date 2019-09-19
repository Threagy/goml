.DEFAULT_GOAL := help

bash: ## Bash에 접속합니다.
	@./scripts/bash-select.sh

build: ## 이미지를 빌드합니다.
	@./scripts/build-select.sh

up: ## 컨테이너를 실행합니다.
	@./scripts/up-select.sh

down: ## 컨테이너를 삭제합니다.
	@./scripts/down-select.sh

# 배포
.PHONY: dist
dist: ## 모든 서비스의 배포할 파일을 생성합니다.
	@./scripts/dist.sh

push: ## 모든 서비스의 Production 이미지를 Push합니다.
	./scripts/push.sh ${REPO} ${SVC} ${VER}

# Etc
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'