# envdocker

`envdocker` is a command-line tool that wraps Docker commands and other container-related tools, automatically setting the `DOCKER_HOST` environment variable. This tool is particularly useful when you want to transparently use a Docker daemon running on a virtual machine without installing the Docker daemon on your host PC. It is intended for use in a development environment.

## Features

- Automatically sets the `DOCKER_HOST` environment variable for each command
- Works with Docker commands and other container-related tools (e.g., kind )

## Installation

Build the binary:

```
$ go build
```

## Docker Daemon Setup

To use `envdocker` with a remote Docker daemon, you need to configure the Docker daemon on your server (in this case, a Debian system) to listen on a TCP port. Follow these steps:

1. Edit the Docker daemon configuration file:
   ```
   $ sudo vi /etc/default/docker
   ```

2. Add the following line to the file:
   ```
   DOCKER_OPTS="-H tcp://0.0.0.0:2375"
   ```

3. Save the file and exit the editor.

4. Restart the Docker service to apply the changes:
   ```
   $ sudo systemctl restart docker
   ```

Note: This configuration was tested on the following Debian version:
```
$ lsb_release -a
No LSB modules are available.
Distributor ID: Debian
Description:    Debian GNU/Linux 12 (bookworm)
Release:        12
Codename:       bookworm
```

With Docker version:
```
$ dpkg -l docker.io
Desired=Unknown/Install/Remove/Purge/Hold
| Status=Not/Inst/Conf-files/Unpacked/halF-conf/Half-inst/trig-aWait/Trig-pend
|/ Err?=(none)/Reinst-required (Status,Err: uppercase=bad)
||/ Name           Version             Architecture Description
+++-==============-===================-============-=================================
ii  docker.io      20.10.24+dfsg1-1+b3 amd64        Linux container runtime
```

## Configuration

By default, `envdocker` sets the `DOCKER_HOST` to `tcp://localhost:2375`. When using SSH tunneling, you need to bind port 2375 on the remote server to your host PC. Ensure that your SSH tunnel is set up to forward this port before using `envdocker`.

## Usage

The basic syntax for using `envdocker` is:

```
envdocker [command] [args...]
```

### Examples

Run a Docker command:

```
$ envdocker -- docker ps
```

Create a Kubernetes cluster using kind:

```
$ envdocker -- kind create cluster --config test-cluster.yml
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/license/mit) for details.
