# Retrospective

- This is the first time I've ever used go
- Going into this particular puzzle I had limitted knowledge of arrays and less knowledge on how slices worked in go and really just a basic understanding of the variable format.


## Self Review

### Where I struggled

- The first portion was solved quickly after a brief search on arrays in go, however I was not aware of the differences between slices and arrays
  - This led to my second answer being wrong at first as the same slice, which at the time I thought was an array, was being passed into the second method after the first.
    - My original solution to get it working how I wanted was to reload the slice from the file. I didn't want to do this long term so that's when I went and watched a video on arrays in go.

### What I learned

- Slices exist and act differently from arrays
- Slices are a reference type and the modifications in later methods will stick
- Arrays are not a reference type
- How to 'convert' from a slice to an array
