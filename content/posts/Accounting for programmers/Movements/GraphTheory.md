---
title: Accounting transactions as graph theory
date: 2025-01-05
mermaid: true
---

A transaction represents a transfer of value, eg investing in a business, the transaction consists of an amount, a source and a destination.  The transaction is a transfer from the source to the destination.

Each node is an account and has a balance.  Each node must represent the same type of stuff in it,
eg money or oranges etc.   The value of a balance at a node is
determined by simple rules:

1. At the beginning of time, the balance is zero.
2. For incoming edge add the value of the edge
3. For each outgoing edge subtract the value of the edge

## Accounts balance

From {{% ref Kleppmann2011%}}:

The account balances have two nice properties:

1. The sum of all account balances is always zero
2. Partition the set of accounts into into any two disjoint sets, and add up all of the balances in each set, then the sum for the one set is always the negative sum of the other set (because, after all, they have to add up to zero).

The sum is zero as every movement has a negative amount in one account and a positive in the other account.


[Parent](/afp/movements/) [Prev](/afp/movements/conventions/) [Next](](/afp/movements/graphtheory/)/afp/movements/oranges/)