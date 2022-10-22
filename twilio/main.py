import os
import grpc
from flask import request, Request, Response, Flask

from twilio.rest import Client
from twilio.twiml.voice_response import VoiceResponse

# import functions_framework
# from protobuf_to_dict import protobuf_to_dict as ptd
# from werkzeug.datastructures import MultiDict
from api.v1.matching_pb2 import ListMatchingRequest
from api.v1.matching_pb2_grpc import MatchingStoreServiceStub
from api.v1.scammer_pb2 import GetScammerRequest
from api.v1.scammer_pb2_grpc import ScammerStoreServiceStub

endpoint_url = os.getenv("ENDPOINT_URL")
assert endpoint_url, "ENDPOINT_URL is empty"
channel = grpc.insecure_channel(endpoint_url)
matching_store = MatchingStoreServiceStub(channel)
scammer_store = ScammerStoreServiceStub(channel)

account_sid = os.getenv("TWILIO_ACCOUNT_SID")
auth_token = os.getenv("TWILIO_AUTH_TOKEN")
client = Client(account_sid, auth_token)

app = Flask(__name__)


@app.route("/conferences", methods=["GET"])
def check() -> Response:
    confereces = client.conferences.list()
    return Response({
        "conferences": [c.sid for c in confereces],
    }, mimetype="application/json")


@app.route("/call", methods=["GET"])
def call() -> Response:
    """/call takes a matching and call them"""
    try:
        args = request.args
        f = args.get("f")
        t = args.get("t")
        assert f and t, "f and t are required"

        confereces = client.conferences.list()
        assert len(confereces) > 0, "Start a new conference first"
        if len(confereces) > 1:
            for c in confereces[1:]:
                c.delete()

        resp = matching_store.ListMatching(ListMatchingRequest(
            from_=f,
            to=t,
        ))

        all_scammers = []

        for index, matching in enumerate(resp.matching):
            scammers = scammer_store.GetScammer(GetScammerRequest(id=list(matching.scammer_id)))
            all_scammers += scammers

            conference = confereces[index]

            for scammer in scammers.scammer:
                # client.conferences(conference.sid).participants.create(
                conference.participants.create(
                    label=f"scammer_talk_{index}",
                    early_media=True,
                    beep=None,
                    # status_callback="https://myapp.com/events",
                    # status_callback_event=["ringing"],
                    record=True,
                    from_=None,
                    to="+81" + scammer.tel,
                )
        return {
            "status": "success",
            "conferences": [c.sid for c in confereces],
            "scammer_tel": [s.tel for s in all_scammers],
        }
    except grpc._channel._InactiveRpcError as e:
        return {"status": "failed", "error": str(e._state.details), "details": str(e.debug_error_string)}
    except Exception as e:
        return {"status": "failed", "error": str(e)}


@app.route("/")
def start_conference():
    resp = VoiceResponse()
    resp.dial().conference(
        "ScammerTalk",
        record="record-from-start",
        beep=False,
        start_conference_on_enter=True,
        end_conference_on_exit=True,
    )

    return Response(str(resp), mimetype="text/xml")


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5050)

# @functions_framework.http
# def main(request: Request):
#     path = request.path
#     args = request.args

#     if path == "/":
#         return start_conference(args)
#     elif path == "/call":
#         return call(args)
#     else:
#         return f"Unknown endpoint {path}"
