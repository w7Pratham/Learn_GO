constants cannot be declared using the := syntax

For

The basic for loop has three components separated by semicolons:

    the init statement: executed before the first iteration (optional)
    the condition expression: evaluated before every iteration
    the post statement: executed at the end of every iteration (optional)

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.