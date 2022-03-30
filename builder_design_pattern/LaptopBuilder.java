package builder_design_pattern;

public interface LaptopBuilder {
    String os = "";
    String processor = "";
    int ram = 0;
    int battery = 0;
    String screenResolution = "";

    // this method actually builds the Laptop (product)
    Laptop getLaptop();
}
