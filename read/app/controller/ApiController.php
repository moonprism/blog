<?php

namespace app\controller;

use app\model\Article;
use app\model\response\ApiResponse;
use app\model\Link;
use app\model\request\LinkRequest;
use kicoe\core\Response;

class ApiController
{
    /**
     * @route get api/articles
     * @param ApiResponse $response
     * @return ApiResponse
     */
    public function articleList(ApiResponse $response)
    {
        $list = Article::where('status', Article::STATUS_PUBLISH)
            ->where('deleted_at is null')
            ->orderBy('created_time', 'desc')
            ->select("id, date_format(created_time, '%Y-%m') as date, title");

        $result = [];
        foreach ($list as $i) {
            if (!isset($result[$i->date])) {
                $result[$i->date] = [];
            }
            $result[$i->date][] = [
                'id' => intval($i->id),
                'title' => $i->title
            ];
        }
        $response->data = $result;
        return $response;
    }

    /**
     * @route get api/article/{id}
     * @param ApiResponse $response
     * @param $id
     * @return Response
     */
    public function article(ApiResponse $response, $id)
    {
        $article = Article::fetchById($id);
        if ($article->status == Article::STATUS_PUBLISH) {
            $response->data = $article->content;
            return $response;
        }
        return $response->setBodyStatus(404, $id.'. article not found');
    }

    /**
     * @route post api/link
     * @param LinkRequest $request
     * @param ApiResponse $response
     * @return Response
     */
    public function link(LinkRequest $request, ApiResponse $response)
    {
        $request->check();
        $link = Link::createByOther($request);
        $link->save();
        return $response;
    }
}
