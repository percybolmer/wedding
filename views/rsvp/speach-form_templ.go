// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.871
package rsvp

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func SpeachForm() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<form hx-post=\"/speach\" hx-swap=\"innerHTML\" hx-target=\"#speach-form\" hx-ext=\"json-enc\" class=\"space-y-4\"><input type=\"text\" required minlength=\"2\" name=\"name\" placeholder=\"Namn på talare\" class=\"input input-bordered w-full\"><div class=\"form-control\"><input type=\"text\" required pattern=\"^[0-9+\\-\\s]{7,}$\" name=\"phone\" placeholder=\"Telefonnummer\" class=\"input input-bordered w-full  peer invalid:border-red-500 focus:invalid:border-red-500\"> <span class=\"mt-1 text-sm text-red-500 invisible peer-invalid:visible\">Ange ett giltigt telefonnummer</span></div><div class=\"form-control\"><input type=\"email\" name=\"email\" placeholder=\"E-post\" required class=\"input input-bordered w-full peer invalid:border-red-500 focus:invalid:border-red-500\"> <span class=\"mt-1 text-sm text-red-500 invisible peer-invalid:visible\">Ange en giltig e-postadress</span></div><textarea name=\"message\" placeholder=\"Meddelande\" class=\"textarea textarea-bordered w-full h-32\"></textarea> <button type=\"submit\" class=\"btn btn-outline btn-primary\">Skicka Anmälan om tal</button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
