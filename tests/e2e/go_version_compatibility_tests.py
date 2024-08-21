#!/usr/bin/env -S python3 -u

"""
Run version compatibility tests
"""
import logging
import os
import subprocess
import sys
from collections import namedtuple
from pathlib import Path

from get_latest_helm_chart_versions import (
    get_supported_helm_chart_versions,
    get_latest_helm_chart_version_for_specific_release,
)


Release = namedtuple("Release", ["major", "minor"])

# start logging
logging.basicConfig(stream=sys.stdout, level=logging.DEBUG)

# set required test parameters
os.environ["ORCHESTRATOR_FLAVOR"] = "k8s"

central_chart_versions, sensor_chart_versions = get_supported_helm_chart_versions()

makefile_path = Path(__file__).parent.parent
latest_tag = subprocess.check_output(
    ["make", "tag", "-C", makefile_path, "--quiet", "--no-print-director"],
    shell=False,
    encoding="utf-8",
).strip()

if len(central_chart_versions) == 0:
    logging.info("Found no older central chart versions to test against according to the product lifecycles API.")
if len(sensor_chart_versions) == 0:
    logging.info("Found no older sensor chart versions to test against according to the product lifecycles API.")
if len(central_chart_versions) == 0 or len(sensor_chart_versions) == 0:
    logging.info("However versions with support exceptions will still be tested against.")

ChartVersions = namedtuple(
    "Chart_versions", ["central_version", "sensor_version"])

# Latest central vs sensor versions in sensor_chart_versions
test_tuples = [
    ChartVersions(central_version=latest_tag,
                  sensor_version=sensor_chart_version)
    for sensor_chart_version in sensor_chart_versions
]
# Latest sensor vs central versions in central_chart_versions
test_tuples.extend(
    [
        ChartVersions(central_version=central_chart_version,
                      sensor_version=latest_tag)
        for central_chart_version in central_chart_versions
    ]
)

# Currently there are no support exceptions, the last one expired on 2024-06-30, see:
# https://issues.redhat.com/browse/ROX-18223
# however a new support exception is being negotiated, add it here when it's ready
support_exceptions = [
    ChartVersions(
        central_version=latest_tag,
        sensor_version=get_latest_helm_chart_version_for_specific_release(
            "stackrox-secured-cluster-services", Release(major=3, minor=74)
        ),
    )
]

test_tuples.extend(
    support_exception
    for support_exception in support_exceptions
    if support_exception not in test_tuples
)

if len(test_tuples) > 0:
    sets = []
    for test_tuple in test_tuples:
        subprocess.run(
            ["run-compatibility.sh", test_tuple.central_version, test_tuple.sensor_version],
            shell=False,
            encoding="utf-8",
        )
else:
    logging.info("There are currently no supported older versions or support exceptions that require compatibility "
                 "testing.")
