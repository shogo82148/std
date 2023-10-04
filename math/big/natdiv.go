// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

Multi-precision division. Here be dragons.

Given u and v, where u is n+m digits, and v is n digits (with no leading zeros),
the goal is to return quo, rem such that u = quo*v + rem, where 0 ≤ rem < v.
That is, quo = ⌊u/v⌋ where ⌊x⌋ denotes the floor (truncation to integer) of x,
and rem = u - quo·v.


Long Division

Division in a computer proceeds the same as long division in elementary school,
but computers are not as good as schoolchildren at following vague directions,
so we have to be much more precise about the actual steps and what can happen.

We work from most to least significant digit of the quotient, doing:

 • Guess a digit q, the number of v to subtract from the current
   section of u to zero out the topmost digit.
 • Check the guess by multiplying q·v and comparing it against
   the current section of u, adjusting the guess as needed.
 • Subtract q·v from the current section of u.
 • Add q to the corresponding section of the result quo.

When all digits have been processed, the final remainder is left in u
and returned as rem.

For example, here is a sketch of dividing 5 digits by 3 digits (n=3, m=2).

	                 q₂ q₁ q₀
	         _________________
	v₂ v₁ v₀ ) u₄ u₃ u₂ u₁ u₀
	           ↓  ↓  ↓  |  |
	          [u₄ u₃ u₂]|  |
	        - [  q₂·v  ]|  |
	        ----------- ↓  |
	          [  rem  | u₁]|
	        - [    q₁·v   ]|
	           ----------- ↓
	             [  rem  | u₀]
	           - [    q₀·v   ]
	              ------------
	                [  rem   ]

Instead of creating new storage for the remainders and copying digits from u
as indicated by the arrows, we use u's storage directly as both the source
and destination of the subtractions, so that the remainders overwrite
successive overlapping sections of u as the division proceeds, using a slice
of u to identify the current section. This avoids all the copying as well as
shifting of remainders.

Division of u with n+m digits by v with n digits (in base B) can in general
produce at most m+1 digits, because:

  • u < B^(n+m)               [B^(n+m) has n+m+1 digits]
  • v ≥ B^(n-1)               [B^(n-1) is the smallest n-digit number]
  • u/v < B^(n+m) / B^(n-1)   [divide bounds for u, v]
  • u/v < B^(m+1)             [simplify]

The first step is special: it takes the top n digits of u and divides them by
the n digits of v, producing the first quotient digit and an n-digit remainder.
In the example, q₂ = ⌊u₄u₃u₂ / v⌋.

The first step divides n digits by n digits to ensure that it produces only a
single digit.

Each subsequent step appends the next digit from u to the remainder and divides
those n+1 digits by the n digits of v, producing another quotient digit and a
new n-digit remainder.

Subsequent steps divide n+1 digits by n digits, an operation that in general
might produce two digits. However, as used in the algorithm, that division is
guaranteed to produce only a single digit. The dividend is of the form
rem·B + d, where rem is a remainder from the previous step and d is a single
digit, so:

 • rem ≤ v - 1                 [rem is a remainder from dividing by v]
 • rem·B ≤ v·B - B             [multiply by B]
 • d ≤ B - 1                   [d is a single digit]
 • rem·B + d ≤ v·B - 1         [add]
 • rem·B + d < v·B             [change ≤ to <]
 • (rem·B + d)/v < B           [divide by v]


Guess and Check

At each step we need to divide n+1 digits by n digits, but this is for the
implementation of division by n digits, so we can't just invoke a division
routine: we _are_ the division routine. Instead, we guess at the answer and
then check it using multiplication. If the guess is wrong, we correct it.

