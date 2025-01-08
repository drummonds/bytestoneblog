---
tile: Accounting Entries
date: 2025-01-05
---
## Accounting records

This is a [term HMRC](https://www.gov.uk/running-a-limited-company/company-and-accounting-records), the UK tax office require companies to keep.  This nicely sums up one of the reasons that accounting records are kept
which is to collect and ensure that taxes are paid.

When running a company you are concerned with the economic activity during a period on which you can pay a loss.
However the [Domesday book](https://en.wikipedia.org/wiki/Domesday_Book) is more a record of what William the 
Conquereor rules over at a point in time, 1086.  In order that taxes can be collected.

However if you are running a business or enterprise it is difficult to know if you are making a profit or loss
unless you keep records.  My father was fascinated by the netting of salmon off the East Coast of Scotland and was convinced that there was money in it.  After purchase of a boat, engines, net and licenses we hauled in a 
few salmon from the sea. It was clear that this was not a success but no records were kept and so the actual
money lost was not known.  Some businesses like shops can run on very small margins and accurate record keeping
is essential to remain profitable.

## The PiggyBank

When you are looking at an accounting system it usually applies to a specific entity.  This is the bounds of the system.  So often this is a limited or public company.  This can range from the smallest micro company to the largest quoted company eg Google or Exxon.

- Summerian
- 
- Restrictions
The piggy bank is the simplest model of an accounting system you have a single box into which you can put money in or take money out.  At any one point in time there is so much money in the piggy.
You could imagine a business selling oranges on the street and supposing you are able to get oranges on a sale or return basis from friendly green grocers.  At the beginning of the day you borrow a box of oranges and go out onto the street.  During the day you sell your oranges and put the money in the piggy bank. You sell your oranges and at the end of the day your return to your friendly green grocers.  You return the unsold oranges, pay the green grocer.  The money left over in your piggy bank at the end of the day is your profit.

There is also a limitation in that you only have the current value of the profit and there is not historical record.  If this was a virtual piggy bank you would have no record of what you had made in a week month or a year.  This leads onto time series data.

## Time Series Data

When you measure something at some point in time the measurement can be of a number of  different types.  A stock will have a price every day.  If you want to know the price for a time period you might either choose the end of the period, (closing price on Friday which is simple to measure) or the average price (a more robust measurement but more difficult to measure).  Over a longer period you can average the data or pick a single point to represent it.

There is another type of data period which is more cummulative.  An example might be world CO2 emissions.  Over a period of time you emit a certain amount of CO2.  At a constant rate if you double the period you emit twice as much.  So to aggregate over a longer period you sum the amounts over a shorter period.
In accounting you have the same type of differences.  So the amount of money in a bank account, an asset, is a point in time data point.  I don’t have 52 times as much money in the bank in a year that I do in a week.   However the amount of profit I make is exactly that.  To work out the profit for a year I would add up the profit (or loss) I made for each week.
In fact the two main financial summaries that are produced, the profit and loss is a period and the balance sheet is mostly point in time data.

## Zero, Single Entry and Double entry accounting systems

Single entry accounting systems represent each fact with a single entry.  This is the simplest method and actually corresponds to the common Python metaphor of DRY “Don’t Repeat Yourself”.  The downside of this is that you have no additional information to check that what you have entered is correct.  For very simple systems this can be ok.  Often you can actually get a check by verifying the information to another source.  So very simple accounting systems that are essentially driven from bank statements – the bank statement provides redundancy.

In computer terms old serial systems used to transmit bytes with parity.  The parity bit was an additional bit added to each byte so that the number of bits that were one was either odd or even.  This additional information gave you information that the information was correct.

## Petty Cash Accounting

## Trial Balance

A trial balance is a critical accounting output from a set of accounts. It is a named vector of account balances that in total should add up to zero.  In a manual set of accounts it was a trial balance as you would test to see that they did add up to zero before going on to the next stage of accounts preparation.  A simple mathematical manipulation will map a trial balance to a profit and loss statement and to a balance sheet.

The profit and loss data is all for a period and all the trial balance data is point in time data of the balance sheet at the end of the period except for the retained profit figure.  The retained profit figure  is an owners equity figure which appears on the balance sheet.

# Invoicing

When you invoice you create an entry against a debtors control account to indicate that your customer owes you money until they pay.  The debtors control account is cleared when money is transferred into your bank account.

# Stock

In a simple business all the stock is represented in a single account.  At the end of the accounting period the stock is revalued and the difference between the beginning plus the purchases and end valuations is then  cost of goods sold.
In a more elaborate stock taking system the value of each stock item will be measured in real time.
