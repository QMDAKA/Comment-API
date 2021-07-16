package apperr

import (
	"fmt"
	"strings"

	"golang.org/x/xerrors"
)

type options struct {
	clientMessages   []string
	internalMessages []string
}

// Option は AppError 生成用のオプション.
type Option interface {
	// このインターフェースを実装したstructの値を渡された options に設定する.
	apply(*options)
}

// AppError にクライアント用メッセージを追加するオプション
// 内部エラーにも追加される. そのため、 OptionClientMessages を指定した場合は、
// OptionInternalMessages は指定不要となる.
type OptionClientMessages []string

func OptCltMsg(messages ...string) OptionClientMessages {
	return messages
}

func (o OptionClientMessages) apply(options *options) {
	options.clientMessages = append(options.clientMessages, o...)
	options.internalMessages = append(options.internalMessages, o...)
}

// AppError に内部にのみ共有したいメッセージを追加するオプション.
type OptionInternalMessages []string

func OptIntMsg(messages ...string) OptionInternalMessages {
	return messages
}

func (o OptionInternalMessages) apply(options *options) {
	options.internalMessages = append(options.internalMessages, o...)
}

// ================================================================================= //

type AppError struct {
	// xerrorsでstacktraceを実現するための項目
	next  error         // wrapされたエラー
	frame xerrors.Frame // xerrorsでstacktraceを実現するためのframe情報

	// 基本エラー情報
	code            Code     // エラーを一意に識別するコード
	internalMessage []string // 内部表示用のメッセージ
	clientMessage   []string // クライアント側に表示可能なメッセージ
}

// Code はエラー発生原因場所を表すコードを返す
func (e *AppError) Code() Code {
	return e.code
}

// ClientMessage はクライアント表示可能なエラーメッセージを返す.
func (e *AppError) ClientMessage() string {
	if len(e.clientMessage) > 0 {
		return strings.Join(e.clientMessage, "\n")
	}
	return ""
}

// InternalMessage は内部用のエラーメッセージを返す.
func (e *AppError) InternalMessage() string {
	if len(e.internalMessage) > 0 {
		return strings.Join(e.internalMessage, "\n")
	}
	return ""
}

// Error .
func (e *AppError) Error() string {
	return fmt.Sprintf("code: %s,\nmessage: %s", e.Code(), e.InternalMessage())
}

// Unwrap .
func (e *AppError) Unwrap() error {
	return e.next
}

// fmt.Formatter を実装
func (e *AppError) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

// xerrors.Formatter を実装
func (e *AppError) FormatError(p xerrors.Printer) error {
	p.Print(e.Error())
	e.frame.Format(p)
	return e.Unwrap()
}

// 共通生成処理
func create(code Code, opts ...Option) *AppError {
	aErr := &AppError{
		code:  code,
		frame: xerrors.Caller(2),
	}

	// オプション値を設定する
	var in options
	for i := range opts {
		opts[i].apply(&in)
	}
	aErr.clientMessage = in.clientMessages
	aErr.internalMessage = in.internalMessages

	return aErr
}

// As .
func As(err error, target interface{}) bool {
	return xerrors.As(err, target)
}

// AsAppError はerrorから AppError への変換を行う.
func AsAppError(err error) (aErr *AppError, ok bool) {
	return aErr, As(err, &aErr)
}

// FIXME New が使用されなくなったらNewを削除後、New_=>Newにリネームする.
func New_(code Code, opts ...Option) error {
	aErr := create(code, opts...)
	return aErr
}

// Wrap_ は Code をerrのエラーコードを引き継いでエラーをラップする.
//
// AppError ではなければ Internal とする.
// FIXME Wrap が使用されなくなったらWrapを削除後、Wrap_=>Wrapにリネームする.
func Wrap_(err error, opts ...Option) error {
	if code := GetCode(err); code != Unknown {
		return WrapWithCode(code, err, opts...)
	}
	return WrapWithCode(Internal, err, opts...)
}

// WrapWithCode は Code を書き換えてエラーをラップする.
func WrapWithCode(code Code, err error, opts ...Option) error {
	aErr := create(code, opts...)
	aErr.next = err
	return aErr
}

// GetCode はcodeを取得する.
func GetCode(err error) Code {
	if err == nil {
		return OK
	}

	if aErr, ok := AsAppError(err); ok {
		return aErr.Code()
	}

	return Unknown
}

// IsCode は AppError に付与されているコードが code と一致しているかを返します.
func IsCode(err error, code Code) bool {
	return GetCode(err) == code
}

// ========================================
// 廃止予定のメソッド

// Wrap は外部エラー(アプリケーション発生ではなく、外部/内部パッケージで発生したエラー/外部APIエラーなど)
// にコードを付与して、 AppError に変換するための関数です.
// この際、詳細コードは外部エラーでハンドリングできない、という意図で DetailCodeInternal が指定されます.
// Deprecated: 廃止予定
func Wrap(code Code, err error, opts ...Option) error {
	aErr := create(code, opts...)
	aErr.next = err
	return aErr
}

// New はAppErrorをコードと詳細コードから生成する
// Deprecated: 廃止予定
func New(code Code, detailCode string, opts ...Option) error {
	// 共通エラー生成処理
	aErr := create(code, opts...)
	// 一時的な措置detailコードが返却されないのでメッセージに追記しておく
	aErr.clientMessage = append(aErr.clientMessage, detailCode)
	return aErr
}

// NewWithMsg は AppError を追加メッセージ付きで生成します.
// 内部、クライアント用の両方にメッセージが追加されます.
// Deprecated: 廃止予定, NewメソッドにOptionを渡すようにしてください.
func NewWithMsg(code Code, detailCode string, msg ...string) error {
	aErr := create(code)
	aErr.clientMessage = msg
	// 一時的な措置detailコードが返却されないのでメッセージに追記しておく
	aErr.clientMessage = append(aErr.clientMessage, detailCode)
	aErr.internalMessage = msg
	return aErr
}
