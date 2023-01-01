#!/bin/bash

set -o errexit

./migrate
./createuser
./app
