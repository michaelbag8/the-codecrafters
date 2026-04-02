# Emmanuel Elaigwu

The goal of the project is to rebuild the go-reloaded project from scratch-up.,but this time split up among team members with individual working solo on each part of the go-reloaded functions. I was given the "fixArticle" part to handle, at first seems easy but along the line i was stuck with a bug which i struggled to find, but with some help from the other team members i was able to debug and finish the part given to me...Its fun and at thesame time tiring.But it's all learning in the making

The part i handle "fixArticle" iterate over an article (which are a sequence strings) look for "a" where not appropriate and replace it with an "an" vice-versa.

And i learnt debugging is patience.


# Samuel Jireh

## -My personal contribution was fixPuntuation and fix Quotes:
    Fixing Punctuation of just a basic was easy but it was not handlying some edge cases like group punctuations so i have to come up with another code that can handly group punctuations, after that i called the code for the group punctuation into the fix punctuation so it can be use as one.
    I also created the fix quotes and add it in the fix punctuation code to act as one function

## -What my function dose:
###    In short:  
    Your program takes messy sentences and fixes punctuation spacing so they look grammatically correct and clean. It handles normal punctuation, grouped punctuation like ... or !?, and quotes 'word' properly.

## -One thing you found hardest today:
    The hardest thing i found out in today work was that without a good plan you can easly get lost in a project and you will want to give up, At the time i ask my fellow for help he made me understand the project and got me back to my feet, after that i started coding non stop tell the time i finaly finshed my work and i started helping my fellow with theis.


## Contributor : Michael Bulus

## Overview

Sentinel Reborn is a Go command line program.

The program reads text from an input file, transforms the content, then writes the result to an output file.

This project tests my skill in:

- file input and output
- string processing
- function design
- edge case handling
- program structure


## What the Program Does

The program:

1. reads content from an input file
2. applies a set of text transformations
3. handles bad input without crashing
4. writes the final result to an output file


## How to run the program
```bash
go run . input.txt output.txt

```

## Transformations functions useds
### 1. Hexadecimal to decimal conversion
Converts a hexadecimal value followed by `(hex)` into decimal.

### 2. Binary to decimal conversion
Converts a binary value followed by `(bin)` into decimal.

### 3. Uppercase transformation
Changes the previous word, or previous words, to uppercase when the `(up)` is used.

### 4. Lowercase transformation
Changes the previous word, or previous words, to lowercase when the `(low)` is used.

### 5. Capitalization transformation
Changes the previous word, or previous words, so the first letter is uppercase and the remaining letters are lowercase when the `(cap)` is used.

### 6. Multi word case transformation
Supports markers with a number, such as `(up, 2)`, `(low, 3)`, or `(cap, 2)`, to apply the change to more than one previous word.


## My Personal Contribution
1. I helped assign responsibilities to group members.
2. I worked on converting hexadecimal and binary input to decimal.
3. I contributed to refactoring other team members' code to cover edge cases.
4. I assisted team members who were having issues with their code.

## One Thing I Found Hardest Today
1. One thing I found hardest today was writing a function to get the previous word.
2. Another hard part was reviewing and fixing code I did not write. I had to first understand the logic, then trace the problem, then fix the issue without breaking other parts of the program.

## One Thing I Understand Now That I Did Not Understand This Morning
1. One thing I understand now that I did not understand this morning is the value of planning.
2. Good planning makes execution easier. Poor planning slows down the whole team.
3. I also understand teamwork better now. No one has all the knowledge. Every team member matters. Each person adds value to the final result.
