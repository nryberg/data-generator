10/26/2016  The ABC's of data

* Apples and Oranges
* Bike Rentals
* Cats and Dogs
* D
* E
* Finance
* G
* H
* I
* J
* K
* M
* N
* O
* P
* Q
* R
* S
* T
* U
* V
* W
* X
* Y
* Z

== Basic structure of data generation ==
Combining random values within a limits.  Quantities are limited, possible items are limited.  Building out those limitations is the hard part - can you do it in a way that's easy to maintain?

Make a pattern that's reasonably predictable:

Dimensions -
* Index Number - 1, 2, 3, 4, 5 -
* Name - Simple Text - Apple, Orange, Banana
* Category - Something to group by
* Value - something to multiply by
*
Fact - only one
* Date - within limits - varying atomicity
* Dimension One Index
* Dimension One Value - quantity
* Dimension Two Index

== General Feel ==
Should _feel_ like real data, but not really - has to straddle the line.  Setting the params is the hardest part.
