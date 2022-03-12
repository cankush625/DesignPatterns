package singleton_design_pattern;

class Singleton {
    static private Singleton instance = null;

    // Make the constructor private, so that the instance cannot be created
    // from outside
    private Singleton() {}

    public static Singleton getInstance() {
        if (instance == null) {
            instance = new Singleton();
        }
        return instance;
    }

    public String sayHello() {
        return "hello";
    }
}
