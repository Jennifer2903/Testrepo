all: build copy

build: main.go
	go build -o bin/main main.go
	make copy

copy:
	cp file.txt bin/
 
run:
	build

  

