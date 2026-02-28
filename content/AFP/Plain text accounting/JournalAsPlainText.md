---
title: using Journal entries as plain text
date: 2025-01-10
show_in_timeline: true
---

# Journal Entries

Pacioli as translated by {{% ref Geisjbeck-1914 %}} has a very nice division of the primary records into:

- Opening Inventory
- Day book
- **Journal**
- Ledger

The journal is a neat list of entries which are cross referenced and entered in the ledger.  Now a computer
is great at animating that data so you only need to enter the journal data.  Each journal entry is a single
transaction in date order.


## Two legs better

I am taken with the idea of the directed graph to inform the way you create plain text accounting.
Most plain text formats follow [ledger from John Wiegley] {{% ref Wiegly-2003 %}} in that they are journals
which support transactions with multiple posting legs:

``` PTA
; The next day, I buy lemonade supplies
2020-06-02 * "Safeway, Inc." "Supplies"
  Assets:BofA:Checking                        CR 52.48 USD
  Expenses:OperatingExpenses:CostOfGoodsSold  DR 52.23 USD
  Expenses:Taxes:SalesTax                     DR  0.25 USD
```

From {{%ref Cooter-2020 %}}, but I haven't yet found a reference for GAAP rules:

```
Here's the GAAP rules for accounts:
Assets: Debit-normal
Liabilities: Credit-normal
Capital/Equity: Credit-normal
Income: Credit-normal
Expenses: Debit-normal
```

#  Accounts

{{% ref Pacioli-1494 %}} has quite long journal entries eg using {{% ref Giesbeck1914 %}} translation for opening 
cash is:

```
Per cash // A — Capital of myself so and so, etc. In cash I have at present, in gold
and coin, silver and copper of different coinage as it appears in the first sheet of the 
Inventory in cash, etc., in total so many gold ducats and so many silver ducats. All this is
our Venetian money ; that is counting 24 grossi per ducat and 32 picioli per grosso in gold
is worth : L (Lire), S (Soldi), G (Grossi), P (Picioli).
```

but actually most of it is commentary.  The accounts are created in the Ledger as needed so are
defined from the journal.

Beancount has account directives which are optional {{% ref Beancount-Accounts %}}

But as a consequence the notation in the text is repetitive when making entries eg:

```
2002-01-17 open Assets:US:BofA:Checking

2002-01-17 * "Opening Balance"
  Assets:US:BofA:Checking        987.34 USD
  Equity:Opening-Balances       -987.34 USD
```

whereas you could describe it as: 

```
2002-01-17 open Assets:US:BofA:Checking
account Assets:US:BofA:Checking
  alias BofA
  commodity USD

2002-01-17 * "Opening Balance"
 Opening-Balances -- 987.34 --> BofA
```

The commodity is described by the account definition and they have to match