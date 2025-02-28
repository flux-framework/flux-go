#!/usr/bin/env python

# This is the ensemble client init, where we can start a client for a particular
# queue backend (like Flux)

import argparse
import sys

import fluxgrpc
import fluxgrpc.defaults as defaults


def get_parser():
    parser = argparse.ArgumentParser(
        description="Flux GRPC Python",
        formatter_class=argparse.RawTextHelpFormatter,
    )
    parser.add_argument(
        "--version",
        dest="version",
        help="show software version.",
        default=False,
        action="store_true",
    )

    description = "actions for Flux GRPC Python"
    subparsers = parser.add_subparsers(
        help="actions",
        title="actions",
        description=description,
        dest="command",
    )

    subparsers.add_parser("version", description="show software version")

    # Submit a job to flux
    submit = subparsers.add_parser(
        "submit",
        description="run the ensemble",
        formatter_class=argparse.RawTextHelpFormatter,
    )
    submit.add_argument(
        "--env",
        help="Environment variable to include",
        action="add",
        default=[],
    )
    submit.add_argument(
        "-N" "--nodes",
        default=1,
        help="Number of nodes for the job",
    )
    submit.add_argument(
        "--debug",
        help="Enable debug logging for the config",
        action="store_true",
        default=False,
    )
    submit.add_argument(
        "--port",
        help=f"Port to run application (defaults to {defaults.port})",
        default=defaults.port,
        type=int,
    )
    submit.add_argument(
        "--host",
        help="Host to run application (defaults to localhost)",
        default="localhost",
    )
    return parser


def run_():
    parser = get_parser()

    def help(return_code=0):
        version = ensemble.__version__

        print("\nFlux Service v%s" % version)
        parser.print_help()
        sys.exit(return_code)

    # If the user didn't provide any arguments, show the full help
    if len(sys.argv) == 1:
        help()

    # If an error occurs while parsing the arguments, the interpreter will exit with value 2
    args, extra = parser.parse_known_args()

    # Show the version and exit
    if args.command == "version" or args.version:
        print(ensemble.__version__)
        sys.exit(0)

    # retrieve subparser (with help) from parser
    helper = None
    subparsers_actions = [
        action for action in parser._actions if isinstance(action, argparse._SubParsersAction)
    ]
    for subparsers_action in subparsers_actions:
        for choice, subparser in subparsers_action.choices.items():
            if choice == args.command:
                helper = subparser
                break

    # Does the user want a shell?
    if args.command == "run":
        from .run import main

    # Pass on to the correct parser
    return_code = 0
    try:
        main(args=args, parser=parser, extra=extra, subparser=helper)
        sys.exit(return_code)
    except UnboundLocalError:
        return_code = 1
    help(return_code)


if __name__ == "__main__":
    run_ensemble()
