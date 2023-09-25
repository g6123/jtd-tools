split = $(word $3,$(subst $2, ,$1))

all: bin/jtd-tools-darwin-arm64 bin/jtd-tools-linux-arm64 bin/jtd-tools-linux-amd64

bin/jtd-tools-%:
	GOOS=$(call split,$*,-,1) GOARCH=$(call split,$*,-,2) \
		go build -trimpath -o $@ ./cmd/

clean:
	rm -vf bin/*
