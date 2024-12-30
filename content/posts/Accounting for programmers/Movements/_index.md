---
title: Accounting movements
date: 2024-12-26
mermaid: true
----

## Movements


Accounting is really the tracking of resources.  [Kleppmann][Kleppmann2011]
 [IA](https://web.archive.org/web/20110717155537/https://martin.kleppmann.com/2011/03/07/accounting-for-computer-scientists.html) really crystallised this concept for me by explaing them as graph theory.  A transaction is about the movement of resources from one place to another.

[Kleppmann2011]: https://martin.kleppmann.com/2011/03/07/accounting-for-computer-scientists.html

If a founder puts money into a business you can represent the transaction as a movement of money:

{{< mermaid >}}
flowchart LR
   I([Investor]) -- £5 --> B([Business])
{{< /mermaid >}}


The founder is hoping that the business will do well and then the following will occur:

{{< mermaid >}}
flowchart LR
   I -- £5 --> B
   B([Business]) -- £100 --> I([Investor])
{{< /mermaid >}}

One of the principles of accounting is that is done from the point of view of single entity.  Here we have two, the investor and the business.  If the investor is self employed and investing in their own business then the boundaries between the two get blurred.  In the UK the tax authority [HMRC](https://www.gov.uk/money/income-tax) will normally require as self employed person to complete a tax return on behalf of you and your business.

From the point of view of transactions there are two separate entities:

{{< mermaid >}}
flowchart LR
   subgraph Investor
    I([Investor])
   end
   subgraph Business
    B([Business])
   end
   I -- £5 --> B
{{< /mermaid >}}

## Business as an entity

In the previous paragraph we were looking from the outside at a transaction between two entities.
One of the features of double entry book keeping.

## Accounts balancing

From [Kleppmann][Kleppmann2011]:

The account balances have two nice properties:

1. The sum of all account balances is always zero
2. Partition the set of accounts into into any two disjoint sets, and add up all of the balances in each set, then the sum for the one set is always the negative sum of the other set (because, after all, they have to add up to zero).

The sum is zero as every movement has a negative amount in one account and a positive in the other account.

## Simplest business transaction

This is a very simple business model.  You invest in your business.  You then buy oranges from an orange farmer, take them to market and sell them at a profit.  You then keep the profit.

So reusing the diagram above you the reader as entrepreneur invests in your own business:

{{< mermaid >}}
flowchart LR
   subgraph Entrepreneur
    E([Entrepreneur])
   end
   subgraph Own Business
    B([Business])
   end
   E -- £25 --> B
{{< /mermaid >}}

If we focus on the accounts for the business as a single entity and treat the entrepreneur as an investor you get the following:

{{< mermaid >}}
flowchart LR
   E([Entrepreneur])
   C([Cash])
   E -- £25 --> C
{{< /mermaid >}}

From now on we are only going to treat this as the accounting records of the business. This is 
one of the main tennets of double entry accounting.  Following
[Kleppmann][Kleppmann2011] each node represents an account and each edge (vector) represents a transaction.  

A transaction represents a transfer of value, eg investing in a business, the transaction consists of an amount, a source and a destination.  The transaction is a transfer from the source to the destination.

Each node is an account and has a balance.  Each node must represent the same type of stuff in it,
eg money or oranges etc.   The value of a balance at a node is
determined by simple rules:

1. At the beginning of time, the balance is zero.
2. For incoming edge add the value of the edge
3. For each outgoing edge subtract the value of the edge


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

### Converting stock purchase as money

The conversion of amounts of things to their monetary value is a key element of accounting records.  This makes it much possible to create accounts
So we are going to represent the stock as money and at the cost price.  This is not present
in the original stock taking accounting of the Sumerians where each type of item eg barley or malt 
was [shown on the transaction](/afp/uruk/).

So now we convert the original stock purchase to its monetary value:


{{< mermaid >}}
flowchart TD
   subgraph As Entities
   direction LR
      C([Cash])
      F([Farmer £])
      F1([Farmer 🍊])
      S([Stock 🍊])
   end
   subgraph As Money
   direction LR
      CM([Cash £])
      FM([Farmer £])
      SM([Stock £])
   end
   C -- £25 --> F
   F1 -- 25 Oranges --> S
   CM -- £25 --> FM
   FM -- £25 --> SM
{{< /mermaid >}}

And looking at the balances for those transactions:

{{< mermaid >}}
flowchart LR
   CM([Cash **£-25**])
   FM([Farmer **£0**])
   SM([Stock **£25**])
   CM -- £25 --> FM
   FM -- £25 --> SM
{{< /mermaid >}}

You get the satisfying transfer of £25 to inventory from cash.  The balance with the farmer is
zero as the goods have been paid for.

### Converting sale of oranges to money equivalent

{{< mermaid >}}
flowchart LR
   subgraph As Entities
   direction LR
      M([Market £]) -- £37.50 --> C([Cash])
      S([Stock 🍊]) -- 25 Oranges --> M1([Market 🍊])
   end
{{< /mermaid >}}

{{< mermaid >}}
flowchart LR
   subgraph As Money
   direction LR
      M([Market £]) -- £37.50 --> C([Cash])
      S([Stock £]) -- £25 --> CG([Cost of goods sold])
   end
{{< /mermaid >}}



### Overall accounts in money terms

So now looking at the overall picture using money values you have:

{{< mermaid >}}
flowchart LR
   E([Entrepreneur **-£25**])
   C([Cash **£37.50**])
   F([Farmer **£0**])
   M([Market **-£12.50**])
   S([Stock **£0**])
   CG([Cost of goods sold **£25**])
   E -- £25 --> C
   C -- £25 --> F
   M -- £37.50 --> C

   F -- £25 --> S
   S -- £25 --> CG
{{< /mermaid >}}

From this you can see that no money is owed to the Farmer nor is there any stock.  However the
overall profit and health of the business is unclear.  At a point in time (eg after trading) the 
value of a business is the value of the assets minus its liabilities.  The value incudes the investments by the entrepreneur.

We should be able to see:

   <!--  -->

{{< mermaid >}}
flowchart LR
   classDef classEquity  fill:#EEB;
   classDef classAssets  fill:#CDA;
   classDef classIncome  fill:#BCE;
   classDef classExpense  fill:#FCC;
   classDef classLiabilities  fill:#DDD;
   E@{ shape: stadium, label: "Entrepreneur **-£25**"}
   class E classEquity;
   C([Cash **£37.50**])
   class C classAssets;
   F([Farmer **£0**])
   class F classLiabilities;
   M([Market **-£37.50**])
   class M classIncome;
   S([Stock **£0**])
   class S classAssets;
   CG([Cost of goods sold **£25**])
   class CG classExpense;
   E -- £25 --> C
   C -- £25 --> F
   M -- £37.50 --> C

   F -- £25 --> S
   S -- £25 --> CG
{{< /mermaid >}}


### Connection to double entry accounting

[Luca Pacioli](https://en.wikipedia.org/wiki/Luca_Pacioli) published of the first accounting book
and the english translation by Jeremy Cripps starts with the following:

```
Three things are needed by anyone who wishes to carry on business carefully:
1) Cash or Credit
2) A good accountant
3) proper internal control
```

In the example above the entrepreneur has used their own cash to fund the business, this blog is 
providing the internal control and after this blog the reader will be a good accountant.





<!-- ### Representation of stock as its purchase price

In order to represent the accounts in a single currency we will convert our oranges to their purchase price:

{{< mermaid >}}
flowchart LR
   C([Cash]) -- £25 --> F([Farmer])
   F([Farmer]) -- £25 for 25 Oranges --> S([Stock])
{{< /mermaid >}}

Now the Farmer has a zero balance and we have £25 in stock.
 -->


