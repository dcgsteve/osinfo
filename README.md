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

# Scheduling daily on Edge Host
1. Ensure that environment variables REG_USER and REG_PASS are exported under user root or globally
```shell
#values will be printed, if values are empty cron job will not work
echo $REG_USER
echo $REG_PASS
```
2. Copy file provided in cron-job folder to /etc/cron.daily/
```shell
sudo cp cron-job/stats /etc/cron.daily
sudo chmod 0755 /etc/cron.daily/stats
```
3. Edit file /etc/cron.daily/stats with any text editor like vi or nano
```shell
sudo vi /etc/cron.daily/stats
#Replace FILL-STAT-PACKAGE-PKG with stats package name, save and exit. 
```
4. To make sure that scripts will work, run it manually
```shell
sudo /etc/cron.daily/stats
```
5. For Testing purposes
```shell
#Note replace xx with size in Gigabyte to create file which will occupy 85% or more of any partition on edge
head -c xxG </dev/urandom >testfile
#if version of head tool not accepting size in G format, provide same amount in bytes to fill 85% disk on edge
head -c 1073741824 </dev/urandom >testfile #example is 1G size in bytes
```
6. Check next day /var/log/messages
```shell
sudo tail /var/log/messages
#you will see log about storage 
```