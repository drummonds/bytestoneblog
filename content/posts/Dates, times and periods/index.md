---
title: Dates, Times and Periods
date: 2026-03-02
---

I want to write about the dates times and periods including how these are represented in text.

ISO8601 or RFC-3330 is a good place to start as UK goverment writes in their article on [dates and time-stamps standard][].

We normally differentiate between periods of time with a start and end date.  Eg a week goes from Mon to Sunday where in a precision of days Mon is the start
and Sunday is the end.
Often we are less precise and just use the whole period, eg this will be done in week 44.

Typcally the maximum period is a year, however sometimes eg for births of ancestors the exact date may not be known (or vary over someones lifetime) eg 1890+-10.

Accounting periods run from the begining of time A to the end of time B, eg income for tax purposes goes from the begining of April 6th to the end of April 5th in the following year.

Some times in plain text accounting you want to say something starts at the beginning where the exact time is unspecified.  In a bittemporal system when something starts can change at a later date when more information becomes clear.

[dates and time-stamps standard](https://www.gov.uk/government/publications/open-standards-for-government/date-times-and-time-stamps-standard)


## curiosities
In tiger beatle documenation user define timestamps must be globally unique `User-defined timestamps must be unique and expressed as nanoseconds since the UNIX epoch. No two objects can have the same timestamp, even different objects like an Account and a Transfer cannot share the same timestamp.`
https://docs.tigerbeetle.com/reference/transfer/
