<?php

namespace app\controller;

use kicoe\core\Controller;
use app\model\Comment as CommentModel;
use kicoe\core\Request;
use kicoe\core\Response;
use kicoe\core\Config;
use kicoe\core\cache\Factory as CacheFactory;

class Comment extends Controller
{
    public function up(Request $request, Response $response)
    {
        if ($request->getMethod() !== 'POST') {
            $response->status('404 Not Found');
            return;
        }
        $comment = new CommentModel;
        $artId = $request->post('art_id');
        $toId = $request->post('to_id') ? $request->post('to_id') : 0;
        $name = $request->validate('name', ['len'=>9]);
        $email = $request->validate('email', ['reg'=>'/^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/']);
        $text = $request->validate('comment', ['len'=>444]);
        if ($name === false || $email === false || $text === false || $artId === false || $toId === false) {
            $response->status('404 Not Found');
        } else {
            $comment->insert(['art_id', 'to_id', 'name', 'email', 'text'],
                [[$artId, $toId, $name, $email, $text]]);
            $this->cacheComment($artId);
        }
        return;
    }

    /**
     * 取得所有评论数据
     */
    public function ls()
    {
        $request = Request::getInstance();
        $artId = $request->post('art_id');
        $comment = new CommentModel();
        $commentList = $comment->set([['art_id',$artId]])->order('id', 'desc')->select();
        for ($i=0; $i<count($commentList); $i++) {
            $commentList[$i]['email'] = ['g', md5($commentList[$i]['email']) ];
        }
        $this->json($commentList);
    }

    /**
     * 单个评论数据
     */
    public function cat(Response $response, $artId)
    {
        $cache = CacheFactory::getInstance('redis');
        $key = Config::prpr('cache_comment_prefix').$artId;
        if (!$cache->has($key)) {
            $this->cacheComment($artId);
        }
        $data = $cache->read($key);
        $response->json($data);
    }

    private function cacheComment($artId)
    {
        $limit = 7;
        $comment = new CommentModel;
        // 首先获取非回复数据
        $comList = $comment->set([['art_id', $artId]])->limit(0, $limit)->order('id', 'desc')->select();
        $idList = array(0);
        foreach ($comList as $key => $value) {
            $idList[] = $value['id'];
        }
        // 循环获取其中需要的to_id
        for ($i=0; $i<count($comList); $i++) {
            if (!in_array($comList[$i]['to_id'], $idList)) {
                $idList[] = $comList[$i]['to_id'];
                $comList[] = $comment->get($comList[$i]['to_id'])->select()[0];
            }
        }
        // 根据id倒序排序，否则js会出问题
        $mySort = function($a,$b) {
            if ($a['id'] === $b['id']) return 0;
            return (intval($a['id'])<intval($b['id']))?1:-1;
        };
        usort($comList, $mySort);
        // 邮箱加密
        for ($i=0; $i<count($comList); $i++) {
            $comList[$i]['email'] = ['g', md5($comList[$i]['email'])];
        }
        $cache = CacheFactory::getInstance('redis');
        $cache->write(Config::prpr('cache_comment_prefix').$artId, $comList);
    }
}
