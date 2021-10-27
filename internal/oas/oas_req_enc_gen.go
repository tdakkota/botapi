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

func encodeAddStickerToSetRequest(req AddStickerToSet) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeAnswerCallbackQueryRequest(req AnswerCallbackQuery) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeAnswerInlineQueryRequest(req AnswerInlineQuery) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeAnswerPreCheckoutQueryRequest(req AnswerPreCheckoutQuery) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeAnswerShippingQueryRequest(req AnswerShippingQuery) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeBanChatMemberRequest(req BanChatMember) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeCopyMessageRequest(req CopyMessage) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeCreateChatInviteLinkRequest(req CreateChatInviteLink) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeCreateNewStickerSetRequest(req CreateNewStickerSet) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeDeleteChatPhotoRequest(req DeleteChatPhoto) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeDeleteChatStickerSetRequest(req DeleteChatStickerSet) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeDeleteMessageRequest(req DeleteMessage) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeDeleteMyCommandsRequest(req DeleteMyCommands) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeDeleteStickerFromSetRequest(req DeleteStickerFromSet) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeDeleteWebhookRequest(req DeleteWebhook) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeEditChatInviteLinkRequest(req EditChatInviteLink) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeEditMessageCaptionRequest(req EditMessageCaption) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeEditMessageLiveLocationRequest(req EditMessageLiveLocation) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeEditMessageMediaRequest(req EditMessageMedia) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeEditMessageReplyMarkupRequest(req EditMessageReplyMarkup) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeEditMessageTextRequest(req EditMessageText) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeExportChatInviteLinkRequest(req ExportChatInviteLink) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeForwardMessageRequest(req ForwardMessage) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetChatRequest(req GetChat) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetChatAdministratorsRequest(req GetChatAdministrators) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetChatMemberRequest(req GetChatMember) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetChatMemberCountRequest(req GetChatMemberCount) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetFileRequest(req GetFile) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetGameHighScoresRequest(req GetGameHighScores) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetMyCommandsRequest(req GetMyCommands) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetStickerSetRequest(req GetStickerSet) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetUpdatesRequest(req GetUpdates) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeGetUserProfilePhotosRequest(req GetUserProfilePhotos) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeLeaveChatRequest(req LeaveChat) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePinChatMessageRequest(req PinChatMessage) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePromoteChatMemberRequest(req PromoteChatMember) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeRestrictChatMemberRequest(req RestrictChatMember) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeRevokeChatInviteLinkRequest(req RevokeChatInviteLink) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendAnimationRequest(req SendAnimation) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendAudioRequest(req SendAudio) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendChatActionRequest(req SendChatAction) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendContactRequest(req SendContact) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendDiceRequest(req SendDice) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendDocumentRequest(req SendDocument) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendGameRequest(req SendGame) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendInvoiceRequest(req SendInvoice) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendLocationRequest(req SendLocation) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendMediaGroupRequest(req SendMediaGroup) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendMessageRequest(req SendMessage) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendPhotoRequest(req SendPhoto) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendPollRequest(req SendPoll) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendStickerRequest(req SendSticker) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendVenueRequest(req SendVenue) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendVideoRequest(req SendVideo) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendVideoNoteRequest(req SendVideoNote) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSendVoiceRequest(req SendVoice) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetChatAdministratorCustomTitleRequest(req SetChatAdministratorCustomTitle) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetChatDescriptionRequest(req SetChatDescription) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetChatPermissionsRequest(req SetChatPermissions) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetChatPhotoRequest(req SetChatPhoto) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetChatStickerSetRequest(req SetChatStickerSet) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetChatTitleRequest(req SetChatTitle) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetGameScoreRequest(req SetGameScore) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetMyCommandsRequest(req SetMyCommands) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetPassportDataErrorsRequest(req SetPassportDataErrors) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetStickerPositionInSetRequest(req SetStickerPositionInSet) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetStickerSetThumbRequest(req SetStickerSetThumb) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeSetWebhookRequest(req SetWebhook) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeStopMessageLiveLocationRequest(req StopMessageLiveLocation) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeStopPollRequest(req StopPoll) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeUnbanChatMemberRequest(req UnbanChatMember) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeUnpinAllChatMessagesRequest(req UnpinAllChatMessages) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeUnpinChatMessageRequest(req UnpinChatMessage) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodeUploadStickerFileRequest(req UploadStickerFile) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}
