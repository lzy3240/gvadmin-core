package basedto

// 时间参数
// 暂未使用
type TimeParams struct {
	BeginTime string `json:"beginTime" form:"beginTime" search:"type:gte"`
	EndTime   string `json:"endTime" form:"beginTime" search:"type:lte"`
}
