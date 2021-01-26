.PHONY: tao tao-cross evm all test clean
.PHONY: tao-linux tao-linux-386 tao-linux-amd64 tao-linux-mips64 tao-linux-mips64le
.PHONY: tao-darwin tao-darwin-386 tao-darwin-amd64

GOBIN = $(shell pwd)/build/bin
GOFMT = gofmt
GO ?= 1.13
GO_PACKAGES = .
GO_FILES := $(shell find $(shell go list -f '{{.Dir}}' $(GO_PACKAGES)) -name \*.go)
GIT = git

tao:
	go run build/ci.go install ./cmd/tao
	@echo "Done building."
	@echo "Run \"$(GOBIN)/tao\" to launch tao."

gc:
	go run build/ci.go install ./cmd/gc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gc\" to launch gc."

bootnode:
	go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch a bootnode."

puppeth:
	go run build/ci.go install ./cmd/puppeth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/puppeth\" to launch puppeth."

all:
	go run build/ci.go install

test: all
	go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

tao-cross: tao-windows-amd64 tao-darwin-amd64 tao-linux
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/tao-*

tao-linux: tao-linux-386 tao-linux-amd64 tao-linux-mips64 tao-linux-mips64le tao-linux-arm
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/tao-linux-*

tao-linux-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/tao
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/tao-linux-* | grep 386

tao-linux-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/tao
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/tao-linux-* | grep amd64

tao-linux-mips:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/tao
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/tao-linux-* | grep mips

tao-linux-mipsle:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/tao
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/tao-linux-* | grep mipsle

tao-linux-mips64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/tao
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/tao-linux-* | grep mips64

tao-linux-mips64le:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/tao
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/tao-linux-* | grep mips64le

tao-linux-arm:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm --dest=build/bin --ldflags '-extldflags "-static"' -v github.com/Tao-Network/tao2/cmd/tao
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/tao-linux-* | grep arm

tao-darwin: tao-darwin-386 tao-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/tao-darwin-*

tao-darwin-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/tao
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/tao-darwin-* | grep 386

tao-darwin-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/tao
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/tao-darwin-* | grep amd64

tao-windows-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/tao
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/tao-windows-* | grep amd64
gofmt:
	$(GOFMT) -s -w $(GO_FILES)
	$(GIT) checkout vendor
