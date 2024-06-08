const std = @import("std");

const ascii = std.ascii;

pub fn isPangram(str: []const u8) bool {
    var mp = std.bit_set.IntegerBitSet(26).initEmpty();

    for (str) |value| {
        const v = ascii.toLower(value);

        if (ascii.isAlphabetic(v)) {
            mp.set(v - 'a');
        }
    }

    return mp.count() == 26;
}
