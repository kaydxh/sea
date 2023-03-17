#!/usr/bin/env bash

# Fail on any error.
set -euo pipefail
# set -o xtrace

function signal_handler()
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

# reset signal handle to signal_handle
trap signal_handler SIGINT SIGQUIT SIGKILL SIGTERM SIGABRT SIGSEGV

./bin/seadate --config ./conf/seadate.yaml
