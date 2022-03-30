package builder_design_pattern_simple;

public class LaptopBuilder {
    private String os;
    private String processor;
    private int ram;
    private int battery;
    private String screenResolution;

    public LaptopBuilder setOs(String os) {
        this.os = os;
        return this;
    }

    public LaptopBuilder setProcessor(String processor) {
        this.processor = processor;
        return this;
    }

    public LaptopBuilder setRam(int ram) {
        this.ram = ram;
        return this;
    }

    public LaptopBuilder setBattery(int battery) {
        this.battery = battery;
        return this;
    }

    public LaptopBuilder setScreenResolution(String screenResolution) {
        this.screenResolution = screenResolution;
        return this;
    }

    // this method actually builds the Laptop (product)
    public Laptop getLaptop() {
        if (this.os == null || this.processor == null) {
            throw new RuntimeException("Required properties are not present");
        }
        return new Laptop(os, processor, ram, battery, screenResolution);
    }
}
