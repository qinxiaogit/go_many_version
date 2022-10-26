<?php

$uint32 = pack("nNa*",10000 ,5,92301);

var_dump($uint32);
var_dump(bin2hex($uint32));