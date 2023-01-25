clean:
	rm -f application

build: clean
	go build -o application main.go

run: build
	./application

image: clean
	sudo docker container stop app || true && sudo docker container rm app || true && sudo docker image rm app || true
	sudo docker image build -t app . && sudo docker run --name app -dp 80:8080 app