#!/usr/bin/env bash
set -eou pipefail

artifacts_dir=$1
collector_image_tag=${2:-}
collector_image_registry=${3:-quay.io/rhacs-eng}

echo "Starting secure cluster services"

export KUBECONFIG=$artifacts_dir/kubeconfig

settings=(
    --namespace stackrox stackrox-secured-cluster-services rhacs/secured-cluster-services
    --values perf-bundle.yml
    --set clusterName=perf-test
    --set enableOpenShiftMonitoring=true
    --set exposeMonitoring=true
    --set collector.collectionMethod=CORE_BPF
    --set collector.forceCollectionMethod=true
)

if [[ -n ${ROX_RESYNC_DISABLED:-} ]]; then
    settings+=(--set customize.envVars.ROX_RESYNC_DISABLED="${ROX_RESYNC_DISABLED}")
fi

if [[ -n ${collector_image_tag:-} ]]; then
    settings+=(--set image.collector.registry="$collector_image_registry")
    settings+=(--set image.collector.name=collector)
    settings+=(--set image.collector.tag="$collector_image_tag")
fi

if [[ -n ${DOCKER_USERNAME:-} ]]; then
    settings+=(--set imagePullSecrets.username="$DOCKER_USERNAME")
fi

if [[ -n ${DOCKER_PASSWORD:-} ]]; then
    settings+=(--set imagePullSecrets.password="$DOCKER_PASSWORD")
fi

if [[ -n ${IMAGE_MAIN_REGISTRY:-} ]]; then
    settings+=(--set image.main.registry="$IMAGE_MAIN_REGISTRY")
fi

if [[ -n ${IMAGE_MAIN_NAME:-} ]]; then
    settings+=(--set image.main.name="$IMAGE_MAIN_NAME")
fi

if [[ -n ${IMAGE_MAIN_TAG:-} ]]; then
    settings+=(--set image.main.tag="$IMAGE_MAIN_TAG")
fi

echo "Running: helm install ${settings[@]}"

helm install "${settings[@]}"
