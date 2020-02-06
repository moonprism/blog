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
    public function getCodeList($limit, $whereSet)
    {
        return $this->set($whereSet)
            ->limit(0, $limit)
            ->order('id', 'desc')
            ->select('id,lang,description,tags,content', 'id');
    }
}