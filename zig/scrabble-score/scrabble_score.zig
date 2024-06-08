const std = @import("std");
const ascii = std.ascii;

pub fn score(s: []const u8) u32 {
    var scrabble_score: u32 = 0;

    for (s) |value| {
        switch (ascii.toLower(value)) {
            'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't' => scrabble_score += 1,
            'd', 'g' => scrabble_score += 2,
            'b', 'c', 'm', 'p' => scrabble_score += 3,
            'f', 'h', 'v', 'w', 'y' => scrabble_score += 4,
            'k' => scrabble_score += 5,
            'j', 'x' => scrabble_score += 8,
            'q', 'z' => scrabble_score += 10,
            else => {
                // do nothing
            },
        }
    }

    return scrabble_score;
}
