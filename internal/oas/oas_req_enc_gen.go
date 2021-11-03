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
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/errors"
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
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
)

func encodeAddStickerToSetRequest(req AddStickerToSet, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeAnswerCallbackQueryRequest(req AnswerCallbackQuery, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeAnswerInlineQueryRequest(req AnswerInlineQuery, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeAnswerPreCheckoutQueryRequest(req AnswerPreCheckoutQuery, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeAnswerShippingQueryRequest(req AnswerShippingQuery, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeBanChatMemberRequest(req BanChatMember, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeCopyMessageRequest(req CopyMessage, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeCreateChatInviteLinkRequest(req CreateChatInviteLink, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeCreateNewStickerSetRequest(req CreateNewStickerSet, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeDeleteChatPhotoRequest(req DeleteChatPhoto, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeDeleteChatStickerSetRequest(req DeleteChatStickerSet, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeDeleteMessageRequest(req DeleteMessage, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeDeleteMyCommandsRequest(req DeleteMyCommands, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeDeleteStickerFromSetRequest(req DeleteStickerFromSet, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeDeleteWebhookRequest(req DeleteWebhook, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeEditChatInviteLinkRequest(req EditChatInviteLink, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeEditMessageCaptionRequest(req EditMessageCaption, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeEditMessageLiveLocationRequest(req EditMessageLiveLocation, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeEditMessageMediaRequest(req EditMessageMedia, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeEditMessageReplyMarkupRequest(req EditMessageReplyMarkup, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeEditMessageTextRequest(req EditMessageText, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeExportChatInviteLinkRequest(req ExportChatInviteLink, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeForwardMessageRequest(req ForwardMessage, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetChatRequest(req GetChat, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetChatAdministratorsRequest(req GetChatAdministrators, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetChatMemberRequest(req GetChatMember, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetChatMemberCountRequest(req GetChatMemberCount, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetFileRequest(req GetFile, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetGameHighScoresRequest(req GetGameHighScores, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetMyCommandsRequest(req GetMyCommands, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetStickerSetRequest(req GetStickerSet, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetUpdatesRequest(req GetUpdates, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetUserProfilePhotosRequest(req GetUserProfilePhotos, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeLeaveChatRequest(req LeaveChat, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePinChatMessageRequest(req PinChatMessage, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePromoteChatMemberRequest(req PromoteChatMember, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeRestrictChatMemberRequest(req RestrictChatMember, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeRevokeChatInviteLinkRequest(req RevokeChatInviteLink, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendAnimationRequest(req SendAnimation, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendAudioRequest(req SendAudio, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendChatActionRequest(req SendChatAction, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendContactRequest(req SendContact, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendDiceRequest(req SendDice, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendDocumentRequest(req SendDocument, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendGameRequest(req SendGame, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendInvoiceRequest(req SendInvoice, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendLocationRequest(req SendLocation, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendMediaGroupRequest(req SendMediaGroup, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendMessageRequest(req SendMessage, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendPhotoRequest(req SendPhoto, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendPollRequest(req SendPoll, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendStickerRequest(req SendSticker, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendVenueRequest(req SendVenue, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendVideoRequest(req SendVideo, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendVideoNoteRequest(req SendVideoNote, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendVoiceRequest(req SendVoice, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetChatAdministratorCustomTitleRequest(req SetChatAdministratorCustomTitle, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetChatDescriptionRequest(req SetChatDescription, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetChatPermissionsRequest(req SetChatPermissions, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetChatPhotoRequest(req SetChatPhoto, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetChatStickerSetRequest(req SetChatStickerSet, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetChatTitleRequest(req SetChatTitle, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetGameScoreRequest(req SetGameScore, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetMyCommandsRequest(req SetMyCommands, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetPassportDataErrorsRequest(req SetPassportDataErrors, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetStickerPositionInSetRequest(req SetStickerPositionInSet, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetStickerSetThumbRequest(req SetStickerSetThumb, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetWebhookRequest(req SetWebhook, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeStopMessageLiveLocationRequest(req StopMessageLiveLocation, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeStopPollRequest(req StopPoll, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeUnbanChatMemberRequest(req UnbanChatMember, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeUnpinAllChatMessagesRequest(req UnpinAllChatMessages, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeUnpinChatMessageRequest(req UnpinChatMessage, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeUploadStickerFileRequest(req UploadStickerFile, span trace.Span) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	e := json.GetEncoder()
	defer json.PutEncoder(e)
	more := json.NewMore(e)
	defer more.Reset()
	more.More()
	req.WriteJSON(e)
	if _, err := e.WriteTo(buf); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}
