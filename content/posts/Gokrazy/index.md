---
title: Gokrazy WSL and photos
date: 2024-04-28
tags: [gophoto,go,gokrazy]
---

Very excited to be getting gokrazy to work as a deployment platform for devices.  My first project
is to use our Samsung smart screen as a [photo display](https://github.com/drummonds/gophoto).

It has taken a little while to work out how to use effectively use gokrazy from WSL and here are some of my notes.

1) When building the initial image you can use `gok -i hello overwrite --full x.img --target_storage_bytes=1258299392` to output to a file rather than 
a device which allows you to export the img from from wsl and burn on Windows.  Attaching an sd card
to the WSL instance I found tricky.
2) To cross compile and update I used `CGO_ENABLED=0 GOARCH=arm64 gok -i hello update` for a raspberry pi 3B+.




## Solved Problems

- Using Task to store required cross compile flag configuration.
- Changing DNS server on router to change DNS entries of DHCP names.

## Current Problems

### Monitor fails after a time on my Samsung 95Q55 screen. 

It works after uploading but then goes to a state overnight (power saving ?)  which the TV doesn't recognise.

### Changing name/PW

Once you have created an image and a password can you change them?


## Learning

### setting /perm

I set a small image and there was a command I used to extend it on my raspberry pi but not yet
sure how to use a USB as a permanent store so when fails doesn't affect main memory.

You can't use the extraFiles in config to write the /perm area so get a file there you either have to 
do breakinglass or add it in the app.

### Using Raspberry pi graphic modes.

I want to display pictures as high resolution photo albums.  At the moment I am just getting the first image to display.  The default image is I thinking HD at 16b its per pixel.  This means only 5 or 6 per colour.  There are two improvements to be done:
- Switch to a higher resolution mode of 32 bits per pixel perhaps at 4K which I think the raspberry pi 4 supports.  This can be done I think by change the config.txt on boot up;
- Use 10 bit colour.  This is a lot harder and may involve:
  -  custom graphics drivers
  -  Switching to RGBA64 as basic colours

One thing that will make it easier is that I don't care about the refresh rate that much so as long as the buffer can out put the data at the required refresh rate, the CPU has plenty of time to calculate the next screen.

See the following links

- 10bit format for YUV420 https://forums.raspberrypi.com/viewtopic.php?t=342120
- good summary https://psychtoolbox.discourse.group/t/raspberry-pi-4-high-bit-depth-output/4824/2

### Linux Frame Buffer

The Kernel frame buffer documentation [is here](https://www.kernel.org/doc/html/latest/fb/api.html) at www.kernel.org and describes the frame buffer api.

## Journey

### Hello world

The most basic program.  you can see that it has run in the online web status
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Gophoto start V0.0.1")
}


```

### Hello world web

Now you can see hello world on Web:

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Gophoto start V0.0.1")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s! from gophoto", r.URL.Path[1:])
}
```

