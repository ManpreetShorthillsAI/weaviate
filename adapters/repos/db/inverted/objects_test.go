package inverted

import (
	"testing"

	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnalyzeObject(t *testing.T) {
	a := NewAnalyzer()

	t.Run("with multiple properties", func(t *testing.T) {
		schema := map[string]interface{}{
			"description": "I am great!",
			"email":       "john@doe.com",
		}

		props := []*models.Property{
			&models.Property{
				Name:     "description",
				DataType: []string{"text"},
			},
			&models.Property{
				Name:     "email",
				DataType: []string{"string"},
			},
		}
		res, err := a.Object(schema, props)
		require.Nil(t, err)

		expected := []Property{
			Property{
				Key: "description",
				Items: []Countable{
					Countable{
						Data:          []byte("i"),
						TermFrequency: float32(1) / 3,
					},
					Countable{
						Data:          []byte("am"),
						TermFrequency: float32(1) / 3,
					},
					Countable{
						Data:          []byte("great"),
						TermFrequency: float32(1) / 3,
					},
				},
			},
			Property{
				Key: "email",
				Items: []Countable{
					Countable{
						Data:          []byte("john@doe.com"),
						TermFrequency: float32(1) / 1,
					},
				},
			},
		}

		assert.ElementsMatch(t, expected, res)
	})
}