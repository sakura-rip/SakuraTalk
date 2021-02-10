import json
from typing import Any, Callable

import grpc
import requests
from grpc_interceptor import ClientInterceptor, ClientCallDetails
from grpc_interceptor.client import ClientInterceptorReturnType

import sakuraTalk_pb2
import sakuraTalk_pb2_grpc
from config import Config


class MetadataClientInterceptor(ClientInterceptor):
    def intercept(
            self,
            method: Callable,
            request_or_iterator: Any,
            call_details: grpc.ClientCallDetails,
    ) -> ClientInterceptorReturnType:
        new_details = ClientCallDetails(
            call_details.method,
            call_details.timeout,
            [("x-sakura-access", firebase_token)],
            call_details.credentials,
            call_details.wait_for_ready,
            call_details.compression,
        )
        return method(request_or_iterator, new_details)


def signUpWithEmailAndPasswd(email: str, password: str):
    uri = f"https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key={Config.FIREBASE_AUTH_KEY}"
    headers = {"Content-type": "application/json"}
    data = {"email": email, "password": password, "returnSecureToken": True}
    result = requests.post(
        url=uri,
        headers=headers,
        data=json.dumps(data)
    )
    return result.json()["idToken"]


firebase_token = signUpWithEmailAndPasswd(email="fadsfads@gamil.com", password="fadsfasdfas")

channel = grpc.insecure_channel('localhost:8806')
channel = grpc.intercept_channel(channel, MetadataClientInterceptor())
stub = sakuraTalk_pb2_grpc.TalkServiceStub(channel)
primary = stub.registerPrimary(
    sakuraTalk_pb2.registerPrimaryRequest(token=firebase_token)
)
tag = sakuraTalk_pb2.Tag(
    tagID="",
    name="name",
    description="description",
    color="color",
    creator="creator",
    createdTime=0
)
response = stub.createTag(
    sakuraTalk_pb2.createTagRequest(tag=tag)
)
print(response.tagID)
