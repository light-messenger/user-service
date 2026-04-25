package mocks

//go:generate mockgen -destination=handler/type.go -package=handlerMock github.com/light-messenger/user-service/internal/handler Service
//go:generate mockgen -destination=service/type.go -package=serviceMock github.com/light-messenger/user-service/internal/service Repository
