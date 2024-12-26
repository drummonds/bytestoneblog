---
title: Accounting movements
date: 2024-12-26
mermaid: true
----

Accounting is really the tracking of resources.  [Kleppmann](https://martin.kleppmann.com/2011/03/07/accounting-for-computer-scientists.html )[IA](https://web.archive.org/web/20110717155537/https://martin.kleppmann.com/2011/03/07/accounting-for-computer-scientists.html) really cyrstallised this for me as graph theory.  A transaction is about the movement of resources from one place to another.

If a founder puts money into a business you can represent the transaction as a movement of money:

{{< mermaid >}}
flowchart LR
   F([Founder]) -- £500 --> BA([Business])
{{< mermaid >}}


{{< mermaid >}}
flowchart LR
   F([Founder]) -- £500 --> B@{ shape: odd, label: "Business" }
{{< mermaid >}}

The founder is hoping that the business will do well and then the following will occur:

{{< mermaid >}}
flowchart LR
    BA([Business]) -- £500 --> F([Founder])
{{< mermaid >}}

One of the principles of accounting is that is done from the point of view of single entity.  Here we have two, the founder and the business.  If the founder is self employed then the boundaries between the two get blurred.  However in the UK the tax authority [HMRC](https://www.gov.uk/money/income-tax) will normally require you to complete a tax return on behalf of you and your business.



