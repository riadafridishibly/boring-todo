package api

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/k0kubun/pp/v3"
	"github.com/riadafridishibly/svelte-todo/db"
)

type contextKey struct {
	key string
}

var TodoCtxKey = &contextKey{"todoItem"}

const todoID = "todoID"

type TodoAPI struct {
	dao *db.Dao
}

func NewTodoAPI(dsn string) (*TodoAPI, error) {
	dao, err := db.NewDao(dsn)
	if err != nil {
		return nil, err
	}
	return &TodoAPI{dao: dao}, nil
}

func (api *TodoAPI) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api/todos", func(r chi.Router) {
		r.Get("/", api.ListTodos)
		r.Post("/", api.CreateTodo) // POST /todo
		r.Route("/{todoID}", func(r chi.Router) {
			r.Use(api.TodoCtx)            // Load the *Todo on the request context
			r.Get("/", api.ReadTodo)      // GET /todo/123
			r.Put("/", api.UpdateTodo)    // PUT /todo/123
			r.Delete("/", api.DeleteTodo) // DELETE /todo/123
		})
	})

	return r
}

func (api *TodoAPI) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := api.dao.ReadAll()
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	pp.Println(len(todos))

	if err := render.RenderList(w, r, NewTodoListResponse(todos)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (api *TodoAPI) TodoCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var todo *db.Todo
		var err error

		if id := chi.URLParam(r, "todoID"); id != "" {
			var idInt int64
			idInt, err = strconv.ParseInt(id, 10, 64)
			if err != nil {
				render.Render(w, r, ErrInvalidID)
				return
			}
			todo, err = api.dao.Read(idInt)
		} else {
			render.Render(w, r, ErrNotFound)
			return
		}

		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), TodoCtxKey, todo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *TodoAPI) CreateTodo(w http.ResponseWriter, r *http.Request) {
	data := TodoRequest{}
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	todo, err := api.dao.Create(ToParam(data))
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewTodoResponse(todo))
}

func (api *TodoAPI) ReadTodo(w http.ResponseWriter, r *http.Request) {
	todo := r.Context().Value(TodoCtxKey).(*db.Todo)

	if err := render.Render(w, r, NewTodoResponse(todo)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (api *TodoAPI) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoModel := r.Context().Value(TodoCtxKey).(*db.Todo)
	if todoModel == nil {
		panic("invariant! todo ctx must resolve before hitting this handler")
	}

	req := TodoRequest{}
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	params := ToParam(req)

	todo, err := api.dao.Update(todoModel.Id, params)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	render.Render(w, r, NewTodoResponse(todo))
}

func (api *TodoAPI) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var err error
	todo := r.Context().Value(TodoCtxKey).(*db.Todo)
	todo, err = api.dao.Delete(todo.Id)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewTodoResponse(todo))
}

// paginate is a stub, but very possible to implement middleware logic
// to handle the request params for handling a paginated request.
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}

type TodoRequest struct {
	Title    string  `json:"title,omitempty"`
	ParentId int64   `json:"parent_id,omitempty"`
	Content  *string `json:"content,omitempty"`
	Done     *bool   `json:"done,omitempty"`
}

func ToParam(req TodoRequest) db.TodoParams {
	params := db.TodoParams{
		Title:    req.Title,
		ParentId: req.ParentId,
	}
	if req.Content != nil {
		params.Content = sql.NullString{String: *req.Content, Valid: true}
	}
	if req.Done != nil {
		params.Done = sql.NullBool{Bool: *req.Done, Valid: true}
	}
	return params
}

func (a *TodoRequest) Bind(r *http.Request) error {
	return nil
}

type TodoResponse struct {
	*db.Todo
	Id string `json:"id"`
}

func NewTodoResponse(todo *db.Todo) *TodoResponse {
	idStr := strconv.FormatInt(todo.Id, 10)
	resp := &TodoResponse{Todo: todo, Id: idStr}

	return resp
}

func (rd *TodoResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewTodoListResponse(todos []*db.Todo) []render.Renderer {
	list := []render.Renderer{}
	for _, todo := range todos {
		list = append(list, NewTodoResponse(todo))
	}
	return list
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Todo not found"}
var ErrInvalidID = &ErrResponse{HTTPStatusCode: 400, StatusText: "Invalid ID format"}
