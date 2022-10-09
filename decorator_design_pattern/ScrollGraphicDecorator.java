package decorator_design_pattern;


/**
 * Scroll Graphic decorator adds scrolling functionality to the graphic
 */
public class ScrollGraphicDecorator extends GraphicDecorator {
    public ScrollGraphicDecorator(Graphic graphic) {
        super(graphic);
    }

    private void addScrolling() {
        System.out.println("Add Scrolling to the graphic");
    }

    @Override
    public void draw() {
        super.draw();
        addScrolling();
    }
}
