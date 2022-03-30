package builder_design_pattern;

public class Main {
    public static void main(String[] args) {
        // Build Apple laptop
        AppleLaptop appleLaptop = new AppleLaptopBuilder().setOs("macOS").setProcessor("M1").getLaptop();
        appleLaptop.runXCode();

        // Build Windows laptop
        WindowsLaptop windowsLaptop = new WindowsLaptopBuilder().setOs("windows").setProcessor("AMD").getLaptop();
        windowsLaptop.runEdgeBrowser();
    }
}
