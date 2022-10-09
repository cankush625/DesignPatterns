package decorator_design_pattern;


/**
 * TextBox implements Graphic and overrides method draw which
 * draws TextBox on the UI
 */
public class TextBox implements Graphic {
    @Override
    public void draw() {
        System.out.println("Draw TextBox graphic");
    }
}
