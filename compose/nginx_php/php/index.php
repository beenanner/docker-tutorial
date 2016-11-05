<?php

$memcached = new Memcached;
$memcached->addServer('memcached', 11211);

$version = $memcached->getVersion()['memcached:11210'];
echo "Server's memcached version: ".$version."<br/>\n";
echo "<br />\n";
echo "Set \"key\" in memcached to \"Hello Memcached!\".<br/>\n";
$memcached->set('key', "Hello Memcached!", 60) or die ("Failed to save data at the server");
echo "Storing data in the cache for 60 seconds.<br/>\n";
echo "<br />\n";
$get_result = $memcached->get('key');
echo "Data from the cache:<br/>\n";
var_dump($get_result);

?>