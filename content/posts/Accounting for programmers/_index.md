---
title: "Accounting for Programmers"
update: "2025-01-15"
date: "2016-10-30"
categories: ["Humphrey"]
tags: ["Humphrey"]
author: "Humphrey Drummond"
summary: "An introduction to accounting for programmers."
mermaid: true
---
# Accounting for Programmers


This is an extended series of posts to introduce the concepts of accounting to programmers.  It is born of my own exploration into the subject and I have started writing in 2016.  It take a somewhat different approach to most books in that we will be looking at accounting
from the point of view of a programmer, me!

There two goals, in the first part I am exploring the nature of double entry book keeping with
the directed graph approach and how accounting emerges from this.  There is a lot of explanation and 
exploration as to the nature of accounting records and why the credit/debit system is isomorphic 
with positive and negative balances.  

If you want to just read one thing then try the [explanation of an orange trading business](/afp/movements/orangesasmoney).

{{< mermaid_fig title="Demo" caption="Business with a liability"  >}}
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


The use of computers allows a freedom to present information in different ways to those bound by
clay or paper.  For instance exact conventions become unimportant as you can display data in any
way you like.


It is currently a work in progress so see the overall [plan](/afp/_plan/).

Current Structure
- [Accounting Records](/afp/accounting-records/)
- [Historical](/afp/uruk/), current state
- [Movements](/afp/movements/)
- Single Entry and [Double entry accounting](/afp/double-entry-bookkeeping/)
- [Plain Text accounting](/afp/plain-text-accounting/)
- [Free and commercial software packages](/afp/free-and-commercial-software-packages/)
- [Accounting tricks](/afp/accounting-tricks/)
- [British and American terminology](/afp/british-american-terminology/)
- [References](/afp/references/)


