package registry

var FactorySingleton *Factory

// Factory 様々なインスタンスを生成する構造体
type Factory struct {
	cache map[string]interface{}
}

// ClearFactory 使いまわしているインスタンスを削除する
func ClearFactory() {
	FactorySingleton = nil
}

// GetFactory Factoryのインスタンスを取得する
func GetFactory() *Factory {
	if FactorySingleton == nil {
		FactorySingleton = &Factory{
		}
	}
	return FactorySingleton
}
