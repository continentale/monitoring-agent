# monitoring-agent

The monitoring agent is an mulit platform web api for monitoring. It has predefined endpoints to support default monitoring and provides a way to execute user defined scripts for a more precise monitoring
The counterpart (the check for nagios) is published here: https://github.com/continentale/monitoring-agent-check

# installation

You can just download the binary from the release page

# running and examples

There is an example config file which is called config.toml. For a more detailed view in the config file please look at the docs below

## example outputs for the endpoints

here are some examples of all outputs on a linux system (Ubuntu 22.04)

### mem

endpoint: http://localhost:20480/api/v2/mem

output

```json
{
  "total": 13359026176,
  "available": 11675070464,
  "used": 1355632640,
  "usedPercent": 10.147690573699494,
  "free": 8672419840,
  "active": 1037271040,
  "inactive": 3243470848,
  "wired": 0,
  "laundry": 0,
  "buffers": 229568512,
  "cached": 3101405184,
  "writeBack": 0,
  "dirty": 4096,
  "writeBackTmp": 0,
  "shared": 73728,
  "slab": 291307520,
  "sreclaimable": 248287232,
  "sunreclaim": 43020288,
  "pageTables": 18882560,
  "swapCached": 0,
  "commitLimit": 10974478336,
  "committedAS": 2348265472,
  "highTotal": 0,
  "highFree": 0,
  "lowTotal": 0,
  "lowFree": 0,
  "swapTotal": 4294967296,
  "swapFree": 4294967296,
  "mapped": 84787200,
  "vmallocTotal": 35184372087808,
  "vmallocUsed": 27443200,
  "vmallocChunk": 0,
  "hugePagesTotal": 0,
  "hugePagesFree": 0,
  "hugePagesRsvd": 0,
  "hugePagesSurp": 0,
  "hugePageSize": 2097152
}
```

### procs

endpoint: http://localhost:20480/api/v2/procs

output

```json
[
  {
    "name": "init",
    "memoryPercent": 0.004077902,
    "cpuPercent": 0.00047838765760734546,
    "exe": "",
    "status": ["sleep"]
  },
  {
    "name": "sh",
    "memoryPercent": 0.0045991377,
    "cpuPercent": 0.00027342042326018614,
    "exe": "/usr/bin/dash",
    "status": ["sleep"]
  }
]
```

### disks

endpoint: http://localhost:20480/api/v2/disks

output

```json
[
  {
    "usage": {
      "path": "/",
      "fstype": "ext2/ext3",
      "total": 269490393088,
      "free": 229497516032,
      "used": 26232205312,
      "usedPercent": 10.257785123346386,
      "inodesTotal": 16777216,
      "inodesUsed": 926115,
      "inodesFree": 15851101,
      "inodesUsedPercent": 5.520075559616089
    },
    "details": {
      "device": "/dev/sdb",
      "mountpoint": "/",
      "fstype": "ext4",
      "opts": ["rw", "relatime"]
    }
  },
  {
    "usage": {
      "path": "/mnt/wsl",
      "fstype": "tmpfs",
      "total": 6679511040,
      "free": 6679511040,
      "used": 0,
      "usedPercent": 0,
      "inodesTotal": 1630740,
      "inodesUsed": 1,
      "inodesFree": 1630739,
      "inodesUsedPercent": 0.00006132185388228658
    },
    "details": {
      "device": "tmpfs",
      "mountpoint": "/mnt/wsl",
      "fstype": "tmpfs",
      "opts": ["rw", "relatime"]
    }
  }
]
```

### load

endpoint: http://localhost:20480/api/v2/load

output

```json
{
  "load1": 0,
  "load5": 0.2,
  "load15": 0.29
}
```

### time

endpoint: http://localhost:20480/api/v2/time

output

```json
{
  "timestamp": 1668715878,
  "formatted": "Thu Nov 17 21:11:18 CET 2022"
}
```

### cpus

endpoint: http://localhost:20480/api/v2/cpus

output

```json
[
  {
    "cpu": "cpu-total",
    "user": 365.98,
    "system": 197.14,
    "idle": 176918.48,
    "nice": 4.1,
    "iowait": 54.78,
    "irq": 0,
    "softirq": 71.52,
    "steal": 0,
    "guest": 0,
    "guestNice": 0
  }
]
```

endpoint: http://localhost:20480/api/v2/cpus?perCPU=true

output

