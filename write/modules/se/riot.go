// ===================
//
//	Search Engine
//
// ===================
package se

import (
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
)

type RiotSE struct {
	engine *riot.Engine
	option types.EngineOpts
}

func NewRiotSE() *RiotSE {
	return &RiotSE{
		engine: &riot.Engine{},
		option: types.EngineOpts{
			UseStore: true,
			IndexerOpts: &types.IndexerOpts{
				// 查询位置
				IndexType: types.LocsIndex,
			},
			StoreEngine:   "ldb",
			StoreFolder:   "./data/index",
			GseDict:       "./data/dictionary.txt",
			StopTokenFile: "./data/stop_tokens.txt",
		},
	}
}

func (r *RiotSE) Init() *RiotSE {
	r.engine.Init(r.option)
	r.engine.Flush()
	return r
}

var once sync.Once

// todo 兼容ES
func InitRiot() {
	once.Do(func() {
		Engine = NewRiotSE().Init()
	})
}

func (r *RiotSE) Insert(index string, id int64, data interface{}) error {
	r.engine.Index(strconv.FormatInt(id, 10), types.DocData{Content: data.(string)}, true)
	r.engine.Flush()
	return nil
}

func (r *RiotSE) Update(index string, id int64, data interface{}) error {
	r.Delete(index, id)
	r.Insert(index, id, data)
	return nil
}

func (r *RiotSE) Delete(index string, id int64) error {
	r.engine.RemoveDoc(strconv.FormatInt(id, 10))
	r.engine.Flush()
	return nil
}

func (r *RiotSE) Search(index, text string, offset, limit int) (result []interface{}, sum int64, err error) {
	output := r.engine.SearchDoc(types.SearchReq{
		Text: text,
		RankOpts: &types.RankOpts{
			OutputOffset: offset,
			MaxOutputs:   limit,
		},
	})
	// 看了文档，这种代码还需要自己实现的
	sum = int64(output.NumDocs)
	ids := make([]int64, 0)
	for _, doc := range output.Docs {
		id, _ := strconv.ParseInt(doc.DocId, 10, 64)
		ids = append(ids, id)
		var locs []int
		for _, v := range doc.TokenLocs {
			locs = append(locs, v...)
		}
		sort.Ints(locs)
		re := []byte(doc.Content[0:locs[0]])
		var endL, tokenL int
		for i, loc := range locs {
			endL = loc
			if i != 0 && loc <= len(doc.Content) {
				re = append(re, doc.Content[locs[i-1]+tokenL:loc]...)
			}
			for _, token := range output.Tokens {
				l := len(token)
				if loc+l <= len(doc.Content) && strings.EqualFold(string(doc.Content[loc:loc+l]), string(token)) {
					re = append(re, "<em>"+doc.Content[loc:loc+l]+"</em>"...)
					tokenL = l
					endL += l
					break
				}
			}
		}
		re = append(re, doc.Content[endL:]...)
		result = append(result, re)
	}
	return
}
