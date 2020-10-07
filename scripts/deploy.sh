#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname "$0")
echo "${SCRIPT_DIR}"

project=$(gcloud secrets versions access latest --secret="project-id")
if [[ -z "${project}" ]]; then
  echo -n "need project"
  exit 1
fi
echo "${project}"

gcloud run deploy go-publisher-fs \
  --image gcr.io/"${project}"/go-publisher-fs:latest \
  --platform managed \
  --project "${project}" \
  --region asia-northeast1 \
  --allow-unauthenticated \
  --set-env-vars PUB_PROJECT="${project}"
