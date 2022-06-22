<?php

namespace app\controller;

use app\model\Comment;
use app\model\Link as LinkModel;
use app\model\Article;
use app\model\request\CommentRequest;
use app\model\response\ApiResponse;
use kicoe\core\DB;
use kicoe\core\Request;
use kicoe\core\Cache;
use kicoe\core\Link;

class CommentController
{
    /**
     * @route post /comment/{art_id}
     * @param CommentRequest $request
     * @param ApiResponse $response
     * @param int $art_id
     * @return ApiResponse
     * @throws \Exception
     */
    public function post(CommentRequest $request, ApiResponse $response, int $art_id)
    {
        // 两种错误处理
        if ($err = $request->filter()) {
            return $response->setBodyStatus(422, $err);
        }
        if (!Article::fetchById($art_id)) {
            throw new \Exception("文章不存在");
        }
        /** @var Cache $cache */
        $cache = Link::make(Cache::class);
        $comment = Comment::createByOther($request);
        DB::transaction(function() use ($comment, $art_id, $cache) {
            if ($comment->to_id) {
                $to_comment = Comment::fetchById($comment->to_id);
                if (!$to_comment || $to_comment->deleted_at != null || $to_comment->art_id != $art_id) {
                    throw new \Exception("数据错误");
                }
                $comment->top_id = $to_comment->top_id ?: $to_comment->id;
            }
            $comment->art_id = $art_id;
            $comment->save();
            $cache->del('comment_preview:'.$art_id);
        });
        $cache->lPush('comment_message', $comment->id);
        return $response;
    }

    /**
     * @route get /comment/{art_id}
     * @param Request $request
     * @param ApiResponse $response
     * @param int $art_id
     * @return ApiResponse
     */
    public function list(Request $request, ApiResponse $response, int $art_id): ApiResponse
    {
        $before_id = intval($request->query('before_id', 0));
        $page_size = intval($request->query('page_size', 2));
        if ($page_size > 10) {
            $page_size = 10;
        }
        $seg = Comment::cWhere($art_id, 0)
            ->limit($page_size)
            ->orderBy('id', 'desc');

        if ($before_id > 0) {
            $seg = $seg->where('comment.id < ?', $before_id);
        }

        $comments = $seg->get();
        /** @var Comment $comment */
        foreach ($comments as &$comment) {
            $comment->email = md5($comment->email);
            $comment->sub_comments = Comment::cWhere($art_id, $comment->id)
                ->limit(20)
                ->orderBy('id', 'asc')
                ->get();
            foreach ($comment->sub_comments as &$sub) {
                $sub->email = md5($sub->email);
            }
        }
        $response->data = [
            'count' => $seg->count(),
            'comments' => $comments,
        ];
        return $response;
    }

    /**
     * @route get /subcomment/{top_id}/{art_id}
     * @param Request $request
     * @param ApiResponse $response
     * @param int $top_id
     * @param int $art_id
     * @return ApiResponse
     */
    public function subList(Request $request, ApiResponse $response, int $top_id, int $art_id)
    {
        $page_size = 5;
        $page = intval($request->query('page', 1));
        $seg = Comment::cWhere($art_id, $top_id)
            ->limit(($page-1)*$page_size, $page_size)
            ->orderBy('id', 'desc');
        $response->data = [
            'page' => $page,
            'page_size' => $page_size,
            'count' => $seg->count(),
            'comments' => $seg->get(),
        ];
        return $response;
    }
}
