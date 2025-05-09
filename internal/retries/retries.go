package retries

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"net"
	"os"
	"syscall"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

var retriablePGCodes = []string{
	pgerrcode.ConnectionException,
	pgerrcode.ConnectionDoesNotExist,
	pgerrcode.SQLServerRejectedEstablishmentOfSQLConnection,
	pgerrcode.ConnectionFailure,
	pgerrcode.SQLClientUnableToEstablishSQLConnection,
	pgerrcode.TransactionResolutionUnknown,
}

var retriableErrors = []error{syscall.EAGAIN, syscall.EWOULDBLOCK, os.ErrDeadlineExceeded, context.DeadlineExceeded, io.ErrUnexpectedEOF}

func WithRetries(fn func() error) (err error) {
	for count, wait := range []time.Duration{1 * time.Second, 3 * time.Second, 5 * time.Second} {
		if err = fn(); err != nil {
			log.Err(err).Msgf("Retriable error - %d try", count)
			if !isRetriableError(err) {
				return
			}
			time.Sleep(wait)
		} else {
			err = nil
			return
		}
	}
	return
}

func isRetriableError(err error) bool {
	if err == nil {
		return false
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {

		for _, retriablePGCode := range retriablePGCodes {
			if pgErr.Code == retriablePGCode {
				return true
			}
		}
	}

	if errors.Is(err, sql.ErrConnDone) {
		return true
	}

	var netErr net.Error
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			return true
		}
	}

	for _, error := range retriableErrors {
		if errors.Is(err, error) {
			return true
		}
	}

	return false
}
