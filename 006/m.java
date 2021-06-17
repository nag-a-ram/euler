public class m {
  public static void main(String[] args) {
    int sumOfSquaresNum = getSumOfSquaresOf(1, 100); // Invoke Function To Get Sum Of Squares From 1, 100
    int squareOfSumNum = (int) getSquareOfSumOf(1, 100); // Invoke Function To Get Square Of The Sum Of 1, 100
    System.out.println(squareOfSumNum - sumOfSquaresNum); // Get Their Difference And Print It Out
  }

  public static int getSumOfSquaresOf(int startNum, int endNum) {
    int sum = 0;
    for (int I = startNum; I <= endNum; I++) {
      sum += Math.pow(I, 2); // Sum Squared Values
    }
    return sum;
  }
  
  public static double getSquareOfSumOf(int startNum, int endNum) {
    int sum = 0;
    for (int I = startNum; I <= endNum; I++) {
      sum += I; // Sum Values
    }
    return Math.pow(sum, 2); // Square Values
  }
}
