<?php

namespace app\controller;

use GuzzleHttp\Client;
use kicoe\core\Config;
use kicoe\core\Link;
use kicoe\core\Request;
use kicoe\core\Response;

/**
 * Class AuthController
 * CAS 服务基础类, 使用示例：
 * if (!$this->isLogin()) return $this->login($request, $response)
 * @package app\controller
 */
class AuthController
{
    // 正常页面不会用到session
    private bool $is_session_start = false;

    public function sessionStart()
    {
        if (!$this->is_session_start) {
            session_start();
            $this->is_session_start = true;
        }
    }

    public function isLogin()
    {
        $this->sessionStart();
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
        $redirect = $request->query('redirect');
        $key = $request->query('key');
        $client = new Client();
        $res = $client->get($config->get('cas.auth_url'), [
            'query' => ['key' => $key]
        ]);
        if ((string)$res->getBody() === 'success') {
            $this->sessionStart();
            $_SESSION['user'] = $key;
            return $response->redirect('http://'.$redirect);
        } else {
            return 'login failed';
        }
    }

    /**
     * @param Request $request
     * @param Response $response
     * @return Response
     */
    public function login(Request $request, Response $response)
    {
        $config = Link::make(Config::class);
        return $response->redirect($config->get('cas.login_url').urlencode($request->url()));
    }
}