SIMPLE_VERSION=v0.4.0
SIMPLE_SYSTEM=aarch64-linux-gnu-gcc-9
SIMPLE_FILE=libsimple.zip
BIN=blog

all: test

get_dict:
	@echo "获取分词插件与词典数据"
	@wget -O $(SIMPLE_FILE) https://github.com/wangfenjin/simple/releases/download/$(SIMPLE_VERSION)/libsimple-$(SIMPLE_SYSTEM).zip
	unzip -j $(SIMPLE_FILE) -d ./dict

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

build: build_f build_go
	@echo "编译 blog 可执行文件 打包前端静态文件 www.tar.gz"
	mkdir www
	cp -r ui/vanilla/dist www/v
	cp -r ui/admin/build www/404
	tar -zcvf www.tar.gz www
	rm -r www
	mkdir blog_sp
	cp www.tar.gz blog_sp
	cp $(BIN) blog_sp
	tar -zcvf blog_sp.tar.gz blog_sp
	rm -r blog_sp
