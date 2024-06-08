---
title: Gokrazy WSL and photos
date: 2024-04-28
---

Very excited to be getting gokrazy to work as a deployment platform for devices.  My first project
is to use our Samsang smart screen as a photo display.

It has taken a little while to work out how to use effectively use gokrazy from WSL and here are some of my notes.

1) When building the initial image you can use `-file x.img` to output to a file rather than 
a device which allows you to export the img fril from wsl and burn on Windows.  Attaching an Sd card
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

