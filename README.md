# SensorMesh

<p align="left">
  <img src="doc/branding/sensormesh-logo.png" width="360" />
</p>

<br>

[![Go](https://github.com/davidjosearaujo/sensor-mesh/actions/workflows/go.yml/badge.svg)](https://github.com/davidjosearaujo/sensor-mesh/actions/workflows/go.yml) ![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/davidjosearaujo/sensor-mesh?label=release)

SensorMesh allows you to direct output from serial devices and internal services into a distributed peer-to-peer database on IPFS.

SensorMesh is built as a layer of abstraction and service compatibility for [OrbitDB](https://github.com/orbitdb/orbit-db/). It is built in [Go](https://go.dev/) and uses the [go-orbit-db](https://github.com/berty/go-orbit-db) module developed by the folks at [Berty](https://berty.tech/). Check all the awesome people out!

Imagine a factory floor with multiple independent nodes, each with its sensors and services. With SensorMesh, each node is capable of publishing its data to a distributed P2P log database, accessible from any node at any time and with a loss of connection awareness between nodes.

<br>

# Table of Contents

- [Install](#install)
  - [GLIBC error](#glibc-error)
  - [Build](#build)
- [Configuration](#configurations)
- [Usage](#usage)

<br>

# Install

To install SensorMesh, simply download the latest binary which can be found at [**Releases**](https://github.com/davidjosearaujo/sensor-mesh/releases).

Or do it from the terminal by running the following command.
``` bash
curl https://github.com/davidjosearaujo/sensor-mesh/blob/main/install.sh && \
chmod +x install.sh && \
sudo ./install.sh
```

You should now be able to call _**sensormesh**_ from any terminal session!
```
$ sensormesh
Usage:
  sensormesh [command]

Available Commands:
  channel     Channel is a palette that contains MQTT addion based commands
  completion  Generate the autocompletion script for the specified shell
  config      Config allows you to see your configurations and edit them
  daemon      Run a OrbitDB sensor logger
  help        Help about any command
  init        Initialize local SensorMesh configuration
  sensor      Sensor is a palette that contains sensor based commands

Flags:
  -h, --help   help for sensormesh

Use "sensormesh [command] --help" for more information about a command.
```

## GLIBC error

If you run into problems regarding **GLIBC**, such as:
``` bash
sensormesh: /lib/x86_64-linux-gnu/libc.so.6: version `GLIBC_x.xx' not found (required by sensormesh)
sensormesh: /lib/x86_64-linux-gnu/libc.so.6: version `GLIBC_x.xx' not found (required by sensormesh)
``` 
don't worry, you'll just have to **compile SensorMesh** yourself, for that, follow the instructions in the [Build](#build) section to see how.

<br>

# Build

If you prefer to compile SensorMesh yourself or you ran into _GLIBC_ errors, run the following command.
``` bash
git clone https://github.com/davidjosearaujo/sensor-mesh.git && \
cd sensor-mesh && \
chmod +x ./build.sh && \
./build.sh
``` 

_The build script will request **sudo** access to move the generated binary to **/usr/local/bin** which should already be in your PATH env variable. If you choose not to provide sudo access, the binary will remain in the **current directory**._

<br>

# Configurations

TD

<br>

# Usage

TD