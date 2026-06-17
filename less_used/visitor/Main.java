// PROBLEM: Jak dodawać nowe operacje na obiektach bez zmiany ich klas?

// Interfejs elementu
interface Document {
    void accept(DocumentVisitor visitor);
}

// Dokument PDF
class PDFDocument implements Document {
    public void accept(DocumentVisitor visitor) {
        visitor.visit(this);
    }
}

// Dokument Word
class WordDocument implements Document {
    public void accept(DocumentVisitor visitor) {
        visitor.visit(this);
    }
}

// Visitor
interface DocumentVisitor {
    void visit(PDFDocument doc);
    void visit(WordDocument doc);
}

// Konkretna operacja: eksport
class ExportVisitor implements DocumentVisitor {

    public void visit(PDFDocument doc) {
        System.out.println("Export PDF");
    }

    public void visit(WordDocument doc) {
        System.out.println("Export Word");
    }
}

// użycie
public class Main {
    public static void main(String[] args) {

        Document doc = new PDFDocument();

        // operacja wykonywana przez Visitor
        doc.accept(new ExportVisitor());
    }
}