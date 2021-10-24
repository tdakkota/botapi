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

func encodeAddStickerToSetResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeAnswerCallbackQueryResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeAnswerInlineQueryResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeAnswerPreCheckoutQueryResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeAnswerShippingQueryResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeBanChatMemberResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeCopyMessageResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeCreateChatInviteLinkResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeCreateNewStickerSetResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeDeleteChatPhotoResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeDeleteChatStickerSetResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeDeleteMessageResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeDeleteMyCommandsResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeDeleteStickerFromSetResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeDeleteWebhookResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeEditChatInviteLinkResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeEditMessageCaptionResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeEditMessageLiveLocationResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeEditMessageMediaResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeEditMessageReplyMarkupResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeEditMessageTextResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeExportChatInviteLinkResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeForwardMessageResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetChatResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetChatAdministratorsResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetChatMemberResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetChatMemberCountResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetFileResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetGameHighScoresResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetMeResponse(response User, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetMyCommandsResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetStickerSetResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetUpdatesResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetUserProfilePhotosResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeLeaveChatResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodePinChatMessageResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodePromoteChatMemberResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeRestrictChatMemberResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeRevokeChatInviteLinkResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendAnimationResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendAudioResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendChatActionResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendContactResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendDiceResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendDocumentResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendGameResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendInvoiceResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendLocationResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendMediaGroupResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendMessageResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendPhotoResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendPollResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendStickerResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendVenueResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendVideoResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendVideoNoteResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendVoiceResponse(response Message, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetChatAdministratorCustomTitleResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetChatDescriptionResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetChatPermissionsResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetChatPhotoResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetChatStickerSetResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetChatTitleResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetGameScoreResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetMyCommandsResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetPassportDataErrorsResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetStickerPositionInSetResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetStickerSetThumbResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetWebhookResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeStopMessageLiveLocationResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeStopPollResponse(response Poll, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeUnbanChatMemberResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeUnpinAllChatMessagesResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeUnpinChatMessageResponse(response bool, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeUploadStickerFileResponse(response string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}
