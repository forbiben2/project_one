GO
    -functions can return multiple values in GO
    -If statements check for parameters
    -nil-no errors
    -randomFormat function-returns randomly selected result
    -init functions-execute on program start
    -string-a slice of bytes
    -maps-group of strings
    -test functions can be added to a package for testing
        -should be seperate from other functions
    -by convention a packages name is the same as the last element in its import path
        ie. math/rand uses files beginning with the statment package rand
    -you can group imports into a parenthesized statment(factored import statement)
    -in Go a name is exported if it begins with a capital letter 
        ie Pi,not pi
    -functions can take 0 or more arguments
    -types are written after the variable name
        ie x int
    -when types are shared you can omit all but the last variables type
        ie y,x int
    -swap functions can return 2 results but swaps them
    Returns 
        -values may be named if so they are treated as variables defined at the top of a function
            -names should be used to document the meaning of the return values
        -a return statement without arguments returns the named values aka "naked" return
            -should only be used for shorter functions as it can harm readability in longer functions
    Typing
        -List of Types
            -bool-true/false
            -string
            -int
                int8 int16 int32 etc
                -int it self usually uses system bit size
                -standard typing for numbers unless a specific use is needed
            -byte-int8
            -rune-int32
            -float
                float32 float64
            -complex64 complex128

        -typing for variables can be ommited with initializers as the variable will take the initializers type
             ie var k int=3 same as k:=3
        -if a value is uninitialzed it is given a value of zero
            -for bool this is false
            -int is 0
            -string is " "
        -Go requires explicit conversion of types 
        -Constants are declared like variables
            -numeric constants are high precisions variables
    Looping Constructs
        -Go has 1 the for loop it has three compontents separated by a ;
            -the init statment before the first iteration
                -often a short variable declaration when declared here only inscope for the for statment
            -the conditional expression that is evaluated before each iteration
            -the post statement executed at the end of every loop
                -continues loop until this boolean is false
                -if ommited and not rectified else where will loop infinitely
            -all three are optional and can be done in other areas if needed
            -for statements use {} not ()
            -example 
                    for i;i++;i<3{

                    }
            -for statements can look over a range of data
            -for can be dynamic
        -math.Sqrt is solved by using a looping statement
        -loops can be used inside loops also called a nested loop
            -each loop must break independently
            -can break other loops from within a loop with a labelled statement
                ie package main

                    func main() {
                         out:
                        for i := 0; i < 10; i++ {
                            for j := 0; j < 10; j++ {
                                 if i + j == 20 {
                                    break out
                                 }
                             }
                         }
                    }
        -break can make loops more efficient by leaving early when neccessary
    If Statements
        -Like For statements can start with a short statement to execute before the condition
        -variables declared inside are only inscope til the end of the if statement
            -will work in related else statements?
        -slower then a switch statement
    Switch Statement
        -simplified if else statement
        -cases-different options depending on input
        -default used if no cases are met
        -runs first case whose value is equal to a condition
        -does not need to be a constant value for a switch statement nor must it be a number
        -evaluates top to bottom
        -switch without a statement is same as switch true
        -faster then an if statement
    Defer Statement
       -calculates immediately but is not returned til the surrounding functions are evaluated 
    