```json
[
  {
    "cpu": "cpu0",
    "user": 30.32,
    "system": 46.6,
    "idle": 14724.38,
    "nice": 0.16,
    "iowait": 4.46,
    "irq": 0,
    "softirq": 50.53,
    "steal": 0,
    "guest": 0,
    "guestNice": 0
  },
  {
    "cpu": "cpu1",
    "user": 28.05,
    "system": 10.56,
    "idle": 14745.15,
    "nice": 1.73,
    "iowait": 29.49,
    "irq": 0,
    "softirq": 7.74,
    "steal": 0,
    "guest": 0,
    "guestNice": 0
  }
]
```

endpoint: http://localhost:20480/api/v2/file?name=environment

output

```json
PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin"
PATH=$PATH:/usr/local/go/bin
```

endpoint: http://localhost:20480/api/v2/file?name=fstab

output

```json
{
  "Path": "/etc/fstab",
  "IsDir": false,
  "ModTime": 1629409395,
  "Mode": "-rw-r--r--",
  "Name": "fstab",
  "Size": 43,
  "Content": "LABEL=cloudimg-rootfs\t/\t ext4\tdefaults\t0 1\n"
}
```

endpoint: http://localhost:20480/api/v2/file?name=fstab

output

### the file endpoint

the file endpoint needs a special config. This config prevents users to read any data they want, so you must make the file accesible.

For this you have to go in the file section and enable the endpoint with

```toml
[file]
enabled = true
```

next you define your entries with two possible formats

```toml
# use this format if you just want to specify the path
the FILE_NAME is just the name for the agent and must not match the real file name
FILE_NAME = path

# the other option is to provide the name directly and pass an path value

full example:
[file]
enabled = true
  [file.entries]
  	fsTab = "/etc/fstab"

    [file.entries.environment]
    path = "/etc/environment"
    contentOnly = true
```

### the exec endpoint

the exec endpoint needs a special config too to prevent the user exec all files on your system

to begin with a exec endpoint you should enable it:

the file endpoint needs a special config. This config prevents users to read any data they want, so you must make the file accesible.

For this you have to go in the file section and enable the endpoint with

```toml
[exec]
enabled = true
shell = "/bin/bash"
```

here you can define a shell if you want in where the file gets executed. If you leave it empty the default will be taken (windows: cmd | linux: /bin/bash)

Then you define a name for your command to reference it in the agent and specify the path to file
If you want to define the entrie directly you can do that by setting a path to your script in the section

full example:

```toml
[exec]
enabled = true
shell = "/bin/bash"
  [exec.entries]
    file_tester = "-c ./scripts/echo"
    [exec.entries.test]
    shell = "/usr/bin/python3"
    path = "./scripts/echo"
```

# configuration flags

```toml
# under server you can change the behavior of the server itself
[server]
# its the protocol used by the server. valid values are: http|https
protocol = "http"
# if https is enabled that are the paths to the certificates
certificate = "certificates/cert.crt"
key = "certificates/priv.key"

# use a secret in the client to get data
secret = "superSecretPassword"
useSecret = false

# set the port for the webserver
port = 20480
# sets the address for the webserver to listen on. If empty all interfaces will be used
address = ""

# under global there are the default values for all endpoints
[global]
timeStringFormat = "Mon Jan _2 15:04:05 MST 2006"

# configure the mem endpoint
[mem]
# disable or enable the endpoint
enabled = true

# configure the procs endpoint
[procs]
# disable or enable the endpoint
enabled = true

# configure the disks endpoint
[disks]
# disable or enable the endpoint
enabled = true

# configure the load endpoint
[load]
# disable or enable the endpoint
enabled = true

# configure the time endpoint
[time]
# disable or enable the endpoint
enabled = true

# configure the cpus endpoint
[cpus]
# disable or enable the endpoint
enabled = true
# turns of the per cpu view and summarizes all cpus in one output
# can configured in the api call as http-get parameter
perCPU = false
´
# configure the file endpoint
[file]
# disable or enable the endpoint
enabled = true
  [file.entries]
  	fsTab = "/etc/fstab"

    [file.entries.environment]
    path = "/etc/environment"
    contentOnly = true

[exec]
# disable or enable the endpoint

enabled = true
shell = "/bin/bash"
  [exec.entries]
    file_tester = "-c ./scripts/echo"
    [exec.entries.test]
    shell = "/usr/bin/python3"
    path = "./scripts/echo"
```

# TODOs or Roadmap

- lots of documentation
- writing tests
- react on feedback
