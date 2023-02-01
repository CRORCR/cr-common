package g_crypto

import (
	"fmt"
	"testing"
)

func TestCryptoClass_Sha256ToHex(t *testing.T) {
	if Crypto.Sha256ToHex(`12123`) != `c492e2a3e4f6cc9c5b3a1ae173333905d4cf6407f1c3b87c50763bbbbadc0dd9` {
		t.Error()
	}
}

func TestCryptoClass_ShiftCryptForInt(t *testing.T) {
	if 79456702 != Crypto.ShiftCryptForInt(79456732, 70) {
		t.Error()
	}
}

func TestCryptoClass_GeneRsaKeyPair(t *testing.T) {
	fmt.Println(Crypto.GeneRsaKeyPair())
}

func TestCryptoClass_AesCbcEncrypt(t *testing.T) {
	if `bj7P4lrG3TyB8KBpCDyGqQ==` != Crypto.AesCbcEncrypt(`1234567890123456`, `haha`) {
		t.Error()
	}
}

func TestCryptoClass_AesCbcDecrypt(t *testing.T) {
	if `haha` != Crypto.AesCbcDecrypt(`1234567890123456`, `bj7P4lrG3TyB8KBpCDyGqQ==`) {
		t.Error()
	}
}

func TestCryptoClass_HmacSha256ToHex(t *testing.T) {
	type args struct {
		str    string
		secret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `test`,
			args: args{
				str:    `625422`,
				secret: `test`,
			},
			want: `9cb8d8c168d20c0bd03782acbda3dfa504fcac3be2b80176134b89f54e376dd5`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Crypto.HmacSha256ToHex(tt.args.str, tt.args.secret); got != tt.want {
				t.Errorf("CryptoClass.HmacSha256ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
