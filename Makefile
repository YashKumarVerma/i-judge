pack:
	go build -o build/i-judge internal/main.go

run:
	go run internal/main.go

clean:
	rm -rf build/*
	rm -rf output/*