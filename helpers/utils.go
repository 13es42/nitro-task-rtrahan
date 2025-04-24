package util

import (
	"math/big"
	"sync"
)

// FactorialCache provides a thread-safe cache for factorial results
type FactorialCache struct {
	cache map[int]*big.Int
	mutex sync.RWMutex
}

// NewFactorialCache creates a new factorial cache
func NewFactorialCache() *FactorialCache {
	return &FactorialCache{
		cache: make(map[int]*big.Int),
	}
}

// Get retrieves a factorial result from the cache
func (fc *FactorialCache) Get(n int) (*big.Int, bool) {
	fc.mutex.RLock()
	defer fc.mutex.RUnlock()

	result, exists := fc.cache[n]
	return result, exists
}

// Set stores a factorial result in the cache
func (fc *FactorialCache) Set(n int, result *big.Int) {
	fc.mutex.Lock()
	defer fc.mutex.Unlock()

	fc.cache[n] = result
}

// CalculateFactorial computes the factorial of n,
// using the cache if available
func (fc *FactorialCache) CalculateFactorial(n int) *big.Int {
	// Check if result is in cache
	if result, found := fc.Get(n); found {
		return result
	}

	// Base cases
	if n <= 1 {
		result := big.NewInt(1)
		fc.Set(n, result)
		return result
	}

	// Recursive calculation
	prev := fc.CalculateFactorial(n - 1)
	result := new(big.Int).Mul(big.NewInt(int64(n)), prev)

	// Cache the result
	fc.Set(n, result)
	return result
}
