package ignore

import (
	"github.com/mengri/utils/utils"
	"sync"
)

var (
	rwLock      = sync.RWMutex{}
	ignorePaths = make(map[string]map[string]utils.Set[string])
)

func IgnorePath(name, method, path string) {
	rwLock.Lock()
	defer rwLock.Unlock()
	if ignorePaths == nil {
		ignorePaths = make(map[string]map[string]utils.Set[string])
	}
	if ignorePaths[name] == nil {
		ignorePaths[name] = make(map[string]utils.Set[string])
	}
	if ignorePaths[name][method] == nil {
		ignorePaths[name][method] = utils.NewSet(path)
	} else {
		ignorePaths[name][method].Set(path)
	}

}

func IsIgnorePath(name, method, path string) bool {
	rwLock.RLock()
	defer rwLock.RUnlock()
	if isIgnorePath(name, method, path) {
		return true
	}
	if method != "*" {
		if isIgnorePath(name, "*", path) {
			return true
		}
	}
	if name != "*" {
		if isIgnorePath("*", method, path) {
			return true
		}
	}
	return false
}

func isIgnorePath(name, method, path string) bool {

	if ignorePaths == nil {
		return false
	}
	if ignorePaths[name] == nil {
		return false
	}
	if ignorePaths[name][method] == nil {
		return false
	}
	return ignorePaths[name][method].Has(path)
}
