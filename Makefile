GOCC = go

CMD_DIR = ./cmd
PKG_DIR = ./pkg

main:
	$(GOCC) build -o v2utils -v $(CMD_DIR)/

test:
	$(GOCC) test -v $(PKG_DIR)
