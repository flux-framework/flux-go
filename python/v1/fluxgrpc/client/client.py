import fluxgrpc.auth as auth
from fluxgrpc.protos import flux_pb2, flux_pb2_grpc


class FluxClient:
    """
    The FluxClient is used to communicate with the grpc service.
    """

    def __init__(self, host="localhost:50051", use_ssl=False):
        self.host = host
        self.use_ssl = use_ssl

    def submit(self, jobspec):
        """
        Submit a job to flux via a json jobspec
        """
        # These are submit variables. A more substantial submit script would have argparse, etc.
        request = flux_pb2.SubmitRequest(jobspec=jobspec)

        with auth.grpc_channel(self.host, self.use_ssl) as channel:
            stub = flux_pb2_grpc.FluxServiceStub(channel)
            return stub.Submit(request)
