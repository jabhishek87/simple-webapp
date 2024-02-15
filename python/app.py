import os
from flask import Flask
from flask import render_template
import logging
import redis
import socket
import random
import os

app = Flask(__name__)

logging.basicConfig(filename='app.log', level=logging.DEBUG)

appColor = os.environ.get("APP_COLOR") or "white"

# connect to redis
redis_host =  os.environ.get("REDIS_HOST") or "localhost"
redis_cache = redis.Redis(host=redis_host, port=6379)

@app.route("/")
def main():
    #return 'Hello'
    print(appColor)
    return render_template(
        'index.html', name=socket.gethostname(), color=appColor
    )

@app.route('/color/<appColor>')
def new_color(appColor):
    return render_template(
        'index.html', name=socket.gethostname(), color=appColor
    )

@app.route('/readlogs')
def read_file():
    with open("app.log") as f:
        log_content = f.read()
    print(log_content)
    return render_template(
        'index.html', name=socket.gethostname(), color=appColor,
        log_content=log_content,
    )

@app.route('/counter')
def some_def():
    try:
        redis_cache.ping()
        data = "You've visited me {} times.\n".format(redis_cache.incr('hits'))
    except redis.ConnectionError as e:
        data = "Could not able to connect REDIS: "+ str(e)

    return render_template(
        'index.html', name=socket.gethostname(), color=appColor,
        data=data
    )

if __name__ == "__main__":
    app.run(host="0.0.0.0", port="8000")