How can this guessing possibly be efficient? It turns out that the following
statement (let's call it the Good Guess Guarantee) is true.

If

 • q = ⌊u/v⌋ where u is n+1 digits and v is n digits,
 • q < B, and
 • the topmost digit of v = vₙ₋₁ ≥ B/2,

then q̂ = ⌊uₙuₙ₋₁ / vₙ₋₁⌋ satisfies q ≤ q̂ ≤ q+2. (Proof below.)

That is, if we know the answer has only a single digit and we guess an answer
by ignoring the bottom n-1 digits of u and v, using a 2-by-1-digit division,
then that guess is at least as large as the correct answer. It is also not
too much larger: it is off by at most two from the correct answer.

Note that in the first step of the overall division, which is an n-by-n-digit
division, the 2-by-1 guess uses an implicit uₙ = 0.

Note that using a 2-by-1-digit division here does not mean calling ourselves
recursively. Instead, we use an efficient direct hardware implementation of
that operation.

Note that because q is u/v rounded down, q·v must not exceed u: u ≥ q·v.
If a guess q̂ is too big, it will not satisfy this test. Viewed a different way,
the remainder r̂ for a given q̂ is u - q̂·v, which must be positive. If it is
negative, then the guess q̂ is too big.

This gives us a way to compute q. First compute q̂ with 2-by-1-digit division.
Then, while u < q̂·v, decrement q̂; this loop executes at most twice, because
q̂ ≤ q+2.


Scaling Inputs

The Good Guess Guarantee requires that the top digit of v (vₙ₋₁) be at least B/2.
For example in base 10, ⌊172/19⌋ = 9, but ⌊18/1⌋ = 18: the guess is wildly off
because the first digit 1 is smaller than B/2 = 5.

We can ensure that v has a large top digit by multiplying both u and v by the
right amount. Continuing the example, if we multiply both 172 and 19 by 3, we
now have ⌊516/57⌋, the leading digit of v is now ≥ 5, and sure enough
⌊51/5⌋ = 10 is much closer to the correct answer 9. It would be easier here
to multiply by 4, because that can be done with a shift. Specifically, we can
always count the number of leading zeros i in the first digit of v and then
shift both u and v left by i bits.

Having scaled u and v, the value ⌊u/v⌋ is unchanged, but the remainder will
be scaled: 172 mod 19 is 1, but 516 mod 57 is 3. We have to divide the remainder
by the scaling factor (shifting right i bits) when we finish.

Note that these shifts happen before and after the entire division algorithm,
not at each step in the per-digit iteration.

Note the effect of scaling inputs on the size of the possible quotient.
In the scaled u/v, u can gain a digit from scaling; v never does, because we
pick the scaling factor to make v's top digit larger but without overflowing.
If u and v have n+m and n digits after scaling, then:

  • u < B^(n+m)               [B^(n+m) has n+m+1 digits]
  • v ≥ B^n / 2               [vₙ₋₁ ≥ B/2, so vₙ₋₁·B^(n-1) ≥ B^n/2]
  • u/v < B^(n+m) / (B^n / 2) [divide bounds for u, v]
  • u/v < 2 B^m               [simplify]

The quotient can still have m+1 significant digits, but if so the top digit
must be a 1. This provides a different way to handle the first digit of the
result: compare the top n digits of u against v and fill in either a 0 or a 1.


Refining Guesses

Before we check whether u < q̂·v, we can adjust our guess to change it from
q̂ = ⌊uₙuₙ₋₁ / vₙ₋₁⌋ into the refined guess ⌊uₙuₙ₋₁uₙ₋₂ / vₙ₋₁vₙ₋₂⌋.
Although not mentioned above, the Good Guess Guarantee also promises that this
3-by-2-digit division guess is more precise and at most one away from the real
answer q. The improvement from the 2-by-1 to the 3-by-2 guess can also be done
without n-digit math.

If we have a guess q̂ = ⌊uₙuₙ₋₁ / vₙ₋₁⌋ and we want to see if it also equal to
⌊uₙuₙ₋₁uₙ₋₂ / vₙ₋₁vₙ₋₂⌋, we can use the same check we would for the full division:
if uₙuₙ₋₁uₙ₋₂ < q̂·vₙ₋₁vₙ₋₂, then the guess is too large and should be reduced.

Checking uₙuₙ₋₁uₙ₋₂ < q̂·vₙ₋₁vₙ₋₂ is the same as uₙuₙ₋₁uₙ₋₂ - q̂·vₙ₋₁vₙ₋₂ < 0,
and

	uₙuₙ₋₁uₙ₋₂ - q̂·vₙ₋₁vₙ₋₂ = (uₙuₙ₋₁·B + uₙ₋₂) - q̂·(vₙ₋₁·B + vₙ₋₂)
	                          [splitting off the bottom digit]
	                      = (uₙuₙ₋₁ - q̂·vₙ₋₁)·B + uₙ₋₂ - q̂·vₙ₋₂
	                          [regrouping]

The expression (uₙuₙ₋₁ - q̂·vₙ₋₁) is the remainder of uₙuₙ₋₁ / vₙ₋₁.
If the initial guess returns both q̂ and its remainder r̂, then checking
whether uₙuₙ₋₁uₙ₋₂ < q̂·vₙ₋₁vₙ₋₂ is the same as checking r̂·B + uₙ₋₂ < q̂·vₙ₋₂.

If we find that r̂·B + uₙ₋₂ < q̂·vₙ₋₂, then we can adjust the guess by
decrementing q̂ and adding vₙ₋₁ to r̂. We repeat until r̂·B + uₙ₋₂ ≥ q̂·vₙ₋₂.
(As before, this fixup is only needed at most twice.)

Now that q̂ = ⌊uₙuₙ₋₁uₙ₋₂ / vₙ₋₁vₙ₋₂⌋, as mentioned above it is at most one
away from the correct q, and we've avoided doing any n-digit math.
(If we need the new remainder, it can be computed as r̂·B + uₙ₋₂ - q̂·vₙ₋₂.)

The final check u < q̂·v and the possible fixup must be done at full precision.
For random inputs, a fixup at this step is exceedingly rare: the 3-by-2 guess
is not often wrong at all. But still we must do the check. Note that since the
3-by-2 guess is off by at most 1, it can be convenient to perform the final
u < q̂·v as part of the computation of the remainder r = u - q̂·v. If the
subtraction underflows, decremeting q̂ and adding one v back to r is enough to
arrive at the final q, r.

That's the entirety of long division: scale the inputs, and then loop over
each output position, guessing, checking, and correcting the next output digit.

For a 2n-digit number divided by an n-digit number (the worst size-n case for
division complexity), this algorithm uses n+1 iterations, each of which must do
at least the 1-by-n-digit multiplication q̂·v. That's O(n) iterations of
O(n) time each, so O(n²) time overall.


Recursive Division

For very large inputs, it is possible to improve on the O(n²) algorithm.
Let's call a group of n/2 real digits a (very) “wide digit”. We can run the
standard long division algorithm explained above over the wide digits instead of
the actual digits. This will result in many fewer steps, but the math involved in
each step is more work.

Where basic long division uses a 2-by-1-digit division to guess the initial q̂,
the new algorithm must use a 2-by-1-wide-digit division, which is of course
really an n-by-n/2-digit division. That's OK: if we implement n-digit division
in terms of n/2-digit division, the recursion will terminate when the divisor
becomes small enough to handle with standard long division or even with the
2-by-1 hardware instruction.

For example, here is a sketch of dividing 10 digits by 4, proceeding with
wide digits corresponding to two regular digits. The first step, still special,
must leave off a (regular) digit, dividing 5 by 4 and producing a 4-digit
remainder less than v. The middle steps divide 6 digits by 4, guaranteed to
produce two output digits each (one wide digit) with 4-digit remainders.
The final step must use what it has: the 4-digit remainder plus one more,
5 digits to divide by 4.

	                       q₆ q₅ q₄ q₃ q₂ q₁ q₀
	            _______________________________
	v₃ v₂ v₁ v₀ ) u₉ u₈ u₇ u₆ u₅ u₄ u₃ u₂ u₁ u₀
	              ↓  ↓  ↓  ↓  ↓  |  |  |  |  |
	             [u₉ u₈ u₇ u₆ u₅]|  |  |  |  |
	           - [    q₆q₅·v    ]|  |  |  |  |
	           ----------------- ↓  ↓  |  |  |
	                [    rem    |u₄ u₃]|  |  |
	              - [     q₄q₃·v      ]|  |  |
	              -------------------- ↓  ↓  |
	                      [    rem    |u₂ u₁]|
	                    - [     q₂q₁·v      ]|
	                    -------------------- ↓
	                            [    rem    |u₀]
	                          - [     q₀·v     ]
	                          ------------------
	                               [    rem    ]

