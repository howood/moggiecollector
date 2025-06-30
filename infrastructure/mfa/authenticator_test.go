package mfa_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/howood/moggiecollector/infrastructure/mfa"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticator_GenerateKey(t *testing.T) {
	t.Parallel()

	type args struct {
		accountID uuid.UUID
		period    uint
	}

	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "正常系: 標準期間で秘密鍵が生成できる",
			args: args{
				accountID: uuid.New(),
				period:    30,
			},
			wantErr: assert.NoError,
		},
		{
			name: "正常系: 60秒期間で秘密鍵が生成できる",
			args: args{
				accountID: uuid.New(),
				period:    60,
			},
			wantErr: assert.NoError,
		},
		{
			name: "正常系: 15秒期間で秘密鍵が生成できる",
			args: args{
				accountID: uuid.New(),
				period:    15,
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authenticator := mfa.NewAuthenticator()

			secret, err := authenticator.GenerateKey(tt.args.accountID, tt.args.period)

			if !tt.wantErr(t, err) {
				return
			}

			assert.NotEmpty(t, secret)
			assert.Len(t, secret, 32)
		})
	}
}

func TestAuthenticator_GenerateKey_Uniqueness(t *testing.T) {
	t.Parallel()

	authenticator := mfa.NewAuthenticator()
	accountID := uuid.New()
	period := uint(30)

	secret1, err1 := authenticator.GenerateKey(accountID, period)
	secret2, err2 := authenticator.GenerateKey(accountID, period)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NotEqual(t, secret1, secret2, "異なる秘密鍵が生成されるべき")
}

//nolint:funlen
func TestAuthenticator_Validate(t *testing.T) {
	t.Parallel()

	type args struct {
		passcode string
		secret   string
		period   uint
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "異常系: 空のパスコード",
			args: args{
				passcode: "",
				secret:   "JBSWY3DPEHPK3PXP",
				period:   30,
			},
			want:    false,
			wantErr: assert.Error,
		},
		{
			name: "異常系: 無効なパスコード",
			args: args{
				passcode: "000000",
				secret:   "JBSWY3DPEHPK3PXP",
				period:   30,
			},
			want:    false,
			wantErr: assert.NoError,
		},
		{
			name: "異常系: 空の秘密鍵",
			args: args{
				passcode: "123456",
				secret:   "",
				period:   30,
			},
			want:    false,
			wantErr: assert.NoError,
		},
		{
			name: "異常系: 無効な秘密鍵",
			args: args{
				passcode: "123456",
				secret:   "invalid",
				period:   30,
			},
			want:    false,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authenticator := mfa.NewAuthenticator()

			got, err := authenticator.Validate(tt.args.passcode, tt.args.secret, tt.args.period)

			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewAuthenticator(t *testing.T) {
	t.Parallel()

	authenticator := mfa.NewAuthenticator()

	assert.NotNil(t, authenticator)
}
