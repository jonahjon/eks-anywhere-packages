#!/usr/bin/env bash
# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -e
set -x
set -o pipefail

export LANG=C.UTF-8

BASE_DIRECTORY=$(git rev-parse --show-toplevel)
HELM_REPO=${HELM_REPO}
echo ${HELM_REPO}

make build
chmod +x ${BASE_DIRECTORY}/generatebundlefile/bin

AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)
aws ecr get-login-password --region us-west-2 | HELM_EXPERIMENTAL_OCI=1 helm registry login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.us-west-2.amazonaws.com
aws ecr-public get-login-password --region us-east-1 | HELM_EXPERIMENTAL_OCI=1 helm registry login --username AWS --password-stdin public.ecr.aws

cd "${BASE_DIRECTORY}/generatebundlefile"

if [[ -n "${PROMOTE_FILE}" ]]; then
    if [[ "${HELM_REPO}" == "eks-anywhere-packages" ]]; then
        # We need to copy the images for the packages helm chart, but no other packages.
        ./bin/generatebundlefile  \
            --promote ${HELM_REPO} --input ${PROMOTE_FILE} --copy-images
    else
        ./bin/generatebundlefile  \
            --promote ${HELM_REPO} --input ${PROMOTE_FILE}
    fi
else
    ./bin/generatebundlefile  \
        --promote ${HELM_REPO}
fi