An alternative would be to look ahead to how well n/2 divides into n+m and
adjust the first step to use fewer digits as needed, making the first step
more special to make the last step not special at all. For example, using the
same input, we could choose to use only 4 digits in the first step, leaving
a full wide digit for the last step:

	                       q₆ q₅ q₄ q₃ q₂ q₁ q₀
	            _______________________________
	v₃ v₂ v₁ v₀ ) u₉ u₈ u₇ u₆ u₅ u₄ u₃ u₂ u₁ u₀
	              ↓  ↓  ↓  ↓  |  |  |  |  |  |
	             [u₉ u₈ u₇ u₆]|  |  |  |  |  |
	           - [    q₆·v   ]|  |  |  |  |  |
	           -------------- ↓  ↓  |  |  |  |
	             [    rem    |u₅ u₄]|  |  |  |
	           - [     q₅q₄·v      ]|  |  |  |
	           -------------------- ↓  ↓  |  |
	                   [    rem    |u₃ u₂]|  |
	                 - [     q₃q₂·v      ]|  |
	                 -------------------- ↓  ↓
	                         [    rem    |u₁ u₀]
	                       - [     q₁q₀·v      ]
	                       ---------------------
	                               [    rem    ]

Today, the code in divRecursiveStep works like the first example. Perhaps in
the future we will make it work like the alternative, to avoid a special case
in the final iteration.

