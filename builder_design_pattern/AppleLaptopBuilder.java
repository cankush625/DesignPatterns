package builder_design_pattern;

public class AppleLaptopBuilder implements LaptopBuilder {
    private String os;
    private String processor;
    private int ram;
    private int battery;
    private String screenResolution;

    public AppleLaptopBuilder setOs(String os) {
        this.os = os;
        return this;
    }

    public AppleLaptopBuilder setProcessor(String processor) {
        this.processor = processor;
        return this;
    }

    public AppleLaptopBuilder setRam(int ram) {
        this.ram = ram;
        return this;
    }

    public AppleLaptopBuilder setBattery(int battery) {
        this.battery = battery;
        return this;
    }

    public AppleLaptopBuilder setScreenResolution(String screenResolution) {
        this.screenResolution = screenResolution;
        return this;
    }

    // this method actually builds the Laptop (product)
    public AppleLaptop getLaptop() {
        if (this.os == null || this.processor == null) {
            throw new RuntimeException("Required properties are not present");
        }
        return new AppleLaptop(os, processor, ram, battery, screenResolution);
    }
}
