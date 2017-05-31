package main

import (
	"github.com/radicalmind/xeon"
	"github.com/radicalmind/xeon/context"
)

type Tag struct {
	Name string
}

type Todo struct {
	Content string `json:"content"`
	Tags    []Tag  `json:"tags"`
}

func main() {
	app := xeon.New()

	app.OnErrorCode(xeon.StatusInternalServerError, func(ctx context.Context) {
		// .Values are used to communicate between handlers, middleware.
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}

		ctx.Writef("(Unexpected) internal server error")
	})

	app.Get("/", func(ctx context.Context) {
		todo1 := Todo{
			Content: "todo1",
			Tags:    []Tag{Tag{Name: "グルメ"}},
		}
		todo2 := Todo{
			Content: "todo2",
			Tags:    []Tag{Tag{Name: "天気がいい日"}, Tag{Name: "ちょっと遠出"}},
		}

		todos := []Todo{todo1, todo2}

		ctx.JSON(todos)
	})

	app.Post("/todo", func(ctx context.Context) {
		var todo Todo
		ctx.ReadJSON(&todo)

		ctx.Writef("todo: %s , tags: %s", todo.Content, todo.Tags)
	})

	if err := app.Run(xeon.Addr(":3000")); err != nil {
		panic(err)
	}
}
