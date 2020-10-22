<?php

namespace app\controller;

use Elasticsearch\ClientBuilder;
use kicoe\core\Controller;
use kicoe\core\Response;
use kicoe\core\Config;
use app\model\Code;

class CodeController
{
    /**
     * @param Response $response
     * @param $text
     */
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

    /**
     * @route get /page/code
     * @route get /page/code/{page}
     * @param Response $response
     * @param int $page
     * @return Response
     */
    public function index(Response $response, $page = 1)
    {
        $limit = 10;
        $code_list = Code::where('deleted_at is null')
            ->orderBy('updated_time', 'desc')
            ->limit(($page-1)*$limit, $limit)
            ->get();
        $next_page = count($code_list) == 10 ? $page+1 : 0;
        return $response->view('pages/code', compact('code_list', 'next_page'));
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
