---
title: Golang CAD versions
date: 2023-02-23
---
I am keen to try out writing and designing in go CAD models.  I am on a drive to simplify and unify my coding styles.  My process for designing  I would much rather be creating my models by programming. So I started with OpenScad. I discovered that as the models got very complicated it would start to have problems particularly with the difference function.

There was a great go library called [ghostscad][] that generated OpenSCAD language but as it transpiles to OpenSCAD it suffers from the same problems as complexity increases.


## Signed distance functions

There are two libraries out there sdfx

{{< github_button button="view"  user="deadsy" repo="sdfx" >}}
{{< github_button button="stars"  user="deadsy" repo="sdfx" >}}

[ghostscad]:https://github.com/ljanyst/ghostscad
[sdfx]:https://github.com/deadsy/sdfx
