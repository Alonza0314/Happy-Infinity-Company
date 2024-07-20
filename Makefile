RED=\033[;41m
GREEN=\033[;42m
YELLOW=\033[;43m
BLUE=\033[;44m

NC=\033[0m

F_BLUE=\033[;34m

BUILD_FILE_NAME=main.exe
BUILD_DIR=build
SOURCE_FILE=main.go
CONFIG_FILE=configs/config.conf
TODO="#TODO"

.PHONY: all
all: build

check-config:
	@echo "${BLUE}==> Checking configuration...${NC}"
	@if grep -q 'addr *= *${TODO}' ${CONFIG_FILE}; then \
		echo "${RED}Error: HICserver.addr is still set to #TODO in ${CONFIG_FILE}${NC}"; \
		echo "${RED}HINT: Set HICserver.addr as your own domain name in ${CONFIG_FILE}${NC}";\
		exit 1; \
	else \
		echo "${GREEN}==> Configuration is valid${NC}"; \
	fi

.PHONY: deps
deps:
	@echo "${BLUE}==> Installing dependencies...${NC}"
	@go mod tidy
	@echo "${GREEN}==> Dependencies installed!${NC}"

.PHONY: run
run: check-config
	@echo "${BLUE}==> Running application...${NC}"
	@go run ${SOURCE_FILE}

.PHONY: build
build: check-config
	@echo "${BLUE}==> Building...${NC}"
	@mkdir -p ${BUILD_DIR}
	@go build -o ${BUILD_DIR}/${BUILD_FILE_NAME} ${SOURCE_FILE}
	@echo "Execution file is under ${F_BLUE}'${BUILD_DIR}'${NC} directory"
	@echo "You can execute the file by this command: ${F_BLUE}./${BUILD_DIR}/${BUILD_FILE_NAME}${NC}"
	@echo "${GREEN}==> Build complete${NC}"

.PHONY: clean
clean:
	@echo "${BLUE}==> Cleaning up...${NC}"
	@rm -rf ${BUILD_DIR}
	@echo "${GREEN}==> Clean complete${NC}"
