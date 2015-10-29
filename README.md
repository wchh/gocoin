# Binaries

Windows or Linux (amd64) binaries of Gocoin release 1.0.0 available for download

 * https://sourceforge.net/projects/gocoin/files/?source=directory

# About Gocoin

**Gocoin** is a full **Bitcoin** solution written in Go language (golang).

The wallet combined with `balio` tool also provides a working solution for **Litecoin**.

The software's architecture is focused on maximum security and good performance.

The **client** (p2p node) is an application independent from the **wallet**.

The **wallet** is deterministic and password seeded.
As long as you remember the password, you do not need any backups of your wallet.

There is additional tool called **downloader** that
can quickly sync (download) the blockchain state from the p2p network.
Use it for the initial blockchain download or to sync your **client** after having it offline for a longer time.

In addition there is also a set of more and less useful tools.
They are all inside the `tools/` folder.
Each source file in that folder is a separate tool.


# Documentation
The official web page of the project is served at <a href="http://gocoin.pl">gocoin.pl</a>
where you can find extended documentation, including **User Manual**.


# Requirements

## Hardware

**client** / **downloader**:

* 64-bit architecture OS and Go compiler.
* File system supporting files larger than 4GB.
* At least 4GB of system memory highly recommended.


**wallet**:

* Any platform that you can make your Go (cross)compiler to build for (Raspberry Pi works).
* For security reasons make sure to use encrypted swap file (if there is a swap file).
* If you decide to store your password in a file, have the disk encrypted (in case it gets stolen).


## Operating System
Having hardware requirements met, any target OS supported by your Go compiler will do.
Currently that can be at least one of the following:

* Windows
* Linux
* OS X
* Free BSD

## Build environment
Since no binaries are provided, in order to build Gocoin yourself, you will need the following tools installed in your system:

* **Go** (version 1.2 or higher) - http://golang.org/doc/install
* **Git** - http://git-scm.com/downloads
* **Mercurial** - http://mercurial.selenic.com/

If the tools mentioned above are all properly installed, you should be able to execute `go`, `git` and `hg`
from your OS's command prompt without a need to specify full path to the executables.


### Optionally

It is recommended to have `gcc` complier installed in your system, to get advantage of performance improvements and memory usage optimizations.

For Windows install `mingw`, or rather `mingw64` since the client node needs 64-bit architecture.

 * http://mingw-w64.sourceforge.net/


# Getting sources

Use `go get` to fecth and install the source code files.
Note that source files get installed within your GOPATH folder.

## Dependencies

Two extra packages are needed, which are not included in the standard set of Go libraries.
You need to install them before building Gocoin.
In order to do this, simply execute the following commands:

	go get golang.org/x/crypto/ripemd160
	go get github.com/golang/snappy/snappy

## Gocoin
Use `go get` to fetch and install Gocoin sources for you:

	go get github.com/wchh/gocoin


# Building

## Client node
Go to the `client/` folder and execute `go build` there.


Not having a compatible `gcc` installed in your system, you will likely see an error like this:

	# github.com/wchh/gocoin/lib/qdb
	exec: "gcc": executable file not found in %PATH%

You can go on without *gcc*, although your running client (and downloader) will then need some more system memory.

To go on without *gcc*, copy file `lib/qdb/no_gcc/membind.go` one folder up (overwriting the original `lib/qdb/membind.go`).

## Wallet and downloader
Go to the `wallet/` folder and execute `go build` there.

Go to the `downloader/` folder and execute `go build` there.

## Tools
Go to the `tools/` folder and execute:

	go build btcversig.go

Repeat the `go build` for each source file of the tool you want to build.

# Development
Although it is an open source project, I am sorry to inform you that I will not merge in any pull requests.
The reason is that I want to stay an explicit author of this software, to keep a full control over its
licensing. If you are missing some functionality, just describe me your needs and I will see what I can do
for you. But if you want your specific code in, please fork and develop your own repo.
