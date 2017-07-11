<?php

function accumulate(array $input, callable $accumulator)
{
    if (count($input) === 0) {
        return [];
    }

    $ret = [];

    foreach ($input as $value) {
	   $ret[] = call_user_func($accumulator, $value);
    }

    return $ret;
}
