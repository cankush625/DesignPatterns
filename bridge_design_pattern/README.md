# Bridge Design Pattern

## Intent
Decouple an abstraction from its implementation so that the two can vary independently.  
The abstraction has an implementation and any implementation can be used with any abstraction.

## Also Known as
Handle/Body

## When to use
- Use this pattern when we want to avoid a permanent binding between an
  abstraction, and it's implementation. This might be the case, for example,
  when the implementation must be selected or switched at run-time.
- Both abstractions and their implementations should be extensible by
  subclassing them independently. We can combine abstractions and implementations
  and extend them independently.
- Changes in the implementation of an abstraction should have no impact
  on client; that is, their code should not have to be recompiled.
- We want to share an implementation among multiple objects (perhaps using
  reference counting), and this fact should be hidden from the client. 
