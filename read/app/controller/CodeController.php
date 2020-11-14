<?php

namespace app\controller;

use app\model\response\ViewResponse;
use Elasticsearch\ClientBuilder;
use kicoe\core\Request;
use kicoe\core\Response;
use kicoe\core\Config;
use app\model\Code;

class CodeController
{
    /**
     * @route get /code/search/{text}
     * @param Config $config
     * @param Response $response
     * @param string $text
     * @return Response
     */
    public function search(Config $config, Response $response, string $text)
    {
        $client = ClientBuilder::create()
            ->setSSLVerification(false)
            ->setHosts(["{$config->get('es.host')}:9200"])
            ->build();

        $options = [
            'index' => $config->get('es.code_index'),
            'type'  => $config->get('es.code_type'),
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
        foreach ($client->search($options)['hits']['hits'] as $hit) {
            $res[] = [
                    'id' => $hit['_id'],
                    'lang' => htmlspecialchars($hit['_source']['lang']),
                    'content' => htmlspecialchars($hit['_source']['content']),
                ] + $this->mergeEsHitHighlightFields($hit, ['description', 'tags']);
        }
        return $response->json($res);
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
     * @param ViewResponse $response
     * @param int $page
     * @return Response
     */
    public function index(ViewResponse $response, int $page = 1)
    {
        $limit = 10;
        $code_list = Code::where('deleted_at is null')
            ->orderBy('updated_time', 'desc')
            ->limit(($page-1)*$limit, $limit)
            ->get();
        $next_page = count($code_list) == 10 ? $page+1 : 0;
        return $response->view('pages/code', compact('code_list', 'next_page'));
    }

    /**
     * @route get /code/preview
     * @route get /code/preview/{lang}
     * @param Response $response
     * @param Request $request
     * @param Config $config
     * @param string $lang
     * @return Response
     */
    public function preview(Response $response, Request $request, Config $config, $lang = 'md')
    {
        // todo cas client
        if (!isset($_SESSION['user'])) {
            $query = [
                'callback' => urlencode($request->url())
            ];
            var_dump($config->get('cas.login_url').'?'.http_build_query($query));
            //return $response->redirect($config->get('cas.url'));
        }
        return '';
        $code = new Code();
        $code->lang = htmlspecialchars($lang);
        $code->description = htmlspecialchars(urldecode($request->query('description')));
        $code->tags = htmlspecialchars(urldecode($request->query('tags')));
        $code->content = htmlspecialchars(urldecode(base64_decode($request->query('content'))));
        $code_list = [$code];
        $next_page = -1;
        return $response->view('pages/code', compact('code_list', 'next_page'));
    }
}
