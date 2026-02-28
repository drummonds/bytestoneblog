---
title: Trading oranges as money - new conventions
date: 2025-01-07
updated: 2025-01-13
mermaid: true
type: afp_single
show_in_timeline: true
---

The first time through I tried using the conventional accounting conventions but the
overall result was unsatisfactory so that here I have reverted to something that makes
more sense and show how the accounting conventions are built from it.

- [Using money and accounting conventions](/afp/movements/orangesasmoney_ac/)
- [Pacioli was wrong and why it doesn't matter](/afp/movements/pacioliogotitwrong/)

So we rework the previous example and treat the oranges as a money equivalent stock item.
This allows complete accounts to be done. In order to make clear the sign difference between a node
in graph terminology and an account in accounting terminology I will refer to nodes in a graph sense (with positive and negative balances) and
accounts in a conventional accounting sense (credits and debits).

### Converting stock purchase as money

The conversion of amounts of things to their monetary value is a key element of accounting records. This makes it possible to create accounts.
So we are going to represent the stock as money and at the cost price. This is not present
in the original stock taking accounting of the Sumerians where each type of item eg barley or malt
was [shown on the transaction](//afp/historical-accounting//).

## Trading

So to recap the scenario is:

- invest in your business.
- buy oranges from an orange farmer
- sell them at a profit at market

### Investment

We start from nothing. The first step is an injection of cash from an external investor. We are going to separate
out boundary modes/accounts from those inside the business.

{{< mermaid_fig title="Figure 1" caption="Investment"  >}}
flowchart LR
{{% account_class %}}
I([Investor])
class I classInvestor;
subgraph Business
direction LR
E([Equity Input **£-25**])
class E classEquity;
subgraph Internal
style Internal fill:#EEE
direction LR
B([Cash Assets **£25**])
class B classAssets;
end
end
I -- £25 --> Business
E -- £25 --> B
{{< /mermaid_fig >}}

### Buying stock / inventory

For each stage there is only the entries of that stage although the balances of any nodes
are also shown.
I have used the term stock as it is a British English term
[Wikipedia has inventory](https://en.wikipedia.org/wiki/Inventory) as the American synonym
see [this comparison]().

Having got investment in the business lets buy 25 oranges from a farmer for one pound for each
orange.
The exchange of cash for oranges is denoted by a single accounting entry.

{{< mermaid_fig title="Figure 2" caption="Buying stock"  >}}
flowchart RL
{{% account_class %}}
F([Farmer</br>-25🍊 **£25**])
subgraph Business
direction LR
E([Equity Input **£-25**])
class E classEquity;
subgraph Internal
style Internal fill:#EEE
direction LR
B([Cash Assets **£0**])
class B classAssets;
end
P([Purchases **£25**])
class P classExpense;
end
Business -- £25 --> F
B -- £25 --> P
F -- 25 🍊 --> Business
{{< /mermaid_fig >}}

### Selling the stock at market

So now we go to market and sell the oranges:

{{< mermaid_fig title="Figure 3" caption="Selling the stock"  >}}
flowchart RL
{{% account_class %}}
M([Market</br>25🍊 **£-37.50**])
subgraph Business
direction LR
E([Equity Input **£-25**])
class E classEquity;
subgraph Internal
style Internal fill:#EEE
direction LR
B([Cash Assets **£37.50**])
class B classAssets;
end
P([Purchases **£25**])
class P classExpense;
S([Sales **£-37.50**])
class S classIncome;
end
S -- £37.50 --> B
M -- £37.50 --> Business
Business -- 25 🍊 --> M
{{< /mermaid_fig >}}

So when you make a sale you transfer money from the Sales revenue node to the cash
node.

### Overall nodes

So now looking at the overall picture using money values you have:

{{< mermaid_fig title="Figure 4" caption="Summary"  >}}
flowchart RL
{{% account_class %}}
F([Farmer</br>-25🍊 **£25**])
I([Investor])
class I classInvestor;
M([Market</br>25🍊 **£-37.50**])
subgraph Business
direction LR
E([Equity Input **£-25**])
class E classEquity;
subgraph Internal
style Internal fill:#EEE
direction LR
B([Cash Assets **£37.50**])
class B classAssets;
end
P([Purchases **£25**])
class P classExpense;
S([Sales **£-37.50**])
class S classIncome;
end
I -- £25 --> Business
E -- £25 --> B
Business -- £25 --> F
B -- £25 --> P
F -- 25 🍊 --> Business
S -- £37.50 --> B
M -- £37.50 --> Business
Business -- 25 🍊 --> M
{{< /mermaid_fig>}}

From this you can see that no money is owed to the Farmer nor is there any stock. There
is also a nice flow of capital from left to right. The total value of the business is the sum
of all the nodes (1) inside the internal box. Since this is a directed graph this is necessarily the
opposite of the sum of the nodes on the outside. On the outside you have the negative equity that was
put in originally and the sum of the trading (sales and purchases).

To make this double entry and to have meaning the inflows must be kept separate from the
outflows. This allows the tracking of profitability. In this case it is easy but with
different stock items you could separate out profitable and unprofitable lines.

From now on we will leave the external boundary and look only what happens in the firm. We will keep the
internal box as this nicely delineates the boundary nodes and the internal ones.

{{< mermaid_fig title="Figure 5" caption="Summary for the business"  >}}
flowchart LR
{{% account_class %}}
E([Equity Input **£-25**])
class E classEquity;
subgraph Internal
style Internal fill:#EEE
direction LR
B([Cash Assets **£37.50**])
class B classAssets;
end
S([Sales **£-37.50**])
class S classIncome;
P([Purchases **£25**])
class P classExpense;
E -- £25 --> B
S -- £37.50 --> B
B -- £25 --> P
{{< /mermaid_fig>}}

## Adding an example liability

If we modify the story a bit we can add a liability. If instead of giving cash to the farmer
at the time of the purchase of the oranges you promise to pay on the way back from the market.
Then when you get back you will owe the farmer for the oranges (and might not need an investor).

{{< mermaid_fig title="Figure 6" caption="Unfinished business with a liability"  >}}
flowchart LR
{{% account_class %}}
E([Equity Input **£-25**])
class E classEquity;
subgraph Internal
style Internal fill:#EEE
direction LR
B([Cash Assets **£62.50**])
class B classAssets;
F([Farmer **£-25**])
class F classLiabilities;
end
S([Sales **£-37.50**])
class S classIncome;
P([Purchases **£25**])
class P classExpense;
E -- £25 --> B
S -- £37.50 --> B
F -- £25 --> P
{{< /mermaid_fig>}}

Looking at the business we have inputs and outputs in external nodes (in bold edges) and internal nodes such
as assets and liabilities (in thin edges).
The total value in the business is the sum of those internal nodes assets + liabilities, as the liabilities are
negative. It is important to note that both [figures 5](#figure-5) and [figure 6](#figure-6) have the same overall
value. This is because in [figure 6](#figure-6) we have not yet paid the farmer.

## Balance sheet interpretation

In fact the definition of an asset is an internal node with a positive balance and a liability as an internal node with
a negative liability. Typically but not always nodes are either assets or liabilities. For instance a pile of
cash will always be positive. However if you store your money in a bank it can seamlessly switch from positive to
negative and back again. So summarising the value in the entity is called the total net assets (net as we have combined the assets and liabilities).

$$
\definecolor{asset}{RGB}{204,221,170}
\definecolor{liability}{RGB}{221,221,221}
totalNetAssets=\colorbox{asset}{internalNodes}=\colorbox{asset}{assetNodes}+\colorbox{liability}{liabilityNodes}=£37.50
$$

We have promised to pay the farmer and have the cash to do so. The value of the totalNetAssets will not change when
we pay the farmer.
Because this is a graph all the external nodes equals the negative of all the internal nodes.

Looking at the position of the entrepreneur, they have invested £25 in the business and so the
rest of the balances will add up to £25. This says nothing about the health of the business
(Balance sheet) or its trading position (Profit and Loss).

### Full balance sheet

We know that all the nodes must add up to zero:

$$
\definecolor{asset}{RGB}{204,221,170}
\definecolor{liability}{RGB}{221,221,221}
\definecolor{equity}{RGB}{238,238,187}
\definecolor{income}{RGB}{187,204,238}
\definecolor{expense}{RGB}{255,204,204}
0=\colorbox{asset}{assetNodes}+\colorbox{liability}{liabilityNodes}
+\colorbox{equity}{equityNodes}+\colorbox{income}{incomeNodes}+\colorbox{expense}{expenseNodes}
$$

So the full equation becomes:

$$
\definecolor{asset}{RGB}{204,221,170}
\definecolor{liability}{RGB}{221,221,221}
\definecolor{equity}{RGB}{238,238,187}
\definecolor{income}{RGB}{187,204,238}
\definecolor{expense}{RGB}{255,204,204}
totalNetAssets=\colorbox{asset}{assetNodes}+\colorbox{liability}{liabilityNodes}
=-(\colorbox{equity}{equityNodes}+\colorbox{income}{incomeNodes}+\colorbox{expense}{expenseNodes})
$$

with liabilities, equity and income being typically negative.

### To an accountants balance sheet

As we have [touched on accountants](/afp/movements/conventions/) avoid negative numbers so we do the following
transformations in a report stage to create accountant friendly reports. This is much the same ways as formatting
cells in Excel:

| Node Account | Usual value | Accountant term | transform |
| ------------ | ----------- | --------------- | --------- |
| Asset        | +ve         | Debit Normal    | +1        |
| Liability    | -ve         | Credit Normal   | -1        |
| Equity       | -ve         | Credit Normal   | -1        |
| Income       | -ve         | Credit Normal   | -1        |
| Expense      | +ve         | Debit Normal    | +1        |

Accountants like to start with the [accounting equation](https://en.wikipedia.org/wiki/Accounting_equation) rather
than the outcome of treating transactions as movements within a single entity.  I am going to present it for a
point in time as:

$$
\definecolor{asset}{RGB}{204,221,170}
\definecolor{liability}{RGB}{221,221,221}
\definecolor{equity}{RGB}{238,238,187}
\colorbox{equity}{Equity}=\colorbox{asset}{Assets}-\colorbox{liability}{Liabilities}
$$

I like this arrangement as it says from an equity holders point of view how much is a business worth.
So using the table above to translate the raw balances to this we get:

$$
\definecolor{asset}{RGB}{204,221,170}
\definecolor{liability}{RGB}{221,221,221}
\definecolor{equity}{RGB}{238,238,187}
\definecolor{income}{RGB}{187,204,238}
\definecolor{expense}{RGB}{255,204,204}
(-\colorbox{equity}{equityNodes})=\colorbox{asset}{assetNodes} - (-\colorbox{liability}{liabilityNodes})
$$

and repeating for the full equations you get:

$$
\definecolor{asset}{RGB}{204,221,170}
\definecolor{liability}{RGB}{221,221,221}
\definecolor{equity}{RGB}{238,238,187}
\definecolor{income}{RGB}{187,204,238}
\definecolor{expense}{RGB}{255,204,204}
\colorbox{equity}{Equity}=\colorbox{asset}{Assets}-\colorbox{liability}{Liabilities}
+\colorbox{income}{Income}
-\colorbox{expense}{Expenses}
$$


and using node balances:

$$
\definecolor{asset}{RGB}{204,221,170}
\definecolor{liability}{RGB}{221,221,221}
\definecolor{equity}{RGB}{238,238,187}
\definecolor{income}{RGB}{187,204,238}
\definecolor{expense}{RGB}{255,204,204}
(-\colorbox{equity}{equityNodes})=\colorbox{asset}{assetNodes} - (-\colorbox{liability}{liabilityNodes})
+\colorbox{income}{incomeNodes}
-(-\colorbox{expense}{expenseNodes})
$$

### Note on income and expenses

These are both boundary nodes.  You could combine them in a single input and output node.  However you
would then lose the ability to track the total volume of sales.  So a low profit business with a large turnover
would look the same as a high profit business with a small turnover.

Inside a business if you have multiple sales lines, a single trading account loses the ability to track the
profitability of say of two lines, eg in this case oranges and lemons.

[Parent](/afp/movements/) [Prev](/afp/movements/oranges/)
