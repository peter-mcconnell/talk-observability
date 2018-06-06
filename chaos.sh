#!/bin/bash - 

docker exec -i $(docker ps --filter "name=talk-observability_app1_" -q) sh -c "echo 'corrupting network ...'; tc qdisc add dev eth0 root netem corrupt 50%; echo 'done...'"