Either way, each step is a 3-by-2-wide-digit division approximated first by
a 2-by-1-wide-digit division, just as we did for regular digits in long division.
Because the actual answer we want is a 3-by-2-wide-digit division, instead of
multiplying q̂·v directly during the fixup, we can use the quick refinement
from long division (an n/2-by-n/2 multiply) to correct q to its actual value
and also compute the remainder (as mentioned above), and then stop after that,
never doing a full n-by-n multiply.

Instead of using an n-by-n/2-digit division to produce n/2 digits, we can add
(not discard) one more real digit, doing an (n+1)-by-(n/2+1)-digit division that
produces n/2+1 digits. That single extra digit tightens the Good Guess Guarantee
to q ≤ q̂ ≤ q+1 and lets us drop long division's special treatment of the first
digit. These benefits are discussed more after the Good Guess Guarantee proof
below.


How Fast is Recursive Division?

For a 2n-by-n-digit division, this algorithm runs a 4-by-2 long division over
wide digits, producing two wide digits plus a possible leading regular digit 1,
which can be handled without a recursive call. That is, the algorithm uses two
full iterations, each using an n-by-n/2-digit division and an n/2-by-n/2-digit
multiplication, along with a few n-digit additions and subtractions. The standard
n-by-n-digit multiplication algorithm requires O(n²) time, making the overall
algorithm require time T(n) where

	T(n) = 2T(n/2) + O(n) + O(n²)

which, by the Bentley-Haken-Saxe theorem, ends up reducing to T(n) = O(n²).
This is not an improvement over regular long division.

When the number of digits n becomes large enough, Karatsuba's algorithm for
multiplication can be used instead, which takes O(n^log₂3) = O(n^1.6) time.
(Karatsuba multiplication is implemented in func karatsuba in nat.go.)
That makes the overall recursive division algorithm take O(n^1.6) time as well,
which is an improvement, but again only for large enough numbers.

It is not critical to make sure that every recursion does only two recursive
calls. While in general the number of recursive calls can change the time
analysis, in this case doing three calls does not change the analysis:

	T(n) = 3T(n/2) + O(n) + O(n^log₂3)

ends up being T(n) = O(n^log₂3). Because the Karatsuba multiplication taking
time O(n^log₂3) is itself doing 3 half-sized recursions, doing three for the
division does not hurt the asymptotic performance. Of course, it is likely
still faster in practice to do two.


Proof of the Good Guess Guarantee

Given numbers x, y, let us break them into the quotients and remainders when
divided by some scaling factor S, with the added constraints that the quotient
x/y and the high part of y are both less than some limit T, and that the high
part of y is at least half as big as T.

	x₁ = ⌊x/S⌋        y₁ = ⌊y/S⌋
	x₀ = x mod S      y₀ = y mod S

	x  = x₁·S + x₀    0 ≤ x₀ < S    x/y < T
	y  = y₁·S + y₀    0 ≤ y₀ < S    T/2 ≤ y₁ < T

