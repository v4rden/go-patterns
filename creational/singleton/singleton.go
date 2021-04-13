package singleton

type singleton struct {
	counter int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) Increment() int {
	s.counter++
	return s.counter
}
