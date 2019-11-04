NAME=template
VERSION=1.0

debug: run


build:
	go build -o ${NAME} .

run:
	go run *.go

clean:
	-@rm ${NAME}
