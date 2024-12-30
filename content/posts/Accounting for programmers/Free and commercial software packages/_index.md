---
title: Free and commercial accounting software packages
date: 2024-12-18
---

This is an eclectic list and is just those which I have been familiar with.

## Gnucash

I have used this both directly and via [Piecash](https://github.com/sdementen/piecash).  Piecash is a very nice SQLAlchemy based ORM which hides many of the idiosyncrasies of GnuCash.  From a database point of view the problem is that the database definition has been frozen for a while and all the changes have been done by adding key value pairs in a generic KVP data.

## Sage line 50

This has an ODBC connection although it is 32bits.  This allows you to connect with it 
programmatically.
and I wrote I python interface for this, [PySage](https://github.com/drummonds/pySage50) but I haven't worked on it for ages.

# Plain text accounting implementations

These are a number of the [plain text accounting](/afp/plain-text-accounting/) implementations.  

## [Beancount](https://beancount.github.io/)

I have used this when I converted from [Xero to Beancount](posts/2019-05-29-conversion-from-xero-to-beancount/).  Martin Blais writes clearly
and in depth.  He also has a cleaner description and version of a plain text accounting grammar.
He also considers in depth FX conversions and stock pricing.

There is also an SQL implementation to query the data to produce reports.

There is a nice gui version of this called [Fava](https://github.com/beancount/fava) which I have used effectively.

## [Coin](https://github.com/mkobetic/coin)

I like this as it is fast, written in go and simplified.  It also uses big int's as a rational 
basis for storing amounts.  It actually abstracts the amount so that different methods can be used.
It also has good test coverage.

One of the snags with it is the syntax follows ledger rather than beancount. 
- Balance assertions come on a single posting and become order dependent.  This is 
particle bad as the order for bank reconciliation may be different from that of entries in the file.
