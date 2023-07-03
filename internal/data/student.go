package data

import (
	"context"
	"student/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type studentRepo struct {
	data *Data
	log  *log.Helper
}

func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (sr *studentRepo) GetStudent(ctx context.Context, id int32) (*biz.Student, error) {
	return nil, nil
}
