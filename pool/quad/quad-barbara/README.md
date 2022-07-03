# Quad

Quad with Matthew, Maye & Barbara

## Structure

We agreed as a group to create 3 separate branches, so that we could each
give our best effort at solving the problems and learning for ourselves whilst also helping out each
other along the way.

master - Matthew's function implementations\
barbara - Barbara's function implementations\
maye - Maye's function implementations*

*N.B. While Barbara and I fared well with writing the code, Maye felt she needed a bit more familiarity with Go (she also had to go in to work!), and the meaning of the various Git commands. Therefore, when we met up, Barbara and I helped her specifically with understanding Go modules,packages, for loops and conditional statements. We also discussed on Discord how we used these
concepts to achieve the correct solutions!

For now, we decided not to merge the branches so that is quick and easy to see the differences in
the way Barbara and myself approached the problem!

## Code Structure

All the functions (QuadA - QuadE), along with any helper functions, may be found in the piscine/quad.go file.

We took slightly different approaches to our solutions, but of course they share many similarities. Barbara used implicit function declarations and maps, for example, while my approach used slices and was very mathematical, probably due to my maths background (I did a year of Maths at University before leaving to focus on other things).

We have included comments so hopefully the purpose of our code is clear!

## Testing Structure

I (Matthew) created a nested loop to iterate through different test conditions for each function,
and separated each test into its own sub-folder within the 'tests' folder, while Barbara submitted
the specific test cases as per the examples given, to be commented/uncommented and run when required!