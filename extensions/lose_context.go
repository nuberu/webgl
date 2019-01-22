package extensions

import "syscall/js"

const LoseContextExtensionName Name = "WEBGL_lose_context"

type LoseContext struct {
	Extension
}

func LoadLoseContextExtension(glContext js.Value) *LoseContext {
	return &LoseContext{
		Extension: Extension{
			js: glContext.Call("getExtension", string(LoseContextExtensionName)),
		},
	}
}

func (lc *LoseContext) LoseContext() {
	lc.js.Call("loseContext")
}
