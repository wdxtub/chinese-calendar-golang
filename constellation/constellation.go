package constellation

import (
	"time"
)

// Constellation 星座
type Constellation struct {
	t *time.Time
}

// NewConstellation 创建星座对象
func NewConstellation(t *time.Time) *Constellation {
	return &Constellation{t: t}
}

// Alias 返回星座名称
func (constellation *Constellation) Alias() string {
	dates := [...]int{20, 19, 21, 20, 21, 21, 22, 23, 23, 23, 22, 22}
	// 水瓶 1月20日~2月18日
	// 双鱼 2月19日~3月20日
	// 白羊 3月21日~4月19日
	// 金牛 4月20日~5月20日
	// 双子 5月21日~6月20日
	// 巨蟹 6月21日~7月21日
	// 狮子 7月22日~8月22日
	// 处女 8月23日~9月22日
	// 天秤 9月23日~10月22日
	// 天蝎 10月23日~11月21日
	// 射手 11月22日~12月21日
	// 魔羯 12月22日~1月19日
	constellations := []rune("水瓶双鱼白羊金牛双子巨蟹狮子处女天秤天蝎射手魔羯")
	from := (constellation.t.Month() - 1) * 2
	// 12 月 22 日之后会越界，所以需要进行处理
	if constellation.t.Day() < dates[constellation.t.Month()-1] {
		from -= 2
	}

	if from < 0 {
		from += 24
	}
	return string(constellations[from:][:2])
}
