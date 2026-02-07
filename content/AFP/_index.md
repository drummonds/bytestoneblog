---
title: "Accounting for Programmers"
update: "2025-09-21"
date: "2016-10-30"
categories: ["Humphrey"]
tags: ["Humphrey"]
author: "Humphrey Drummond"
summary: "An introduction to accounting for programmers."
mermaid: true
---

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
way you like.  A good example of this is an early Ada computer that stored
programs as abstract syntax trees and then presented to the programmer
in their preferred format.  This is an alternative approach to formatting all
the code in the same the Black does with Python.



It is currently a work in progress so see the overall plan further down this page.

## Current Structure
- [Movements](/afp/movements/)
- [Accounting Records](/afp/accounting-records/)
- Single Entry and [Double entry accounting](/afp/double-entry-bookkeeping/)
- [Plain Text accounting](/afp/plain-text-accounting/)
- [Free and commercial software packages](/afp/free-and-commercial-software-packages/)
- [Accounting tricks](/afp/accounting-tricks/)
- [British and American terminology](/afp/british-american-terminology/)
- [Slide presentations](/slides/)
- [Historical Accounting](/afp/historical-accounting/)
- [References](/afp/references/)

### Current plan of work

 - Sort current structiure
 - Generate PDF's and epub from it
 - Complete historical accounting
 - Complete why Pacioli was wrong and why it does not matter
 -  Complete Balance sheet presentation
 - Do profit and loss calculations and presentation
- Do profit and lost calculation

### Future subjects

  - End of period accounts
  - Expand slides to include notes on banks
    -  Safety deposit
    -  Ponzi
    -  Central bank
  - Hierachical accounts
    - input and output
  -VAT
   - Government tax collection
  - Rework Accounting conventions
  - Double entry requires separate input and output accounts
  - Account types - correspondence of Personal real and nominal

## Mindmap

{{< mermaid >}}
mindmap
  afp((accounting for programmers))
    Accounting Records
      historical context
        Summerian Uruk
        Chinese?
        Indian
    Single Entry and Double entry accounting
      The PiggyBank
    Entity
    Counting
    M))Movements((
      Balance to Zero
      Credit and debits
      Negative numbers
    Balance sheet
    Profit and loss
    Banking
    Time series data
    Traditional Account topics
      VAT
        Government tax collection
      Division of labour
      Petty cash accounting
      The trial balance (and the accounting equation)
      Reporting for Taxes
      Reporting for management info
      Stock
      Invoicing
    Plain Text accounting
    Free and commercial software packages
{{< /mermaid >}}
