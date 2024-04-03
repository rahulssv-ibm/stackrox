#!/usr/bin/env bash

# Compatibility test installation of ACS using separate version arguments for
# central and secured cluster.

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")"/../.. && pwd)"
# shellcheck source=../../scripts/ci/gcp.sh
source "$ROOT/scripts/ci/gcp.sh"
# shellcheck source=../../scripts/ci/lib.sh
source "$ROOT/scripts/ci/lib.sh"
# shellcheck source=../../scripts/ci/sensor-wait.sh
source "$ROOT/scripts/ci/sensor-wait.sh"
# shellcheck source=../../scripts/ci/create-webhookserver.sh
source "$ROOT/scripts/ci/create-webhookserver.sh"
# shellcheck source=../../tests/e2e/lib.sh
source "$ROOT/tests/e2e/lib.sh"
# shellcheck source=../../tests/scripts/setup-certs.sh
source "$ROOT/tests/scripts/setup-certs.sh"
# shellcheck source=../../qa-tests-backend/scripts/lib.sh
source "$ROOT/qa-tests-backend/scripts/lib.sh"

set -euo pipefail

compatibility_test() {
    if [[ "$#" -ne 2 ]]; then
        die "missing args. usage: compatibility_test <central_version> <sensor_version>"
    fi

    local central_version="$1"
    local sensor_version="$2"

    require_environment "ORCHESTRATOR_FLAVOR"
    require_environment "KUBECONFIG"

    short_central_tag="$(shorten_tag "${central_version}")"
    short_sensor_tag="$(shorten_tag "${sensor_version}")"

    junit_wrap CentralSensorVersionCompatibility "central: ${short_central_tag}, sensor: ${short_sensor_tag}" "" \
        _compatibility_test "${central_version}" "${sensor_version}" "${short_central_tag}" "${short_sensor_tag}"
}

_compatibility_test() {
    info "Starting test (compatibility test Central version - ${central_version}, Sensor version - ${sensor_version})"

    local central_version="$1"
    local sensor_version="$2"
    local short_central_tag="$3"
    local short_sensor_tag="$4"

    export_test_environment

    if [[ "${SKIP_DEPLOY:-false}" = "false" ]]; then
        if [[ "${CI:-false}" = "true" ]]; then
            setup_gcp
        else
            info "Not running on CI: skipping cluster setup make sure cluster is already available"
        fi

        setup_deployment_env false false
        setup_podsecuritypolicies_config
        remove_existing_stackrox_resources
        setup_default_TLS_certs

        deploy_stackrox_with_custom_central_and_sensor_versions "${central_version}" "${sensor_version}"
        echo "Stackrox deployed"
        kubectl -n stackrox get deploy,ds -o wide

        deploy_webhook_server
        get_ECR_docker_pull_password
    fi

    rm -f FAIL
    remove_qa_test_results

    info "Running compatibility tests"

    if [[ "${ORCHESTRATOR_FLAVOR}" == "openshift" ]]; then
        oc get scc qatest-anyuid || oc create -f "${ROOT}/qa-tests-backend/src/k8s/scc-qatest-anyuid.yaml"
    fi

    export CLUSTER="${ORCHESTRATOR_FLAVOR^^}"

    make -C qa-tests-backend compatibility-test || touch FAIL

    update_junit_prefix_with_central_and_sensor_version "${short_central_tag}" "${short_sensor_tag}"

    store_qa_test_results "compatibility-test-central-v${short_central_tag}-sensor-v${short_sensor_tag}"
    [[ ! -f FAIL ]] || die "compatibility-test-central-v${short_central_tag}-sensor-v${short_sensor_tag}"
}

update_junit_prefix_with_central_and_sensor_version() {
    local short_central_tag="$1"
    local short_sensor_tag="$2"

    result_folder="${ROOT}/qa-tests-backend/build/test-results/testCOMPATIBILITY"
    info "Updating all test in $result_folder to have \"Central-v${short_central_tag}_Sensor-v${short_sensor_tag}_\" prefix"
    for f in "$result_folder"/*.xml; do
        sed -i "s/testcase name=\"/testcase name=\"[Central-v${short_central_tag}_Sensor-v${short_sensor_tag}] /g" "$f"
    done
}

shorten_tag() {
    if [[ "$#" -ne 1 ]]; then
        die "Expected a version tag as parameter in shorten_tag: shorten_tag <tag>"
    fi

    long_tag="$1"

    short_tag_regex='([0-9]+\.[0-9]+\.[0-9xX]+)'

    if [[ $long_tag =~ $short_tag_regex ]]; then
        echo "${BASH_REMATCH[1]}"
    else
        echo "${long_tag}"
        >&2 echo "Failed to shorten tag ${long_tag} as it did not match the regex: \"${short_tag_regex}\""
        exit 1
    fi
}

compatibility_test "$@"
