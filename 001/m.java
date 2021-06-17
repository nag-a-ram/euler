public class m {
  public static void main(String[] args) {
    int[] arr = new int[1000]; // Initialize An Empty Array
    int sum = 0;
    for (int I = 1; I < arr.length; I++) {
      arr[I] = I; // Populate The Array
      if (arr[I] % 3 == 0 || arr[I] % 5 == 0) {
        sum += arr[I]; // Add Numbers That Are Multiples Of 3 And 5 From 0 - 1000
      }
    }
    System.out.println(sum);
  }
}
