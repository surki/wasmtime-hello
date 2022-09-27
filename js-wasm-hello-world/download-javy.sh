#!/bin/bash

set -euo pipefail

cd $( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

mkdir -p javy
pushd javy >/dev/null
JAVY_VERSION="v0.3.0"
OUT="javy.gz"
if [[ "$(uname -s)" == "Darwin" ]]; then
    curl --silent --fail --location https://github.com/Shopify/javy/releases/download/${JAVY_VERSION}/javy-x86_64-macos-${JAVY_VERSION}.gz -o $OUT
else
    curl --silent --fail --location https://github.com/Shopify/javy/releases/download/${JAVY_VERSION}/javy-x86_64-linux-${JAVY_VERSION}.gz -o $OUT
fi
gzip -fd $OUT
chmod +x ./javy
