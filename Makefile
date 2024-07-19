BLUE=\033[;44m
GREEN=\033[;42m
NC=\033[0m

F_BLUE=\033[;34m

BUILD_FILE_NAME=main.exe
BUILD_DIR=build
SOURCE_FILE=main.go

.PHONY: all
all: build

.PHONY: run
run:
	@echo "${BLUE}==> Running application...${NC}"
	@go run ${SOURCE_FILE}

.PHONY: deps
deps:
	@echo "${BLUE}==> Installing dependencies...${NC}"
	@go mod tidy
	@echo "${GREEN}==> Dependencies installed!${NC}"

.PHONY: build
build:
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
