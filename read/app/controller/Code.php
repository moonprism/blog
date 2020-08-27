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

    public function index($page = 1)
    {
        $limit = 10;
        $whereSet = [['deleted_at', 'is null']];
        $codeModel = New CodeModel();
        $codeList = $codeModel->getCodeList($page, $limit, $whereSet);
        $this->assign('code_list', $codeList);
        $this->assign('next_page', count($codeList) == 10 ? $page+1 : 0);
        $this->show('Page/code');
    }

    public function preview($lang = 'md', $code = '', $title = '', $tags = '')
    {
        $lang = htmlspecialchars($lang);
        $code = htmlspecialchars(urldecode(base64_decode($code)));
        $this->assign('code_list', [[
            'id' => 0,
            'tags' => htmlspecialchars(urldecode($tags)),
            'description' => htmlspecialchars(urldecode($title)),
            'lang' => $lang,
            'content' => $code
        ]]);
        $this->assign('next_page', -1);
        $this->show('Page/code');
    }
}
