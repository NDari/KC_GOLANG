// NoiseMaker.java
public interface NoiseMaker {
  String speak();
}
// Dog.Java
public class Dog implements NoiseMaker {
  public String speak() {
    return "Arf! Arf!";
  }
}
// Cat.Java
public class Cat implements NoiseMaker {
  public String speak() {
    return "Meow!";
  }
}
// MakeNoise.java
class talk {
  public static void makeNoise(NoiseMaker n) {
    n.speak();
  }
}
