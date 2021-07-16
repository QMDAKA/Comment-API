package apperr

import (
	"net/http"
)

// Code はコード.
// エラーの種別を一意に識別するために使用されるコードです.
//
// コードの表現する対象はエラーの発生場所によって変化します.
//
// 例えば
// * handler層で NotFound がエラーに付与された => APIでリクエストされたリソースが見つからない
// * infrastructure層(リポジトリ)で NotFound がエラーに付与された場合 => DBに対象のレコードが見つからない
// というように、発生場所によってそのコードの意味は異なります。
//
// そのため、各層で IsCode を使用して、エラーを判別し、
// WrapWithCode を使用してエラーコードを変換する必要があります.
type Code string

func (c Code) String() string {
	return string(c)
}

// Code の定義時にプレフィックスとして付与するようにする.
//
// エラーコードの分類のために使用される.
const (
	// 共通エラーコード
	CodeCategoryCommon = ""
)

// 下記にエラーコードを定義します.
// エラーコードはプレフィックつに分類コードを必ず付与するようにします.
// エラーコードに対する説明コメントは必須で記述してください.
// また、エラーコードに対応する CodeInfo をcodeInfoMapに必ず追加してください.
const (
	OK Code = CodeCategoryCommon + "ok"
	// AppError 以外のerrorが渡された場合のハンドリング用に使用されます.
	Unknown  Code = CodeCategoryCommon + "e999"
	NotFound Code = CodeCategoryCommon + "resource_not_found"
	// 認証エラーを表現します.
	Unauthorized Code = CodeCategoryCommon + "unauthorized"
	Forbidden    Code = CodeCategoryCommon + "forbidden"
	// 一般的なシステムエラー
	Internal   Code = CodeCategoryCommon + "e000" // アプリケーションで発生したハンドリング不可能なエラー
	Database   Code = CodeCategoryCommon + "e001" // データベースで発生したハンドリング不可能なエラー
	BadRequest Code = CodeCategoryCommon + "bad_request"
)

// CodeInfo は エラーコードに関する情報.
type CodeInfo struct {
	// httpステータス
	httpStatus int
}

// CodeInfo のコンストラクタ.
func newCodeInfo(httpStatus int) CodeInfo {
	return CodeInfo{
		httpStatus: httpStatus,
	}
}

// エラーコードに対応する情報
var codeInfoMap = map[Code]CodeInfo{
	// 共通エラーコード
	OK:           newCodeInfo(http.StatusOK),
	Unknown:      newCodeInfo(http.StatusInternalServerError),
	NotFound:     newCodeInfo(http.StatusNotFound),
	Unauthorized: newCodeInfo(http.StatusUnauthorized),
	Internal:     newCodeInfo(http.StatusInternalServerError),
	Database:     newCodeInfo(http.StatusInternalServerError),
	BadRequest:   newCodeInfo(http.StatusBadRequest),
	Forbidden:    newCodeInfo(http.StatusForbidden),
}

// エラーコードに対する情報を取得する.
func getCodeInfo(code Code) CodeInfo {
	if info, ok := codeInfoMap[code]; ok {
		return info
	}
	return codeInfoMap[Unknown]
}

// ToHTTPStatus はエラーに対応するステータスコードを返す.
func ToHTTPStatus(err error) int {
	return getCodeInfo(GetCode(err)).httpStatus
}
