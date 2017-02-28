<?php

define("VALID_CONSTANTS", [3,5,7]);

function raindrops(int $num) : string
{

    $reply = "";

    for ($i = 1; $i <= $num; $i++) {
        if ($num % $i === 0) {
            switch ($i) {
                case VALID_CONSTANTS[0]:
                    $reply .= "Pling";
                    break;
                case VALID_CONSTANTS[1]:
                    $reply .= "Plang";
                    break;
                case VALID_CONSTANTS[2]:
                    $reply .= "Plong";
                    break;
                default:
                    continue;
            }
        }
    }

    if (strlen($reply) === 0) {
        $reply .= $num;
    }

    return $reply;
}
