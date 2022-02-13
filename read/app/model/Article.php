<?php

namespace app\model;

use kicoe\core\Model;
use kicoe\core\DB;

class Article extends Model
{
    public int $id;
    public string $title;
    public int $status;
    public string $image;
    public string $summary;
    public string $content;
    public string $updated_time;
    public string $created_time;
    // public ?string $deleted_at;

    protected array $tags;

    const STATUS_DRAFT = 1;
    const STATUS_PUBLISH = 2;

    public function getList()
    {
        return self::setTagsByList($this->removeColumns('content')->get());
    }

    /**
     * @param int $tag_id
     * @param int $page
     * @param int $limit
     * @return Model|self
     */
    public static function listByTagId(int $tag_id, int $page, int $limit)
    {
        return (self::list($page, $limit))
            ->join('article_tag at', 'article.id = at.art_id and at.tag_id = ?', $tag_id);
    }

    /**
     * @param int $page
     * @param int $limit
     * @return Model|self
     */
    public static function list(int $page, int $limit)
    {
        return (new self())
            ->where('status = ?', self::STATUS_PUBLISH)
            ->where('deleted_at is null')
            ->orderBy('id', 'desc')
            ->limit(($page-1)*$limit, $limit);
    }

    public function setTags(...$tags)
    {
        $this->tags = $tags;
    }

    public function getTags()
    {
        return $this->tags ?? [];
    }

    /**
     * 和 Model 无关的设置全体tags方法
     * @param array $art_list
     * @return array
     */
    public static function setTagsByList(array $art_list)
    {
        // php7
        $ids = array_column($art_list, 'id');
        if (!$ids) {
            return $art_list;
        }

        // 获取集合中的所有文章id
        $tag_list = DB::select('
            select 
            at.art_id, at.tag_id id, t.color, t.name 
            from tag t 
            inner join article_tag at 
            on t.id = at.tag_id and at.art_id in (?)
        ', $ids);

        // 多对多关联的处理
        $tag_map = [];
        foreach ($tag_list as $tag) {
            if (isset($tag_map[$tag->art_id])) {
                $tag_map[$tag->art_id][] = $tag;
            } else {
                $tag_map[$tag->art_id] = [$tag];
            }
            unset($tag->art_id);
        }

        /** @var Article $art */
        foreach ($art_list as $art) {
            if ($tags = $tag_map[$art->id] ?? false) {
                $art->setTags(...$tags);
            }
        }

        return $art_list;
    }
}
