package builder_design_pattern;

public class AppleLaptop extends Laptop {

    public AppleLaptop(String os, String processor, int ram, int battery, String screenResolution) {
        super(os, processor, ram, battery, screenResolution);
    }

    public void runXCode() {
        System.out.println("Starting XCode");
    }
}
