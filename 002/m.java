// Import ArrayList From java.util
import java.util.*;

public class m {
  public static void main(String[] args) {
    // Initialize An ArrayList Of Integers To Store All Even Fibonacii Numbers
    ArrayList<Integer> evenFib = new ArrayList<Integer>();
    int c = 0; // Counter
    int sumFib = 0; // Sum Of Even Fibonacii Numbers
    while (Fib(c) < 4000000) {
      if (Fib(c) % 2 == 0) // Checking If Fibonacii Number Is Even
        evenFib.add(Fib(c)); // Appending To evenFib ArrayList
      c++; // Incrementing Counter
    }
    for (int evenNum : evenFib)
      sumFib += evenNum; // Accumulation
  }
  // Function For Calculating Fibonacii Numbers Via Recursion
  public static int Fib(int Idx) {
    if (Idx == 0 || Idx == 1)
      return 0;
    else if (Idx == 2 || Idx == 3)
      return 1;
    else
      return Fib(Idx - 1) + Fib(Idx - 1);
  } 
}
