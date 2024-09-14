FROM ubuntu

RUN apt-get update
RUN apt-get install python3 -y
RUN python3 --version ; pip --version

RUN pip install flask
RUN pip install flask-mysql

COPY . /opt/source-code

ENTRYPOINT FLASK_APP=/opt/source-code/app.py flask run
