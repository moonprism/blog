<?php

namespace app\controller;

use app\model\Article;
use kicoe\core\Cache;
use kicoe\core\DB;
use kicoe\core\Response;

class ArticleController
{

    /**
     * @route get /
     * @route get /article/page/{page}
     * @param Response $response
     * @param int $page
     * @return Response
     */
    public function list(Response $response, int $page = 1)
    {
        $page_size = 10;
        $art = Article::list($page, $page_size);
        $article_list = $art->getList();
        Article::setTagsByList($article_list);
        return $response->view('article/list', [
            'article_list' => $article_list,
            'count' => $art->count(),
            'page' => $page,
            'url' => '/article/page/%d',
            'tag' => null,
            'limit' => $page_size,
        ]);
    }

    /**
     * @route get /article/tag/{tag_id}
     * @route get /article/tag/{tag_id}/page/{page}
     * @param Response $response
     * @param int $tag_id
     * @param int $page
     * @return Response
     */
    public function tag(Response $response, int $tag_id, int $page = 1)
    {
        $page_size = 10;
        $art = Article::listByTagId($tag_id, $page, $page_size);
        $article_list = $art->getList();
        $count = $art->count();
        Article::setTagsByList($article_list);
        $tag = DB::table('tag')->fetchById($tag_id);
        $tag->count = $count;
        return $response->view('article/list', [
            'article_list' => $article_list,
            'count' => $count,
            'page' => $page,
            'url' => "/article/tag/$tag_id/page/%d",
            // 再查一次不要紧~~
            'tag' => $tag,
            'limit' => $page_size,
        ]);
    }

    /**
     * @route get /article/id/{id}
     * @param Response $response
     * @param int $id
     * @return Response
     */
    public function detail(Response $response, int $id)
    {
        $article = Article::fetchById($id);
        if ($article->status == Article::STATUS_DRAFT) {
            $article->title = '*****';
            $article->content = ">[danger] 这篇文章还没发布呢 ~~";
        }
        Article::setTagsByList([$article]);
        return $response->view('article/detail', [
            'article' => $article
        ]);
    }

    /**
     * @route get /art/{art_id}/comments
     * @param Cache $cache
     * @param int $art_id
     * @return array
     */
    public function comments(Cache $cache, int $art_id)
    {
        return $cache->getArr('art:'.$art_id) ?? [];
    }

    /**
     * @route get /page/link
     * @param Response $response
     * @return Response
     */
    public function linkDetail(Response $response)
    {
        $article = Article::fetchById(1);
        return $response->view('pages/link', compact('article'));
    }

    /**
     * @route get /page/about
     */
    public function aboutDetail(Response $response)
    {
        $article = Article::fetchById(2);
        return $response->view('pages/about', compact('article'));
    }
}
