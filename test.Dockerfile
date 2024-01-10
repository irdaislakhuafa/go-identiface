# use ubuntu:24.04 as image to run
FROM ubuntu:24.04 AS builder
WORKDIR /app
COPY . .

# update local repository with cloud repo and install libs to run/compile https://github.com/irdaislakhuafa/go-identiface.git
# RUN apt update && \
# 	apt install -y libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg-turbo8-dev \
# 	curl tar ca-certificates golang-go && \
# 	apt autoclean && \
# 	go version

# test
CMD [ "go", "test", "-v", ".../.." ]