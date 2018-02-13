zip:
	GOOS=linux go build -o main main.go
	zip lambda.zip main
