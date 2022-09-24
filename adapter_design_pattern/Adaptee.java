package adapter_design_pattern;

public class Adaptee {
    // This class represents a dependency/library whose function
    // we want to use as is with same signature but the function we
    // will call have different signature.

    public void specialRequest(String req, boolean sslVerify) {
        System.out.println("Invoking special request " + req + " with sslVerify " + sslVerify);
    }
}
