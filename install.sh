#!/bin/bash

echo "[+] Downloading latest version of SensorMesh"

curl -s https://api.github.com/repos/davidjosearaujo/sensor-mesh/releases/latest \
| grep "browser_download_url.*tar.gz" \
| cut -d : -f 2,3 \
| tr -d \" \
| wget --show-progress -qi  -

echo "[+] Extracting binary to /usr/local/bin"
tar -C /usr/local/bin -zxf sensormesh*.tar.gz
chmod +x /usr/local/bin/sensormesh

rm sensormesh*.tar.gz
echo "[+] All set!"