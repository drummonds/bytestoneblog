---
title: Pacioli got it wrong and why it doesn't matter.
date: 2025-08-23
updated: 2025-09-07
---

Pacioli

After doing a worked example of trading using normal accounting conventions with graphs and it
 just seems wrong and `a bit abitrary` as Tiger Beatle puts it in
[their documentation](https://github.com/tigerbeetle/tigerbeetle/blob/main/docs/coding/financial-accounting.md).  Luca Pacioli's wrote the original textbook in 1494 {{% ref Pacioli-1494 %}} and everyone has copied
 the conventions since.

I think it is wrong rather than just a matter of convention.    There are other good examples in
Science where for instance the electron has a negative charge so the flow of electrons is in the
opposite sent to which current is meant to flow.  The negative charge was discovered in
[1897 by JJ Thompson](https://en.wikipedia.org/wiki/Electron)

I am theorising that there are two problems which lead to this:

- Not using negative numbers
- Starting with a going concern rather than a fresh one

## Negative numbers

The Chinese ca 200 BCE first [invented and used negative numbers](https://nrich.maths.org/articles/history-negative-numbers).  A rod system
represented positive numbers as red and negative numbers as black.  In the 7th Century negative numbers were well established in India eg [Brahmagupta](https://en.wikipedia.org/wiki/Brahmagupta).
Mathematics was further transmitted to the Arabs with the invention of Arabic numerals and then to the Italians
where it replaced Roman numerals.

It would seem likely that this transmission was a diffusion of ideas along the silk road.  However there
are people like
[Srinivasa Ramanujan](https://en.wikipedia.org/wiki/Srinivasa_Ramanujan) who was an Indian genius.  One of the amazing
things about him is that he created and taught himself mathematics with intellectual leaps.  So much so that we are
still catching up with his inventions.  What this means is it is very possible for an individual genius to reinvent things
rather than assume the transfer of knowledge from the Chinese to the Indian hindu culture.  Both routes could be
possible.


They used the same convention as Pacioli, with amount sold as positive and amount purchased as negative.
However as it seems to be a single entry system a sale corresponds to an increase in cash rather than a movement.


The problem comes in the transmission of Hindu mathematics to Italy via the Arabs.  The Hindus had
negative numbers which they used for debts.  Along the way the Hindu/Arabic numerals were
transmitted but negative numbers were dropped.

In early Renaissance Venice ["Pacioli, like other mathematicians of his time, did his utmost to avoid even the use of a symbol for minus, let alone a negative number"](https://www.jstor.org/stable/2490577).  However the use of negative numbers for balances with the direction gives the graph representation meaning.
The use of negative numbers with the direction gives the graph representation meaning.
Without negative numbers you move in Pacioli's language from shall give `in dare` and
shall receive "in habere" and by "per" and to "a" in the journal.  In the ledger you have the
left and right hand side which substitute for positive and negative numbers.

There is an advantage of separating out the positive and negative numbers in that it becomes easier
to error check your balances.  Making simple errors is very easy however you do it espeically if writing
all the numbers out 4 times by hand.

At least two people who have used graph theory have both swapped the convention as it just
makes more sense.  So perhaps {{% ref Kleppmann-2011 %}} did it and he just said
```
Now, by convention, accountants flip the sign on all of the blue and pink nodes’ balances, which means that the two sums end up being equal. And that’s why it’s called a balance sheet.```

## Going concern

I think the reason Pacioli went for his interpretation is that he started with an inventory of a going
concern.  The first movement is cash to equity and once you have made that choice everything has to
follow. If he had started with an empty balance sheet I think he might have gone in a different
direction.


In the end it becomes a matter of presentation and computers are adept at reformatting and
redisplaying in your own personal choice.
