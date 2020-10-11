<?php

namespace app\controller;

use kicoe\core\Controller;
use app\model\Article as ArticleModel;

class Api extends Controller 
{
    public function articleList()
    {
        $articleModel = new ArticleModel;
        $idList = $articleModel->getIdGroupByDate();
        $result = [];
        foreach ($idList as $i) {
            if (!isset($result[$i['date']])) {
                $result[$i['date']] = [];
            }
            $result[$i['date']][] = [
                'id' => intval($i['id']),
                'title' => $i['title']
            ];
        }
        return $result;
    }

    public function article($id)
    {
        $article = new ArticleModel();
        $article->get($id);
        if ($article->status == ArticleModel::STATUS_PUBLISH) {
            return $article->content;
        }
    }
}
