#!/bin/bash

home=/home/sberm/yagoo
cd home
nohup ./yagoo &
cd $home/web/web-app
nohup npm run dev &


