# go-eqfeed
This will pull USGS earthquake JSON and output a RSS feed to the console - there's more to do with this and that will start to happen soon.

## Update: 9/4/2023
I found an old Pi Zero W. After a few feeble attempts to load 1.21.0 on it (what this code was originally written in) I gave up and used the `sudo apt install golang` command.
That installed 1.15.15 so I had a bit of refactoring - see the commented code for the 1.15 bits.

This project was based upon "#315 (2023-08-11) Weekend Project Edition" of [The Daily Drop](https://dailyfinds.hrbrmstr.dev/archive)
There were a couple of changes that I had to make from Bob's code to get it to run in v1.21.0 but I didn't document them.

To run this on your system I suggest removing go.mod and go.sum then run the go init command that's appropriate for your setup.