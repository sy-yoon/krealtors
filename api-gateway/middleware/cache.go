package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var (
	targetCache = map[string]bool{
		"/v1/region/countries":  true,
		"/v1/region/provincies": true,
	}

	expires, _ = time.ParseDuration("12h")
)

//NewResponseHeader constructs a new ResponseHeader middleware handler
func NewCacheHandler(handlerToWrap http.Handler, client *redis.Client) *CacheHandler {
	return &CacheHandler{handlerToWrap, client}
}

type CacheHandler struct {
	handler http.Handler
	client  *redis.Client
}

func (me *CacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !targetCache[r.URL.String()] || r.Method != "GET" {
		me.handler.ServeHTTP(w, r)
		return
	}

	ctx := context.Background()
	val, err := me.client.Get(ctx, r.URL.String()).Bytes()
	if err != redis.Nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(val)
		return
	}

	cw := NewCacheWriter(w)
	me.handler.ServeHTTP(cw, r)
	if cw.StatusCode == http.StatusOK {
		me.client.Set(ctx, r.URL.String(), cw.Body, expires)
	}
}

func NewCacheWriter(w http.ResponseWriter) *CacheWriter {
	return &CacheWriter{w, -1, nil}
}

type CacheWriter struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func (me *CacheWriter) Header() http.Header {
	me.StatusCode = me.ResponseWriter.(gin.ResponseWriter).Status()
	return me.ResponseWriter.Header()
}

func (me *CacheWriter) WriteHeader(statusCode int) {
	me.ResponseWriter.WriteHeader(statusCode)
	me.StatusCode = statusCode
}

func (me *CacheWriter) Write(data []byte) (int, error) {
	me.Body = make([]byte, len(data))
	copy(me.Body, data)
	return me.ResponseWriter.Write(data)
}
