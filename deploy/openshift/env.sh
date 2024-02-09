#!/usr/bin/env bash
set -e

export ORCH="openshift"
export ORCH_CMD="oc"
export ORCH_FULLNAME="openshift"

export ADMISSION_CONTROLLER_POD_EVENTS="${ADMISSION_CONTROLLER_POD_EVENTS:-false}"
echo "ADMISSION_CONTROLLER_POD_EVENTS set to ${ADMISSION_CONTROLLER_POD_EVENTS}"

export ROX_OPENSHIFT_VERSION="${ROX_OPENSHIFT_VERSION:-4}"
echo "ROX_OPENSHIFT_VERSION set to ${ROX_OPENSHIFT_VERSION}"

export CLUSTER_TYPE="OPENSHIFT_CLUSTER"
if [[ "${ROX_OPENSHIFT_VERSION:-}" == "4" ]]; then
    export CLUSTER_TYPE="OPENSHIFT4_CLUSTER"
fi
