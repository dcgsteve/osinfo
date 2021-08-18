# osinfo
Tool is responsible to check edge nodes CPU, Memory and other stats and in case of treshhold violation generate critical log message to local Edge host syslog server via umix 
socket

# Required Environment Variables during package build
| Variables | Description |
|---|---|
| ART_REG_URL | Artisan registry URL  |
| ART_REG_USR | Artisan registry user |
| ART_REG_PWD | Artisan registry password |
| PKG_NAME | package name, for example: stat  |

# Build app, package and push to registry
```shell
art run build-package
```
><b>Note: </b> First fill in build.yaml environment vars

# How to use:
```shell
art exe <PACKAGE_NAME> run-stats
```
