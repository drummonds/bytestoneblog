---
title: "Claude, Singularity Postgres and Gokrazy"
date: "2026-01-07"
---

My son wrote a great piece about [Lectures in London](https://taydrummond.com/lectures_in_london) which I really like.
To create the lecture scanning page idea he enlisted ChatGPT.
As he was writing ChatGPT went from 5.2 to 5.3
and whereas 5.2 was having difficulties executing the task the latest version 5.3 was very successful.

He had been chatting to one of his friends at Anthropic who was saying they and openAI were using their coding models for about 70% of the code.  It feels like the singularity although it is happening relatively slowly as we go through it eg months and years rather than
seconds that for instance Nick Bostrom talked about.

For my own code I tried working with Gokrazy and trying to find a nice way to use databases.  go-sqlite was great but using
a different dialect of SQL leads to trouble and complexity.  I thought about CockroachDB as it is written in GO but it is
quite a large project without an easy build.  So I decided to create a translation layer to convert postgres to sqlite and
then use the pure go sqlite project.  For simple cases this worked well and in a few hours I had a working library 
[github.com/drummonds/go-postgres](https://github.com/drummonds/go-postgres) and in no time at all a working demo
todo list that uses it [github.com/drummonds/pglike_todo](https://github.com/drummonds/pglike_todo) and this installed
happily in to gokrazy.  A solid win!

One of the things that has come out of this is the idea that I might concentrate on in home projects being on simple databases and
using gokrazy on Raspberry pi 5.  When I want to grow them I will deploy on cloud architecture rather than deploy baby versions of
cloud at home.
I started this particular journey when Paperless-NGX crashed due
to a postgres database update and I think that I am getting much
closer to actually getting it done with go-postgres making it simpler
and Claude helping write the code.

