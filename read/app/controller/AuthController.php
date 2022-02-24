<?php

namespace app\controller;

use GuzzleHttp\Client;
use kicoe\core\Config;
use kicoe\core\Link;
use app\model\Code;
use kicoe\core\Request;
use kicoe\core\Response;
use app\model\Article;
use app\model\response\ViewResponse;

/**
 * Class AuthController
 * @package app\controller
 */
class AuthController
{

    public function __construct()
    {
        session_start();
    }

    public function isLogin()
    {
        return isset($_SESSION['user']);
    }

    /**
     * cas auth
     * 本地请求 | -  php前台服务                                  golang后台服务
     *               | - 通过上面的函数未登录则使用下面的login跳转后台系统
     *                                                             | - 登录后台拿到key,再次跳转前台到当前函数
     *               | - 当前函数通过guzzle传递拿到的key去后台系统验证
     *                                                             | - 后台验证key通过，返回success
     *               | - 登录成功
     *  跳转 redirect
     * @route get /cas/auth
     * @param Request $request
     * @param Response $response
     * @param Config $config
     * @return Response|string
     * @throws \GuzzleHttp\Exception\GuzzleException
     */
    public function auth(Request $request, Response $response, Config $config)
    {
        $redirect = base64_decode($request->query('redirect'));
        $key = $request->query('key');
        $client = new Client(['timeout' => 3]);
        $res = $client->get($config->get('cas.auth_url'), [
            'query' => ['key' => $key]
        ]);
        if ((string)$res->getBody() === 'success') {
            $_SESSION['user'] = $key;
            return $response->redirect('http://'.urldecode($redirect));
        }
    }

    /**
     * @param $url
     * @return void
     */
    public function login($url)
    {
        $config = Link::make(Config::class);
        header("Location: ".$config->get('cas.login_url').urlencode($url));
        die();
    }

    /**
     * @route get /preview
     * @route get /preview/{type}
     * @param Request $request
     * @param ViewResponse $response
     * @param string $type
     * @return ViewResponse
     */
    public function preview(Request $request, ViewResponse $response, $type = 'code')
    {
        if (!$this->isLogin()) {
            $this->login($request->url());
        }
        if ($type == 'code') {
            $code = new Code();
            $code->lang = $request->query('lang');
            $code->description = urldecode($request->query('description'));
            $code->tags = urldecode($request->query('tags'));
            $code->content = urldecode(base64_decode($request->query('content')));
            return $response->view('pages/code', [
                'code_list' => [$code],
                'next_page' => -1,
            ]);
        } else if ($type == 'article') {
            $article = Article::fetchById($request->query('id'));
            return $response->view('article/detail', compact('article'));
        }
    }
}
