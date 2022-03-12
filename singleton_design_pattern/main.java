package singleton_design_pattern;

public class Main {
    public static void main(String[] args) {
        Singleton sing = Singleton.getInstance();
        System.out.println(sing.sayHello());
   }
}
