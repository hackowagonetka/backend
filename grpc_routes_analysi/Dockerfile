FROM python:3.8-slim

WORKDIR /app

RUN apt-get -y update --fix-missing
RUN apt-get -y install apt-utils
RUN apt-get -y dist-upgrade
RUN apt-get -y install gcc

COPY ./grpc_routes_analysis/requirements.txt ./grpc_routes_analysis/requirements.txt
RUN pip3 install -r ./grpc_routes_analysis/requirements.txt

RUN apt-get -y clean
COPY . /app

ENTRYPOINT ["python", "./grpc_routes_analysis/src/main.py"]