# coding=utf-8
import requests
import os
import time

isFirst = True


def handler(event, context):
    global isFirst
    if isFirst:
        print("first sleep 1s")
        time.sleep(1)
        isFirst = False
    url = os.environ['KEEP_WARM_URL']
    res = requests.get(url)
    print(res.status_code)
