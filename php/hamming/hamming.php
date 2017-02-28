<?php

function distance(string $a, string $b)
{

    if (mb_strlen($a, 'UTF-8') !== mb_strlen($b, 'UTF-8')) {
        throw new InvalidArgumentException(
            'DNA strands must be of equal length.'
        );
    }

    $hammingDistance = 0;

    for ($i = 0; $i < mb_strlen($a); $i++) {
        if (strcmp($a[$i],$b[$i]) !== 0) {
            $hammingDistance++;
        }
    }

    return $hammingDistance;
}
