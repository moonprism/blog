<?php

namespace app\model;

use kicoe\core\Model;

class Code extends Model
{

    /**
     * 获取最新的code list
     * @param int $limit 数量
     * @return array
     * @throws \kicoe\core\Exception
     */
    public function getCodeList($page, $limit, $whereSet)
    {
        return $this->set($whereSet)
            ->limit(($page-1)*$limit, $limit)
            ->order('updated_time', 'desc')
            ->select('id,lang,description,tags,content', 'id');
    }
}