And consider the two truncated quotients:

	q = ⌊x/y⌋
	q̂ = ⌊x₁/y₁⌋

We will prove that q ≤ q̂ ≤ q+2.

The guarantee makes no real demands on the scaling factor S: it is simply the
magnitude of the digits cut from both x and y to produce x₁ and y₁.
The guarantee makes only limited demands on T: it must be large enough to hold
the quotient x/y, and y₁ must have roughly the same size.

To apply to the earlier discussion of 2-by-1 guesses in long division,
we would choose:

	S  = Bⁿ⁻¹
	T  = B
	x  = u
	x₁ = uₙuₙ₋₁
	x₀ = uₙ₋₂...u₀
	y  = v
	y₁ = vₙ₋₁
	y₀ = vₙ₋₂...u₀

These simpler variables avoid repeating those longer expressions in the proof.

Note also that, by definition, truncating division ⌊x/y⌋ satisfies

	x/y - 1 < ⌊x/y⌋ ≤ x/y.

This fact will be used a few times in the proofs.

Proof that q ≤ q̂:

	q̂·y₁ = ⌊x₁/y₁⌋·y₁                      [by definition, q̂ = ⌊x₁/y₁⌋]
	     > (x₁/y₁ - 1)·y₁                  [x₁/y₁ - 1 < ⌊x₁/y₁⌋]
	     = x₁ - y₁                         [distribute y₁]

	So q̂·y₁ > x₁ - y₁.
	Since q̂·y₁ is an integer, q̂·y₁ ≥ x₁ - y₁ + 1.

	q̂ - q = q̂ - ⌊x/y⌋                      [by definition, q = ⌊x/y⌋]
	      ≥ q̂ - x/y                        [⌊x/y⌋ < x/y]
	      = (1/y)·(q̂·y - x)                [factor out 1/y]
	      ≥ (1/y)·(q̂·y₁·S - x)             [y = y₁·S + y₀ ≥ y₁·S]
	      ≥ (1/y)·((x₁ - y₁ + 1)·S - x)    [above: q̂·y₁ ≥ x₁ - y₁ + 1]
	      = (1/y)·(x₁·S - y₁·S + S - x)    [distribute S]
	      = (1/y)·(S - x₀ - y₁·S)          [-x = -x₁·S - x₀]
	      > -y₁·S / y                      [x₀ < S, so S - x₀ < 0; drop it]
	      ≥ -1                             [y₁·S ≤ y]

	So q̂ - q > -1.
	Since q̂ - q is an integer, q̂ - q ≥ 0, or equivalently q ≤ q̂.

Proof that q̂ ≤ q+2:

	x₁/y₁ - x/y = x₁·S/y₁·S - x/y          [multiply left term by S/S]
	            ≤ x/y₁·S - x/y             [x₁S ≤ x]
	            = (x/y)·(y/y₁·S - 1)       [factor out x/y]
	            = (x/y)·((y - y₁·S)/y₁·S)  [move -1 into y/y₁·S fraction]
	            = (x/y)·(y₀/y₁·S)          [y - y₁·S = y₀]
	            = (x/y)·(1/y₁)·(y₀/S)      [factor out 1/y₁]
	            < (x/y)·(1/y₁)             [y₀ < S, so y₀/S < 1]
	            ≤ (x/y)·(2/T)              [y₁ ≥ T/2, so 1/y₁ ≤ 2/T]
	            < T·(2/T)                  [x/y < T]
	            = 2                        [T·(2/T) = 2]

	So x₁/y₁ - x/y < 2.

	q̂ - q = ⌊x₁/y₁⌋ - q                    [by definition, q̂ = ⌊x₁/y₁⌋]
	      = ⌊x₁/y₁⌋ - ⌊x/y⌋                [by definition, q = ⌊x/y⌋]
	      ≤ x₁/y₁ - ⌊x/y⌋                  [⌊x₁/y₁⌋ ≤ x₁/y₁]
	      < x₁/y₁ - (x/y - 1)              [⌊x/y⌋ > x/y - 1]
	      = (x₁/y₁ - x/y) + 1              [regrouping]
	      < 2 + 1                          [above: x₁/y₁ - x/y < 2]
	      = 3

	So q̂ - q < 3.
	Since q̂ - q is an integer, q̂ - q ≤ 2.

