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

This follows simply from an external investor.

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

So when you make a sale you transfer money from the Sales revenue account to the cash
account.

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