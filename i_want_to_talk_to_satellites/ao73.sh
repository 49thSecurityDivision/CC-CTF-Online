#!/bin/bash

timeout=15
maxchildren=10
backlog=5

socat -v -T "$timeout" TCP-LISTEN:4533,reuseaddr,fork,max-children="$maxchildren",backlog="$backlog" EXEC:"python3 azimuth.py"
