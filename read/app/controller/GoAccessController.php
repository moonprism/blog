<?php

namespace app\controller;

use app\model\response\ViewResponse;
use kicoe\core\Request;

class GoAccessController extends AuthController
{
    /**
     * @route get /goaccess/report
     * @param Request $request
     * @param ViewResponse $response
     * @return ViewResponse
     */
    public function report(Request $request, ViewResponse $response)
    {
        if (!$this->isLogin()) {
            return $this->login($request, $response);
        }
        $response->setSuffix('.html');
        return $response->view('goaccess/index');
    }
}
