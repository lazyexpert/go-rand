clean: 
	rm -rf bin
compile: 
	go build -o bin/random-int
all:
	make clean && mkdir bin && make compile