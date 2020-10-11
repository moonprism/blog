<?php
namespace app\model;

use \kicoe\Core\Model;

/**
* tags表的模型类
*/
class Tag extends Model
{
    
    /**
     * 根据id获取对应tag信息 Map
     * @param string $ids
     * @return array
     */
    public static function getArticleTagsMap(string $ids)
    {
        $articleTagsMap = [];
        $tagList = self::query("
            select t.id, t.name, t.color, at.art_id from article_tag at inner join tag t on t.id = at.tag_id where at.art_id in ({$ids}) 
        ");

        foreach ($tagList as $tag) {
            $articleTagsMap[$tag['art_id']][] = $tag;
        }

        return $articleTagsMap;
    }

    public static function loadByID(int $id)
    {
        return (new self)->get($id);
    }
}
