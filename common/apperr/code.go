package apperr

import (
	"net/http"

	"github.com/team-lab/enoteca-ec/internal/common/logger"
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
	CodeCategoryCommon = "common_"
	// 受注関連のエラーコード
	CodeCategoryOrder = "order_"
	// 配送/出荷関連のエラーコード
	CodeCategoryShipment = "shipment_"
)

// 下記にエラーコードを定義します.
// エラーコードはプレフィックつに分類コードを必ず付与するようにします.
// エラーコードに対する説明コメントは必須で記述してください.
// また、エラーコードに対応する CodeInfo をcodeInfoMapに必ず追加してください.
const (
	// ============================================================================== //
	// 特殊ケースのハンドリング用コード
	// ============================================================================== //
	// 正常系を表すコードです.
	// エラー生成時には使用されず、渡されたerrがnilの場合の特殊ケース処理の際に使用されます.
	OK Code = CodeCategoryCommon + "ok"
	// AppError 以外のerrorが渡された場合のハンドリング用に使用されます.
	Unknown Code = CodeCategoryCommon + "e999"

	// ============================================================================== //
	// 共通エラーコード
	// ============================================================================== //
	// 取得しようとしたデータが存在しないことを表現します. (DBのレコードが存在しないなど)
	NotFound Code = CodeCategoryCommon + "resource_not_found"
	// 認証エラーを表現します.
	Unauthorized Code = CodeCategoryCommon + "unauthorized"

	// 一般的なシステムエラー
	Internal      Code = CodeCategoryCommon + "e000" // アプリケーションで発生したハンドリング不可能なエラー
	Database      Code = CodeCategoryCommon + "e001" // データベースで発生したハンドリング不可能なエラー
	ExternalAPI   Code = CodeCategoryCommon + "e002" // 外部APIで発生したハンドリング不可能なエラー
	AWS           Code = CodeCategoryCommon + "e003" // AWSで発生したハンドリング不可能なエラー¬
	Inconsistency Code = CodeCategoryCommon + "e004" // データ不整合で発生したハンドリング不可能なエラー

	// 一般的なリクエスト不正系のエラー
	NotFoundPaymentMethod Code = CodeCategoryCommon + "not_found_payment_method" // リクエストされた支払方法が見つからない

	// ============================================================================== //
	// 受注関連のエラーコード
	// ============================================================================== //
	// 受注生成時にカートに存在しない商品が指定されている
	OrderSpecifiedNotExistsItemInCart Code = CodeCategoryOrder + "specified_not_exists_item_in_cart"
	// 受注生成時にカート明細と受注明細が一致しない
	OrderCartItemAndOrderItemNotMatch Code = CodeCategoryOrder + "cart_and_order_item_not_match"
	// 受注に対して有効でない支払方法が選択された
	OrderInvalidPaymentMethod Code = CodeCategoryOrder + "invalid_payment_method"

	// ============================================================================== //
	// 配送関連のエラーコード
	// ============================================================================== //
	// 配送に指定可能な配送方法が存在しない場合のエラー
	ShipmentNotExistsSelectableShippingMethod Code = CodeCategoryShipment + "not_exists_selectable_shipping_method"
	// 配送に指定可能な配送日が存在しない場合のエラー
	ShipmentNotExistsSelectableShippingDate Code = CodeCategoryShipment + "not_exists_selectable_shipping_date"
	// 配送希望日が不正な場合のエラー
	ShipmentInvalidDesiredDeliveryDate Code = CodeCategoryShipment + "invalid_desired_delivery_date"
	// 配送希望日時が不正な場合のエラー
	ShipmentInvalidDesiredDeliveryTime Code = CodeCategoryShipment + "invalid_desired_delivery_time"
	// 配送に対して出荷可能な日が存在しない
	ShipmentNotExistsShippableDate Code = CodeCategoryShipment + "not_exists_shippable_date"

	// ============================================================================== //
	// 廃止予定
	// ============================================================================== //
	// 渡された引数や、リクエスト内容が不正な場合のエラーを表現します.
	// Deprecated: 廃止予定
	BadRequest Code = CodeCategoryCommon + "bad_request"
)

// CodeInfo は エラーコードに関する情報.
type CodeInfo struct {
	// httpステータス
	httpStatus int
	// ログレベル
	logLevel logger.LogLevel
}

// CodeInfo のコンストラクタ.
func newCodeInfo(httpStatus int, logLevel logger.LogLevel) CodeInfo {
	return CodeInfo{
		httpStatus: httpStatus,
		logLevel:   logLevel,
	}
}

// エラーコードに対応する情報
var codeInfoMap = map[Code]CodeInfo{
	// 共通エラーコード
	OK:                    newCodeInfo(http.StatusOK, logger.LevelDebug),
	Unknown:               newCodeInfo(http.StatusInternalServerError, logger.LevelError),
	NotFound:              newCodeInfo(http.StatusNotFound, logger.LevelDebug),
	Unauthorized:          newCodeInfo(http.StatusUnauthorized, logger.LevelWarn),
	Internal:              newCodeInfo(http.StatusInternalServerError, logger.LevelError),
	Database:              newCodeInfo(http.StatusInternalServerError, logger.LevelError),
	ExternalAPI:           newCodeInfo(http.StatusInternalServerError, logger.LevelError),
	AWS:                   newCodeInfo(http.StatusInternalServerError, logger.LevelError),
	Inconsistency:         newCodeInfo(http.StatusInternalServerError, logger.LevelError),
	NotFoundPaymentMethod: newCodeInfo(http.StatusBadRequest, logger.LevelInfo),
	// 受注関連エラーコード
	OrderSpecifiedNotExistsItemInCart: newCodeInfo(http.StatusBadRequest, logger.LevelInfo),
	OrderCartItemAndOrderItemNotMatch: newCodeInfo(http.StatusBadRequest, logger.LevelInfo),
	OrderInvalidPaymentMethod:         newCodeInfo(http.StatusBadRequest, logger.LevelInfo),
	// 配送関連エラーコード
	ShipmentNotExistsSelectableShippingMethod: newCodeInfo(http.StatusBadRequest, logger.LevelInfo),
	ShipmentNotExistsSelectableShippingDate:   newCodeInfo(http.StatusBadRequest, logger.LevelInfo),
	ShipmentInvalidDesiredDeliveryDate:        newCodeInfo(http.StatusBadRequest, logger.LevelInfo),
	ShipmentInvalidDesiredDeliveryTime:        newCodeInfo(http.StatusBadRequest, logger.LevelInfo),
	ShipmentNotExistsShippableDate:            newCodeInfo(http.StatusBadRequest, logger.LevelInfo),

	// 廃止予定のエラーコード
	BadRequest: newCodeInfo(http.StatusBadRequest, logger.LevelInfo),
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

// ToLogLevel はエラーに対応するログレベルを返す.
func ToLogLevel(err error) logger.LogLevel {
	return getCodeInfo(GetCode(err)).logLevel
}
