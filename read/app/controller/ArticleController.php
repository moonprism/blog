<?php

namespace app\controller;

use app\model\Article;
use app\model\response\ViewResponse;
use kicoe\core\DB;
use kicoe\core\Response;
use kicoe\core\Request;

class ArticleController
{
    /**
     * @route get /
     * @route get /page/{page}
     * @param Request $request
     * @param ViewResponse $response
     * @param int $page
     * @return Response
     */
    public function list(Request $request, ViewResponse $response, int $page = 1)
    {
        $page_size = 10;
        $tag = null;
        $url = '/page/%d';

        if ($tag_id = intval($request->query('tag_id'))) {
            $art = Article::listByTagId($tag_id, $page, $page_size);
            $tag = DB::table('tag')->fetchById($tag_id);
            $url .= "?tag_id={$tag_id}";
        } else {
            $art = Article::list($page, $page_size);
        }

        return $response->view('article/list', [
            'article_list' => $art->getList(),
            'count' => $art->count(),
            'page' => $page,
            'url' => $url,
            'tag' => $tag,
            'limit' => $page_size,
        ]);
    }

    /**
     * @route get /article/{id}
     * @route get /article/id/{id}
     * @param ViewResponse $response
     * @param int $id
     * @return Response
     * @throws \Exception
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
            if ($id == 1) {
                return $this->linkDetail($response);
            } else if ($id == 2) {
                return $this->aboutDetail($response);
            }
            Article::setTagsByList([$article]);
        }

        return $response->view('article/detail', compact('article'));
    }

    /**
     * @route get /link
     * @param ViewResponse $response
     * @return Response
     */
    public function linkDetail(ViewResponse $response)
    {
        $article = Article::fetchById(1);
        return $response->view('pages/link', compact('article'));
    }

    /**
     * @route get /about
     * @param ViewResponse $response
     * @return Response
     */
    public function aboutDetail(ViewResponse $response)
    {
        $article = Article::fetchById(2);
        return $response->view('pages/about', compact('article'));
    }
}
