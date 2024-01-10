# use ubuntu:24.04 as image to run
FROM ubuntu:24.04 AS builder
WORKDIR /app
COPY . .

# update local repository with cloud repo and install libs to run/compile https://github.com/irdaislakhuafa/go-identiface.git
RUN apt update && \
	echo "tzdata tzdata/Areas select Asia" | debconf-set-selections && \
	echo "tzdata tzdata/Zones/Asia select Jakarta" | debconf-set-selections && \
	apt-get install -y tzdata && \
	apt install -y libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg-turbo8-dev \
	curl tar ca-certificates golang-go && \
	apt autoclean && \
	go version

# build/compile your code into binary
RUN go build -o app main.go

# then run your app
CMD [ "./app" ]

