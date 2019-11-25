go-dns-match
===============

This project aims to match domain names in logs from dictionnaries of well known domain names.

For example, dns queries from dns server Bind is a good candidate.

̀Script `update.sh` is an easy way to fetch famous dictionnaries maintained by Toulouse 1 Capitole University.

File `main.go` is the entrypoint.

## Prerequisites
- docker
- make
- rsync

## Quick start

Command `make` to build arm binary and fetch updates from dinctionnaries.

See the target build-amd64 in the `Makefile`, if you want to build for this architecture. 
Documentation here : https://hub.docker.com/_/golang/