package server

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"google.golang.org/grpc"
	"log"
)

type RateLimiter struct {
	max         int64
	methods     []string
	sem         *semaphore.Weighted
	grpcPackage string
}

func NewRateLimiter(methods []string, maxRequests int64) *RateLimiter {
	return &RateLimiter{
		max:         maxRequests,
		methods:     methods,
		sem:         semaphore.NewWeighted(maxRequests),
		grpcPackage: "fileserver.v1.FileStore",
	}

}
func (u *RateLimiter) inMethods(mtd string) bool {
	for _, m := range u.methods {
		if mtd == fmt.Sprintf("/%s/%s", u.grpcPackage, m) {
			return true
		}
	}
	return false
}

func (u *RateLimiter) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		if u.inMethods(info.FullMethod) {
			err := u.sem.Acquire(ctx, 1)
			if err != nil {
				return nil, err
			}
			log.Println(info.FullMethod)
			resp, err := handler(ctx, req)
			u.sem.Release(1)
			return resp, err
		}
		return handler(ctx, req)
	}
}

func (u *RateLimiter) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		if u.inMethods(info.FullMethod) {
			err := u.sem.Acquire(ss.Context(), 1)
			if err != nil {
				return err
			}
			log.Println(info.FullMethod)
			err = handler(srv, ss)
			u.sem.Release(1)
			return err
		}
		return handler(srv, ss)
	}
}
