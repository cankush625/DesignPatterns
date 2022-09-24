package adapter_design_pattern;

public class Client {
    public boolean getResponse(String req) {
        Adaptee adaptee = new Adaptee();
        Adapter adapter = new Adapter(adaptee);
        return adapter.request(req);
    }
}
