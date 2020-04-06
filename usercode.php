<?php

require 'vendor/autoload.php';

use ASK\XmlBuilder\XmlBuilder;

function handle($data) {
    $xmlBuilder = XmlBuilder::create();
    $xmlBuilder->element('root')->end();

    echo $xmlBuilder->getXml(true);

    $data['foo'] = 123;

    return $data;
}
