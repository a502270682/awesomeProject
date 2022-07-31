package simple_factory

import "testing"

func TestSimpleFactory(t *testing.T) {
	factory := &CacheFactory{}
	redisC, err := factory.Create(CacheTypeOfRedis)
	if err != nil {
		t.Fatal(err)
	}
	redisC.Set("k1", 1)
	rel, err := redisC.Get("k1")
	if err != nil {
		t.Fatal(err)
	}
	s1, _ := rel.(int)
	if s1 != 1 {
		t.Fatal("get wrong")
	}
	t.Log("redis success", s1)

	memC, err := factory.Create(CacheTypeOfMemcache)
	if err != nil {
		t.Fatal(err)
	}
	memC.Set("k2", 2)
	rel, err = memC.Get("k2")
	if err != nil {
		t.Fatal(err)
	}
	s2, _ := rel.(int)
	if s2 != 2 {
		t.Fatal("get wrong")
	}
	t.Log("mem success", s2)
}
