# Abstract Factory Design Pattern

## Example Implementation

We want to make a device. This device can be a mobile phone, or laptop or tablet.  
These devices have some common parts like camera and display but the mobile camera  
can be used for mobile phone and laptop camera can be used for laptop. We can't use  
mobile camera for laptop, this constraint we want to enforce.  

Here, we have two products, Camera and Display. ICamera and IDisplay are two interfaces  
through which we have implemented MobileCamera and LaptopCamera, and, MobileDisplay and  
LaptopDisplay concrete products.

IDeviceAbstractFactory is an interface through which we have implemented different devices.  
Mobile and Laptop are two device that share IDeviceAbstractFactory interface. Mobile device  
have MobileCamera and MobileDisplay while Laptop device have LaptopCamera and LaptopDisplay.  

In the program, we have to pass the device using IDeviceAbstractFactory interface. This way  
the implementation need not know about the device type but what it will know is that, there  
will be one device that will have Camera and Display methods implemented.

## Use the Abstract Factory pattern when
- A system should be independent of how its products are created, composed,
  and represented.
- A system should be configured with one of multiple families of products.
- A family of related product objects is designed to be used together, and
  you need to enforce this constraint.
- You want to provide a class library of products, and you want to reveal
  just their interfaces, not their implementations.
