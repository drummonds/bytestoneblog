---
title: Accounting sign conventions
date:  2025-01-05
mermaid: true
type: afp_single
---

There are at least two of ways of modelling the account movements.  The notation of
{{% ref Kleppmann2011%}} chooses the opposite notation although slight more obvious is the
opposite of accounting notation as
published by Paciolo. 

This is the overall position:

{{< mermaid >}}
flowchart LR
   {{% account_class %}}
   class I classInvestor;
   I([Investor])
   subgraph Business
    B([...Accounts...])
    class B classAssets;
  end
   I -- £5 --> Business
{{< /mermaid >}}

and if you had a single entry accounting system you might represent this as a positive balance.
For a double entry accounting system you are representing not just the value but a movement from
one account to another.

There are now choices on how to represent this transaction inside the business when you have two accounts/nodes:

{{< mermaid >}}
flowchart LR
   {{% account_class %}}
   I([Investor])
   class I classInvestor;
   subgraph Business
   direction LR
      E([Equity Input])
      class E classEquity;
      B([Cash Assets])
      class B classAssets;
   end
   I -- £5 --> Business
   E x-- ?? --x B
{{< /mermaid >}}

So following Kleppman you can do the following with positive balances for assets and negative
balances for equity:

{{< mermaid >}}
flowchart LR
   {{% account_class %}}
   I([Investor])
   class I classInvestor;
   subgraph Business
   direction LR
      E([Equity Input **£-5**])
      class E classEquity;
      B([Cash Assets **£5**])
      class B classAssets;
   end
   I -- £5 --> Business
   E -- £5 --> B
{{< /mermaid >}}

or following conventional accounting and reversing direction:

{{< mermaid >}}
flowchart LR
   {{% account_class %}}
   I([Investor])
   class I classInvestor;
   subgraph Business
   direction LR
      E([Equity Input **£5**])
      class E classEquity;
      B([Cash Assets **£-5**])
      class B classAssets;
   end
   I -- £5 --> Business
   B --  £5 --> E
{{< /mermaid >}}

This seems slightly odd in that the cash is negative when you add cash, but the equity account 
representing the investor is positive and does not change as the cash is used in the entity, eg
purchasing stock.  It may make more sense in that the investor is not giving cash but rather
becoming a creditor by providing liquidity.  So the focus is on the creditor relationship
rather than on the actual money.
  Using conventional accounting terminology the same arrangement can be looked at like:

{{< mermaid >}}
flowchart LR
   {{% account_class %}}
   I([Investor])
   class I classInvestor;
   subgraph Business
   direction LR
      E([Equity Input<br/> Dr ‡ Cr<br/>-- ‡ **£5**])
      class E classEquity;
      B([Cash Assets<br/> Dr ‡ Cr<br/>**£5** ‡ -- ])
      class B classAssets;
   end
   I -- £5 --> Business
   B -- Dr £5 Cr --> E
{{< /mermaid >}}

The nice thing about this is that the Equity account is a creditor which is the relationship
of the investor to the entity.

Paciolo's example journal entry for
accounting for inventory cash is:

`November 8, MCCCCLXXXXIII in Venice.
Per cash // A - Capital 24 gold ducats...`

where per is translated as italian for By and A Italian for to.  This this argues for an
arrow going from cash to Capital / Equity.  Also there is a nice relationship between the lefthand
side of the journal and the left hand side of the ledger.

The ledger entries the left hand side is `in dare` "shall give" and for the right hand side
`in havare` "shall receive".  Benedetto Cotrugli was writing a little earlier than Paciolo in 1458
in the `Art of Trade` and explains ledgers with 
`the right-hand side of the book under “sums owed” and on the left “sums owing".`  
He keeps the cash as a debtor column as money owed.  This seems the same as current practice.

I am going to quote directly from [geisbeck14](/afp/references/#Geisjbeck1914) **Chapter 11**

``` quote
CHAPTER 11.
THE TWO EXPRESSIONS USED IN THE JOURNAL, ESPECIALLY IN VENICE, THE ONE
CALLED "PER," AND THE OTHER "A," AND WHAT IS UNDERSTOOD BY THEM.
As we have said, there are two expressions {termini) used in the said Journal; the one is called "per,"
and the other is called "a," each of which has a meaning of its own. "Per" indicates the debtor {debitore)
one or more as the case may be, and "a," creditor {creditore) , one or more as the case may be. Never is
any item entered in the Journal which also is to be entered in the Ledger, without preceding it by one of
the two expressions. At the beginning of each entry, we always provide "per," because, first, the debtor
must be given, and immediately after the creditor, the one separated from the other by two little slanting
parallels {virgolette) , thus, //, as the example below will show.
```

This is a tricky area as [Cripps modern translation](/afp/references/#Cripps1994) gets it the wrong way round see the foot notes on page 20:

```
1. "Per," that is "To"
2. "A," that is "By"
```


Paciolo has the left column marked per "in dare" and the right "in havare".  Later the left hand
side became a debit column mark dr and the right hand side a credit column marked cr.

[Parent](/afp/movements/) [Prev](/afp/movements/boundaries/) [Next](/afp/movements/graphtheory/)