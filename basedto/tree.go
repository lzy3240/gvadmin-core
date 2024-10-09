package basedto

// 树形结构
type SysCommonTree struct {
	Id       int             `json:"id"`
	Label    string          `json:"label"`              /** 节点名称 */
	Children []SysCommonTree `json:"children,omitempty"` /** 子节点 */
}
