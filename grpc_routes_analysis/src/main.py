from datetime import datetime
from concurrent import futures
import logging

import grpc

from grpc_generated_files import (
    RoutesAnalysis_pb2,
    RoutesAnalysis_pb2_grpc

)
from tmodel.time_model import Tmodel


class Routes(RoutesAnalysis_pb2_grpc.RoutesAnalysisServicer):
    def Analyse(self, request, context):
        date = request.timestamp
        date = datetime.fromtimestamp(date)
        week = date.isocalendar()[1]
        model = Tmodel()
        if request.cargo_filled * 2 > request.cargo_total:
            is_train_filled = 1
        else:
            is_train_filled = 0
        params = [
                date.year,
                date.month,
                week,
                date.day,
                date.hour,
                is_train_filled,
                request.distance,
            ]
        time_spent = model.predict(params)
        return RoutesAnalysis_pb2.AnalyseResponse(
            time_spent=time_spent[0]
        )


def serve():
    port = '7878'
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    RoutesAnalysis_pb2_grpc.add_RoutesAnalysisServicer_to_server(
        Routes(), server)
    server.add_insecure_port('[::]:' + port)
    server.start()
    print(f"server listening on port {port}")
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
