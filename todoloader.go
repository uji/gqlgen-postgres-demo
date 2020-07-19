package main

import (
	"context"
	"database/sql"
	"gqlgen-postgres-demo/graph/model"
	"log"
	"net/http"
	"time"
)

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"todoLoader"}

// DataloaderMiddleware is middleware
func DataloaderMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fetch := func(keys []int) ([][]*model.Todo, []error) {
			resultSet := make([][]*model.Todo, len(keys))
			errors := make([]error, len(keys))
			var todos []*model.Todo

			rows, err := db.Query("SELECT t.text, t.done FROM todos t WHERE user_id IN %d", 2)
			if err != nil {
				log.Fatalf("%+v\n", err)
			}

			for rows.Next() {
				var t *model.Todo
				if err := rows.Scan(&t.Text, &t.Done); err != nil {
					return nil, []error{err}
				}
				todos = append(todos, t)
			}

			// NOTE: DBから取得したデータを各親項目に分配する
			// メモリ上でやるからはやいぜ
			for i, key := range keys {
				for _, todo := range todos {
					if todo.ID == string(key) {
						resultSet[i] = append(resultSet[i], todo)
					}
				}
			}

			return resultSet, errors
		}

		conf := TodoLoaderConfig{
			Wait:     250 * time.Microsecond,
			MaxBatch: 100,
			Fetch:    fetch,
		}

		// detail loader - 1:N
		todoLoader := NewTodoLoader(conf)
		ctx := context.WithValue(r.Context(), ctxKey, &todoLoader)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
