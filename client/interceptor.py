from typing import Callable, Any, Tuple, List

import grpc
from grpc_interceptor import ClientInterceptor, ClientCallDetails
from grpc_interceptor.client import ClientInterceptorReturnType


class MetadataClientInterceptor(ClientInterceptor):
    @property
    def metadata(self):
        return self._metadata

    @metadata.setter
    def metadata(self, v: List[Tuple[str, str]]):
        self._metadata = v

    def intercept(
            self,
            method: Callable,
            request_or_iterator: Any,
            call_details: grpc.ClientCallDetails,
    ) -> ClientInterceptorReturnType:
        new_details = ClientCallDetails(
            call_details.method,
            call_details.timeout,
            self._metadata,
            call_details.credentials,
            call_details.wait_for_ready,
            call_details.compression,
        )
        return method(request_or_iterator, new_details)


def newMetadataClientInterceptor(*metadata: Tuple[str, str]) -> MetadataClientInterceptor:
    interceptor = MetadataClientInterceptor()
    interceptor.metadata = [*metadata]
    return interceptor
