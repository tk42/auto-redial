import os
import grpc
from flask import Request, Response

# from twilio.twiml.voice_response import VoiceResponse, Dial

import functions_framework
from protobuf_to_dict import protobuf_to_dict as ptd
from werkzeug.datastructures import MultiDict
from api.v1.matching_pb2 import ListMatchingRequest
from api.v1.matching_pb2_grpc import MatchingStoreServiceStub
from api.v1.scammer_pb2 import GetScammerRequest
from api.v1.scammer_pb2_grpc import ScammerStoreServiceStub

endpoint_url = os.getenv("ENDPOINT_URL")
assert endpoint_url, "ENDPOINT_URL is empty"
channel = grpc.insecure_channel(endpoint_url)
matching_store = MatchingStoreServiceStub(channel)
scammer_store = ScammerStoreServiceStub(channel)


def call(args: MultiDict):
    """/call takes a matching and call them"""
    try:
        resp = matching_store.ListMatching(ListMatchingRequest())
        scammer_set = set()
        for matching in resp.matching:
            scammer_set |= set(matching.scammer_id)
        scammers = scammer_store.GetScammer(GetScammerRequest(id=list(scammer_set)))
        tels = [s.tel for s in scammers.scammer]

    except grpc._channel._InactiveRpcError as e:
        return {"error": str(e._state.details), "details": str(e.debug_error_string)}
    except Exception as e:
        return {"error": str(e)}
    return {"tels": tels}


@functions_framework.http
def main(request: Request):
    path = request.path
    args = request.args

    if path == "/":
        with open("./twilio.xml", encoding="utf8") as f:
            content = f.read()
        return Response(content, mimetype="text/xml")
    elif path.startswith("/call"):
        return call(args)
    else:
        return f"Unknown endpoint {path}"
