---
title: Accounting movements
date: 2025-01-01
mermaid: true
---

## Accounting Movements is just graph theory

Accounting is the tracking of resources in order to answer question.. This article 
was inspired by a post explaining accounting for computer scientists by
{{% ref Kleppmann2011%}}. The key insight is to view accounting transactions as
a [graph](<https://en.wikipedia.org/wiki/Graph_(discrete_mathematics)>) which is a set of nodes
(accounts) joined by edges (transactions).

A transaction is simply the movement of resources from one place to another (one account to another) at a particular time.
For example if an investor puts money into a business you can represent the transaction as a movement of money:

{{< mermaid >}}
flowchart LR
I([Investor]) -- £5 --> B([Business])
{{< /mermaid >}}

The investor is hoping that the business will do well and then the following will occur:

{{< mermaid >}}
flowchart LR
I -- £5 --> B
B -. Trading .-> B
B([Business]) -- £100 --> I([Investor])
{{< /mermaid >}}

Accounting records are within the business and have to match the external movements:

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

See [Boundaries of entities](/afp/movements/boundaries/) for more details

[Accounting conventions](/afp/movements/conventions/) can go either way and can get confusing.
We will stick to long standing (500+ years) accounting conventions.

This is an index to colour coding each of the accounts/nodes and giving them a are more expressive
title:

| Colour                                                                         | Meaning                                       |
| ------------------------------------------------------------------------------ | --------------------------------------------- |
| $$\definecolor{equity}{RGB}{238,238,187}\colorbox{equity}{Equity}$$            | Money invested by the owners or their profits |
| $$\definecolor{asset}{RGB}{204,221,170}\colorbox{asset}{Asset}$$               | Things that the entity has or is owing to it  |
| $$\definecolor{liability}{RGB}{221,221,221}\colorbox{liability}{Liabilities}$$ | Things that the entity has or is owing to it  |
| $$\definecolor{income}{RGB}{187,204,238}\colorbox{income}{Income}$$            | Income earned by the business                 |
| $$\definecolor{expense}{RGB}{255,204,204}\colorbox{expense}{Expense}$$         | Expenses incurred by the business             |

The [link to graph theory](/afp/movements/graphtheory/) is that each node is an account and each edge a transaction.

## Detailed simplest business transaction

This is a very simple business model. You invest in your business. You then buy oranges
from an orange farmer, take them to market and sell them at a profit.   You then keep the profit. I quite like
orange trading as my wifes family in eighteenth century {{% ref Farrell2020 %}}.

This is shown with accounts denominated in oranges [Orange trading example](/afp/movements/oranges/) and then [reworked using money](/afp/movements/orangesasmoney/)  to give the result below:

### Input and outputs

We can model what the investor hopes will happening by adding trading and the revenue out:

{{< mermaid >}}
flowchart LR
classDef classEquity fill:#EEB,stroke-width:4px;
classDef classAssets fill:#CDA;
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


[Parent](/afp/) [Prev](/afp/uruk/) [Next](/afp/double-entry-bookkeeping/)