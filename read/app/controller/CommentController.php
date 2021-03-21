<?php

namespace app\controller;

use app\model\Comment;
use app\model\request\CommentRequest;
use app\model\response\ApiResponse;
use kicoe\core\Cache;
use kicoe\core\Link;

class CommentController
{
    /**
     * @route post /comment/up
     * @param CommentRequest $request
     * @param ApiResponse $response
     * @return ApiResponse
     */
    public function up(CommentRequest $request, ApiResponse $response)
    {
        if ($err = $request->filter()) {
            return $response->setBodyStatus(422, 'ValidationError: '.$err);
        }
        $comment = Comment::createByOther($request);
        $comment->save();
        $response->data = $comment;
        // 邮件推送队列
        if ($comment->to_id) {
            /** @var Cache $cache */
            $cache = Link::make(Cache::class);
            $cache->lPush('comment_message', $comment->id);
        }
        return $response;
    }

    /**
     * 下面历史代码就不用新的 api response 了
     * @route get /comment/list/{art_id}
     * @param $art_id
     * @return array
     */
    public function list(int $art_id)
    {
        $comments = Comment::where('art_id', $art_id)->orderBy('id', 'desc')->get();

        // 邮箱加密
        /** @var Comment $comment */
        foreach ($comments as &$comment) {
            $comment->email = md5($comment->email);
        }
        return $comments;
    }

    /**
     * @route get /json/comment/{art_id}
     * @param $art_id
     * @return array
     */
    public function preview(int $art_id)
    {
        $limit = 7;
        $comments = Comment::where('art_id', $art_id)
            ->limit(0, $limit)
            ->orderBy('id', 'desc')
            ->get();
        $id_list = array_column($comments, 'id');

        // 循环获取其中需要的to_id
        /** @var Comment $comment */
        foreach ($comments as $comment) {
            if (!in_array($comment->to_id, $id_list)) {
                $id_list[] = $comment->to_id;
                $to_id = $comment->to_id;
                // 啊这...
                hoho:
                if ($reply_comment = Comment::fetchById($to_id)) {
                    $comments[] = $reply_comment;
                    if ($to_id = $reply_comment->to_id) {
                        if (array_search($to_id, array_column($comments, 'id')) === false) {
                            goto hoho;
                        }
                    }
                }
            }
        }

        // 根据id倒序排序，否则js会出问题
        $reverse_id_sort = function($a, $b) {
            if ($a->id === $b->id) return 0;
            return intval($a->id) < intval($b->id) ? 1 : -1;
        };
        usort($comments, $reverse_id_sort);

        // 邮箱加密
        foreach ($comments as &$comment) {
            $comment->email = md5($comment->email);
        }

        return $comments;
    }
}
