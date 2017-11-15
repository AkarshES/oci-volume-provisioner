#!/bin/bash

RELEASE=${RELEASE:=0.1.0}

IMAGE=wcr.io/oracle/oci-volume-provisioner

if git rev-parse "$RELEASE" >/dev/null 2>&1; then
    echo "Tag $RELEASE already exists. Doing nothing"
else
    SHA=$(git describe --always HEAD)

    echo "Creating new release $RELEASE for SHA $SHA"

    git tag -a "$RELEASE" -m "Release version: $RELEASE"
    git push --tags

    docker pull $IMAGE:$SHA
    docker tag $IMAGE:$SHA $IMAGE:$RELEASE
    docker push $IMAGE:$RELEASE
fi
