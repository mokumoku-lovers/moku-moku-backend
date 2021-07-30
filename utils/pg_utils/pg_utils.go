package pg_utils

import (
	errors2 "errors"
	"fmt"
	"github.com/jackc/pgconn"
	"moku-moku/utils/errors"
	"strconv"
	"strings"
)

const (
	errorEmailExists    = "users_email_key"
	errorUsernameExists = "users_username_key"
)

func ParseError(err error, message string) *errors.RestErr {
	// Check if the err type is of Pgconn.PgError
	var pgErr *pgconn.PgError
	if errors2.As(err, &pgErr) {
		codeSQL, _ := strconv.ParseInt(pgErr.Code, 10, 64)
		switch codeSQL {
		case 23505:
			if strings.Contains(pgErr.Message, errorEmailExists) {
				return errors.UniqueConstraintViolation("email already exists")
			}
			if strings.Contains(pgErr.Message, errorUsernameExists) {
				return errors.UniqueConstraintViolation("username already exists")
			}
		}
	}

	return errors.InternalServerError(fmt.Sprintf("%s: %s", message, err.Error()))
}
