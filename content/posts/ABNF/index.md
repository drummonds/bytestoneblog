---
title: "Augmented Backus Naur Format"
date: 2025-02-15
updated: 2026-02-27
tags: ["ABNF", "EBNF"]
summary: A comparison of ABNF and EBNF notation systems
mermaid: true
---

# Comparing ABNF and EBNF

Backus Naur formats are metalanguages used to describe other languages and structured text. They have a longish history.  While formal language description systems date back to [Pāṇini's][Panini] Sanskrit grammar from 2,500 years ago, modern variants evolved from John Backus's work in the 1950s.

[Panini](https://en.wikipedia.org/wiki/P%C4%81%E1%B9%87ini)
{{< mermaid_fig title="Demo" caption="Variation history"  >}}
graph TB
pan[<b>-350</b>: <a href='https://en.wikipedia.org/wiki/P%C4%81%E1%B9%87ini'>Pāṇini</a>]
bal[<b>1958</b>: <a href='https://en.wikipedia.org/wiki/ALGOL_58'>Backus Algol</a>]
bnf[<b>1963</b>: <a href='https://en.wikipedia.org/wiki/Backus%E2%80%93Naur_form'>Backus Naur Form</a>]
wsn[<b>1970</b>: <a href='https://en.wikipedia.org/wiki/Wirth_syntax_notation'>Wirth Syntax Notation</a>]
ebnf[<b>1977</b>: <a href='https://dl.acm.org/doi/pdf/10.1145/359863.359883'>Wirth</a> <a href='https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form'>EBNF</a>]
xml[<b>2006</b>: <a href='https://www.w3.org/TR/xml/'>XML EBNF</a>]
gon[<b>2009</b>: <a href='https://pkg.go.dev/golang.org/x/exp/ebnf'>Go EBNF</a>]
abnf[<b>2010</b>: <a href='https://en.wikipedia.org/wiki/Augmented_Backus%E2%80%93Naur_form'>Augmented BNF</a>]
iso[<b>1996</b>: <a href='https://en.wikipedia.org/wiki/Wirth_syntax_notation'>ISO BNF</a>]

pan --> bal
bal --> bnf
bnf --> wsn
bnf --> xml
wsn --> ebnf
ebnf --> xml
ebnf --> gon
ebnf --> abnf
ebnf --> iso

     classDef green fill:#9f6,stroke:#333,stroke-width:2px;
     classDef orange fill:#f96,stroke:#333,stroke-width:4px;
     class abnf green
     class pan orange

{{< /mermaid_fig>}}

## Evolution of Notation Systems

### BNF and Wirth Syntax Notation (1960s-1970s)
The original Backus-Naur Form (BNF) was developed for defining ALGOL 58's syntax. Niklaus Wirth later created the Wirth Syntax Notation (WSN), which added important features like repetition using `{...}`.

### EBNF Development (1977)
EBNF emerged from Wirth's [beautifully short article](http://pascal.hansotten.com/uploads/ebnf/What%20can%20we%20do%20Notation%20EBNF.pdf) in ACM. Key developments in Pascal's documentation:
- 1969: Original BNF notation
- 1971: Introduction of WSN
- 1973-74: Refinements in Pascal reports
- 1977: Formal EBNF proposal
- 1985: EBNF adoption in Pascal (3rd Edition)

### Modern Variants
- **Go EBNF**: Simplified version that leverages Go language conventions
- **ISO EBNF**: Formal but more complex standard
- **ABNF**: Standardized through IETF RFCs with precise rules

## Syntax Comparison

| Feature      | BNF     | EBNF    | ABNF    | Notes |
|--------------|---------|---------|---------|-------|
| Definition   | ::=     | =       | =       | Rule declaration |
| Alternatives | \|      | \|      | \|      | Choice between options |
| Optional     | -       | [...]   | [...]   | Optional elements |
| Repetition   | -       | {...}   | n*m     | ABNF allows count ranges |
| Grouping     | -       | (...)   | (...)   | Combining elements |
| Comments     | -       | (*...*) | ;       | ABNF uses line comments |
| Termination  | -       | ;       | ;       | End of rule marker |

## Modern Considerations

### Case Sensitivity
ABNF added case sensitivity support through [RFC 7405](https://datatracker.ietf.org/doc/html/rfc7405), making it more precise for modern usage.

### Unicode Support
While ABNF officially uses ASCII, there are proposals like [Leonard and Kyzivat's draft](https://datatracker.ietf.org/doc/html/draft-seantek-unicode-in-abnf-00) to standardize Unicode handling.

[Panini]: https://en.wikipedia.org/wiki/P%C4%81%E1%B9%87ini
