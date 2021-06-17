import java.util.*;

public class m {
  public static void main(String[] args) {
    System.out.println(getSmallestNumDivisibleBy(1, 20));
  }

  public static int getSmallestNumDivisibleBy(int startNum, int endNum) {
    int isBest = 0;
    int bestNum = 0;
    while (isBest != endNum) {
      if (bestNum == 0) {
        bestNum += 1;
        continue;
      }
      isBest = 0;
      for (int I = startNum; I <= endNum; I++) {
        if (bestNum % I == 0)
          isBest++;
      }
      bestNum += 1;
    }
    return bestNum - 1;
  }
}
