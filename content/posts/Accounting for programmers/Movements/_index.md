---
title: Accounting movements
date: 2025-01-01
mermaid: true
----

## Movements


Accounting is really the tracking of resources.  This article was inspired by a post explaining accounting for computer scientists by [Martin Kleppmann][Kleppmann2011]
 [📚](https://web.archive.org/web/20110717155537/https://martin.kleppmann.com/2011/03/07/accounting-for-computer-scientists.html).
  This crystallised the concept for me by explaining accounting transactions as graph
 theory.  A transaction is simply the movement of resources from one place to another (one account to another) at a particular time.

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
   B -. Trading .-> B
   B([Business]) -- £100 --> I([Investor])
{{< /mermaid >}}

## Entity and the boundary of the firm

One of the principles of accounting is that is done from the point of view of single entity.  Here
we have two, the investor and the business.  If the investor is self employed and investing in
 their own business then the boundaries between the two get blurred.  In the UK the tax authority
  [HMRC](https://www.gov.uk/money/income-tax) will normally require as self employed person to 
  complete a tax return on behalf of you and your business.

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

## Business as a single separate entity

In the previous paragraph we were looking from the outside at a transaction between two entities.
The great Venetian addition was to consider the business entity as separate from the owners and
this is a core feature of double entry book keeping practice.

So in the diagram above we are only considering transactions within the business.

{{< mermaid >}}
flowchart LR
   I([Investor])
   classDef classAssets  fill:#CDA;
   subgraph Business
    B([Business])
    class B classAssets;
  end
   I -- £5 --> B
{{< /mermaid >}}

This presents a problem in that we have movements of resources from outside the business to inside
the business.  This is dealt with by keeping matching input and output accounts which account for
all the flows in and out of the business.

{{< mermaid >}}
flowchart LR
   classDef classEquity  fill:#EEB;
   classDef classAssets  fill:#CDA;
   I([Investor])
   class I classEquity;
   subgraph Business
   direction LR
      E([Equity Input])
      class E classEquity;
      B([Assets])
      class B classAssets;
   end
   I -- £5 --> Business
   E -- £5 --> B
{{< /mermaid >}}


### Input and outputs

We can model what the investor hopes will happening by adding trading and the revenue out:

{{< mermaid >}}
flowchart LR
   classDef classEquity  fill:#EEB;
   classDef classAssets  fill:#CDA;
   I([Investor])
   I1([Investor])
   class I,I1 classEquity;
   subgraph Business
   direction LR
      E([Equity Input])
      D([Dividend output])
      class E,D classEquity;
      B([Assets])
      class B classAssets;
      R([Revenue])
      Ex([Expense])
   end
   I -- £5 --> Business
   E -- £5 --> B
   R -. Sales .-> B
   B -- £1000 --> D
   B -. Expense .-> Ex
   Business -- £1000 --> I1
{{< /mermaid >}}


## Accounts balance

From [Kleppmann][Kleppmann2011]:

The account balances have two nice properties:

1. The sum of all account balances is always zero
2. Partition the set of accounts into into any two disjoint sets, and add up all of the balances in each set, then the sum for the one set is always the negative sum of the other set (because, after all, they have to add up to zero).

The sum is zero as every movement has a negative amount in one account and a positive in the other account.

## Detailed simplest business transaction

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

This change of viewpoint from looking at multiple entities to the interactions of a single
entity is conceptually a big change.  There is a parallel with the first law of thermodynamics 
where energy is preserved in a closed system.  Arguably it is this change in view point that
the Northern Italian merchants made in there large trading companies that gave rise to double
entry book keeping.  In the initial diagrams each entity is regarded by itself as a single entry
book keeping system.

In contrast to the system for thermodynamics the system overall is not closed.  The entrepreneur node denotes input to the system.  You can also have exit nodes as we will see later.

From now on we are only going to look at the accounting records of a single business. This is 
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

You get the satisfying transfer of £25 to stock from cash.  The balance with the farmer is
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
   M([Market **-£37.50**])
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
value of a business is the value of the assets minus its liabilities.  The value incudes the
investments by the entrepreneur.

So by colouring in the various boxes and giving them a are more expressive title:

Colour  | Meaning
---|----
$$\definecolor{equity}{RGB}{238,238,187}\colorbox{equity}{Equity}$$ | Money invested by the owners or their profits
$$\definecolor{asset}{RGB}{204,221,170}\colorbox{asset}{Asset}$$ | Things that the entity has or is owing to it
$$\definecolor{liability}{RGB}{221,221,221}\colorbox{liability}{Liabilities}$$ | Things that the entity has or is owing to it
$$\definecolor{income}{RGB}{187,204,238}\colorbox{income}{Income}$$ | Income earned by the business
$$\definecolor{expense}{RGB}{255,204,204}\colorbox{expense}{Expense}$$ | Expenses incurred by the business

If we now label the various accounts above we should be able to see:

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

The farmer is either an asset if you pay before receiving the goods, or a liability if you pay
after receiving the goods.  So in this example the balance is zero as we paid at the time of 
receiving the goods and have marked as an example of a liability.

In order to show a liability we have shown the position where you have agreed to pay the farmer on
the way back from the market. 

{{< mermaid >}}
flowchart LR
   classDef classEquity  fill:#EEB;
   classDef classAssets  fill:#CDA;
   classDef classIncome  fill:#BCE;
   classDef classExpense  fill:#FCC;
   classDef classLiabilities  fill:#DDD;
   E@{ shape: stadium, label: "Entrepreneur **-£25**"}
   class E classEquity;
   C([Cash **£62.50**])
   class C classAssets;
   F([Farmer **£-25**])
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

Looking at the position of the entrepreneur, they have invested £25 in the business and so the
rest of the balances will add up to £25.  This says nothing about the health of the business
(Balance sheet) or its trading position (Profit and Loss).


The balance sheet is the sum of all the assets minus the liabilities. (Green minus the grey)  The 
value of the balance sheet is the total net assets:

$$\definecolor{asset}{RGB}{204,221,170}
\definecolor{liability}{RGB}{221,221,221}
totalNetAssets=\colorbox{asset}{Assets}+\colorbox{liability}{Liabilities}$$

We add the liabilities to the assets but the liabilities are expressed as negative numbers.

So the full equation becomes:

$$\definecolor{asset}{RGB}{204,221,170}
\definecolor{liability}{RGB}{221,221,221}
\definecolor{equity}{RGB}{238,238,187}
\definecolor{income}{RGB}{187,204,238}
\definecolor{expense}{RGB}{255,204,204}
totalNetAssets=\colorbox{asset}{Assets}+\colorbox{liability}{Liabilities}
=\colorbox{equity}{Equity}+\colorbox{income}{Income}+\colorbox{expense}{Expense}$$

with liabilities, equity and income being typically negative.

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


