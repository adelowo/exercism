<?php

define("ADDITION", "plus");
define("SUBSTRACTION", "minus");
define("MULTIPLICATION", "multiplied");
define("DIVISION", "divided");

function calculate(string $word): int
{

    $total = 0;

    $validConstructs = preg_grep(
        "/(plus|multiplied|divided|minus)|\d+/",
        explode(" ", rtrim($word, "?"))
    );

    if (count($validConstructs) <= 2) {
        //There must be at least 2 ints with an operation to be carried at on them
        throw new InvalidArgumentException("Error Processing Verbal statement");
    }

    $total += $validConstructs[2]; //add the initial value

    foreach ($validConstructs as $key => $value) {
        switch ($value) {
            case ADDITION:
                $total += $validConstructs[$key + 1];
                break;

            case SUBSTRACTION:
                $total -= $validConstructs[$key + 1];
                break;

            //Increment offset by 2 since the verbal processing is like "multiplied by n"
            //This also applies to division

            case MULTIPLICATION:
                $total *= $validConstructs[$key + 2];
                break;

            case DIVISION:
                $total /= $validConstructs[$key + 2];
                break;

            default:
                continue; //We only care about operations

        }
    }

    return $total;
}
