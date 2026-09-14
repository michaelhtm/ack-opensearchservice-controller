# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

import time
import pytest
from typing import Dict, Any
from pathlib import Path

from botocore.config import Config
from botocore.exceptions import ClientError

from acktest.resources import load_resource_file

SERVICE_NAME = "opensearchservice"
CRD_GROUP = "opensearchservice.services.k8s.aws"
CRD_VERSION = "v1alpha1"

# Adaptive mode adds client-side rate limiting on top of backoff, which the
# standard mode lacks.
RETRY_CONFIG = Config(retries={"max_attempts": 20, "mode": "adaptive"})

THROTTLE_ERROR_CODES = frozenset({
    "ThrottlingException",
    "Throttling",
    "ThrottledException",
    "TooManyRequestsException",
    "RequestLimitExceeded",
    "RequestThrottled",
    "RequestThrottledException",
})

RETRY_ON_THROTTLE_TIMEOUT_SECONDS = 60*10
RETRY_ON_THROTTLE_INITIAL_DELAY_SECONDS = 5
RETRY_ON_THROTTLE_MAX_DELAY_SECONDS = 60


def is_throttling_error(err: BaseException) -> bool:
    return (
        isinstance(err, ClientError)
        and err.response.get("Error", {}).get("Code") in THROTTLE_ERROR_CODES
    )


def retry_on_throttle(
        fn,
        *args,
        timeout_seconds: int = RETRY_ON_THROTTLE_TIMEOUT_SECONDS,
        **kwargs,
    ):
    deadline = time.monotonic() + timeout_seconds
    delay = RETRY_ON_THROTTLE_INITIAL_DELAY_SECONDS
    while True:
        try:
            return fn(*args, **kwargs)
        except ClientError as e:
            if not is_throttling_error(e) or time.monotonic() >= deadline:
                raise
            time.sleep(delay)
            delay = min(delay*2, RETRY_ON_THROTTLE_MAX_DELAY_SECONDS)

# PyTest marker for the current service
service_marker = pytest.mark.service(arg=SERVICE_NAME)
bootstrap_directory = Path(__file__).parent
resource_directory = Path(__file__).parent / "resources"

def load_opensearch_resource(resource_name: str, additional_replacements: Dict[str, Any] = {}):
    """ Overrides the default `load_resource_file` to access the specific resources
    directory for the current service.
    """
    return load_resource_file(resource_directory, resource_name, additional_replacements=additional_replacements)
