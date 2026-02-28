---
title: Accounting movements - trading oranges
date: 2025-01-01
mermaid: true
type: afp_single
show_in_timeline: true
---


Using the convention developed in [Accounting Conventions](/afp/movements/conventions/) and
looking solely at the accounts with a start from nothing. 

## The trading

   - invest in your business.  
   - buy oranges from an orange farmer
   - sell them at a profit at market

### Investment

So reusing the diagram above you the reader as entrepreneur invests in your own business:

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
   I -- £25 --> Business
   B --  £25 --> E
{{< /mermaid >}}


### Buying stock / inventory

I have used the term stock as it is a British English term 
[Wikipedia has inventory](https://en.wikipedia.org/wiki/Inventory) as the American synonym.

Having got investment in the business lets buy 25 oranges from a farmer for pound for each orange.
This has two transactions to represent the transfer of oranges and cash.  Here we are showing both
the movement of the cash and the oranges

{{< mermaid >}}
flowchart RL
   {{% account_class %}}
   F([Farmer</br>-25🍊 **£25**])
   subgraph Business
   direction LR
      E([Equity Input **£25**])
      class E classEquity;
      B([Cash Assets **£0**])
      Ex([Expense **£-25**])
      S([Stock 25🍊])
      class B,S classAssets;
      class Ex classExpense;
   end
   Business -- £25 --> F 
   Ex -- £25 --> B
   F -- 25 Oranges --> Business
{{< /mermaid >}}

You can make things more complicated by realising that the stock is perishable and so gets
less valuable the older it is but we are going to ignore that here.

### Sale of oranges

We now take the oranges to market and sell them at £1.50 each.  We can represent this as:

{{< mermaid >}}
flowchart RL
   {{% account_class %}}
   M([Market</br>25🍊 **£-37.50**])
   subgraph Business
   direction LR
      E([Equity Input **£25**])
      class E classEquity;
      B([Cash Assets **£-37.50**])
      Ex([Expense **£-25**])
      I([Income **£37.50**])
      S([Stock 0🍊])
      class B,S classAssets;
      class Ex classExpense;
      class I classIncome;
   end
   M -- £37.50 --> Business
   Business -- 25 Oranges --> M
   B -- £37.50 --> I
{{< /mermaid >}}

### Overall accounts

So now you can have all the transactions together:

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
      Ex([Expense **£-25**])
      I([Income **£37.50**])
      S([Stock 0🍊])
      class B,S classAssets;
      class Ex classExpense;
      class I classIncome;
   end
   I -- £25 --> Business
   B --  £25 --> E
   Business -- £25 --> F 
   Ex -- £25 --> B
   F -- 25 Oranges --> Business
   M -- £37.50 --> Business
   Business -- 25 Oranges --> M
   B -- £37.50 --> I
{{< /mermaid >}}

This has led to a problem in that it is complicated to calculate sums as we have two types
of things as accounts/nodes (oranges and cash).  This is particularly apparent after the buying
stage
where the business has no cash and 25 oranges, so the value of the business has gone from £25 
to 25🍊 Oranges.

Additionally we have no information left in the account balances about the amount of
oranges being bought and sold.  The flow of money in and out can be seen but the oranges
are separate.  This is how you handle single entry accounting (for the oranges)
rather than double (for the money).

[Parent](/afp/movements/) [Prev](/afp/movements/graphtheory/) [Next](/afp/movements/orangesasmoney_ac/)
