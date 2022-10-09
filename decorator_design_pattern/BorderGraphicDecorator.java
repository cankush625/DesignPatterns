package decorator_design_pattern;


/**
 * Border Graphic decorator applies border to the graphic
 */
public class BorderGraphicDecorator extends GraphicDecorator {
    public BorderGraphicDecorator(Graphic graphic) {
        super(graphic);
    }

    private void drawBorder() {
        System.out.println("Draw Border to the graphic");
    }

    @Override
    public void draw() {
        super.draw();
        drawBorder();
    }
}
