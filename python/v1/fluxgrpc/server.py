import argparse
import logging
import sys
from concurrent import futures

import grpc

import fluxgrpc.defaults as defaults
from fluxgrpc.protos import flux_pb2_grpc


def get_parser():
    parser = argparse.ArgumentParser(
        description="Flux Service Endpoint",
        formatter_class=argparse.RawTextHelpFormatter,
    )
    subparsers = parser.add_subparsers(
        help="actions",
        title="actions",
        description="actions",
        dest="command",
    )
    start = subparsers.add_parser(
        "start",
        description="Start the server.",
        formatter_class=argparse.RawTextHelpFormatter,
    )
    start.add_argument(
        "--workers",
        help=f"Number of workers (defaults to {defaults.workers})",
        default=defaults.workers,
        type=int,
    )
    start.add_argument(
        "--port",
        help=f"Port to run application (defaults to {defaults.port})",
        default=defaults.port,
        type=int,
    )
    start.add_argument(
        "--host",
        help="Host to run application (defaults to localhost)",
        default="localhost",
    )
    return parser


class FluxService(flux_pb2_grpc.FluxServiceServicer):
    """
    A FluxService is a grpc service to run jobs with Flux.
    """

    def SubmitJob(self, request, context):
        """
        Submit a job
        """
        print(request)
        response = flux_pb2_grpc.SubmitResponse()
        print(response)
        return response


def serve(args):
    """
    start the Flux Service
    """
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=args.workers))
    endpoint = FluxService()
    flux_pb2_grpc.add_FluxServiceServicer_to_server(endpoint, server)
    host = f"{args.host}:{args.port}"
    server.add_insecure_port(f"{host}")
    print(f"🥞️ Starting flux service at {host}")
    server.start()
    server.wait_for_termination()


def main():
    """
    Light wrapper main to provide a parser with port/workers
    """
    parser = get_parser()

    # If the user didn't provide any arguments, show the full help
    if len(sys.argv) == 1:
        help()

    # If an error occurs while parsing the arguments, the interpreter will exit with value 2
    args, _ = parser.parse_known_args()
    logging.basicConfig()
    serve(args)


if __name__ == "__main__":
    main()
