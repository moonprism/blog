<?php

namespace app\controller;

use app\model\Comment;
use kicoe\core\Cache;
use kicoe\core\Link;
use kicoe\core\Request;

class CommentController
{
    /**
     * @route post /comment/up
     * @param Request $request
     */
    public function up(Request $request)
    {
        $comment = new Comment();
        $comment->art_id = $request->input('art_id');
        $comment->to_id = $request->input('to_id');
        $comment->name = htmlspecialchars($request->input('name'));
        $comment->email = htmlspecialchars($request->input('email'));
        $comment->text = htmlspecialchars($request->input('comment'));
        $comment->save();
        if ($comment->to_id) {
            /** @var Cache $cache */
            $cache = Link::make(Cache::class);
            $cache->lPush('comment_message', $comment->id);
        }
    }

    /**
     * @route get /comment/list/{art_id}
     * @param $art_id
     * @return array
     */
    public function list(int $art_id)
    {
        $comments = Comment::where('art_id', $art_id)->orderBy('id', 'desc')->get();
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
                if ($reply_comment = Comment::fetchById($comment->to_id)) {
                    $comments[] = $reply_comment;
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
