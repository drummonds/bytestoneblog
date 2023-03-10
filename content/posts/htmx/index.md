---
title:  HTMX applications
date: 2023-03-10
tags: ["htmx","cli","golang","python"]
--- 

I have been wondering about finding a simpler development environment that allows me to do graphic programming.  My first successful tool was Delphi which allowed me to generate tools with a simple Pascal like language and a great IDE.  The interface was the simple VCL ([Visual Control Library][]).

[Visual Control Library]; https://docwiki.embarcadero.com/RADStudio/Sydney/en/VCL_Overview

Over time I have moved languages and currently I like writing in Python for expressiveness and Golang for simplicity and robustness.

I have been using [PyWebIO][] as a way of writing simple visualisations using markdown and HTML rather than plain text.  I had a brief forray in [Streamlit][] which was very sick but retreated when looking at the potential telemetry and focus on external hosting plumbed in.

[PyWebIO]: https://www.pyweb.io/
[Streamlit]: https://streamlit.io/

## Command line tools

I next was thinking about turning command line tools into guis using this type of format.  I have been using [argparse][] (and [akamensky/argparse][] as a golang equivalent) to hold the user interface of the command line tools.  There are a number of examples which do this using this eg [argparseui][] and [gooey][] but they both focus on the command line.  I used a long time ago [wooey][] which was based on Django and worked quite well.

[argparse]:https://docs.python.org/3/library/argparse.html
[akamensky/argparse]:https://github.com/akamensky/argparse
[wooey]:https://github.com/wooey/Wooey
[argparseui]:https://github.com/shimpe/argparseui
[gooey]:https://github.com/chriskiehl/Gooey

## Dynamic reporting

The other application is dynamic reporting.  I was wondering if an approach using htmx could be used.  With python it gets complicated and my experience has been poor when I have tried that.  With Go it has been great as I think the abstractions are firmly founded in Tony Hoares, communicating sequential processes.  

A possible solution would be redo pywebio using a golang underlying server.  The python would then just communicate with putting HTML or getting queries by blocking sockets to a local go server which would do the interaction with the browser and [htmx][]

[htmx]: https://htmx.org/

## Native deployment

Another strand of thinking has been around static web pages and wasm.  There is an example using Rust by [Richard Anaya][]

[Richard Anaya]:https://github.com/richardanaya/wasm-service

This could utilise the golang / htmx interaction although a bit trickier for python.  This also allows the possibliy of semi statically hosting secure content.  You just encrypt and embed in the wasm.  You can then authenticate the user to be able to decrypt it. The authentication service would need to be live rather than static.  I was thinking of using it for geneological data.
