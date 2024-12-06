SIMPLE_VERSION := v0.4.0
SIMPLE_OS := linux-ubuntu-20.04 osx-x64 windows-x64
GOOS := linux darwin windows
BIN := main
GOARCH := amd64
BUILD_DIR := ./dist

# pacman -S mingw-w64-gcc 用于交叉编译 windows cgo
.PHONY: build
build: init_dict
	$(foreach os, $(GOOS), \
		$(eval s_os := $(word $(shell echo $(os) | awk '{print NR}'), $(SIMPLE_OS))) \
		$(eval blog_dir := $(BUILD_DIR)/blog-$(os)) \
		$(eval bin_file := $(if $(filter windows,$(os)), $(BIN).exe, $(BIN))) \
		$(eval cc := $(if $(filter windows,$(os)), CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1, )) \
		mkdir -p $(blog_dir)/www; \
		$(cc) GOOS=$(os) GOARCH=$(GOARCH) go build -tags fts5 -o $(bin_file); \
		mv $(bin_file) $(blog_dir); \
		unzip -j $(BUILD_DIR)/libsimple-$(s_os).zip -d $(blog_dir)/dict; \
		cp -r ui/vanilla/dist $(blog_dir)/www/v; \
		cp -r ui/admin/build $(blog_dir)/www/404; \
		cp app.toml $(blog_dir); \
		(cd $(BUILD_DIR) && tar -zcvf blog-$(os).tar.gz blog-$(os)); \
		rm -r ./$(blog_dir); \
	)

$(BUILD_DIR):
	mkdir $(BUILD_DIR)

SIMPLE_FILES := $(foreach os, $(SIMPLE_OS), $(BUILD_DIR)/libsimple-$(os).zip)

init_dict: $(BUILD_DIR) $(SIMPLE_FILES)

$(SIMPLE_FILES):
	@echo "下载分词插件 $@"
	wget -O $@ https://github.com/wangfenjin/simple/releases/download/$(SIMPLE_VERSION)/$(basename $(notdir $@)).zip

init:
	@echo "安装依赖"
	(go mod tidy)
	(cd ui/vanilla && npm install)
	(cd ui/admin && npm install)

build_vanilla:
	@echo "编译 vanilla 静态文件"
	(cd ui/vanilla && npm run build)

build_admin:
	@echo "编译 admin 静态文件"
	(cd ui/admin && npm run build-pro)

build_f: build_vanilla build_admin

build_admin_gh:
	@echo "编译 github pages 静态文件"
	(cd ui/admin && npm run build-pages)

build_go_embed:
	@echo "编译附带前端 embed 的可执行文件"
	go build -tags fts5,embed -o $(BIN)

build_go:
	@echo "编译可执行文件"
	go build -tags fts5 -o $(BIN)

test: build_vanilla
	@echo "测试 vanilla"
	go run -tags fts5,embed main.go serve

docs-dev:
	@echo "测试文档"
	(cd docs && npm run docs:dev)
