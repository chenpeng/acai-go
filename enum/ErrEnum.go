package enum

type State int

const (
	NotLogged State = iota
	NoData
)

func (state State) String() string {
	switch state {
	case NotLogged:
		return "未登录"
	case NoData:
		return "没有查到数据"
	default:
		return "未知异常"
	}
}
