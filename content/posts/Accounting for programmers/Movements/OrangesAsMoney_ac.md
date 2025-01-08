---
title: Trading oranges as money - accounting conventions
date: 2025-01-07
mermaid: true
type: afp_single
---


So we rework the previous example and treat the oranges as a money equivalent stock item.
This allows complete accounts to be done.

### Converting stock purchase as money

The conversion of amounts of things to their monetary value is a key element of accounting records.  This makes it much possible to create accounts
So we are going to represent the stock as money and at the cost price.  This is not present
in the original stock taking accounting of the Sumerians where each type of item eg barley or malt 
was [shown on the transaction](/afp/uruk/).

## Trading

  So to recap the scenario is:

   - invest in your business.  
   - buy oranges from an orange farmer
   - sell them at a profit at market

### Investment

This has no change as we only deal with money.

{{< mermaid >}}
flowchart LR
   {{% account_class %}}
   I([Investor])
   class I classInvestor;
   subgraph Business
   direction LR
      E([Equity Input **£25**])
      class E classEquity;
      B([Cash Assets **£-25**])
      class B classAssets;
   end
   I == £25 ==> Business
   B ==  £25 ==> E
{{< /mermaid >}}


### Buying stock / inventory

I have used the term stock as it is a British English term 
[Wikipedia has inventory](https://en.wikipedia.org/wiki/Inventory) as the American synonym.

Having got investment in the business lets buy 25 oranges from a farmer for one pound for each
orange.
The exchange of cash for oranges is denoted by a single accounting entry.

{{< mermaid >}}
flowchart RL
   {{% account_class %}}
   F([Farmer</br>-25🍊 **£25**])
   subgraph Business
   direction LR
      E([Equity Input **£25**])
      class E classEquity;
      B([Cash Assets **£0**])
      class B classAssets;
      P([Purchases **£-25**])
      class P classExpense;
   end
   Business == £25 ==> F 
   P == £25 ==> B
   F == 25 🍊 ==> Business
{{< /mermaid >}}


### Converting sale of oranges to money equivalent

{{< mermaid >}}
flowchart RL
   {{% account_class %}}
   M([Market</br>25🍊 **£-37.50**])
   subgraph Business
   direction LR
      E([Equity Input **£25**])
      class E classEquity;
      B([Cash Assets **£-37.50**])
      class B classAssets;
      P([Purchases **£-25**])
      class P classExpense;
      S([Sales **£37.50**])
      class S classIncome;
   end
   B == £37.50 ==> S
   M == £37.50 ==> Business
   Business == 25 🍊 ==> M
{{< /mermaid >}}

### Overall accounts in money terms

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
      E([Equity Input **£25**])
      class E classEquity;
      B([Cash Assets **£-37.50**])
      class B classAssets;
      P([Purchases **£-25**])
      class P classExpense;
      S([Sales **£37.50**])
      class S classIncome;
   end
   I == £25 ==> Business
   B ==  £25 ==> E
   Business == £25 ==> F 
   P == £25 ==> B
   F == 25 🍊 ==> Business
   B == £37.50 ==> S
   M == £37.50 ==> Business
   Business == 25 🍊 ==> M
{{< /mermaid >}}

From this you can see that no money is owed to the Farmer nor is there any stock.  However the
overall profit and health of the business is unclear.  At a point in time (eg after trading) the 
value of a business is the value of the assets minus its liabilities.  The value incudes the
investments by the entrepreneur.



[Parent](/afp/movements/) [Prev](/afp/movements/oranges/) 