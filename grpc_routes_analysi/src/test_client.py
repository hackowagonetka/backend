from __future__ import print_function

import logging

import grpc
from grpc_generated_files import (
    RoutesAnalysis_pb2,
    RoutesAnalysis_pb2_grpc
)


def run():
    print("Trying to send a request")
    with grpc.insecure_channel('localhost:7878') as channel:
        stub = RoutesAnalysis_pb2_grpc.RoutesAnalysisStub(channel)
        response = stub.Analyse(
            RoutesAnalysis_pb2.AnalyseRequest(
                cargo_filled=6,
                cargo_total=10,
                distance=15000,
                timestamp=1671349762
            )
        )
    print(f"Client received: {response}")


if __name__ == '__main__':
    logging.basicConfig()
    run()