Note that when x/y < T/2, the bounds tighten to x₁/y₁ - x/y < 1 and therefore
q̂ - q ≤ 1.

Note also that in the general case 2n-by-n division where we don't know that
x/y < T, we do know that x/y < 2T, yielding the bound q̂ - q ≤ 4. So we could
remove the special case first step of long division as long as we allow the
first fixup loop to run up to four times. (Using a simple comparison to decide
whether the first digit is 0 or 1 is still more efficient, though.)

Finally, note that when dividing three leading base-B digits by two (scaled),
we have T = B² and x/y < B = T/B, a much tighter bound than x/y < T.
This in turn yields the much tighter bound x₁/y₁ - x/y < 2/B. This means that
⌊x₁/y₁⌋ and ⌊x/y⌋ can only differ when x/y is less than 2/B greater than an
integer. For random x and y, the chance of this is 2/B, or, for large B,
approximately zero. This means that after we produce the 3-by-2 guess in the
long division algorithm, the fixup loop essentially never runs.

In the recursive algorithm, the extra digit in (2·⌊n/2⌋+1)-by-(⌊n/2⌋+1)-digit
division has exactly the same effect: the probability of needing a fixup is the
same 2/B. Even better, we can allow the general case x/y < 2T and the fixup
probability only grows to 4/B, still essentially zero.


References

There are no great references for implementing long division; thus this comment.
Here are some notes about what to expect from the obvious references.

Knuth Volume 2 (Seminumerical Algorithms) section 4.3.1 is the usual canonical
reference for long division, but that entire series is highly compressed, never
repeating a necessary fact and leaving important insights to the exercises.
For example, no rationale whatsoever is given for the calculation that extends
q̂ from a 2-by-1 to a 3-by-2 guess, nor why it reduces the error bound.
The proof that the calculation even has the desired effect is left to exercises.
The solutions to those exercises provided at the back of the book are entirely
calculations, still with no explanation as to what is going on or how you would
arrive at the idea of doing those exact calculations. Nowhere is it mentioned
that this test extends the 2-by-1 guess into a 3-by-2 guess. The proof of the
Good Guess Guarantee is only for the 2-by-1 guess and argues by contradiction,
making it difficult to understand how modifications like adding another digit
or adjusting the quotient range affects the overall bound.

All that said, Knuth remains the canonical reference. It is dense but packed
full of information and references, and the proofs are simpler than many other
presentations. The proofs above are reworkings of Knuth's to remove the
arguments by contradiction and add explanations or steps that Knuth omitted.
But beware of errors in older printings. Take the published errata with you.

Brinch Hansen's “Multiple-length Division Revisited: a Tour of the Minefield”
starts with a blunt critique of Knuth's presentation (among others) and then
presents a more detailed and easier to follow treatment of long division,
including an implementation in Pascal. But the algorithm and implementation
work entirely in terms of 3-by-2 division, which is much less useful on modern
hardware than an algorithm using 2-by-1 division. The proofs are a bit too
focused on digit counting and seem needlessly complex, especially compared to
the ones given above.

Burnikel and Ziegler's “Fast Recursive Division” introduced the key insight of
implementing division by an n-digit divisor using recursive calls to division
by an n/2-digit divisor, relying on Karatsuba multiplication to yield a
sub-quadratic run time. However, the presentation decisions are made almost
entirely for the purpose of simplifying the run-time analysis, rather than
simplifying the presentation. Instead of a single algorithm that loops over
quotient digits, the paper presents two mutually-recursive algorithms, for
2n-by-n and 3n-by-2n. The paper also does not present any general (n+m)-by-n
algorithm.

The proofs in the paper are remarkably complex, especially considering that
the algorithm is at its core just long division on wide digits, so that the
usual long division proofs apply essentially unaltered.
*/

package big
