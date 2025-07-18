---
title: "rclone version"
description: "Show the version number."
versionIntroduced: v1.33
# autogenerated - DO NOT EDIT, instead edit the source code in cmd/version/ and as part of making a release run "make commanddocs"
---
# rclone version

Show the version number.

## Synopsis

Show the rclone version number, the go version, the build target
OS and architecture, the runtime OS and kernel version and bitness,
build tags and the type of executable (static or dynamic).

For example:

    $ rclone version
    rclone v1.55.0
    - os/version: ubuntu 18.04 (64 bit)
    - os/kernel: 4.15.0-136-generic (x86_64)
    - os/type: linux
    - os/arch: amd64
    - go/version: go1.16
    - go/linking: static
    - go/tags: none

Note: before rclone version 1.55 the os/type and os/arch lines were merged,
      and the "go/version" line was tagged as "go version".

If you supply the --check flag, then it will do an online check to
compare your version with the latest release and the latest beta.

    $ rclone version --check
    yours:  1.42.0.6
    latest: 1.42          (released 2018-06-16)
    beta:   1.42.0.5      (released 2018-06-17)

Or

    $ rclone version --check
    yours:  1.41
    latest: 1.42          (released 2018-06-16)
      upgrade: https://downloads.rclone.org/v1.42
    beta:   1.42.0.5      (released 2018-06-17)
      upgrade: https://beta.rclone.org/v1.42-005-g56e1e820

If you supply the --deps flag then rclone will print a list of all the
packages it depends on and their versions along with some other
information about the build.


```
rclone version [flags]
```

## Options

```
      --check   Check for new version
      --deps    Show the Go dependencies
  -h, --help    help for version
```

See the [global flags page](/flags/) for global options not listed here.

## See Also

* [rclone](/commands/rclone/)	 - Show help for rclone commands, flags and backends.

