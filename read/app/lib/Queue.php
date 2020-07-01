<?php

namespace app\lib;
use kicoe\core\Config;

class Queue {
    protected static $redisClient = null;

    public static function client() 
    {
        if (!self::$redisClient) {
            $redisConf = Config::prpr('redis');
            self::$redisClient = new \Redis();
            self::$redisClient->connect($redisConf['host'], $redisConf['port'] ?? 6379);
        }
        return self::$redisClient;
    }

    /**
     * @param string $key
     * @param array|string $data
     */
    public static function push($key, $data)
    {
        return self::client()->lPush($key, is_array($data)?json_encode($data):$data);
    }
}