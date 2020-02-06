<?php

namespace app\controller;

use app\model\Tags;
use kicoe\core\Controller;
use app\model\Article as ArticleModel;

class Article extends Controller 
{
    public function index()
    {
        $this->fetchList(1);
    }

    public function fetchList($page)
    {
        $pageSize = 10;
        $whereSet = [['status', ArticleModel::STATUS_PUBLISH], ['deleted_at', 'is null']];
        $articleModel = new ArticleModel();
        $articleList = $articleModel->getArticleList($page, $pageSize, $whereSet);
        $articleIds = [];
        $articleTagsMap = [];

        foreach ($articleList as $article) {
            $articleIds[] = $article['id'];
        }

        if ($articleIds !== []) {
            // 整合tag
            $articleTagsMap = Tags::getArticleTagsMap(implode($articleIds, ','));
        }

        foreach ($articleList as $key => $article) {
            $articleList[$key]['tags'] = $articleTagsMap[$article['id']] ?? [];
        }

        $this->assign([
            'article_list' => $articleList,
            'page' => $page,
            'total_page' => $articleModel->getTotalPage(10, $whereSet)
        ]);
        $this->show('Article/list');
    }

    public function detail($id)
    {
        $article = new ArticleModel();
        $article->get($id);
        // 整合tag
        $articleTagsMap = Tags::getArticleTagsMap($id);
        $article->tags = $articleTagsMap[$id] ?? [];
        $this->assign([
            'article_obj' => $article,
        ]);
        $this->show();
    }

    public function linkDetail()
    {
        $article = new ArticleModel();
        $article->get(1);
        $this->assign([
            'article_obj' => $article
        ]);
        $this->show('Page/link');
    }

    public function aboutDetail()
    {
        $article = new ArticleModel();
        $article->get(2);
        $this->assign([
            'article_obj' => $article
        ]);
        $this->show('Page/about');
    }

}