package controllers

import (
	// "github.com/astaxie/bee/go/orm"
	// "github.com/astaxie/beego"
	m "hnacenter/models"
)

type ResourceController struct {
	LoginController
}
type Tree struct {
	Id         int64
	Fid        int64
	Reskey     string `json:"reskey"` // 资源Key
	Level      int64  `json:"level"`  // 资源级别
	Url        string `json:"url"`    // 资源链接
	Status     int64  `json:"status"` // 0启用 1隐藏
	Sort       int64  `json:"sort"`   // 资源排序
	Ico        string `json:"ico"`    // 图标
	Isfunction int64  // 是否是功能资源
	TwoLevel   bool
	Children   []Tree
}

//取得资源并分级
func (this *ResourceController) GetTree() []Tree {
	user := this.GetSession("user")
	adminUser := user.(m.User).Username
	var resTree []Tree
	if adminUser != "admin" {
		users := user.(m.User)
		var resourceId int64
		resources, _ := m.GetResourceByRid(users.Rid)
		if resources != nil {
			res := make([]m.Resource, len(resources))
			for k1, v1 := range resources {
				resourceId = v1["Resource"].(int64)
				resource := m.GetResourceById(resourceId)
				res[k1].Id = resource.Id
				res[k1].Fid = resource.Fid
				res[k1].Reskey = resource.Reskey
				res[k1].Ico = resource.Ico
				res[k1].Level = resource.Level
				res[k1].Url = resource.Url
				res[k1].Status = resource.Status
				res[k1].Sort = resource.Sort
				res[k1].Isfunction = resource.Isfunction
			}
			r := m.GetTreeAndLv(res, 0, 1)
			var cnt, length int = 0, 0
			for _, v := range r {
				if v.Fid == 0 {
					length = length + 1
				}
			}

			tree := make([]Tree, length)

			for k, v := range r {
				if v.Fid == 0 {
					k = cnt
					cnt = cnt + 1
					// 赋值
					tree[k].Id = v.Id
					tree[k].Fid = v.Fid
					tree[k].Reskey = v.Reskey
					tree[k].Ico = v.Ico
					tree[k].Url = v.Url
					var childCnt int = 0
					children := make([]Tree, 8)
					for _, v3 := range res {
						if v3.Fid == v.Id {
							children[childCnt].Id = v3.Id
							children[childCnt].Fid = v3.Fid
							children[childCnt].Reskey = v3.Reskey
							children[childCnt].Ico = v3.Ico
							children[childCnt].Url = v3.Url
							childCnt++
						}
					}
					// 存值
					tree[k].Children = make([]Tree, childCnt)
					var count int = 0
					for k1, v1 := range r {
						if v.Id == v1.Fid {
							k1 = count
							count = count + 1
							tree[k].Children[k1].Id = v1.Id
							tree[k].Children[k1].Fid = v1.Fid
							tree[k].Children[k1].Reskey = v1.Reskey
							tree[k].Children[k1].Ico = v1.Ico
							tree[k].Children[k1].Url = v1.Url

							var childcount int = 0
							child := make([]Tree, 4)
							for _, v4 := range res {
								if v4.Fid == v1.Id {
									child[childcount].Id = v4.Id
									child[childcount].Fid = v4.Fid
									child[childcount].Reskey = v4.Reskey
									child[childcount].Ico = v4.Ico
									child[childcount].Url = v4.Url
									childcount++
								}
							}
							tree[k].Children[k1].Children = make([]Tree, childcount)
							var ccount int = 0
							for k2, v2 := range r {
								if v1.Id == v2.Fid {
									k2 = ccount
									ccount = ccount + 1
									tree[k].Children[k1].TwoLevel = true
									tree[k].Children[k1].Children[k2].Id = v2.Id
									tree[k].Children[k1].Children[k2].Fid = v2.Fid
									tree[k].Children[k1].Children[k2].Reskey = v2.Reskey
									tree[k].Children[k1].Children[k2].Ico = v2.Ico
									tree[k].Children[k1].Children[k2].Url = v2.Url
								}
							}
						}
					}
				}
			}
			resTree = tree
		}

	} else {
		resources := m.GetTree(0)
		resource := m.GetTreeAndLv(resources, 0, 1)
		var cnt, length int = 0, 0
		for _, v := range resource {
			if v.Fid == 0 {
				length = length + 1
			}
		}
		tree := make([]Tree, length)
		for k, v := range resource {
			if v.Fid == 0 {
				k = cnt
				cnt = cnt + 1
				// 赋值
				tree[k].Id = v.Id
				tree[k].Fid = v.Fid
				tree[k].Reskey = v.Reskey
				tree[k].Ico = v.Ico
				tree[k].Url = v.Url
				var childCnt int = 0
				children := make([]Tree, 8)
				for _, v3 := range resources {
					if v3.Fid == v.Id {
						children[childCnt].Id = v3.Id
						children[childCnt].Fid = v3.Fid
						children[childCnt].Reskey = v3.Reskey
						children[childCnt].Ico = v3.Ico
						children[childCnt].Url = v3.Url
						childCnt++
					}
				}
				// 存值
				tree[k].Children = make([]Tree, childCnt)
				var count int = 0
				for k1, v1 := range resource {
					if v.Id == v1.Fid {
						k1 = count
						count = count + 1
						tree[k].Children[k1].Id = v1.Id
						tree[k].Children[k1].Fid = v1.Fid
						tree[k].Children[k1].Reskey = v1.Reskey
						tree[k].Children[k1].Ico = v1.Ico
						tree[k].Children[k1].Url = v1.Url
						var childcount int = 0
						child := make([]Tree, 4)
						for _, v4 := range resources {
							if v4.Fid == v1.Id {
								child[childcount].Id = v4.Id
								child[childcount].Fid = v4.Fid
								child[childcount].Reskey = v4.Reskey
								child[childcount].Ico = v4.Ico
								child[childcount].Url = v4.Url
								childcount++
							}
						}
						tree[k].Children[k1].Children = make([]Tree, childcount)
						var ccount int = 0
						for k2, v2 := range resource {
							if v1.Id == v2.Fid {
								k2 = ccount
								ccount = ccount + 1
								tree[k].Children[k1].TwoLevel = true
								tree[k].Children[k1].Children[k2].Id = v2.Id
								tree[k].Children[k1].Children[k2].Fid = v2.Fid
								tree[k].Children[k1].Children[k2].Reskey = v2.Reskey
								tree[k].Children[k1].Children[k2].Ico = v2.Ico
								tree[k].Children[k1].Children[k2].Url = v2.Url
							}
						}
					}
				}
			}
		}
		resTree = tree
	}

	return resTree
}
