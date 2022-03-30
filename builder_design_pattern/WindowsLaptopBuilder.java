package builder_design_pattern;

public class WindowsLaptopBuilder implements LaptopBuilder {
    private String os;
    private String processor;
    private int ram;
    private int battery;
    private String screenResolution;

    public WindowsLaptopBuilder setOs(String os) {
        this.os = os;
        return this;
    }

    public WindowsLaptopBuilder setProcessor(String processor) {
        this.processor = processor;
        return this;
    }

    public WindowsLaptopBuilder setRam(int ram) {
        this.ram = ram;
        return this;
    }

    public WindowsLaptopBuilder setBattery(int battery) {
        this.battery = battery;
        return this;
    }

    public WindowsLaptopBuilder setScreenResolution(String screenResolution) {
        this.screenResolution = screenResolution;
        return this;
    }

    // this method actually builds the Laptop (product)
    public WindowsLaptop getLaptop() {
        if (this.os == null || this.processor == null) {
            throw new RuntimeException("Required properties are not present");
        }
        return new WindowsLaptop(os, processor, ram, battery, screenResolution);
    }
}
