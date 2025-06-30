package model_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/howood/moggiecollector/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestBaseModel(t *testing.T) {
	t.Parallel()

	id := uuid.New()
	baseModel := model.BaseModel{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	assert.Equal(t, id, baseModel.ID)
	assert.False(t, baseModel.CreatedAt.IsZero())
	assert.False(t, baseModel.UpdatedAt.IsZero())
}
