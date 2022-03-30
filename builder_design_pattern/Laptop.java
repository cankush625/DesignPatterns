package builder_design_pattern;

public class Laptop {
    private String os;
    private String processor;
    private int ram;
    private int battery;
    private String screenResolution;

    public Laptop(String os, String processor, int ram, int battery, String screenResolution) {
        this.os = os;
        this.processor = processor;
        this.ram = ram;
        this.battery = battery;
        this.screenResolution = screenResolution;
    }

    @Override
    public String toString() {
        return "Laptop{" +
                "os='" + os + '\'' +
                ", processor='" + processor + '\'' +
                ", ram=" + ram +
                ", battery=" + battery +
                ", screenResolution='" + screenResolution + '\'' +
                '}';
    }
}
