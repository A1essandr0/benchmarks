import time
from typing import Any, Callable
import random

import grpc
import grpc.experimental.gevent as grpc_gevent
from grpc_interceptor import ClientInterceptor

from locust import User, task
from locust.exception import LocustError

import common.proto.rawevent_pb2 as rawevent_pb2
import common.proto.server_pb2_grpc as server_pb2_grpc


# patch grpc so that it uses gevent instead of asyncio
grpc_gevent.init_gevent()


class LocustInterceptor(ClientInterceptor):
    def __init__(self, environment, *args, **kwargs):
        super().__init__(*args, **kwargs)

        self.env = environment

    def intercept(
        self,
        method: Callable,
        request_or_iterator: Any,
        call_details: grpc.ClientCallDetails,
    ):
        response = None
        exception = None
        start_perf_counter = time.perf_counter()
        response_length = 0
        try:
            response = method(request_or_iterator, call_details)
            response_length = response.result().ByteSize()
        except grpc.RpcError as e:
            exception = e

        self.env.events.request.fire(
            request_type="grpc",
            name=call_details.method,
            response_time=(time.perf_counter() - start_perf_counter) * 1000,
            response_length=response_length,
            response=response,
            context=None,
            exception=exception,
        )
        return response


class GrpcUser(User):
    abstract = True
    stub_class = None

    def __init__(self, environment):
        super().__init__(environment)
        for attr_value, attr_name in ((self.host, "host"), (self.stub_class, "stub_class")):
            if attr_value is None:
                raise LocustError(f"You must specify the {attr_name}.")

        self._channel = grpc.insecure_channel(self.host)
        interceptor = LocustInterceptor(environment=environment)
        self._channel = grpc.intercept_channel(self._channel, interceptor)

        self.stub = self.stub_class(self._channel)



def generate_random_value():
    return f"{random.randint(100,999)}"

class SendingEventGrpc(GrpcUser):
    host = "localhost:5006"
    stub_class = server_pb2_grpc.CollectorServiceStub

    @task
    def send_event(self):
        event_object = rawevent_pb2.RawEvent(
            source=generate_random_value(),
            event_name="GRPC_EVENT",
            payout=generate_random_value()
        )
        self.stub.PostRawEvent(event_object)