// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
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
)

// Server handles operations described by OpenAPI v3 specification.
type Server interface {
	// AddStickerToSet implements addStickerToSet operation.
	AddStickerToSet(ctx context.Context, req AddStickerToSet) (Result, error)
	// AnswerCallbackQuery implements answerCallbackQuery operation.
	AnswerCallbackQuery(ctx context.Context, req AnswerCallbackQuery) (Result, error)
	// AnswerInlineQuery implements answerInlineQuery operation.
	AnswerInlineQuery(ctx context.Context, req AnswerInlineQuery) (Result, error)
	// AnswerPreCheckoutQuery implements answerPreCheckoutQuery operation.
	AnswerPreCheckoutQuery(ctx context.Context, req AnswerPreCheckoutQuery) (Result, error)
	// AnswerShippingQuery implements answerShippingQuery operation.
	AnswerShippingQuery(ctx context.Context, req AnswerShippingQuery) (Result, error)
	// BanChatMember implements banChatMember operation.
	BanChatMember(ctx context.Context, req BanChatMember) (Result, error)
	// CopyMessage implements copyMessage operation.
	CopyMessage(ctx context.Context, req CopyMessage) (Result, error)
	// CreateChatInviteLink implements createChatInviteLink operation.
	CreateChatInviteLink(ctx context.Context, req CreateChatInviteLink) (Result, error)
	// CreateNewStickerSet implements createNewStickerSet operation.
	CreateNewStickerSet(ctx context.Context, req CreateNewStickerSet) (Result, error)
	// DeleteChatPhoto implements deleteChatPhoto operation.
	DeleteChatPhoto(ctx context.Context, req DeleteChatPhoto) (Result, error)
	// DeleteChatStickerSet implements deleteChatStickerSet operation.
	DeleteChatStickerSet(ctx context.Context, req DeleteChatStickerSet) (Result, error)
	// DeleteMessage implements deleteMessage operation.
	DeleteMessage(ctx context.Context, req DeleteMessage) (Result, error)
	// DeleteMyCommands implements deleteMyCommands operation.
	DeleteMyCommands(ctx context.Context, req DeleteMyCommands) (Result, error)
	// DeleteStickerFromSet implements deleteStickerFromSet operation.
	DeleteStickerFromSet(ctx context.Context, req DeleteStickerFromSet) (Result, error)
	// DeleteWebhook implements deleteWebhook operation.
	DeleteWebhook(ctx context.Context, req DeleteWebhook) (Result, error)
	// EditChatInviteLink implements editChatInviteLink operation.
	EditChatInviteLink(ctx context.Context, req EditChatInviteLink) (Result, error)
	// EditMessageCaption implements editMessageCaption operation.
	EditMessageCaption(ctx context.Context, req EditMessageCaption) (Result, error)
	// EditMessageLiveLocation implements editMessageLiveLocation operation.
	EditMessageLiveLocation(ctx context.Context, req EditMessageLiveLocation) (Result, error)
	// EditMessageMedia implements editMessageMedia operation.
	EditMessageMedia(ctx context.Context, req EditMessageMedia) (Result, error)
	// EditMessageReplyMarkup implements editMessageReplyMarkup operation.
	EditMessageReplyMarkup(ctx context.Context, req EditMessageReplyMarkup) (Result, error)
	// EditMessageText implements editMessageText operation.
	EditMessageText(ctx context.Context, req EditMessageText) (Result, error)
	// ExportChatInviteLink implements exportChatInviteLink operation.
	ExportChatInviteLink(ctx context.Context, req ExportChatInviteLink) (Result, error)
	// ForwardMessage implements forwardMessage operation.
	ForwardMessage(ctx context.Context, req ForwardMessage) (Message, error)
	// GetChat implements getChat operation.
	GetChat(ctx context.Context, req GetChat) (Result, error)
	// GetChatAdministrators implements getChatAdministrators operation.
	GetChatAdministrators(ctx context.Context, req GetChatAdministrators) (Result, error)
	// GetChatMember implements getChatMember operation.
	GetChatMember(ctx context.Context, req GetChatMember) (Result, error)
	// GetChatMemberCount implements getChatMemberCount operation.
	GetChatMemberCount(ctx context.Context, req GetChatMemberCount) (Result, error)
	// GetFile implements getFile operation.
	GetFile(ctx context.Context, req GetFile) (Result, error)
	// GetGameHighScores implements getGameHighScores operation.
	GetGameHighScores(ctx context.Context, req GetGameHighScores) (Result, error)
	// GetMe implements getMe operation.
	GetMe(ctx context.Context) (User, error)
	// GetMyCommands implements getMyCommands operation.
	GetMyCommands(ctx context.Context, req GetMyCommands) (Result, error)
	// GetStickerSet implements getStickerSet operation.
	GetStickerSet(ctx context.Context, req GetStickerSet) (Result, error)
	// GetUpdates implements getUpdates operation.
	GetUpdates(ctx context.Context, req GetUpdates) (Result, error)
	// GetUserProfilePhotos implements getUserProfilePhotos operation.
	GetUserProfilePhotos(ctx context.Context, req GetUserProfilePhotos) (Result, error)
	// LeaveChat implements leaveChat operation.
	LeaveChat(ctx context.Context, req LeaveChat) (Result, error)
	// PinChatMessage implements pinChatMessage operation.
	PinChatMessage(ctx context.Context, req PinChatMessage) (Result, error)
	// PromoteChatMember implements promoteChatMember operation.
	PromoteChatMember(ctx context.Context, req PromoteChatMember) (Result, error)
	// RestrictChatMember implements restrictChatMember operation.
	RestrictChatMember(ctx context.Context, req RestrictChatMember) (Result, error)
	// RevokeChatInviteLink implements revokeChatInviteLink operation.
	RevokeChatInviteLink(ctx context.Context, req RevokeChatInviteLink) (Result, error)
	// SendAnimation implements sendAnimation operation.
	SendAnimation(ctx context.Context, req SendAnimation) (Message, error)
	// SendAudio implements sendAudio operation.
	SendAudio(ctx context.Context, req SendAudio) (Result, error)
	// SendChatAction implements sendChatAction operation.
	SendChatAction(ctx context.Context, req SendChatAction) (Result, error)
	// SendContact implements sendContact operation.
	SendContact(ctx context.Context, req SendContact) (Message, error)
	// SendDice implements sendDice operation.
	SendDice(ctx context.Context, req SendDice) (Message, error)
	// SendDocument implements sendDocument operation.
	SendDocument(ctx context.Context, req SendDocument) (Message, error)
	// SendGame implements sendGame operation.
	SendGame(ctx context.Context, req SendGame) (Message, error)
	// SendInvoice implements sendInvoice operation.
	SendInvoice(ctx context.Context, req SendInvoice) (Message, error)
	// SendLocation implements sendLocation operation.
	SendLocation(ctx context.Context, req SendLocation) (Message, error)
	// SendMediaGroup implements sendMediaGroup operation.
	SendMediaGroup(ctx context.Context, req SendMediaGroup) (Result, error)
	// SendMessage implements sendMessage operation.
	SendMessage(ctx context.Context, req SendMessage) (Message, error)
	// SendPhoto implements sendPhoto operation.
	SendPhoto(ctx context.Context, req SendPhoto) (Message, error)
	// SendPoll implements sendPoll operation.
	SendPoll(ctx context.Context, req SendPoll) (Message, error)
	// SendSticker implements sendSticker operation.
	SendSticker(ctx context.Context, req SendSticker) (Message, error)
	// SendVenue implements sendVenue operation.
	SendVenue(ctx context.Context, req SendVenue) (Message, error)
	// SendVideo implements sendVideo operation.
	SendVideo(ctx context.Context, req SendVideo) (Message, error)
	// SendVideoNote implements sendVideoNote operation.
	SendVideoNote(ctx context.Context, req SendVideoNote) (Message, error)
	// SendVoice implements sendVoice operation.
	SendVoice(ctx context.Context, req SendVoice) (Message, error)
	// SetChatAdministratorCustomTitle implements setChatAdministratorCustomTitle operation.
	SetChatAdministratorCustomTitle(ctx context.Context, req SetChatAdministratorCustomTitle) (Result, error)
	// SetChatDescription implements setChatDescription operation.
	SetChatDescription(ctx context.Context, req SetChatDescription) (Result, error)
	// SetChatPermissions implements setChatPermissions operation.
	SetChatPermissions(ctx context.Context, req SetChatPermissions) (Result, error)
	// SetChatPhoto implements setChatPhoto operation.
	SetChatPhoto(ctx context.Context, req SetChatPhoto) (Result, error)
	// SetChatStickerSet implements setChatStickerSet operation.
	SetChatStickerSet(ctx context.Context, req SetChatStickerSet) (Result, error)
	// SetChatTitle implements setChatTitle operation.
	SetChatTitle(ctx context.Context, req SetChatTitle) (Result, error)
	// SetGameScore implements setGameScore operation.
	SetGameScore(ctx context.Context, req SetGameScore) (Result, error)
	// SetMyCommands implements setMyCommands operation.
	SetMyCommands(ctx context.Context, req SetMyCommands) (Result, error)
	// SetPassportDataErrors implements setPassportDataErrors operation.
	SetPassportDataErrors(ctx context.Context, req SetPassportDataErrors) (Result, error)
	// SetStickerPositionInSet implements setStickerPositionInSet operation.
	SetStickerPositionInSet(ctx context.Context, req SetStickerPositionInSet) (Result, error)
	// SetStickerSetThumb implements setStickerSetThumb operation.
	SetStickerSetThumb(ctx context.Context, req SetStickerSetThumb) (Result, error)
	// SetWebhook implements setWebhook operation.
	SetWebhook(ctx context.Context, req SetWebhook) (Result, error)
	// StopMessageLiveLocation implements stopMessageLiveLocation operation.
	StopMessageLiveLocation(ctx context.Context, req StopMessageLiveLocation) (Result, error)
	// StopPoll implements stopPoll operation.
	StopPoll(ctx context.Context, req StopPoll) (Result, error)
	// UnbanChatMember implements unbanChatMember operation.
	UnbanChatMember(ctx context.Context, req UnbanChatMember) (Result, error)
	// UnpinAllChatMessages implements unpinAllChatMessages operation.
	UnpinAllChatMessages(ctx context.Context, req UnpinAllChatMessages) (Result, error)
	// UnpinChatMessage implements unpinChatMessage operation.
	UnpinChatMessage(ctx context.Context, req UnpinChatMessage) (Result, error)
	// UploadStickerFile implements uploadStickerFile operation.
	UploadStickerFile(ctx context.Context, req UploadStickerFile) (Result, error)
}
