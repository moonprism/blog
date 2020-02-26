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

    public function tag($tagId, $page = 1)
    {
        $pageSize = 10;
        $whereSql = "a.status = ".ArticleModel::STATUS_PUBLISH." and a.deleted_at is null";
        $articleModel = new ArticleModel;
        $articleList = $articleModel->getArticleListByTagId($tagId, $page, $pageSize, $whereSql);
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
            // todo 这里要优化成自动需要重构下路由
            'url'  => "/article/tag/{$tagId}/",
            'total_page' => $articleModel->getTotalPageByTagId($tagId, $pageSize, $whereSql)
        ]);
        $this->show('Article/list');
    }

    public function fetchList($page, $where = [])
    {
        $pageSize = 10;
        $whereSet = [['status', ArticleModel::STATUS_PUBLISH], ['deleted_at', 'is null']];
        $whereSet = array_merge($whereSet, $whereSet);
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
            'url'  => '/article/page/',
            'total_page' => $articleModel->getTotalPage($pageSize, $whereSet)
        ]);
        $this->show('Article/list');
    }

    public function detail($id)
    {
        $article = new ArticleModel();
        $article->get($id);
        if ($article->status == ArticleModel::STATUS_DRAFT) {
            $article->content = ">[danger] 这篇文章还没发布呢 ~ \n （￣へ￣） \n\n *** \n *** \n ---";
        }
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