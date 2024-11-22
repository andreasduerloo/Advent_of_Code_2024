#!/bin/bash

podman build -t advent .
podman run --rm advent:latest /bin/advent $1