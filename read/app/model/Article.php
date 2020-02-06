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
            ->order('updated_time', 'desc')
            ->select('id,title,image,summary,updated_time', 'id');
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
}