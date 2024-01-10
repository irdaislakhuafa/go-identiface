# STAGE BUILDER
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
	echo "tzdata tzdata/Areas select Asia" | debconf-set-selections && \
	echo "tzdata tzdata/Zones/Asia select Jakarta" | debconf-set-selections && \
	apt-get install -y tzdata && \
	apt install -y libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg-turbo8-dev \
	curl tar ca-certificates && \
	apt autoclean

# copy used files from builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/assets/models ./assets/models

# then run your app
CMD [ "./app" ]

