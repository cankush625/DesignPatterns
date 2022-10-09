package decorator_design_pattern;

public class Main {
    public static void main(String[] args) {
        // Draw TextBox with Border
        Graphic BorderedTextBox = new BorderGraphicDecorator(new TextBox());
        BorderedTextBox.draw();

        System.out.println("------------------------------------");

        // Draw TextBox with Scrolling functionality
        Graphic ScrollableTextBox = new ScrollGraphicDecorator(new TextBox());
        ScrollableTextBox.draw();

        System.out.println("------------------------------------");

        // Draw TextBox that has Border and Scrolling functionality
        Graphic ScrollableBorderedTextBox = new ScrollGraphicDecorator(
            new BorderGraphicDecorator(
                new TextBox()
            )
        );
        ScrollableBorderedTextBox.draw();
    }
}
