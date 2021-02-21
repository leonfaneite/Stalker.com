FROM python:3.9.0

RUN mkdir /app
WORKDIR /app

RUN apt-get -y update
RUN apt-get -y upgrade
RUN pip install --upgrade pip



COPY requirements.txt .
COPY . .
RUN pip install -r requirements.txt
RUN pip install tweepy
RUN pip install pymongo


ENV PYTHONUNBUFFERED 1

CMD ["python3", "twiiter.py"]
