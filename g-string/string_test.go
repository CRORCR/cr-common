package g_string

import (
	"testing"
)

func TestStringClass_DesensitizeMobile(t *testing.T) {
	result := String.DesensitizeMobile(`15212341234`)
	if result != `152****1234` {
		t.Error(`15212341234 error`)
	}

	result1 := String.DesensitizeMobile(`4042248`)
	if result1 != `40****48` {
		t.Error(`4042248 error`)
	}
}

func TestStringClass_DesensitizeEmail(t *testing.T) {
	result := String.DesensitizeEmail(`1234@qq.com`)
	if result != `1234****@qq.com` {
		t.Error(`1234@qq.com error`)
	}

	result1 := String.DesensitizeEmail(`abc@qq.com`)
	if result1 != `abc****@qq.com` {
		t.Error(`abc@qq.com error`)
	}

	result2 := String.DesensitizeEmail(`a@163.com`)
	if result2 != `a****@163.com` {
		t.Error(`a@163.com error`)
	}
}

func TestStringClass_SpanLeft(t *testing.T) {
	if String.SpanLeft(`1234`, 6, `0`) != `001234` {
		t.Error()
	}
}
