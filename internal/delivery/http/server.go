package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tanya-lyubimaya/date_difference_calculator/internal/domain"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	origins = map[string]bool{
		"http://localhost": true,
		"http://127.0.0.1": true,
	}
)

type Server struct {
	srv    *http.Server
	uc     domain.UseCase
	logger *logrus.Logger
}

func New(uc domain.UseCase, logger *logrus.Logger) (*Server, error) {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Use(PingPongMiddleware())
	server := &Server{uc: uc, srv: &http.Server{Handler: router}, logger: logger}

	router.GET("/when/:year", server.CalculateDateDifference)
	gin.LoggerWithWriter(logger.Writer())
	return server, nil
}

func PingPongMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("X-PING") == "ping" {
			c.Writer.Header().Set("X-PONG", "pong")
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Request.Referer()) != 0 {
			for v := range origins {
				if strings.Index(v, c.Request.Referer()) > -1 {
					if c.Request.Referer()[len(c.Request.Referer())-1:] == "/" {
						c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Referer()[0:len(c.Request.Referer())-1])
					} else {
						c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Referer())
					}
					break
				}
			}
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Referer")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *Server) CalculateDateDifference(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		s.logger.WithContext(c).Errorln(err)
		c.IndentedJSON(http.StatusBadRequest, "cannot parse input parameter as integer value")
		return
	}
	s.logger.WithContext(c).Tracef("input year: %d", year)
	output, err := s.uc.CalculateDateDifference(c, year)
	if err != nil && err == domain.ErrInvalidYear {
		s.logger.WithContext(c).Errorln(err)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	} else if err != nil {
		s.logger.WithContext(c).Errorln(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, output)
}

func (s *Server) Serve(port string) error {
	s.srv.Addr = port
	return s.srv.ListenAndServe()
}

func (s *Server) GracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	go func() {
		err := s.srv.Shutdown(ctx)
		if err != nil {
			s.logger.Fatalln(err)
		}
	}()
	<-ctx.Done()
	s.uc.Close()
}
