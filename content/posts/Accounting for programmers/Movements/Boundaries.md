---
title: Boundaries
date: 2025-01-05
mermaid: true
---

## Entity and the boundary of the entity

One of the principles of accounting is that is done from the point of view of single entity.  Here
we have two, the investor and the business.  If the investor is self employed and investing in
 their own business then the boundaries between the two get blurred.  In the UK the tax authority
  [HMRC](https://www.gov.uk/money/income-tax) will normally require a self employed person to 
  complete a tax return on behalf of you and your business, ie as a single entity.

From the point of view of transactions there are two separate entities with a transaction between them (an investment into the business):

{{< mermaid >}}
flowchart LR
   {{% account_class %}}
   I([Investor])
   class I classInvestor;
   subgraph Business
    B([Accounting Entries])
    class B classAssets;
  end
   I -- £5 --> B
{{< /mermaid >}}

## Business as a single separate entity

The great Venetian addition was to consider the business entity as separate from the owners and
to record the transactions that happen within the business entity.  This is at the heart of
double entry book keeping practice.

The diagram above presents a problem in that we have movements of resources from outside the business to inside
the business.  This is dealt with by keeping matching input and output accounts which account for
all the flows in and out of the business.

{{< mermaid >}}
flowchart LR
   {{% account_class %}}
   I([Investor])
   class I classEquity;
   subgraph Business
   direction LR
      E([Equity Input])
      class E classEquity;
      B([Assets Cash])
      class B classAssets;
   end
   I -- £5 --> Business
   B -- £5 --> E
{{< /mermaid >}}

The details of the accounts used will come in the [next section](/afp/movements/conventions/).  The important thing is that
Equity Input node/account corresponds to the Investor account.  This follows Venetian practice in the the most important is the owners equity.  It also shows that we have
accounts that represent state within the business and not outside it, eg the assets.

This change of viewpoint from looking at multiple entities to the interactions of a single
entity is conceptually a big change.  There is a parallel with the first law of thermodynamics 
where energy is preserved in a closed system.  Arguably it is this change in view point that
the Northern Italian merchants made in there large trading companies that gave rise to double
entry book keeping.  In the initial diagrams each entity is regarded by itself as a single entry
book keeping system.

In contrast to the system for thermodynamics the system overall is not closed.  The entrepreneur node denotes input to the system.  You can also have exit nodes as we will see later.

From now on we are only going to look at the accounting records of a single business. This is 
one of the main tennets of double entry accounting.  Following
{{% ref Kleppmann2011%}} each node represents an account and each edge (vector) represents a
transaction.  


[Parent](/afp/movements/) [Next](/afp/movements/conventions/)