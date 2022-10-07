#!/usr/bin/env bash

cd $(dirname $0)

bass build.bass | bass --export | tar -xf - -C bin
