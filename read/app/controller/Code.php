<?php

namespace app\controller;

use Elasticsearch\ClientBuilder;
use kicoe\core\Controller;
use kicoe\core\Response;
use kicoe\core\Config;
use app\model\Code as CodeModel;

class Code extends Controller
{
    public function search(Response $response, $text)
    {
        $elasticConf = Config::prpr('elastic');
        $client = ClientBuilder::create()->setSSLVerification(false) ->setHosts(["{$elasticConf['host']}:9200"])->build();
        $searchParams = [
            'index' => Config::prpr('elastic_code_index'),
            'type'  => Config::prpr('elastic_code_type'),
            'body' => [
                'query' => [
                    'multi_match' => [
                        'query' => $text,
                        'fields' => ['description', 'lang', 'tags'],
                    ]
                ],
                'from' => 0,
                'size' => 10,
                "_source" => ['description', 'lang', 'tags', 'content'],
                'highlight' => [
                    'fields' => [
                        'description' => new \stdClass(),
                        'lang' => new \stdClass(),
                        'tags' => new \stdClass(),
                    ]
                ]
            ]
        ];
        $res = [];
        $esResult = $client->search($searchParams);
        foreach ($esResult['hits']['hits'] as $hit) {
            $res[] = [
                    'id' => $hit['_id'],
                    'lang' => htmlspecialchars($hit['_source']['lang']),
                    'content' => htmlspecialchars($hit['_source']['content']),
                ] + $this->mergeEsHitHighlightFields($hit, ['description', 'tags']);
        }

        $response->json($res);
    }

    private function mergeEsHitHighlightFields(&$hit, $fields)
    {
        $res = [];
        foreach ($fields as $field) {
            $res[$field] = isset($hit['highlight'][$field]) ? $hit['highlight'][$field][0] : $hit['_source'][$field];
        }
        return $res;
    }

    public function index()
    {
        $aquaSearch = [
            'description' => '一个简单的代码搜索框',
            'lang' => 'css',
            'tags' => 'aqua, unable search',
            'content' => "input{
    width: 44.5%;
}
input:focus {
    border-color: aqua;
}",
        ];
        $limit = 11;
        $whereSet = [['deleted_at', 'is null']];
        $codeModel = New CodeModel();
        $codeList = $codeModel->getCodeList($limit, $whereSet);
        $codeList = ['114514' => $aquaSearch] + $codeList;
        $this->assign('code_list', $codeList);
        $this->show('Page/code');
    }
}