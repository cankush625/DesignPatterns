package adapter_design_pattern;

public class Adapter implements ITarget{
    Adaptee adaptee;

    public Adapter(Adaptee adaptee) {
        this.adaptee = adaptee;
    }

    @Override
    public boolean request(String req) {
        this.adaptee.specialRequest(req, false);
        return true;
    }
}
