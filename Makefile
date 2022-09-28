.PHONY: install
install:
	go build \
		-o  ~/.steampipe/plugins/local/clevercloud/clevercloud.plugin \
		main.go