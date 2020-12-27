<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Protodata;

/**
 */
class CodeClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Protodata\SearchRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Search(\Protodata\SearchRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/protodata.Code/Search',
        $argument,
        ['\Protodata\SearchResponse', 'decode'],
        $metadata, $options);
    }

}
