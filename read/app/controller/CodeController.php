<?php

namespace app\controller;

use app\model\response\ViewResponse;
use Elasticsearch\ClientBuilder;
use GuzzleHttp\Client;
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
     * @param ViewResponse $response
     * @param Request $request
     * @param Config $config
     * @param string $lang
     * @return Response
     */
    public function preview(ViewResponse $response, Request $request, Config $config, $lang = 'md')
    {
        session_start();
        if (!isset($_SESSION['user'])) {
            return $response->redirect($config->get('cas.login_url').urlencode($request->url()));
        }
        $code = new Code();
        $code->lang = htmlspecialchars($lang);
        $code->description = htmlspecialchars(urldecode($request->query('description')));
        $code->tags = htmlspecialchars(urldecode($request->query('tags')));
        $code->content = htmlspecialchars(urldecode(base64_decode($request->query('content'))));
        $code_list = [$code];
        $next_page = -1;
        return $response->view('pages/code', compact('code_list', 'next_page'));
    }

    /**
     * cas login
     * 本地请求 | -  php前台服务                                  golang后台服务
     *               | - 通过上面的函数未登录则跳转后台系统
     *                                                             | - 登录后台拿到key,再次跳转前台到当前函数
     *               | - 当前函数通过guzzle传递拿到的key去后台系统验证
     *                                                             | - 后台验证key通过，返回success
     *               | - 登录成功
     *  跳转 redirect
     * @route get /login
     * @param Request $request
     * @param Response $response
     * @param Config $config
     * @return Response
     * @throws \GuzzleHttp\Exception\GuzzleException
     */
    public function login(Request $request, Response $response, Config $config)
    {
        $redirect = $request->query('redirect');
        $key = $request->query('key');
        $client = new Client();
        $res = $client->get($config->get('cas.auth_url'), [
            'query' => ['key' => $key]
        ]);
        if ((string)$res->getBody() === 'success') {
            session_start();
            $_SESSION['user'] = $key;
        }
        return $response->redirect('http://'.$redirect);
    }
}