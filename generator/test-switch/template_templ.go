// Code generated by templ DO NOT EDIT.

package testswitch

import "github.com/a-h/templ"
import "context"
import "io"

func render(input string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		ctx, _ = templ.RenderedCSSClassesFromContext(ctx)
		ctx, _ = templ.RenderedScriptsFromContext(ctx)
		switch input {
		case "a":
			_, err = io.WriteString(w, templ.EscapeString("it was 'a'"))
			if err != nil {
				return err
			}
		default:			_, err = io.WriteString(w, templ.EscapeString("it was something else"))
			if err != nil {
				return err
			}
		}
		return err
	})
}

