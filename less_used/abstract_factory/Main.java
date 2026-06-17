// PROBLEM: Jak tworzyć rodziny powiązanych obiektów bez ujawniania ich konkretnych klas?

// Interfejs produktu
interface Button {
    void render();
}

// Implementacje produktów
class WindowsButton implements Button {
    public void render() {
        System.out.println("Windows Button");
    }
}

class MacButton implements Button {
    public void render() {
        System.out.println("Mac Button");
    }
}

// Fabryka abstrakcyjna
interface GUIFactory {
    Button createButton();
}

// Konkretna fabryka Windows
class WindowsFactory implements GUIFactory {
    public Button createButton() {
        return new WindowsButton();
    }
}

// Konkretna fabryka Mac
class MacFactory implements GUIFactory {
    public Button createButton() {
        return new MacButton();
    }
}

// użycie
public class Main {
    public static void main(String[] args) {

        // wybór rodziny obiektów
        GUIFactory factory = new WindowsFactory();

        // tworzenie bez znajomości konkretnej klasy
        Button button = factory.createButton();

        button.render();
    }
}