package providers

type Mysql struct {}

func (s *Mysql) Close() {}
func (s *Mysql) SaveService(service service.Service) (*service.Service) {}
func (s *Mysql) LoadService() {}
func (s *Mysql) CreateSchemas() (error) {return nil}
