# Adapter Design Pattern

## Intent
Convert the interface of a class into another interface clients expect. Adapter
lets classes work together that couldn't otherwise because of incompatible
interfaces.

## Also Known as
Wrapper

## When to use
- When we want to use an existing class, but its interface does not match the
    one we need.
- When we want to create a reusable class  the cooperates with the unreleated and
    unforeseen classes, that is the classes that don't necessarily have compatible
    interfaces.
- When we need to use several existing subclasses, but it's impracticle to adapt
    their interfaces by subclassing everyone. An object interface can adapt the 
    interface of its parent class.
