package operations

import (
	"context"
	"errors"
	"fmt"
	api "operationProject/pkg/api/github.com/example/path/gen"
	"regexp"
	"strconv"
)

type GRPCServer struct {
}

func (s *GRPCServer) Operation(ctx context.Context, req *api.OpRequest) (*api.OpResponse, error) {
	r := regexp.MustCompile(`(-?\d+(?:\,\d+)?)\s*([-+*:\/])\s*(-?\d+(?:\,\d+)?)`)
	group := r.FindStringSubmatch(req.GetInput())

	if len(group) < 3 {
		return &api.OpResponse{}, errors.New("некорректный ввод")
	}
	var a, b int64
	var err error
	if a, err = strconv.ParseInt(group[1], 10, 64); err != nil {
		return &api.OpResponse{}, fmt.Errorf("ошибка конвертации: %w", err)
	}
	if b, err = strconv.ParseInt(group[3], 10, 64); err != nil {
		return &api.OpResponse{}, fmt.Errorf("ошибка конвертации: %w", err)
	}
	var res int64
	switch group[2] {
	case "+":
		res = Sum(a, b)
	case "-":
		res = Diff(a, b)
	case "*":
		res = Multiply(a, b)
	case ":":
		res = Split(a, b)
	case "/":
		res = Split(a, b)
	default:
		err = errors.New("оператор не определен")
	}
	if err != nil {
		return &api.OpResponse{}, err
	}
	return &api.OpResponse{Result: res}, err
}

func Sum(a, b int64) int64 {
	return a + b
}

func Diff(a, b int64) int64 {
	return a - b
}

func Multiply(a, b int64) int64 {
	return a * b
}

func Split(a, b int64) int64 {
	return a / b
}
