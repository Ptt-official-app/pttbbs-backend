package mand

import (
	"github.com/Ptt-official-app/pttbbs-backend/metrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Init(isLocked bool) (err error) {
	if !isLocked {
		cliLock.Lock()
		defer cliLock.Unlock()
	}

	if Cli != nil {
		return ErrCliAlreadyInit
	}

	if IsTest {
		mockConn := NewMockClientConn()
		Cli = NewManServiceClient(mockConn)
	} else {
		conn, err = grpc.NewClient(
			GRPC_HOST,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(metrics.UnaryClientInterceptor()),
		)
		if err != nil {
			return err
		}

		Cli = NewManServiceClient(conn)
	}

	return nil
}

func Finalize(isLocked bool) {
	if !isLocked {
		cliLock.Lock()
		defer cliLock.Unlock()
	}

	defer func() {
		if conn != nil {
			conn.Close()
			conn = nil
		}
	}()

	defer func() {
		Cli = nil
	}()
}

func Reset() (err error) {
	cliLock.Lock()
	defer cliLock.Unlock()

	Finalize(true)

	return Init(true)
}
