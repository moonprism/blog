<?php
namespace app\model;

use kicoe\core\Model;

class Article extends Model
{
    const STATUS_DRAFT = 1;
    const STATUS_PUBLISH = 2;

    /**
     * 获取文章列表
     * @param int $page 要获取的页数
     * @param int $limit 每页数量
     * @return array 该页的数据集
     * @throws \kicoe\core\Exception
     */
    public function getArticleList($page, $limit, $whereSet)
    {
        return $this->set($whereSet)
            ->limit(($page-1)*$limit, $limit)
            ->order('created_time', 'desc')
            ->select('id,title,image,summary,created_time', 'id');
    }

    /**
     * 获取文章总页数
     * @param int $limit 每页数量
     * @return int 总页数
     * @throws \kicoe\core\Exception
     */
    public function getTotalPage($limit, $whereSet)
    {
        $pages = $this->set($whereSet)->select('count(id) as total')[0]["total"];
        return ceil($pages/number_format($limit, 1));
    }

    public function getArticleListByTagId($tagId, $page, $limit, $whereSql)
    {
        $start = ($page-1)*$limit;
        // todo 框架里的 build sql 方法都应该抽出来
        return $this->query("select a.id, a.title, a.image, a.summary, a.created_time 
            from article as a 
            inner join article_tag as at 
            on a.id = at.art_id and at.tag_id = ? where 
            {$whereSql} order by created_time desc limit {$start}, {$limit}", [$tagId]);
    }

    public function getTotalPageByTagId($tagId, $limit, $whereSql)
    {
        $pages = $this->query("select count(a.id) as total from article as a inner join article_tag as at on a.id = at.art_id and at.tag_id = ?", [$tagId])[0]["total"];
        return ceil($pages/number_format($limit, 1));
    }
}