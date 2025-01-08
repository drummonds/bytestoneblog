---
title: Trading oranges as money - new conventions
date: 2025-01-07
mermaid: true
type: afp_single
---

The first time through I tried using the conventional accounting conventions but the
overall result was unsatisfactory so that here I have reverted to something that makes 
more sense and show how the accounting conventions are built from it.

- [Using money and accounting conventions](/afp/movements/orangesasmoney_ac/)
- [Pacioli was wrong and why it doesn't matter](/afp/movements/pacioliogotitwrong/)

So we rework the previous example and treat the oranges as a money equivalent stock item.
This allows complete accounts to be done.  In order to make clear the sign difference between a node
in graph terminology and an account in accounting terminology I will refer to nodes in a graph sense (with positive and negative balances) and 
accounts in a conventional accounting sense (credits and debits).

### Converting stock purchase as money

The conversion of amounts of things to their monetary value is a key element of accounting records.  This makes it much possible to create accounts.
So we are going to represent the stock as money and at the cost price.  This is not present
in the original stock taking accounting of the Sumerians where each type of item eg barley or malt 
was [shown on the transaction](/afp/uruk/).

## Trading

  So to recap the scenario is:

   - invest in your business.  
   - buy oranges from an orange farmer
   - sell them at a profit at market

### Investment

We start from nothing.  The first step is an injection of cash from an external investor.

{{< mermaid >}}
flowchart LR
   {{% account_class %}}
   I([Investor])
   class I classInvestor;
   subgraph Business
   direction LR
      E([Equity Input **£-25**])
      class E classEquity;
      B([Cash Assets **£25**])
      class B classAssets;
   end
   I -- £25 --> Business
   E --  £25 --> B
{{< /mermaid >}}


### Buying stock / inventory

For each stage there is only the entries of that stage although the balances of any nodes
are also shown.
I have used the term stock as it is a British English term 
[Wikipedia has inventory](https://en.wikipedia.org/wiki/Inventory) as the American synonym
see [this comparison]().

Having got investment in the business lets buy 25 oranges from a farmer for one pound for each
orange.
The exchange of cash for oranges is denoted by a single accounting entry.

{{< mermaid >}}
flowchart RL
   {{% account_class %}}
   F([Farmer</br>-25🍊 **£25**])
   subgraph Business
   direction LR
      E([Equity Input **£-25**])
      class E classEquity;
      B([Cash Assets **£0**])
      class B classAssets;
      P([Purchases **£25**])
      class P classExpense;
   end
   Business -- £25 --> F 
   B -- £25 --> P
   F -- 25 🍊 --> Business
{{< /mermaid >}}


### Converting sale of oranges to money equivalent

{{< mermaid >}}
flowchart RL
   {{% account_class %}}
   M([Market</br>25🍊 **£-37.50**])
   subgraph Business
   direction LR
      E([Equity Input **£-25**])
      class E classEquity;
      B([Cash Assets **£37.50**])
      class B classAssets;
      P([Purchases **£25**])
      class P classExpense;
      S([Sales **£-37.50**])
      class S classIncome;
   end
   S -- £37.50 --> B
   M -- £37.50 --> Business
   Business -- 25 🍊 --> M
{{< /mermaid >}}

So when you make a sale you transfer money from the Sales revenue node to the cash
node.

### Overall nodes in money terms

So now looking at the overall picture using money values you have:

{{< mermaid >}}
flowchart RL
   {{% account_class %}}
   I([Investor])
   class I classInvestor;
   F([Farmer</br>-25🍊 **£25**])
   M([Market</br>25🍊 **£-37.50**])
   subgraph Business
   direction LR
      E([Equity Input **£-25**])
      class E classEquity;
      B([Cash Assets **£37.50**])
      class B classAssets;
      P([Purchases **£25**])
      class P classExpense;
      S([Sales **£-37.50**])
      class S classIncome;
   end
   I -- £25 --> Business
   E --  £25 --> B
   Business -- £25 --> F 
   B -- £25 --> P
   F -- 25 🍊 --> Business
   S -- £37.50 --> B
   M -- £37.50 --> Business
   Business -- 25 🍊 --> M
{{< /mermaid >}}

From this you can see that no money is owed to the Farmer nor is there any stock.  There
is also a nice flow of capital from left to right.

To make this double entry and to have meaning the inflows must be kept separate from the
outflows.  This allows the tracking of profitability.  In this case it is easy but with
different stock items you could separate out profitable and unprofitable lines.


From now on we will leave the external boundary and look only what happens in the firm.

{{< mermaid >}}
flowchart LR
   {{% account_class %}}
   E([Equity Input **£-25**])
   class E classEquity;
   B([Cash Assets **£37.50**])
   class B classAssets;
   S([Sales **£-37.50**])
   class S classIncome;
   P([Purchases **£25**])
   class P classExpense;
   E --  £25 --> B
   S -- £37.50 --> B
   B -- £25 --> P
{{< /mermaid >}}

## Adding an example liability

If we modify the story a bit we can add a liability.  If instead of giving cash to the farmer
at the time of the purchase of the oranges you promise to pay on the way back from the market.
Then when you get back you will owe the farmer for the oranges (and might not need an investor).


{{< mermaid >}}
flowchart LR
   {{% account_class %}}
   E([Equity Input **£-25**])
   class E classEquity;
   B([Cash Assets **£62.50**])
   class B classAssets;
   S([Sales **£-37.50**])
   class S classIncome;
   P([Purchases **£25**])
   class P classExpense;
   F([Farmer **£-25**])
   class F classLiabilities;
   E --  £25 --> B
   S -- £37.50 --> B
   F -- £25 --> P
{{< /mermaid >}}

Looking at the business we have inputs and outputs in external nodes (in bold edges) and internal nodes such
as assets and liabilities (in thin edges).
The total value in the business is the sum of those internal nodes assets + liabilities, as the liabilities are
negative.

## Balance sheet inerpretation

In fact the definition of an asset is a core node with a positive balance and a liability as a core node with
a negative liability.  Typically but not always nodes are either assets or liabilities.  For instance a pile of
cash will always be positive.  However if you store your money in a bank it can seemlessly switch from positive to
negative and back again. So summarising the value in the entity is called the total net assets (net as we have combined the assets and liabilities).

$$\definecolor{asset}{RGB}{204,221,170}
\definecolor{liability}{RGB}{221,221,221}
totalNetAssets=\colorbox{asset}{coreNodes}=\colorbox{asset}{assetNodes}+\colorbox{liability}{liabilityNodes}=£37.50$$

We have promised to pay the farmer and have the cash to do so.  The value of the totalNetAssets will not change when
we pay the farmer.
Because this is a graph all the external nodes equals the negative of all the internal nodes.



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

[Parent](/afp/movements/) [Prev](/afp/movements/oranges/) 