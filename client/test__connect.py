import json

import grpc
import requests

import sakuraTalk_pb2
import sakuraTalk_pb2_grpc


def signUpWithEmailAndPasswd(email: str, password: str):
    uri = f"https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=AIzaSyAcBzQrmliUXRAxuut3b7j2a1mEXCRpmAs"
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
stub = sakuraTalk_pb2_grpc.TalkServiceStub(channel)
primary, call = stub.registerPrimary.with_call(
    sakuraTalk_pb2.registerPrimaryRequest(token=firebase_token),
    metadata=(("x-sakura-access", firebase_token),)
)
tag = sakuraTalk_pb2.Tag(
    tagID="",
    name="name",
    description="description",
    color="color",
    creator="creator",
    createdTime=0
)
response, call = stub.createTag.with_call(
    sakuraTalk_pb2.createTagRequest(tag=tag),
    metadata=(("x-sakura-access", firebase_token),)
)
print(response.tagID)
for k, v in call.trailing_metadata():
    print(f"k:{k}, v:{v}")
