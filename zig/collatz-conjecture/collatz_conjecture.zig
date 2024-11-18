const std = @import("std");

pub const ComputationError = error{
    IllegalArgument,
};

const defaultNumberOfSteps: usize = 0;

pub fn steps(n: usize) error{IllegalArgument}!usize {
    var numberOfSteps: usize = 0;

    var number = n;

    switch (number) {
        0 => return error.IllegalArgument,
        1 => return 0,
        else => {
            while (number != 1) {
                numberOfSteps += 1;

                if ((number % 2) == 0) {
                    number = number / 2;
                    continue;
                }

                number = (3 * number) + 1;
            }
        },
    }

    return numberOfSteps;
}
