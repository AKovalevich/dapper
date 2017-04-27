package providers

type MongoDB struct {}

func (s *MongoDB) Close() {}
func (s *MongoDB) SaveService(service service.Service) (*service.Service) {}
func (s *MongoDB) LoadService() {}
func (s *MongoDB) CreateSchemas() (error) {return nil}
