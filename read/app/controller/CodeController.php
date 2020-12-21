<?php

namespace app\controller;

use app\model\response\ViewResponse;
use GuzzleHttp\Client;
use kicoe\core\Request;
use kicoe\core\Response;
use kicoe\core\Config;
use app\model\Code;
use Protodata\CodeClient;
use Protodata\CodeDetail;
use Protodata\SearchRequest;

class CodeController
{
    /**
     * @route get /code/search/{text}
     * @param Config $config
     * @param string $text
     * @return array
     */
    public function search(Config $config, string $text)
    {
        $client = new CodeClient($config->get('grpc.host').':'.$config->get('grpc.port'), [
            'credentials' => \Grpc\ChannelCredentials::createInsecure(),
        ]);
        $request = new SearchRequest();
        $request->setText($text);
        list($response, $status) = $client->Search($request)->wait();
        $res = [];
        if ($status->code === 0) {
            /** @var CodeDetail $data */
            foreach ($response->getData() as $data) {
                $res[] = [
                    'id' => $data->getId(),
                    'lang' => $data->getLang(),
                    'content' => $data->getContent(),
                    'description' => $data->getTitle(),
                    'tags' => $data->getTags(),
                ];
            }
        }
        return $res;
    }

    /**
     * @route get /page/code
     * @route get /page/code/{page}
     * @param ViewResponse $response
     * @param int $page
     * @return ViewResponse
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
     * @return ViewResponse
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
     * @return Response|string
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