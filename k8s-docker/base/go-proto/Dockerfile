FROM nvidia/cuda:10.0-cudnn7-devel-ubuntu18.04

# FROM python:3.6-stretch as base

ARG PROTOBUF_VERSION

ENV GO111MODULE=auto
ENV LANG C.UTF-8

RUN APT_INSTALL="apt-get install -y --no-install-recommends" && \
    PIP_INSTALL="python -m pip --no-cache-dir install --upgrade" && \
    GIT_CLONE="git clone --depth 10" && \

    rm -rf /var/lib/apt/lists/* \
           /etc/apt/sources.list.d/cuda.list \
           /etc/apt/sources.list.d/nvidia-ml.list && \

    apt-get update && \

# ==================================================================
# tools
# ------------------------------------------------------------------

    DEBIAN_FRONTEND=noninteractive $APT_INSTALL \
        build-essential \
        apt-utils \
        ca-certificates \
        wget \
        git \
        vim \
        curl \
        unzip \
        unrar \
        && \

# CUSTOM ---------------------------------------------------------------------->

    DEBIAN_FRONTEND=noninteractive $APT_INSTALL \
        sudo \
        zip \
        nano \
        tmux \
        tree \
        mc \
        nodejs \
        npm \
        openssh-server \
        rsync \
        sshfs \
        libsndfile1 \
        # python3-dev \
        # python3-pip \
        # python3-tk \
        && \

# CUSTOM ----------------------------------------------------------------------<

    $GIT_CLONE https://github.com/Kitware/CMake ~/cmake && \
    cd ~/cmake && \
    ./bootstrap && \
    make -j"$(nproc)" install && \

# ==================================================================
# python
# ------------------------------------------------------------------

    DEBIAN_FRONTEND=noninteractive $APT_INSTALL \
        software-properties-common \
        && \
    add-apt-repository ppa:deadsnakes/ppa && \
    apt-get update && \
    DEBIAN_FRONTEND=noninteractive $APT_INSTALL \
        python3.6 \
        python3.6-dev \
        python3-distutils-extra \
        && \
    wget -O ~/get-pip.py \
        https://bootstrap.pypa.io/get-pip.py && \
    python3.6 ~/get-pip.py && \
    ln -s /usr/bin/python3.6 /usr/local/bin/python3 && \
    ln -s /usr/bin/python3.6 /usr/local/bin/python && \
    $PIP_INSTALL \
        setuptools \
        && \
    $PIP_INSTALL \
        numpy \
        scipy \
        pandas \
        cloudpickle \
        scikit-learn \
        matplotlib \
        Cython


#############################################################################
# go install start
# gcc for cgo
RUN apt-get update && apt-get install -y --no-install-recommends \
	g++ \
	gcc \
	libc6-dev \
	make \
	ca-certificates \
	curl \
	netbase \
	wget \
	pkg-config \
	&& rm -rf /var/lib/apt/lists/*

RUN set -ex; \
	if ! command -v gpg > /dev/null; then \
		apt-get update; \
		apt-get install -y --no-install-recommends \
			gnupg \
			dirmngr \
		; \
		rm -rf /var/lib/apt/lists/*; \
	fi

ENV GOLANG_VERSION 1.12.6

RUN set -eux; \
	\
# this "case" statement is generated via "update.sh"
	dpkgArch="$(dpkg --print-architecture)"; \
	case "${dpkgArch##*-}" in \
		amd64) goRelArch='linux-amd64'; goRelSha256='dbcf71a3c1ea53b8d54ef1b48c85a39a6c9a935d01fc8291ff2b92028e59913c' ;; \
		armhf) goRelArch='linux-armv6l'; goRelSha256='0708fbc125e7b782b44d450f3a3864859419b3691121ad401f1b9f00e488bddb' ;; \
		arm64) goRelArch='linux-arm64'; goRelSha256='8f4e3909c74b4f3f3956715f32419b28d32a4ad57dbd79f74b7a8a920b21a1a3' ;; \
		i386) goRelArch='linux-386'; goRelSha256='7aaf25164a9ab5e1005c15535ed16ee122df50ac192c2d79b7940315c2b74f2c' ;; \
		ppc64el) goRelArch='linux-ppc64le'; goRelSha256='67eacb68c1e251c1428e588776c5a02e287a508e3d44f940d31d8ff5d57f0eef' ;; \
		s390x) goRelArch='linux-s390x'; goRelSha256='c14baa10b87a38e56f28a176fae8a839e9052b0e691bdc0461677d4bcedea9aa' ;; \
		*) goRelArch='src'; goRelSha256='c96c5ccc7455638ae1a8b7498a030fe653731c8391c5f8e79590bce72f92b4ca'; \
			echo >&2; echo >&2 "warning: current architecture ($dpkgArch) does not have a corresponding Go binary release; will be building from source"; echo >&2 ;; \
	esac; \
	\
	url="https://golang.org/dl/go${GOLANG_VERSION}.${goRelArch}.tar.gz"; \
	wget -O go.tgz "$url"; \
	echo "${goRelSha256} *go.tgz" | sha256sum -c -; \
	tar -C /usr/local -xzf go.tgz; \
	rm go.tgz; \
	\
	if [ "$goRelArch" = 'src' ]; then \
		echo >&2; \
		echo >&2 'error: UNIMPLEMENTED'; \
		echo >&2 'TODO install golang-any from jessie-backports for GOROOT_BOOTSTRAP (and uninstall after build)'; \
		echo >&2; \
		exit 1; \
	fi; \
	\
	export PATH="/usr/local/go/bin:$PATH"; \
	go version

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH
# go install end
#############################################################################

#
RUN apt-get update && apt-get install -y --no-install-recommends \
	g++ \
	gcc \
	git \
	libc6-dev \
	make \
	wget \
	zip \
	unzip \
	pkg-config \
  && rm -rf /var/lib/apt/lists/*

# Install proto compiler
RUN mkdir -p /usr/share/protobuf && \
	cd /usr/share/protobuf && \
	wget --quiet "https://github.com/google/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip" && \
	unzip protoc-${PROTOBUF_VERSION}-linux-x86_64.zip && \
	rm protoc-${PROTOBUF_VERSION}-linux-x86_64.zip

RUN go get \
	google.golang.org/grpc \
	github.com/golang/protobuf/protoc-gen-go \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
	github.com/cespare/reflex \
	github.com/golang/mock/gomock \
	github.com/golang/mock/mockgen \
	github.com/onsi/ginkgo/ginkgo

# nvm environment variables
ENV NVM_DIR /usr/local/nvm
ENV NODE_VERSION 10.15.3

# install nvm
# https://github.com/creationix/nvm#install-script
RUN curl --silent -o- https://raw.githubusercontent.com/creationix/nvm/v0.31.2/install.sh | bash

# install node and npm
RUN . $NVM_DIR/nvm.sh \
	&& nvm install $NODE_VERSION \
	&& nvm alias default $NODE_VERSION \
	&& nvm use default

# add node and npm to path so the commands are available
ENV NODE_PATH $NVM_DIR/v$NODE_VERSION/lib/node_modules
ENV PATH $NVM_DIR/versions/node/v$NODE_VERSION/bin:$PATH

RUN npm install -g nodemon @vue/cli

# additional package
RUN apt-get update && apt-get install -y --no-install-recommends \
	vim \
	tree \
	&& rm -rf /var/lib/apt/lists/*

VOLUME ["/go"]
