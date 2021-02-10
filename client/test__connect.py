import json

import grpc
import requests

from config import Config
from interceptor import newMetadataClientInterceptor
from protoc import sakuraTalk_pb2_grpc, sakuraTalk_pb2


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
channel = grpc.intercept_channel(channel, newMetadataClientInterceptor(("x-sakura-access", firebase_token)))

client = sakuraTalk_pb2_grpc.TalkServiceStub(channel)
primary = client.registerPrimary(
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
response = client.createTag(
    sakuraTalk_pb2.createTagRequest(tag=tag)
)
print(response.tagID)
