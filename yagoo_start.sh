#!/bin/bash

nohup ./yagoo &
cd ./web/web-app
nohup npm run dev &
