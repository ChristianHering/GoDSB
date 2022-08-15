GoDSB
===========

This repository is a fully functional [DWM](https://dwm.suckless.org/) statusbar.

It provides:
  * A fast and extensible way to create a [DWM](https://dwm.suckless.org/) statusbar in [Go](https://golang.org/).
  * An alternative to [dwmblocks](https://github.com/torrinfail/dwmblocks) for clickable statusbar sections.
  * A way to update sections of your statusbar without running all scripts.

Table of Contents:

  * [About](#about)
  * [Installing and Compiling from Source](#installing-and-compiling-from-source)
  * [Contributing](#contributing)
  * [License](#license)

About
-----

GoDSB implements similar functionality to dwmblocks while giving you full access to Go's vast standard library. This allows easy and fast REST calls, interprocess calls, shell access, and much more. GoDSB doesn't intend on replacing dwmblocks; it only seeks to expand the options for those trying out DWM.

Installing and Compiling from Source
------------

The easiest way to try out GoDSB is to download and run the latest release from the [releases](https://github.com/ChristianHering/GoDSB/releases) tab or download and compile the source.


If you're looking to compile from source, you'll need:

  * [Go](https://golang.org) installed and [configured](https://golang.org/doc/install)
  * A little patience :)

As well as port 1058 available on the host.

Contributing
------------

Contributions are always welcome. If you're interested in contributing, send me an email or submit a PR.

License
-------

This project is currently licensed under GPLv3. This means you may use our source for your own project, so long as it remains open source and is licensed under GPLv3.

Please refer to the [license](/LICENSE) file for more information.
