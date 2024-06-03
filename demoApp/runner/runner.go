package runner

type MultiService []Service

type Service interface {
	Run()
}

func (s *MultiService) Start() {
	for _, service := range *s {
		go service.Run()
	}
}
