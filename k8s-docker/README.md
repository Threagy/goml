# k8s-docker

## Docker

### Build Base

```bash
$ cd k8s-docker/base

$ pwd
.../goml/k8s-docker/base

$ docker-compose build
Successfully built ...
```

### Build Workspace

```bash
$ cd k8s-docker

$ make build
1) Workspace
2) Quit
Please enter your choice: 1

Successfully built ...
```

### Run Workspace

```bash
docker volume create goml_repos_data
docker volume create goml_vscode_data
docker volume create goml_vscode_insider_data
docker volume create goml_go_data
docker volume create goml_py_data

$ make up
```

### Run gomlet

```bash
# ssh로 접속하지 말고, 다음 커맨드로 쉘 열 것.
$ make bash
1) Workspace
2) Quit
Please enter your choice: 1

$ cd ~/repos

$ git clone ssh://git@git.korea.ncsoft.corp:7999/metallica/goml.git

$ cd ~/repos/goml/gomlet

$ make run
```

### Run gomlctl

```
$ cd gomlctl

$ go run main.go get all
```

### Run Rest-Client (VSCode)

```
(edit settings.json)

> Rest Client: Switch Environment
```

# Tmux Project

## Install

```bash
$ pip install --user tmuxp
```

## Usage
```bash
$ pwd
/home/$(user)/repos/goml

$ tmuxp load tmux.yaml
```