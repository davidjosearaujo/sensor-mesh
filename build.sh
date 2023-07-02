#!/bin/bash

echo "[+] Building SensorMesh binary"
go build -o sensormesh main.go

echo "[!] To move SensorMesh to /usr/local/bin, it will require sudo access"
sudo mv sensormesh /usr/local/bin/sensormesh