public class m {
  public static void main(String[] args) {
    System.out.println(primeFactorOf(600851475143L)); // Get Prime Largest Factor Of Long Number
  }
  
  public static long primeFactorOf(long n) {
    long num = n;
    int p = 2;
    while (num > Math.pow(p, 2)) {
      if (num % p == 0) // Math?
        num /= p;
      else
        p++;
    }
    return num;
  }
}
