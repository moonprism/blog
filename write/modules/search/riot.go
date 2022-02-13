// ===================
//   Search Engine
// ===================
package search

// import (
// 	"sort"
// 	"strconv"
// 	"strings"
//
// 	"github.com/go-ego/riot"
// 	"github.com/go-ego/riot/types"
// )
//
// var (
// 	s = &riot.Engine{}
//
// 	opts = types.EngineOpts{
// 		UseStore: true,
// 		IndexerOpts: &types.IndexerOpts{
// 			// 查询位置
// 			IndexType: types.LocsIndex,
// 		},
// 		StoreEngine:   "ldb",
// 		StoreFolder:   "./data/index",
// 		GseDict:       "./data/dict.txt",
// 		StopTokenFile: "./data/stop_tokens.txt",
// 	}
// )
//
// func Init() {
// 	s.Init(opts)
// 	s.Flush()
// }
//
// func InsertDoc(id, content string) {
// 	s.Index(id, types.DocData{Content: content}, true)
// 	s.Flush()
// }
//
// func UpdateDoc(id, content string) {
// 	RemoveDoc(id)
// 	InsertDoc(id, content)
// }
//
// func RemoveDoc(id string) {
// 	s.RemoveDoc(id)
// 	s.Flush()
// }
//
// func SearchDoc(text string, offset, max int) (ids []int64, result map[int64][]byte, count int) {
// 	output := s.SearchDoc(types.SearchReq{
// 		Text: text,
// 		RankOpts: &types.RankOpts{
// 			OutputOffset: offset,
// 			MaxOutputs:   max,
// 		},
// 	})
// 	result = make(map[int64][]byte)
// 	// 看了文档，这种代码还需要自己实现的
// 	count = output.NumDocs
// 	for _, doc := range output.Docs {
// 		id, _ := strconv.ParseInt(doc.DocId, 10, 64)
// 		ids = append(ids, id)
// 		var locs []int
// 		for _, v := range doc.TokenLocs {
// 			locs = append(locs, v...)
// 		}
// 		// 这里不想去验证 tokens 和 locs 是否对应
// 		// 直接排序搜索了事
// 		sort.Ints(locs)
// 		re := []byte(doc.Content[0:locs[0]])
// 		var endL, tokenL int
// 		for i, loc := range locs {
// 			endL = loc
// 			if i != 0 && loc <= len(doc.Content) {
// 				re = append(re, doc.Content[locs[i-1]+tokenL:loc]...)
// 			}
// 			for _, token := range output.Tokens {
// 				l := len(token)
// 				if loc+l <= len(doc.Content) && strings.EqualFold(string(doc.Content[loc:loc+l]), string(token)) {
// 					re = append(re, "<em>"+doc.Content[loc:loc+l]+"</em>"...)
// 					tokenL = l
// 					endL += l
// 					break
// 				}
// 			}
// 		}
// 		re = append(re, doc.Content[endL:]...)
// 		result[id] = re
// 	}
// 	return
// }
