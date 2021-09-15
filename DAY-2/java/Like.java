public class Like {
    private int count = 0;

    public synchronized void add() {
        this.count++;
        System.out.println("like count: " + this.count);
    }
}