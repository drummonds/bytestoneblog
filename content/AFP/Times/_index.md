---
title: Times and dates
draft: true
date: 2024-11-28
---

In most plain text accounting systems a transaction is represented as:

``` text
DATE [DESCRIPTION]
   ACCOUNT    AMOUNT
   ACCOUNT    [AMOUNT]
```

The date components is just the date and not a datetime.  This can be a standardised format like YYYY-MM-DD in [beancount](https://beancount.github.io/docs/beancount_language_syntax.html#directives) or more flexible [plain text account.org](https://plaintextaccounting.org/quickref/#h.xpr0dgy4pyj3)
for various formats.  All of these represent a single day.  The ordering of the transactions with the day may or may not be important.

In [beancount balance assertions](https://beancount.github.io/docs/beancount_language_syntax.html#balance-assertions) are at the beginning of the day and it is stated that the ordering of directives within the day is not important.  Martin Blais also makes the point that balance assertions
are only relevant on Assets and Liabilities.

In [coin](https://github.com/mkobetic) you can have balance assertions attached to a transaction which makes them dependent on the order of transactions.


If you want to put your directives into a Merkle Tree then the ordering becomes important.  Also if you
run a real time accounting system then the date time of the transaction becomes important.  For
instance if you run a bank and only do date based transactions and need to reconcile with an
external payments system 3 times a day then the payment is important.

## Arrow of time

A merkle tree locking of a series of transactions ensures that the arrow of time goes from the past to the future.
