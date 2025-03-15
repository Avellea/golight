# golight
Basic userspace driver for controlling display backlight on legacy Nvidia based MacBooks.

`golight` has been tested on a MacBook 5,2 running Arch Linux.

## Installation
Installation is a simple one-liner.

```bash
$ make clean; make; sudo make install
```

It is possible you may need to change the group/owner of the `nv_backlight/brightness` file to a group you have access to.

## Usage

```bash
$ golight <inc/dec> # Current implementation will increase/decrease by 5. Configuration coming soon.
```
