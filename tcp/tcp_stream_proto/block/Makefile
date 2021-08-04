all: server client

bin_dir := ./bin

dir:
    $(shell if [ ! -d $(bin_dir) ]; then mkdir -p $(bin_dir); fi)

server: dir
	go build -o $(bin_dir)/server ./cmd/server

client: dir
	go build -o $(bin_dir)/client ./cmd/client

clean:
	rm -rf $(bin_dir)