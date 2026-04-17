package time_based_kv

type TimeMap struct{}

func NewTimeMap() TimeMap {
	return TimeMap{}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {}

func (t *TimeMap) Get(key string, timestamp int) string {
	return ""
}
