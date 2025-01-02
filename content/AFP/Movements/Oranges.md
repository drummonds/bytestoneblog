---
title: Accounting movements - trading oranges
date: 2025-01-01
mermaid: true
----

[Prev]()

## Buying stock / inventory

I have used the term stock as it is a British English term 
[Wikipedia has inventory](https://en.wikipedia.org/wiki/Inventory) as the American synonym.

Having got investment in the business lets buy 25 oranges from a farmer for pound for each orange.
This has two transactions to represent the transfer of oranges and cash.  Here we are showing both
the movement of the cash and the oranges

{{< mermaid >}}
flowchart LR
   C([Cash]) -- £25 --> F([Farmer £])
   F1([Farmer 🍊]) -- 25 Oranges --> S([Stock 🍊])
{{< /mermaid >}}

You can make things more complicated by realising that the stock is perishable and so gets
less valuable the older it is.  

### Sale of oranges

We now take the oranges to market and sell them at £1.50 each.  We can represent this as:

{{< mermaid >}}
flowchart LR
   M([Market £]) -- £37.50 --> C([Cash])
   S([Stock 🍊]) -- 25 Oranges --> M1([Market 🍊])
{{< /mermaid >}}


### Overall accounts

I have now put in the end balances for each account:

{{< mermaid >}}
flowchart LR
   E([Entrepreneur **-£25**])
   C([Cash **£37.50**])
   F([Farmer **£25**])
   F1([Farmer **-25🍊**])
   M([Market **-£37.50**])
   M1([Market **-25🍊**])
   S([Stock **0 🍊**])
   E -- £25 --> C
   C -- £25 --> F
   M -- £37.50 --> C

   F1 -- 25 Oranges --> S
   S -- 25 Oranges --> M1
{{< /mermaid >}}

This has led to a problem in that it is complicated to calculate sums as we have two types
of things as nodes (oranges and cash).  This is particularly apparent after the buying stage
where the business has no cash and 25 oranges, so the value of the business has gone from £25 
to 25🍊 Oranges.

