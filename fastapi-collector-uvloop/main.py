import logging
import json

import uvicorn
from fastapi import FastAPI, Request, HTTPException, status as HTTP_status

from config import config
from event import Event
from kafka_client import KafkaClient


logger = logging.getLogger(__name__)
app = FastAPI()
kafka_client = KafkaClient(config=config)


@app.get("/")
async def index(request: Request):
    args = dict(request.query_params)
    logger.info(f"{args=}")
    event = Event(
        source=args.get("source", None),
        event_name=args.get("event_name", None),
        event_status=args.get("event_status", None),
        created=args.get("created", None),
        payout=args.get("payout", None),
    )
    serialized_event = json.dumps(event.dict()).encode()

    try:
        await kafka_client.send_message(
            topic=config["KAFKA_EVENTS_TOPIC"],
            message=serialized_event
        )
        logger.info(f"{event=} sent successfully")

    except Exception as exc:
        logger.error(f"Sending {event=} to {config['KAFKA_EVENTS_TOPIC']} failed; {str(exc)=}")
        raise HTTPException(
            status_code=HTTP_status.HTTP_500_INTERNAL_SERVER_ERROR, detail=str(exc)
        ) from exc

    return args


def main():
    logger.info(f"Running app: {config['HOST']}:{config['PORT']}")
    uvicorn.run(app, host=config['HOST'], port=int(config['PORT']))

if __name__ == "__main__":
    main()