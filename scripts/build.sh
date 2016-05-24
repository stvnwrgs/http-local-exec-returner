#! /bin/bash

if ! which gox > /dev/null; then
  echo "==> Installing gox...";
  go get -u github.com/mitchellh/gox;
fi

echo "==> Removing old dirs"
rm -rf ./bin/*
mkdir -p ./bin/
echo "==> Building..."
gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -output "bin/{{.OS}}_{{.Arch}}/${CMD}"