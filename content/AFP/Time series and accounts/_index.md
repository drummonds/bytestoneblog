---
title: "Time series and accounts"
date: "2024-12-28"
---

## Time series functions

A time series is a function that varies over time. For example the amount of money in a piggy bank.
This has a value at a point in time as well as changes. This corresponds to a derivative although typically the changes are stochastic not a regular time intervals. There is also an integral version which is the change during a period eg how much was added during 1 week or some other period. Here the period can be regular.

The difference in the underlying data become apparent when you change the period of the time
 series.
So for instance if you have annual balance sheets and you want to estimate bi-annual versions you 
could just take the last value.  However for profit and losses you add them over the same period.

This is not a property of the list of transactions but of the timeseries (eg annual 
reports) of that data.