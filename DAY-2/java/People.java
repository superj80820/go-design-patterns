public class People extends Thread {
    private final Like like;

    public People(Like like) {
        this.like = like;
    }
    public void run() {
        for (int i = 0; i < 10000; i++) {
            like.add();
        }
    }
}