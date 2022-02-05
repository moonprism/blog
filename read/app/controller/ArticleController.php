<?php

namespace app\controller;

use app\model\Article;
use app\model\response\ViewResponse;
use kicoe\core\DB;
use kicoe\core\Response;

class ArticleController
{
    /**
     * @route get /
     * @route get /article/page/{page}
     * @param ViewResponse $response
     * @param int $page
     * @return Response
     */
    public function list(ViewResponse $response, int $page = 1)
    {
        $page_size = 10;
        $art = Article::list($page, $page_size);
        $article_list = $art->getList();

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
     * @param ViewResponse $response
     * @param int $tag_id
     * @param int $page
     * @return Response
     */
    public function tag(ViewResponse $response, int $tag_id, int $page = 1)
    {
        $page_size = 10;
        $art = Article::listByTagId($tag_id, $page, $page_size);
        $article_list = $art->getList();
        $count = $art->count();
        // 再查一次不要紧~~
        $tag = DB::table('tag')->fetchById($tag_id);
        $tag->count = $count;

        return $response->view('article/list', [
            'article_list' => $article_list,
            'count' => $count,
            'page' => $page,
            'url' => "/article/tag/$tag_id/page/%d",
            'tag' => $tag,
            'limit' => $page_size,
        ]);
    }

    /**
     * @route get /article/id/{id}
     * @param ViewResponse $response
     * @param int $id
     * @return Response
     */
    public function detail(ViewResponse $response, int $id)
    {
        $article = Article::fetchById($id);
        if (!$article) {
            throw new \Exception('文章不存在', 404);
        } else if ($article->status == Article::STATUS_DRAFT) {
            $article->title = '*****';
            $article->content = ">[danger] 这篇文章还没发布呢 ~~";
        } else {
            Article::setTagsByList([$article]);
        }

        return $response->view('article/detail', compact('article'));
    }

    /**
     * @route get /page/link
     * @param ViewResponse $response
     * @return Response
     */
    public function linkDetail(ViewResponse $response)
    {
        $article = Article::fetchById(1);
        return $response->view('pages/link', compact('article'));
    }

    /**
     * @route get /page/about
     * @param ViewResponse $response
     * @return Response
     */
    public function aboutDetail(ViewResponse $response)
    {
        $article = Article::fetchById(2);
        return $response->view('pages/about', compact('article'));
    }
}
