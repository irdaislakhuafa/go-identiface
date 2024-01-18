# STAGE BUILDER
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
	# install optional pkgs, go-face required g++ as C++ compiler and gcc as C compiler for compile, you can customize other packages as you need.
	curl tar ca-certificates golang-go g++ gcc --no-install-recommends && \
	apt autoclean && \
	go version

# get assets if not exists
RUN if ! [ -d ./assets ]; then \
	make get-assets; \
	fi


# build/compile your code into binary
RUN go build -o app main.go

# STAGE RUNNER
FROM ubuntu:24.04 AS runner
WORKDIR /app
RUN apt update && \
	# install and confgure timezone data
	echo "tzdata tzdata/Areas select Asia" | debconf-set-selections && \
	echo "tzdata tzdata/Zones/Asia select Jakarta" | debconf-set-selections && \
	apt-get install -y tzdata && \
	# install dlib
	apt install -y libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg-turbo8-dev \
	# install optional pkgs, you can customize this package as you need
	curl tar ca-certificates --no-install-recommends && \
	apt autoclean

# copy used files from builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/assets/models ./assets/models

# then run your app
CMD [ "./app" ]

