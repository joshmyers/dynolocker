#!/usr/bin/env bash
#
# Pushes a release to Github via the Github releases API

set -e

if [[ -z "$PKG" ]] \
  || [[ -z "$VERSION" ]] \
  || [[ -z "$GITHUB_OAUTH" ]]
then
  echo "Please don't run this script directly, instead run 'make release'"
  exit 1
fi

AUTH="Authorization: token ${GITHUB_OAUTH}"
GH_API="https://api.github.com"
GH_ASSET_UPLOAD="https://uploads.github.com"
GH_REPO="${GH_API}/repos/alphagov/${PKG}"
RELEASE_FILE=$(mktemp)

# Validate token.
curl -o /dev/null -sH "$AUTH" $GH_REPO || { echo "Error: Invalid repo, token or network issue!";  exit 1; }

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd $DIR/bin

git tag v${VERSION}
git push origin --tags

cat << EOF > "${RELEASE_FILE}"
{
  "tag_name": "v${VERSION}",
  "target_commitish": "$(git rev-parse HEAD)",
  "name": "v${VERSION}",
  "body": "Release of ${PKG} v${VERSION}",
  "draft": false,
  "prerelease": false
}
EOF

release_id=$(curl -X POST \
  --data @"${RELEASE_FILE}" \
  -H "Authorization: token ${GITHUB_OAUTH}" \
  -H "Content-Type: application/json" \
  "${GH_REPO}/releases" | jq -r .id)

rm -rf ${RELEASE_FILE}

for bin in $(ls *"${PKG}"* | grep -v zip); do
  echo "==> Zipping artifact... ${bin}" >&2
  zip -r ${bin}.zip ${bin}
  echo "==> Uploading artifact... ${bin}.zip" >&2
  curl -X POST \
    --data-binary @"${bin}.zip" \
    -H "${AUTH}" \
    -H "Content-Type: application/zip" \
    "${GH_ASSET_UPLOAD}/repos/alphagov/${PKG}/releases/${release_id}/assets?name=${bin}.zip"
done
