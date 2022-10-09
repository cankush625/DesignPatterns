package decorator_design_pattern;


/**
 * Graphic decorator that conforms to the Graphic interface
 */
public class GraphicDecorator implements Graphic {
    protected Graphic graphic;

    public GraphicDecorator(Graphic graphic) {
        this.graphic = graphic;
    }

    @Override
    public void draw() {
        this.graphic.draw();
    }
}
