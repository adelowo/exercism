const std = @import("std");

const math = std.math;

pub fn squareOfSum(number: usize) usize {
    var sumTotal: usize = 0;

    var i: usize = 0;
    while (i <= number) : (i += 1) {
        sumTotal += i;
    }

    return math.pow(usize, sumTotal, 2);
}

pub fn sumOfSquares(number: usize) usize {
    var sumTotal: usize = 0;

    var i: usize = 0;

    while (i <= number) : (i += 1) {
        sumTotal += math.pow(usize, i, 2);
    }

    return sumTotal;
}

pub fn differenceOfSquares(number: usize) usize {
    return squareOfSum(number) - sumOfSquares(number);
}
