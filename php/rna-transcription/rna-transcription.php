<?php

define("DNA_TO_RNA_MAPPINGS", [
    "G" => "C",
    "C" => "G",
    "T" => "A",
    "A" => "U"
]);

function toRna(string $dna) : string
{
    $rna = '';

    foreach (str_split($dna) as $char) {
        $rna .= DNA_TO_RNA_MAPPINGS[$char];
    }

    return $rna;
}
