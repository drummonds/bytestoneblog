---
title: ABNF and Plain Text Accounting
date: 2026-02-26
---

Plain text accounting is a useful way to describe a machine and human readable accounting journal.  Some formats
are single file format and some are multiple file formats.  F

The software primarily beancount and ledger transform the data into useful reports as well as validating that the
data is correct in the first place.

The reports that can be run are:
- Summary balances at varying times and selections.  Hierachical accounts can be consolidated
- Ledger entries by account
- dealing with multiple commodities eg multiple currencies and stocks and shares
- preparing balance sheets
- preparing profit and loss
- SQL like query languages
- gui visualisation eg FAVA for beancount.
- integration of documents into the accounts.

Beancount documents that the clear process of converting text based records to an internal memory model is: 
- [Parsing and loading][] Read text file into memory structures
- [Realisation][] is the process of transforming a flat list of directives (parsed and processed entries) into a hierarchical tree of accounts with computed balances

[Parsing and loading]:https://deepwiki.com/beancount/beancount/3-parsing-and-loading
[Realisation]:https://deepwiki.com/beancount/beancount/6-realization-and-account-trees

I am hoping to compare the syntax and show how they relate to each other as well as to alternative notations such
as accounting movements.

# Data formats

This is a real time format I have used:

```ABNF
plain-text-accounting = *(transaction / posting / balance-check / data-point)

balance-check = datetime ["%" knowledge-date-time] "balance" account amount [comment] EOL

data-point = datetime ["%" knowledge-date-time] "data" name value [comment] EOL

posting = datetime ["%" knowledge-date-time] ("*" / "transaction") *2comment EOL *(metadata / posting-leg)

metadata = "  " parameter [comment] EOL

posting-leg = " " account amount [comment] EOL

posting-id-link = " " parameter EOL

comment = ";" comment-text EOL
```

![PTA Railroad Diagram](pta-railroad.svg)

This is a partial description of the BeanCount Journal format:

```ABNF
journal = *entry

entry = comment / transaction

comment = *WSP ";" *VCHAR CRLF

transaction = header postings

postings = 2*posting

directive = value-date ["^" knowledge-date] ("*" "transaction" / "open" open-directive / "balance" balance-directive)

blank-line = CRLF

header = value WSP id CRLF

value-date = full-date / date-time

knowledge-date = full-date / date-time

posting = account [asset-class] CRLF

balance-posting = account CRLF

id = string

string = DQUOTE *VCHAR DQUOTE

account = ("Assets" / "Liabilities" / "Equity" / "Income" / "Expenses") ":" name

name = TEXT / TEXT ":" name

full-date = date-fullyear "-" date-month "-" date-mday

full-time = partial-time time-offset

date-time = full-date "T" full-time
```

![BeanCount Journal Structure](bc-journal.svg)

![BeanCount Directives](bc-directives.svg)

![BeanCount Postings](bc-postings.svg)

![BeanCount Accounts and Dates](bc-accounts.svg)

## go-luca - Movements

go-luca replaces transactions and postings with movements, inspired by Pacioli's Credit // Debit notation.
A movement transfers an amount from one account to another using an arrow operator.
Linked movements (prefixed with `+`) group related transfers.

```ABNF
journal = *(movement / linked-movement)

movement = date WSP flag CRLF WSP WSP from-account WSP arrow WSP to-account [WSP description] WSP amount WSP commodity CRLF

linked-movement = movement 1*("+" from-account WSP arrow WSP to-account [WSP description] WSP amount [WSP commodity] CRLF)

date = 4DIGIT "-" 2DIGIT "-" 2DIGIT

flag = "*" / "!"

arrow = ">" / "->" / "//"

from-account = account

to-account = account

account = account-type *(":" name)

account-type = "Assets" / "Liabilities" / "Equity" / "Income" / "Expenses"

name = 1*VCHAR

description = DQUOTE *VCHAR DQUOTE

amount = 1*DIGIT ["." 1*DIGIT]

commodity = 1*ALPHA
```

![go-luca Movement Structure](luca-movements.svg)

![go-luca Accounts and Values](luca-accounts.svg)

## Multiple files - Corpus of accounts

It may not convenient or practical to have a single file represent the whole of the accounting journal.

Pacioli describes having multiple books, the first one marked with a * and then A etcera.  The ledgers openbalance is copied from the 
previous one. 


## Periods

I don't think either hledger or beancount consider periods.