---
title: "Plain text accounting"
date: "2024-04-04"
---

This will be the introduction to cover a range of subjects:

- EBNF
- Time formats
- Beancount flavour of plain text accounting
- Text version of t side accounting
- History of accounting
- direct graphs for transactions
- Negative numbers or debits and credits
- Time series data

- Ledger modelling
  - Ledgers
  - Balances
    - Open and close balances
  - Transactions
  - Posting Legs
  - split files
  - time series data

- Transaction lots

## Plain Text Accounting formats

What is common is that this is a list of directives in one or more text files.  There are a number of directirves.   The order of the directives is not that important.

There are a number of definitions:

- [Beancount](https://beancount.github.io/docs/beancount_language_syntax.html)
- [Hledger]()

### Directive types

A directive is one or more lines

Type |   | Beancount | Hledger
---|---|---|
Comment | ; | Pure comment | ;,# may contain notes


- 

### Ordering of directives

Hledger recommonds that directives should be ordered in date order.  A unique ordering is important if you would like to have a consistent hash that is the same if the order of the directives is changed or
the commnents.  