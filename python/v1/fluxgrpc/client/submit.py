import argparse
import sys

import flux
from flux.cli import base

import fluxgrpc.defaults as defaults

from .client import FluxClient


class SubmitCmd(base.SubmitBulkCmd):
    """
    This is the same as the Flux SubmitCmd, but we add the host and port
    of the grpc service

    Usage: flux submit [OPTIONS] cmd ...
    """

    def __init__(self, prog, usage=None, description=None):
        super().__init__(prog, usage, description)
        self.parser.add_argument(
            "command", nargs=argparse.REMAINDER, help="Job command and arguments"
        )
        self.parser.add_argument(
            "--port",
            help=f"Port to run application (defaults to {defaults.port})",
            default=defaults.port,
            type=int,
        )
        self.parser.add_argument(
            "--host",
            help="Host to run application (defaults to localhost)",
            default="localhost",
        )

    def submit_async(self, args, jobspec=None):
        """
        Override submit async to always return the jobspec.
        """
        if jobspec is None:
            jobspec = self.jobspec_create(args)
        return jobspec.dumps()

    def submit(self, args, jobspec=None):
        return self.submit_async(args, jobspec)


def main():
    sys.stdout = open(sys.stdout.fileno(), "w", encoding="utf8", errors="surrogateescape")
    sys.stderr = open(sys.stderr.fileno(), "w", encoding="utf8", errors="surrogateescape")

    # Prepare the submit parser
    submit = SubmitCmd(
        "flux submit",
        description="enqueue a job",
    )
    parser = submit.get_parser()
    parser.set_defaults(func=submit.submit)
    args = parser.parse_args()

    # This is a json string
    jobspec = args.func(args)
    cli = FluxClient(f"{args.host}:{args.port}")
    response = cli.submit(jobspec)
    print(f"Submit response: {response.status}")


if __name__ == "__main__":
    main()
