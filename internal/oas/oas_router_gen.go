// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func skipSlash(p []byte) []byte {
	if len(p) > 0 && p[0] == '/' {
		return p[1:]
	}
	return p
}

// nextElem return next path element from p and forwarded p.
func nextElem(p []byte) (elem, next []byte) {
	p = skipSlash(p)
	idx := bytes.IndexByte(p, '/')
	if idx < 0 {
		idx = len(p)
	}
	return p[:idx], p[idx:]
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := []byte(r.URL.Path)
	if len(p) == 0 {
		s.notFound(w, r)
		return
	}

	var (
		elem []byte            // current element, without slashes
		args map[string]string // lazily initialized
	)

	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "POST":
		// Root edge.
		elem, p = nextElem(p)
		switch string(elem) {
		case "addStickerToSet": // -> 1
			// POST /addStickerToSet
			s.handleAddStickerToSetRequest(args, w, r)
			return
		case "answerCallbackQuery": // -> 2
			// POST /answerCallbackQuery
			s.handleAnswerCallbackQueryRequest(args, w, r)
			return
		case "answerInlineQuery": // -> 3
			// POST /answerInlineQuery
			s.handleAnswerInlineQueryRequest(args, w, r)
			return
		case "answerPreCheckoutQuery": // -> 4
			// POST /answerPreCheckoutQuery
			s.handleAnswerPreCheckoutQueryRequest(args, w, r)
			return
		case "answerShippingQuery": // -> 5
			// POST /answerShippingQuery
			s.handleAnswerShippingQueryRequest(args, w, r)
			return
		case "approveChatJoinRequest": // -> 6
			// POST /approveChatJoinRequest
			s.handleApproveChatJoinRequestRequest(args, w, r)
			return
		case "banChatMember": // -> 7
			// POST /banChatMember
			s.handleBanChatMemberRequest(args, w, r)
			return
		case "banChatSenderChat": // -> 8
			// POST /banChatSenderChat
			s.handleBanChatSenderChatRequest(args, w, r)
			return
		case "close": // -> 9
			// POST /close
			s.handleCloseRequest(args, w, r)
			return
		case "copyMessage": // -> 10
			// POST /copyMessage
			s.handleCopyMessageRequest(args, w, r)
			return
		case "createChatInviteLink": // -> 11
			// POST /createChatInviteLink
			s.handleCreateChatInviteLinkRequest(args, w, r)
			return
		case "createNewStickerSet": // -> 12
			// POST /createNewStickerSet
			s.handleCreateNewStickerSetRequest(args, w, r)
			return
		case "declineChatJoinRequest": // -> 13
			// POST /declineChatJoinRequest
			s.handleDeclineChatJoinRequestRequest(args, w, r)
			return
		case "deleteChatPhoto": // -> 14
			// POST /deleteChatPhoto
			s.handleDeleteChatPhotoRequest(args, w, r)
			return
		case "deleteChatStickerSet": // -> 15
			// POST /deleteChatStickerSet
			s.handleDeleteChatStickerSetRequest(args, w, r)
			return
		case "deleteMessage": // -> 16
			// POST /deleteMessage
			s.handleDeleteMessageRequest(args, w, r)
			return
		case "deleteMyCommands": // -> 17
			// POST /deleteMyCommands
			s.handleDeleteMyCommandsRequest(args, w, r)
			return
		case "deleteStickerFromSet": // -> 18
			// POST /deleteStickerFromSet
			s.handleDeleteStickerFromSetRequest(args, w, r)
			return
		case "deleteWebhook": // -> 19
			// POST /deleteWebhook
			s.handleDeleteWebhookRequest(args, w, r)
			return
		case "editChatInviteLink": // -> 20
			// POST /editChatInviteLink
			s.handleEditChatInviteLinkRequest(args, w, r)
			return
		case "editMessageCaption": // -> 21
			// POST /editMessageCaption
			s.handleEditMessageCaptionRequest(args, w, r)
			return
		case "editMessageLiveLocation": // -> 22
			// POST /editMessageLiveLocation
			s.handleEditMessageLiveLocationRequest(args, w, r)
			return
		case "editMessageMedia": // -> 23
			// POST /editMessageMedia
			s.handleEditMessageMediaRequest(args, w, r)
			return
		case "editMessageReplyMarkup": // -> 24
			// POST /editMessageReplyMarkup
			s.handleEditMessageReplyMarkupRequest(args, w, r)
			return
		case "editMessageText": // -> 25
			// POST /editMessageText
			s.handleEditMessageTextRequest(args, w, r)
			return
		case "exportChatInviteLink": // -> 26
			// POST /exportChatInviteLink
			s.handleExportChatInviteLinkRequest(args, w, r)
			return
		case "forwardMessage": // -> 27
			// POST /forwardMessage
			s.handleForwardMessageRequest(args, w, r)
			return
		case "getChat": // -> 28
			// POST /getChat
			s.handleGetChatRequest(args, w, r)
			return
		case "getChatAdministrators": // -> 29
			// POST /getChatAdministrators
			s.handleGetChatAdministratorsRequest(args, w, r)
			return
		case "getChatMember": // -> 30
			// POST /getChatMember
			s.handleGetChatMemberRequest(args, w, r)
			return
		case "getChatMemberCount": // -> 31
			// POST /getChatMemberCount
			s.handleGetChatMemberCountRequest(args, w, r)
			return
		case "getFile": // -> 32
			// POST /getFile
			s.handleGetFileRequest(args, w, r)
			return
		case "getGameHighScores": // -> 33
			// POST /getGameHighScores
			s.handleGetGameHighScoresRequest(args, w, r)
			return
		case "getMe": // -> 34
			// POST /getMe
			s.handleGetMeRequest(args, w, r)
			return
		case "getMyCommands": // -> 35
			// POST /getMyCommands
			s.handleGetMyCommandsRequest(args, w, r)
			return
		case "getStickerSet": // -> 36
			// POST /getStickerSet
			s.handleGetStickerSetRequest(args, w, r)
			return
		case "getUpdates": // -> 37
			// POST /getUpdates
			s.handleGetUpdatesRequest(args, w, r)
			return
		case "getUserProfilePhotos": // -> 38
			// POST /getUserProfilePhotos
			s.handleGetUserProfilePhotosRequest(args, w, r)
			return
		case "getWebhookInfo": // -> 39
			// POST /getWebhookInfo
			s.handleGetWebhookInfoRequest(args, w, r)
			return
		case "leaveChat": // -> 40
			// POST /leaveChat
			s.handleLeaveChatRequest(args, w, r)
			return
		case "logOut": // -> 41
			// POST /logOut
			s.handleLogOutRequest(args, w, r)
			return
		case "pinChatMessage": // -> 42
			// POST /pinChatMessage
			s.handlePinChatMessageRequest(args, w, r)
			return
		case "promoteChatMember": // -> 43
			// POST /promoteChatMember
			s.handlePromoteChatMemberRequest(args, w, r)
			return
		case "restrictChatMember": // -> 44
			// POST /restrictChatMember
			s.handleRestrictChatMemberRequest(args, w, r)
			return
		case "revokeChatInviteLink": // -> 45
			// POST /revokeChatInviteLink
			s.handleRevokeChatInviteLinkRequest(args, w, r)
			return
		case "sendAnimation": // -> 46
			// POST /sendAnimation
			s.handleSendAnimationRequest(args, w, r)
			return
		case "sendAudio": // -> 47
			// POST /sendAudio
			s.handleSendAudioRequest(args, w, r)
			return
		case "sendChatAction": // -> 48
			// POST /sendChatAction
			s.handleSendChatActionRequest(args, w, r)
			return
		case "sendContact": // -> 49
			// POST /sendContact
			s.handleSendContactRequest(args, w, r)
			return
		case "sendDice": // -> 50
			// POST /sendDice
			s.handleSendDiceRequest(args, w, r)
			return
		case "sendDocument": // -> 51
			// POST /sendDocument
			s.handleSendDocumentRequest(args, w, r)
			return
		case "sendGame": // -> 52
			// POST /sendGame
			s.handleSendGameRequest(args, w, r)
			return
		case "sendInvoice": // -> 53
			// POST /sendInvoice
			s.handleSendInvoiceRequest(args, w, r)
			return
		case "sendLocation": // -> 54
			// POST /sendLocation
			s.handleSendLocationRequest(args, w, r)
			return
		case "sendMediaGroup": // -> 55
			// POST /sendMediaGroup
			s.handleSendMediaGroupRequest(args, w, r)
			return
		case "sendMessage": // -> 56
			// POST /sendMessage
			s.handleSendMessageRequest(args, w, r)
			return
		case "sendPhoto": // -> 57
			// POST /sendPhoto
			s.handleSendPhotoRequest(args, w, r)
			return
		case "sendPoll": // -> 58
			// POST /sendPoll
			s.handleSendPollRequest(args, w, r)
			return
		case "sendSticker": // -> 59
			// POST /sendSticker
			s.handleSendStickerRequest(args, w, r)
			return
		case "sendVenue": // -> 60
			// POST /sendVenue
			s.handleSendVenueRequest(args, w, r)
			return
		case "sendVideo": // -> 61
			// POST /sendVideo
			s.handleSendVideoRequest(args, w, r)
			return
		case "sendVideoNote": // -> 62
			// POST /sendVideoNote
			s.handleSendVideoNoteRequest(args, w, r)
			return
		case "sendVoice": // -> 63
			// POST /sendVoice
			s.handleSendVoiceRequest(args, w, r)
			return
		case "setChatAdministratorCustomTitle": // -> 64
			// POST /setChatAdministratorCustomTitle
			s.handleSetChatAdministratorCustomTitleRequest(args, w, r)
			return
		case "setChatDescription": // -> 65
			// POST /setChatDescription
			s.handleSetChatDescriptionRequest(args, w, r)
			return
		case "setChatPermissions": // -> 66
			// POST /setChatPermissions
			s.handleSetChatPermissionsRequest(args, w, r)
			return
		case "setChatPhoto": // -> 67
			// POST /setChatPhoto
			s.handleSetChatPhotoRequest(args, w, r)
			return
		case "setChatStickerSet": // -> 68
			// POST /setChatStickerSet
			s.handleSetChatStickerSetRequest(args, w, r)
			return
		case "setChatTitle": // -> 69
			// POST /setChatTitle
			s.handleSetChatTitleRequest(args, w, r)
			return
		case "setGameScore": // -> 70
			// POST /setGameScore
			s.handleSetGameScoreRequest(args, w, r)
			return
		case "setMyCommands": // -> 71
			// POST /setMyCommands
			s.handleSetMyCommandsRequest(args, w, r)
			return
		case "setPassportDataErrors": // -> 72
			// POST /setPassportDataErrors
			s.handleSetPassportDataErrorsRequest(args, w, r)
			return
		case "setStickerPositionInSet": // -> 73
			// POST /setStickerPositionInSet
			s.handleSetStickerPositionInSetRequest(args, w, r)
			return
		case "setStickerSetThumb": // -> 74
			// POST /setStickerSetThumb
			s.handleSetStickerSetThumbRequest(args, w, r)
			return
		case "setWebhook": // -> 75
			// POST /setWebhook
			s.handleSetWebhookRequest(args, w, r)
			return
		case "stopMessageLiveLocation": // -> 76
			// POST /stopMessageLiveLocation
			s.handleStopMessageLiveLocationRequest(args, w, r)
			return
		case "stopPoll": // -> 77
			// POST /stopPoll
			s.handleStopPollRequest(args, w, r)
			return
		case "unbanChatMember": // -> 78
			// POST /unbanChatMember
			s.handleUnbanChatMemberRequest(args, w, r)
			return
		case "unbanChatSenderChat": // -> 79
			// POST /unbanChatSenderChat
			s.handleUnbanChatSenderChatRequest(args, w, r)
			return
		case "unpinAllChatMessages": // -> 80
			// POST /unpinAllChatMessages
			s.handleUnpinAllChatMessagesRequest(args, w, r)
			return
		case "unpinChatMessage": // -> 81
			// POST /unpinChatMessage
			s.handleUnpinChatMessageRequest(args, w, r)
			return
		case "uploadStickerFile": // -> 82
			// POST /uploadStickerFile
			s.handleUploadStickerFileRequest(args, w, r)
			return
		default:
			s.notFound(w, r)
			return
		}
	default:
		s.notFound(w, r)
		return
	}
}
