# use ubuntu:24.04 as image to run
FROM ubuntu:24.04 AS builder
WORKDIR /app
COPY . .

# update local repository with cloud repo and install libs to run/compile https://github.com/irdaislakhuafa/go-identiface.git
RUN apt update && \
	# install and confgure timezone data
	echo "tzdata tzdata/Areas select Asia" | debconf-set-selections && \
	echo "tzdata tzdata/Zones/Asia select Jakarta" | debconf-set-selections && \
	apt-get install -y tzdata && \
	# install dlib
	apt install -y libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg-turbo8-dev \
	# install optional pkgs, you can customize this package as you need
	curl tar ca-certificates golang-go make git && \
	apt autoclean && \
	go version

# get assets not exists
RUN if ! [ -d ./assets ]; then \
	make get-assets; \
	fi

# test
CMD [ "go", "test", "-v", ".../.." ]