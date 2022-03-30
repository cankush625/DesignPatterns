package builder_design_pattern;

public class WindowsLaptop extends Laptop {

    public WindowsLaptop(String os, String processor, int ram, int battery, String screenResolution) {
        super(os, processor, ram, battery, screenResolution);
    }

    public void runEdgeBrowser() {
        System.out.println("Starting Edge browser");
    }
}
