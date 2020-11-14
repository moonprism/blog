<?php

namespace app\controller;

use app\model\Article;
use kicoe\core\Response;

class ApiController
{
    /**
     * @route get api/articles
     * @return array
     */
    public function articleList()
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
        return $result;
    }

    /**
     * @route get api/article/{id}
     * @param Response $response
     * @param $id
     * @return Response
     */
    public function article(Response $response, $id)
    {
        $article = Article::fetchById($id);
        if ($article->status == Article::STATUS_PUBLISH) {
            return $response->text($article->content);
        }
        return $response;
    }
}
