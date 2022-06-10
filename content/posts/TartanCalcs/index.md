---
title: "Drummond of Megginch Tartan thread counts"
date: 2022-06-10
tags: ["humphrey","jupyter"]
---

# Given thread counts produce table
Wait for this to come out

```python
test = "R14 DB2 R4 DB4 R70 LB4 R4 DB20 R4 G4 R4 G74 R6 DB4 R12"

    
```

```python
import re

PATTERN = re.compile("([a-zA-Z]*)(\d*)")

def split_thread(threadcount):
    """Given a thread count such as R14, will split to R and integer 14"""
    match = re.search(PATTERN, threadcount)
    if match:
        return match.group(1),match.group(2)
    return threadcount,""
    

class Stripe:
    def __init__(self, short_colour, count):
        self.short_colour = short_colour
        self.count = int(count)

    def __str__(self):
        return f"{self.short_colour}{self.count}"

class Tartan:
    def __init__(self):
        self.stripes = []
        
    @classmethod
    def from_space_threadcount(cls, threadcount):
        tartan = cls()
        threads = threadcount.split(" ")
        for thread in threads:
            colour,count = split_thread(thread)
            tartan.stripes.append(Stripe(colour, count))
        return tartan

    def __str__(self):
        s= ""
        for stripe in self.stripes:
            s += f"{stripe} "
        return s

    @property
    def threadcount(self):
        tc = 0
        for stripe in self.stripes:
            tc += stripe.count
        return tc

a = Tartan.from_space_threadcount(test)
print(f"Pattern = {a}")
print(f"Num threads = {a.threadcount}")
```

    Pattern = R14 DB2 R4 DB4 R70 LB4 R4 DB20 R4 G4 R4 G74 R6 DB4 R12 
    Num threads = 230
