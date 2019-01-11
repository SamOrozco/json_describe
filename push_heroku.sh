#!/bin/bash

docker build -t json_describe .

heroku container:push json_describe --app json-describe

heroku container:release json_describe --app json-describe

