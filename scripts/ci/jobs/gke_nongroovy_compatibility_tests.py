#!/usr/bin/env -S python3 -u

"""
Run version compatibility tests
"""
import logging
import os
import sys
from collections import namedtuple

from pre_tests import (
    PreSystemTests,
    CollectionMethodOverridePreTest
)
from ci_tests import QaE2eGoCompatibilityTest
from post_tests import PostClusterTest, FinalPost
from runners import ClusterTestSetsRunner
from clusters import GKECluster
from get_compatibility_test_tuples import (
    get_compatibility_test_tuples,
)

# start logging
logging.basicConfig(stream=sys.stdout, level=logging.DEBUG)

Release = namedtuple("Release", ["major", "minor"])
ChartVersions = namedtuple(
    "Chart_versions", ["central_version", "sensor_version"])

# set required test parameters
os.environ["ORCHESTRATOR_FLAVOR"] = "k8s"

# Get the test tuples (central_version, sensor_version) for supported versions with available helm charts
test_tuples = get_compatibility_test_tuples()

if len(test_tuples) > 0:
    sets = []
    for test_tuple in test_tuples:
        os.environ["ROX_TELEMETRY_STORAGE_KEY_V1"] = 'DISABLED'
        test_versions = f'{test_tuple.central_version}--{test_tuple.sensor_version}'

        # expected version string is like 74.x.x for ACS 3.74 versions
        is_3_74_sensor = test_tuple.sensor_version.startswith('74')

        sets.append(
            {
                "name": f'version compatibility tests: {test_versions}',
                "test": QaE2eGoCompatibilityTest(test_tuple.central_version, test_tuple.sensor_version),
                "post_test": PostClusterTest(
                    collect_collector_metrics=not is_3_74_sensor,
                    check_stackrox_logs=True,
                    artifact_destination_prefix=test_versions,
                ),
                # Collection not supported on 3.74
                "pre_test": CollectionMethodOverridePreTest("NO_COLLECTION" if is_3_74_sensor else "core_bpf")
            },
        )
    ClusterTestSetsRunner(
        cluster=GKECluster("nongroovy-compat-test",
                           machine_type="e2-standard-8", num_nodes=2),
        initial_pre_test=PreSystemTests(),
        sets=sets,
        final_post=FinalPost(
            store_qa_tests_data=True,
        ),
    ).run()
else:
    logging.info("There are currently no supported older versions or support exceptions that require compatibility "
                 "testing.")
