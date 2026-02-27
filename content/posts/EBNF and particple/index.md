---
title: EBNF and Participle
date: 2023-11-14
summary: Using Extended Backus-Naur Form (EBNF) notation and how the Participle parsing library leverages it for elegant language processing.
---

I have been using the https://github.com/alecthomas/participle library for parsing
a plain text accounting format and have learnt something.  I decided to use the library
rather than write my own parser as I am planning to extend the language with a 
number of new directives and wanted a good framework to describe and parse these.

I have learnt a number of things:

- indented syntax
- debugging
- strings

## Indented syntax

Plain text account format uses indentation to describe blocks.  In order to make
this work with a Lexer and a Parser, a good solution is to have a pre lexer that converts the indentation to IDENT and DEDENT tokens. 

For an example see the indentLexer in https://github.com/drummonds/luca/blob/master/internal/parser/parser.go.

## Debugging

I found it worth creating a debugging parser that would turn on debugging and
 crucially let me look at the result of the tokenizer.

## Strings

My first string parser return the whole string including the delimiters which was
quite ugly.  