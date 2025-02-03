/**
 * monthsToPayOff
 *
 * Uses the standard amortization formula for a single debt:
 *    n = log(A / (A - P*i)) / log(1 + i)
 *
 * where:
 *   P = principal (current debt)
 *   i = monthly interest rate (APR / 100 / 12)
 *   A = monthly payment
 *
 * Returns:
 *   - The number of months (rounded up) to pay off the debt
 *   - Infinity if the payment doesn't cover interest or if result is non-finite
 */
export function monthsToPayOff(
  principal: number,
  annualInterestRate: number,
  monthlyPayment: number,
): number {
  if (principal <= 0) {
    return 0; // nothing owed
  }
  if (monthlyPayment <= 0) {
    return Infinity; // no payment => never
  }

  const i = annualInterestRate / 100 / 12; // monthly interest
  if (i === 0) {
    // No interest => simple division
    return Math.ceil(principal / monthlyPayment);
  }

  // If monthlyPayment is too small to cover monthly interest
  if (monthlyPayment <= principal * i) {
    return Infinity;
  }

  const numerator = Math.log(monthlyPayment / (monthlyPayment - principal * i));
  const denominator = Math.log(1 + i);

  let n = numerator / denominator;
  // Round up partial months
  n = Math.ceil(n);

  // if not finite => Infinity
  if (!isFinite(n)) {
    return Infinity;
  }
  return n;
}

/**
 * getMaximumMonthsToPayOffAllDebts
 *
 * For each debt, compute monthsToPayOff using the formula above,
 * then return the maximum of all (the "worst case" payoff time).
 */
export function getMaximumMonthsToPayOffAllDebts(
  debts: {
    principal: number;
    annualInterestRate: number;
    monthlyPayment: number;
  }[],
): number {
  let maxMonths = 0;
  debts.forEach((d) => {
    const months = monthsToPayOff(
      d.principal,
      d.annualInterestRate,
      d.monthlyPayment,
    );
    if (months === Infinity) {
      maxMonths = Infinity;
      return;
    }
    if (months > maxMonths) {
      maxMonths = months;
    }
  });
  return maxMonths;
}
