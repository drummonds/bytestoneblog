---
title: Go CAD versions
date: 2023-02-23
updated: 2025-01-10
mermaid: true
---
I am keen to try out writing and designing in go CAD models.  I am on a drive to simplify and unify my coding styles.  My process for designing  I would much rather be creating my models by programming. So I started with
OpenScad. I discovered that as the models got very complicated it would start to have problems particularly with
the difference function.

There was a great go library called [ghostscad][] that generated OpenSCAD language but as it transpiles to OpenSCAD
it suffers from the same problems as complexity increases.  As of Jan 2025 OpenScad has more or less stopped being
updated with the latest release in 2021.  There is some work still going on though.


## Signed distance functions

There is a family of projects using signed distance functions for CAD in go.

{{< mermaid >}}
flowchart LR
  sdfx(**sdfx**<br> 548⭐ 13웃 52 ⑂ 2017-01)
  sdf(**sdf**<br> 93⭐ 3웃 9 ⑂ 2022-04)
  classDef classArchived  fill:#DDD;
  class sdf classArchived
  gsdf(**gsdf**<br> 18⭐ 1웃 1 ⑂ 2022-04)
  T(Today)
  sdfx -- 2022 --> sdf
  sdf -- 2024 --> gsdf
  sdfx --> T
  gsdf --> T
{{< /mermaid >}}

## sdfx
- {{< github_button button="view"  user="deadsy" repo="sdfx" >}}
{{< github_button button="star"  user="deadsy" repo="sdfx" count="true">}}

This is the repo at [github.com/deadsy/sdfx](https://github.com/deadsy/sdfx).
Created in 2017-01-09T00:23:49Z and updated yesterday.  When I started I compared with sdf and thought
sdf looked cleaner.

The repo has learned from sdf and implement its own vectors.

## sdf

- {{< github_button button="view"  user="soypat" repo="sdf" >}}
{{< github_button button="star"  user="soypat" repo="sdf" count="true">}}

The repo is https://github.com/soypat/sdf and created in 2022-04-25T02:40:59Z.  I used this as it was meant to be an improvement, eg here is 
project for creating bowls https://github.com/drummonds/go3dp/tree/master/cmd/bowlstacker.  However it has been
archived on
2024-08-12 as soypat is working on his successor project.

The main thing about sdf is that it is focused on go, is faster by using better primitives.

## gsdf

- {{< github_button button="view"  user="soypat" repo="gsdf" >}}
{{< github_button button="star"  user="soypat" repo="gsdf" count="true">}}

This is soypats replacement for sdf at the repo [github.com/soypat/gsdf](https://github.com/soypat/gsdf).
The main focus of this is rendering with a GPU and an integrated viewer.
I am not sure this is a direction that I am very keen on.  I did use it for my Brabantia pin replacement
which is working well.

There are two libraries out there which I have come across for doing CAD with signed distance functions [sdfx][] and [sdf][].  I originally went with sdf as it was under current development or found it first. 

The main thing about sdf is that it is focused on go, is faster by using better primitives.  Support form 3mf is emerging which will allow multi colour materials.

## Others

There are other use cases like https://github.com/daveagill/go-sdf which is about morphing 2 d images from one to another.  This is not really applicable to 3d printing.

[ghostscad]:https://github.com/ljanyst/ghostscad
[sdfx]:https://github.com/deadsy/sdfx
[sdf]:https://github.com/soypat/sdf

## Next steps

To build my hex plug covers for floorboard screws and compare the different viewers the old go one
 https://github.com/Yeicor/sdfx-ui? or the new rust version https://github.com/Yeicor/sdf-viewer.

 

## Research

Some of this I got from this great list of gui packages [go-gui-projects][] but not in curated [awesome-go][] packages:

- {{< github_button button="view"  user="go-graphics" repo="go-gui-projects" >}}
  {{< github_button button="star"  user="go-graphics" repo="go-gui-projects" count="true">}}
- {{< github_button button="view"  user="avelino" repo="awesome-go" >}}
{{< github_button button="star"  user="avelino" repo="awesome-go" count="true">}}

[go-gui-projects]: https://github.com/go-graphics/go-gui-projects
[awesome-go]:https://github.com/avelino/awesome-go