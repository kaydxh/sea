#!/usr/bin/env bash

# Fail on any error.
set -euo pipefail
# set -o xtrace

function signal_handle()
{
    # get task process id
    local pids=`jobs -p`
    if [[ "x${pids}" != "x" ]]; then
      for pid in "${pids}"
      do
        kill -SIGTERM "${pid}" 
      done
    fi
}

# signal handle
trap signal_handle SIGINT SIGQUIT SIGKILL SIGTERM SIGABRT SIGSEGV

./bin/sealet --config ./conf/sealet.yaml
