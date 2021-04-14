# SIMPLE REDIS CLI

## HOW TO USE
Edit `.env` File, replace `REDIS_HOST`,`REDIS_PORT`,`REDIS_PASSWORD` and `REDIS_DB` value with your REDIS configuration
```
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
```

* Linux
    ```
    $ ./redis-linux-64

    ```
* Mac
    ```
    $ ./redis-mac-64

    ```
* Windows
    ```
    $ redis-win-64.exe

    ```

## HELP
```
USAGE:
  apps [arguments...]
COMMANDS:
  SET [redis-key] [data-type] [value]
  GET [redis-key] [data-type]
DATA TYPES:
  STRING escape strings
  TIME RFC3339 format, ex: 2021-04-14T08:09:47Z
```
