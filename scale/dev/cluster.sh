#! /bin/bash

if [[ -z "$1" ]]; then
  >&2 echo "usage: $0 <cluster name>"
  exit 1
fi


infractl create gke-default $1 --arg machine-type=e2-standard-16 --arg nodes=6
