package middlewares

type loggingMiddleware struct {
}

func NewloggingMiddleware() *loggingMiddleware {
	return &loggingMiddleware{}
}

func (lm *loggingMiddleware) LogRequest() {

}

func (lm *loggingMiddleware) LogResponse() {

}

func (lm *loggingMiddleware) LogError() {

}

func (lm *loggingMiddleware) LogInfo() {

}

func (lm *loggingMiddleware) LogWarning() {

}

func (lm *loggingMiddleware) LogDebug() {

}

func (lm *loggingMiddleware) LogFatal() {

}
