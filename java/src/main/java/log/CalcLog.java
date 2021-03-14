package log;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.math.RoundingMode;

// Calculate the natural log
public class CalcLog {


    public static void main(String args[]) throws Exception {
        run(2, 20);
        run(3, 20);
    }

    private static void run(final double z, final int numIterations) {
        CalcLog calc = new CalcLog(z);

        for (int i=0; i < numIterations; i++) {
            System.out.println("Ln(" + z +") with " + i + " iterations, using Double: " + calc.calcualteLnDouble(i));
            System.out.println("Ln(" + z +") with " + i + " iterations, using BigDec: " + calc.calcualteLnBig(i));
        }
    }


    private BigDecimal zBig;
    private double zDouble;
    private int bigDecScale = 100;


    private CalcLog(final BigDecimal z, final int bigDecScale) {
        this.zBig = z;
        this.zDouble = z.doubleValue();
        this.bigDecScale = bigDecScale;
    }

    private CalcLog(final double z) {
        this.zDouble = z;
        this.zBig = BigDecimal.valueOf(z);
    }

    /**
     * Calculate   2 * SUM from 0 to numberOfIterations (1 / (2k+1)  * ( (z-1)/(z+1) )^(2k+1))
     *
     * @param numberOfIterations The number of iterations through the equation.
     */
    private BigDecimal calcualteLnBig(final int numberOfIterations) {
        BigDecimal sum = BigDecimal.ZERO;

        for (int i=0; i < numberOfIterations; i++) {
            sum = sum.add(oneIterationBig(i));
        }
        return sum.multiply(BigDecimal.valueOf(2));

    }

    /**
     * Calculate   2 * SUM from 0 to numberOfIterations (1 / (2k+1)  * ( (z-1)/(z+1) )^(2k+1))
     *
     * @param numberOfIterations The number of iterations through the equation.
     */
    private double calcualteLnDouble(final int numberOfIterations) {
        double sum = 0.0;

        for (int i=0; i < numberOfIterations; i++) {
            sum += oneIterationDouble(i);
        }
        return 2.0 * sum;
    }


    /**
     * Calculate   1 / (2k+1)  * ( (z-1)/(z+1) )^(2k+1)
     *
     * @param k is the iteration to calculate. It is the k parameter in the equation.
     */
    private BigDecimal oneIterationBig(final int k) {
        int exp = 2 * k + 1;

        // double product = (this.zDouble - 1.0) / (this.zDouble + 1.0);
        BigDecimal zLessOne = this.zBig.subtract(BigDecimal.ONE);
        BigDecimal zPlusOne = this.zBig.add(BigDecimal.ONE);
        BigDecimal product = zLessOne.divide(zPlusOne, this.bigDecScale, RoundingMode.HALF_DOWN);

        // product = Math.pow(product, exp);
        product = product.pow(exp);

        // product = 1.0 / exp * product;
        BigDecimal oneOnExp = BigDecimal.ONE.divide(BigDecimal.valueOf(exp), this.bigDecScale, RoundingMode.HALF_DOWN);
        product = oneOnExp.multiply(product);

        return product;
    }

    /**
     * Calculate   1 / (2k+1)  * ( (z-1)/(z+1) )^(2k+1)
     *
     * @param k is the iteration to calculate. It is the k parameter in the equation.
     */
    private double oneIterationDouble(final int k) {
        int exp = 2 * k + 1;
        double product = (this.zDouble - 1.0) / (this.zDouble + 1.0);
        product = Math.pow(product, exp);
        product = 1.0 / exp * product;
        return product;
    }

}
