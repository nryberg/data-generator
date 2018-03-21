## The ABC's of data

* [Apples and Oranges](./Samples/Apples_and_Oranges/fruit.csv)
* Bike Rentals
* Cats and Dogs
* D
* E
* Finance
* G
* Hotels
* I g
* J
* Kangaroos in Australia
* Margaritas and beers
* N
* O
* Plane Trips
* Q
* R
* Ships at sea
* T
* U
* V
* W
* X
* Y
* Z

### Basic structure of data generation
Combining random values within a limits.  Quantities are limited, possible items are limited.  Building out those limitations is the hard part - can you do it in a way that's easy to maintain?

Make a pattern that's reasonably predictable:

Dimensions -

* Two Dimensions preferably
* Index Number - 1, 2, 3, 4, 5 -
* Name - Simple Text - Apple, Orange, Banana
* Category - Something to group by
* Value - something to multiply by

A related Dimension - 

* Just one 
* Tied to the category of the first dimension

Fact - only one

* Date - within limits - varying atomicity
* Dimension One Index
* Dimension One Value - quantity
* Dimension Two Index

### General Feel
Should _feel_ like real data, but not really - has to straddle the line.  Setting the params is the hardest part.
