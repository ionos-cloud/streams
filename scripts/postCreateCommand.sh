#!/bin/bash
# This script is executed after the creation of a new project.

# goreleaser (see https://goreleaser.com/install/)
echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | sudo tee /etc/apt/sources.list.d/goreleaser.list
sudo apt update
sudo apt install -y goreleaser direnv

# k9s (see https://k9scli.io/)
curl -sS https://webinstall.dev/k9s | bash

