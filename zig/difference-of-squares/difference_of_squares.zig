const std = @import("std");

const math = std.math;

pub fn squareOfSum(number: usize) usize {
    var sumTotal: usize = 0;

    for (0..number + 1) |i| {
        sumTotal += i;
    }

    return math.pow(usize, sumTotal, 2);
}

pub fn sumOfSquares(number: usize) usize {
    var sumTotal: usize = 0;

    for (0..number + 1) |i| {
        sumTotal += math.pow(usize, i, 2);
    }

    return sumTotal;
}

pub fn differenceOfSquares(number: usize) usize {
    return squareOfSum(number) - sumOfSquares(number);
}
