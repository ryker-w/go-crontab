package schedule

import "sync"

type jobList struct {
	mu   sync.RWMutex
	data map[string]*job
}

// Set 设置数据
func (t *jobList) Set(key string, val *job) {
	t.mu.Lock()
	t.data[key] = val
	t.mu.Unlock()
}

// Get 获取数据
func (t *jobList) Get(key string) *job {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.data[key]
}

// GetAll 获取数据
func (t *jobList) GetAll() map[string]*job {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.data
}

// Del 设置数据
func (t *jobList) Del(key string) {
	t.mu.Lock()
	delete(t.data, key)
	t.mu.Unlock()
}

// Len 长度
func (t *jobList) Len() int {
	return len(t.data)
}

// Exists Key是否存在
func (t *jobList) Exists(key string) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	_, ok := t.data[key]
	return ok
}
