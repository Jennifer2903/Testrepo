all: build copy


build: copy
	go build -o bin/main main.go
	make copy

copy:
	cp file.txt bin/
	

update_extensions:
	for file in *.text; do mv "$$file" "$${file%.text}.txt" ; done 



clean: all

  

