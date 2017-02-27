<?php

function accumulate(array $input, callable $accumulator)
{
    if (count($input) === 0) {
        return [];
    }

    return array_map($accumulator, $input);
}
