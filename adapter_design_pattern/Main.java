package adapter_design_pattern;

public class Main {
    public static void main(String[] args) {
        Client client = new Client();
        boolean isSuccess = client.getResponse("https://example.com");
        System.out.println(isSuccess);
    }
}
