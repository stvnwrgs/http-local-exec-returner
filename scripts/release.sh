#! /bin/bash
echo "==> Cleaning release dir"
rm -rf ./release/*
mkdir -p ./release/

if ! which github-release > /dev/null; then \
  echo "==> Installing github-release..."; \
  go get -u github.com/aktau/github-release; \
fi

echo "==> Zipping bins:"
for PLATFORM in $(find ./bin -mindepth 1 -maxdepth 1 -type d); do
  OSARCH=$(basename ${PLATFORM})
  echo "=======> ${OSARCH}"
  pushd $PLATFORM >/dev/null 2>&1
  zip $PWD/../../release/${OSARCH}.zip ./*
  popd >/dev/null 2>&1
done

# echo "==> Creating release: ${VERSION}"
# github-release release \
#    --user ${USER} \
#    --repo ${REPOSITORY} \
#    --tag ${VERSION}\
#    --name "${VERSION}" \
#    --description "${VERSION} See: [CHANGELOG.md](CHANGELOG.md)"
#
# for PLATFORM in $(find ./bin -mindepth 1 -maxdepth 1 -type d); do
#     github-release upload \
#         --user ${USER} \
#         --repo ${REPOSITORY} \
#         --tag ${VERSION}\
#     --name "${REPOSITORY}-${VERSION}-osx-amd64" \
#     --file bin/darwin/amd64/gofinance
# done