import os
import json
import grpc
from datetime import datetime
from flask import request, Response, Flask

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


def not_completed():
    return [c for c in client.conferences.list() if c.status != "completed"]


@app.route("/conferences", methods=["GET"])
def conferences() -> Response:
    args = request.args
    num = args.get("n", None)

    confereces = not_completed()

    if num and len(confereces) > int(num):
        for c in confereces[int(num):]:
            c.update(status='completed')
        confereces = not_completed()

    return Response(json.dumps([{
        "account_sid": c.account_sid,
        "api_version": c.api_version,
        "date_created": str(c.date_created),
        "date_updated": str(c.date_updated),
        "friendly_name": c.friendly_name,
        "region": c.region,
        "sid": c.sid,
        "status": c.status,
        "uri": c.uri,
        "reason_conference_ended": c.reason_conference_ended,
    } for c in confereces]), mimetype="application/json")


@app.route("/tel", methods=["GET"])
def tel() -> Response:
    """/tel takes telephone numbers from matchings"""
    try:
        args = request.args
        f = args.get("f", None)
        t = args.get("t", None)
        assert f and t, "f and t are required"

        f = datetime.strptime(f, "%Y%m%d%H%M%S")
        t = datetime.strptime(t, "%Y%m%d%H%M%S")

        resp = matching_store.ListMatching(ListMatchingRequest(
            start= {
                "year": f.year,
                "month": f.month,
                "day": f.day,
                "hours": f.hour,
                "minutes": f.minute,
                "seconds": f.second,
                "nanos": 0,
            },
            end= {
                "year": t.year,
                "month": t.month,
                "day": t.day,
                "hours": t.hour,
                "minutes": t.minute,
                "seconds": t.second,
                "nanos": 0,
            },
        ))

        result = {}
        for matching in resp.matching:
            scammers = scammer_store.GetScammer(GetScammerRequest(id=list(matching.scammer_id)))
            result[matching.id] = [{
                "id": s.id,
                "name": s.name,
                "tel": s.tel,
                "tags": [t for t in s.tags],
                "calls": [c for c in s.calls],
                "is_active": s.is_active
            } for s in scammers.scammer]
        return {
            "status": "success",
            "result": result
        }

    except Exception as e:
        return {"error": str(e)}, 400


@app.route("/call", methods=["GET"])
def call() -> Response:
    """/call takes a matching and call them"""
    try:
        args = request.args
        sid = args.get('sid', None)
        tel1 = args.get('tel1', None)
        tel2 = args.get('tel2', None)

        assert sid, "sid is empty"
        assert tel1, "tel1 is empty"
        assert tel2, "tel2 is empty"

        client.conferences(sid).participants.create(
            label="scammer_talk",
            early_media=True,
            beep=None,
            record=True,
            from_=None,
            to="+81" + tel1,
        )
        client.conferences(sid).participants.create(
            label="scammer_talk",
            early_media=True,
            beep=None,
            record=True,
            from_=None,
            to="+81" + tel2,
        )

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
