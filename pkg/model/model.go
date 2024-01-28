package model

type Blog struct {
	Title           string
	Content         string
	Author          string
	PublicationDate string
	Tags            string
}
type AutoIncrementMap struct {
	Data    map[int64]Blog
	counter int64
}

func NewAutoIncrementMap() *AutoIncrementMap {
	return &AutoIncrementMap{
		Data:    make(map[int64]Blog),
		counter: 0,
	}
}

func (m *AutoIncrementMap) Add(value1 string, value2 string, value3 string, value4 string, value5 string) int64 {
	obj := Blog{
		Title:           value1,
		Content:         value2,
		Author:          value3,
		PublicationDate: value4,
		Tags:            value5,
	}
	m.counter++
	m.Data[m.counter] = obj
	return m.counter
}
func (m *AutoIncrementMap) Update(key int64, value1 string, value2 string, value3 string, value4 string) {
	if obj, exists := m.Data[key]; exists {
		obj.Title = value1
		obj.Content = value2
		obj.Author = value3
		obj.Tags = value4
		m.Data[key] = obj
	}
}

func (m *AutoIncrementMap) Delete(key int64) {
	delete(m.Data, key)
}

func (m *AutoIncrementMap) GetValueForKey(key int64) (Blog, bool) {
	obj, exists := m.Data[key]
	return obj, exists
}
