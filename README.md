go-bind9masq
===============

This project aims to match domain names in logs from dictionnaries of well known domain names.

For example, dns queries from dns server Bind is a good candidate.

̀Script `update.sh` is an easy way to fetch famous dictionnaries maintained by Toulouse 1 Capitole University.

File `main.go` is the entrypoint.

## Prerequisites
- docker or go
- make
- rsync

## Quick start

### Build
Command `make` to build arm or amd64 binary and fetch updates from dictionaries.
```
make build-amd64
# or
make build-amr64
```

### (Un)Install

```
make install
make uninstall
```

### Run
Edit config.yml to fit your needs.
```yaml
# Refers to your log queries.
bind9:
  queries: "/var/log/named/queries.log"
categories:
   # set categories you want to check in bind9.queries log file
  toCheck:
  # set categories you want to be protected
  toProtect:
```

Then execute go-bind9masq binary : 
- `go-bind9masq s` : to show domains categories you wanted to check
- `go-bind9masq u` : to sinkhole your dns queries with toProtect property

### Bind9 configuration for sinkhole
Include blaklisted.zones file to named.conf.local
```
include "/etc/bind/blacklisted.zones";
```

Create blacklisted.db file, replace here 8.8.8.8 with your sinkhole web api !
```
;
; BIND data file for local loopback interface
;
$TTL    604800
@       IN      SOA     local. root.local. (
                        3041023         ; Serial
                         604800         ; Refresh
                          86400         ; Retry
                        2419200         ; Expire
                         604800 )       ; Negative Cache TTL
;
@       IN      NS      local.
@       IN      A       8.8.8.8
@       IN      AAAA    ::1
* IN A 8.8.8.8
* IN AAAA ::1
```

And let `go-bind9masq` binary create zones file.

## TODO

Maybe use these lists
```
StevenBlack -> https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts
MalwareDom  -> https://mirror1.malwaredomains.com/files/justdomains
DisconTrack -> https://s3.amazonaws.com/lists.disconnect.me/simple_tracking.txt
DisconAd    -> https://s3.amazonaws.com/lists.disconnect.me/simple_ad.txt
```