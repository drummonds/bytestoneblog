---
title: "Plain text accounting"
date: "2024-11-22"
---

## Time series functions

A time series is a function that varies over time. For example the amount of money in a piggy bank.
This has a value at a point in time as well as changes. This corresponds to a derivative although typically the changes are stochastic not a regular time intervals. There is also an integral version which is the change during a period eg how much was added during 1 week or some other period. Here the period can be regular.

Accounting is about the transfer of money or other quantity from one place to another.

This has been the case for a [very long time](https://posit-dev.github.io/great-tables/blog/design-philosophy/):

When agriculture became more widespread (ca. 10,000 years ago), there was the need to document and manage economic transactions to do with farming, livestock, and the division of labour. In the fourth millennium BC, Mesopotamian cities that traded with far way kingdoms needed to keep such records. Clay tablets recovered from the ancient Sumerian city of Uruk show early yet sophisticated tables. Here is a drawing of one of the recovered tablets, which contains an accounting of deliveries of barley and malt from two individuals for the production of beer.

![Uruk table with annotations](uruk_tablet_with_annotations.png)

Drawing of clay tablet from Sumerian city of Uruk, circa 3200-3000 BC. Uruk III Tablet (MSVO 3, 51, Louvre Museum, Paris, France). Annotated with the meanings of the columns, rows, and cells.

Note that the recovered tablet is meant to be read from right to left. Inside each box is an ideogram (a symbol that represented a word or idea) and a numerical value representing a quantity.

Its structure is where things get super interesting:

    Rows: there are roughly two rows, each corresponding to an individual.
    Columns: the first two columns from the right contain counts of malt (rightmost column) and barley (second rightmost column).
    Subtotals: the third column from the right sums barley and malt within each individual, and the left-most column displays the grand total.

As a bonus, the table has a footer, since the bottom row contains the name of the official in charge.

## Plain text accounting

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

What is common is that this is a list of directives in one or more text files. There are a number of directirves. The order of the directives is not that important.

There are a number of definitions:

- [Beancount](https://beancount.github.io/docs/beancount_language_syntax.html)
- [Hledger]()
- Ledger the original
- My own

### Chart of accounts

An account has the following syntax

```ebnf
account = {type} name {subaccount}
subaccount = ":" name {subaccount}
```

### Directive types

A directive is one or more lines:

directive = date type

| Type      | Beancount              | Hledger                 |
| --------- | ---------------------- | ----------------------- |
| Posting   | Pure txn,\*,!          | Info rich               |
| Comment   | ; Pure comment         | ;,# may contain notes   |
| unknown   | Ignores                |                         |
| open      | account must be opened |
| close     |
| commodity |
| balance   | Balance assertion      | folded into posting leg |
| pad       |
| note      |
| document  |
| price     |
| event     |
| query     |
| custom    |

There is also metadata and options

#### Posting

A posting is a set of posting legs that balance and happen at a single point in time.

In beancount a posting is called a transaction.

A posting consists of a posting instruction description followed by posting legs.

The posting directive is a multi line directive which must be kept together. Some reordering eg of the posting legs is possible without information being lost.

##### Posting leg

This is indented to denote that it belongs to the preceeding posting instruction.

#### Comment

In beancount all text specified as comment is ignored. In hledger further parsing may reveal
more information

-

#### Open

Beancount requires you to have an open account directive before an account is used. However you can
infer the accounts

Account life is OPEN -> CLOSED However this is not the only state transition possible eg


#### Close

#### Balance

An account is s series of postings and at any one time you have a balance on that account.

A complete list of balances has the identical information to a complete list of postings. There
is cross over between a database value (balance) as against the transaction log (postings) which
contain similar information.

However balances have a couple of useful properties:

- At a point in time all you need is a balance to make decisions (do I have enough money for a fluffy toy)
- If you divide a series of balances in two you can report on the separate halves without loss of information
  whereas for a list of postings you need to be able to sum previous postings.

Kafka also has the same characteristics in that a stream of changes is equivalent to the current state.

#### Balance assertions

Coin the goledger has added in balance transactions to a posting leg:

```pta
2024/01/02 AMZNMktplace ON 29 DEC BCC
  Unbalanced               34.20 GBP
  Assets:Bank:H3Personal  -34.20 GBP = 224.69 GBP
```

This is then used in abbreviated form for instance in Ledger:

```
2020-01-01 Reconciliation
    Assets:Cash at Bank  $0 = $100
```

And in beancount:

```pta
2012-02-04 balance Assets:CA:Bank:Checking    417.61 CAD
```

The Beancount syntax is superior as the balance assertions are clear and distinct from 
the postings and order of postings.  For balance reconciliation it is also superior.  When doing the accounting
for a customer the bank record is one of the independent bits of evidence so the reconciliation is
important.

### Ordering of directives

Hledger recommends that directives should be ordered in date order. A unique ordering is important if you would like to have a consistent hash that is the same if the order of the directives is changed or
the commnents.

# My plain text accounting format

- I want to be able to used date times not only dates
- I want to be able to create ledger files (books) by different methods and to be
  able to compare them.
- I want to segment time series so that old data can be "closed" and connected in a Merkle tree
  to help check it hasn't been changed.
- For there to be only one way to represent things so that you get the same result
- To use parameter time series as a general concept for all data
- To be able to use unchanged two segments

# Golang ledger implementations

This is an abstract and an enhancement of the [PTA apps](https://plaintextaccounting.org/#pta-apps)
table but only focused on golang implementations.

As of 2024-11-21

| Name                 | Start | Last Release | Contributors | Stars | Forks | Comment                 |
| -------------------- | ----- | ------------ | ------------ | ----- | ----- | ----------------------- |
| [Ledger (Go)][go1]   | 2017  | 2024         | 12           | 462   | 43      | Single currency |
| [knut][go2]          | 2020  | 2024         | 3            | 58    | 10    | Multicurrency, accruals |
| [tn47 goledger][go3] | 2017  | 2018         | 1            | 35    | 13    | Float64?
| [Coin][go4]          | 2019  | 2024         | 2            | 10    | 0     | big Int |

[go1]: https://github.com/howeyc/ledger
[go2]: https://github.com/sboehler/knut
[go3]: https://github.com/tn47/goledger
[go4]: https://github.com/mkobetic/coin

## [Ledger (Go)][go1]

Has a [manual](https://howeyc.github.io/ledger/Introduction.html). Much faster than the
other ledger implementations (10 to 100 times) on its own benchmarks.

### Representation

Uses int64 for numbers with a fixed 0.001 resolution.

A book is listed as a collection of transactions.

## [knut][go2]

## [goledger][go3]


## [coin][go4]

### Representation

Uses bigint plus variable number of decimal places
