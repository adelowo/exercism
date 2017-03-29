<?php

function parse_binary($binary)
{

    $splitted = str_split($binary);

    foreach ($splitted as $item => $value) {
        if ($value >= 2 || ctype_alpha($value)) {
            throw new InvalidArgumentException(
                "Only zeros and ones"
            );
        }
    }

    $decimal = 0;

    $len = count($splitted);

    foreach ($splitted as $key => $value) {
        $len--;
        $decimal += $value * pow(2, $len);
    }

    return $decimal;
}