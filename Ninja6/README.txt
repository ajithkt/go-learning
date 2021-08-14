
Syntax
    func, receiver, identifier, params, returns  


variadic parameters and arguments
    func sample(s ...int){
        fmt.Println(s)
    }

Unfolding a slice
    x:={1,2,3,4,5,6,7}
    test(x...)
    func test (s ...int)m{
        fmt.Println(s)
    }

defers :- ececuting at last; it will execute when the func containing the difeeref function exits.

Anonymous function 
    func() { fmt.Println("HELLO FROM ANONYMOUS FUNCTION") }() //Anonymous function and it execution using () at last

Methods :- use receiver in a function & call that function using values of the receiver

returns
    multiple returns
    named returns - yuck!
func expressions
    assigning a func to a variable
callbacks
    passing a func into another func as an argument
closure 
    one scope enclosing another
    variables declared in the outer scope are accessible in inner scopes
    closure helps us limit the scope of variables
recursion
    a func that calls itself 
    factorial
