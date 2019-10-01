NAME=template
VERSION=1.0

debug:
	echo 1

build:
	go build -o ${NAME} .

run:
	go run *.go

clean:
	-@rm ${NAME}
