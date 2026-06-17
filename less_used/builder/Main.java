// PROBLEM: How to create an object with many optional fields without using many constructors? And allow to set fields in any order?
// javac Main.java
// java Main

class User {
    private final String name;
    private final int age;
    private final String email;
    private final String address;
    private final String phone;

    // Constructor with required field only.
    public User(String name) {
        this(name, 0, null, null, null);
    }

    // Constructor with some optional fields.
    public User(String name, int age) {
        this(name, age, null, null, null);
    }

    // Constructor with more fields.
    public User(String name, int age, String email) {
        this(name, age, email, null, null);
    }

    // Full constructor.
    public User(String name, int age, String email, String address, String phone) {
        this.name = name;
        this.age = age;
        this.email = email;
        this.address = address;
        this.phone = phone;
    }

    @Override
    public String toString() {
        return "User{name='" + name + "', age=" + age +
                ", email='" + email + "', address='" + address +
                "', phone='" + phone + "'}";
    }

    // BUILDER
    public static class Builder {
        private String name;
        private int age;
        private String email;
        private String address;
        private String phone;

        public Builder setName(String name) {
            this.name = name;
            return this;
        }

        public Builder setAge(int age) {
            this.age = age;
            return this;
        }

        public Builder setEmail(String email) {
            this.email = email;
            return this;
        }

        public Builder setAddress(String address) {
            this.address = address;
            return this;
        }

        public Builder setPhone(String phone) {
            this.phone = phone;
            return this;
        }

        public User build() {
            return new User(name, age, email, address, phone);
        }
    }
}

// DEMO
public class Main {
    public static void main(String[] args) {

        // Using multiple constructors (problem)
        User u1 = new User("John");
        User u2 = new User("John", 25);

        // Using Builder (solution)
        User u3 = new User.Builder()
                .setName("John")
                .setAge(25)
                .setEmail("john@example.com")
                .setAddress("Warsaw")
                .setPhone("123456789")
                .build();

        System.out.println(u1);
        System.out.println(u2);
        System.out.println(u3);
    }
}