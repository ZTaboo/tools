package tool

import "time"

// Time1 几种时间格式化方式
const (
	Time1 = "20060102 15:04:05"
	Time2 = "2006-01-02 15:04:05"
	Time3 = "2006/01/02 15:04:05"
	Time4 = "20060102"
	Time5 = "2006-01-02"
	Time6 = "2006-01-02"
	Time7 = "15:04:05"
)

type FormatTime struct {
	Data string
}
type ParseTime struct {
	DataFormat string
	StrData    string
}

// FormatTimeFun 获取当前格式化的日期
func (data *FormatTime) FormatTimeFun() string {
	t := time.Now()
	return t.Format(data.Data)
}

// ParseData 根据格式化时间样式解析字符串返回 time.Time
// 解析格式：ParseTime.DataFormat 必须与 ParseTime.StrData 格式一致
func (data *ParseTime) ParseData() (time.Time, error) {
	res, err := time.Parse(data.DataFormat, data.StrData)
	if err != nil {
		return res, err
	}
	return res, err
}
