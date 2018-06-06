#!/bin/bash - 
#===============================================================================
#
#          FILE: chaos.sh
# 
#         USAGE: ./chaos.sh 
# 
#   DESCRIPTION: This script will corrupt the network on one of the golang
#                apps in this project. The purpose of this is to inspect the
#                impact in prometheus
# 
#       OPTIONS: ---
#  REQUIREMENTS: ---
#          BUGS: ---
#         NOTES: ---
#        AUTHOR: Peter McConnell (me@petermcconnell.com), 
#  ORGANIZATION: 
#       CREATED: 06/06/2018 21:57
#      REVISION:  ---
#===============================================================================

set -o nounset                              # Treat unset variables as an error
set -e

docker exec -i $(docker ps --filter "name=talk-observability_app1_" -q) sh -c "echo 'corrupting network ...'; tc qdisc add dev eth0 root netem corrupt 50%; echo 'done...'"
