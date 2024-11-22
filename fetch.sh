#!/bin/bash

token=$(cat ./token)
curl -s --cookie "session=$token" https://adventofcode.com/2024/day/$1/input -o ./inputs/$2.txt