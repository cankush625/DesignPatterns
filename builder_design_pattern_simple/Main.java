package builder_design_pattern_simple;

public class Main {
    public static void main(String[] args) {
        /*
          Use Builder Design Pattern when:
          1. we have classes with optional properties and creation of objects is complex.
          2. the construction process must allow different representations for the object that's constructed.
          Builder pattern allows us to make the object immutable.
          The Laptop object is immutable here. We can't set/change the properties after construction.
          Here, setters are provided on the Builder class and not on the actual product (Laptop) class.
          3. Builder pattern allows us to create immutable objects by having only one constructs. We don't need
          to create different constructor for different combination of required fields.
         */
        Laptop laptop = new LaptopBuilder().setOs("macOS").setProcessor("M1").setRam(32).getLaptop();
        System.out.println(laptop);
    }
}
