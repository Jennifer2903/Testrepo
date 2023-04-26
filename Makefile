all: build 

build:
	go build main.go 
	go build -o prg2.go
	

run: build
  

