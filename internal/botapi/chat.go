package botapi

import (
	"context"

	"github.com/go-faster/errors"

	"github.com/gotd/botapi/internal/oas"
)

// ApproveChatJoinRequest implements oas.Handler.
func (b *BotAPI) ApproveChatJoinRequest(ctx context.Context, req oas.ApproveChatJoinRequest) (oas.Result, error) {
	return oas.Result{}, &NotImplementedError{}
}

// DeclineChatJoinRequest implements oas.Handler.
func (b *BotAPI) DeclineChatJoinRequest(ctx context.Context, req oas.DeclineChatJoinRequest) (oas.Result, error) {
	return oas.Result{}, &NotImplementedError{}
}

// DeleteChatPhoto implements oas.Handler.
func (b *BotAPI) DeleteChatPhoto(ctx context.Context, req oas.DeleteChatPhoto) (oas.Result, error) {
	return oas.Result{}, &NotImplementedError{}
}

// DeleteChatStickerSet implements oas.Handler.
func (b *BotAPI) DeleteChatStickerSet(ctx context.Context, req oas.DeleteChatStickerSet) (oas.Result, error) {
	return oas.Result{}, &NotImplementedError{}
}

// GetChat implements oas.Handler.
func (b *BotAPI) GetChat(ctx context.Context, req oas.GetChat) (oas.ResultChat, error) {
	return oas.ResultChat{}, &NotImplementedError{}
}

// SetChatAdministratorCustomTitle implements oas.Handler.
func (b *BotAPI) SetChatAdministratorCustomTitle(ctx context.Context, req oas.SetChatAdministratorCustomTitle) (oas.Result, error) {
	return oas.Result{}, &NotImplementedError{}
}

// SetChatDescription implements oas.Handler.
func (b *BotAPI) SetChatDescription(ctx context.Context, req oas.SetChatDescription) (oas.Result, error) {
	p, err := b.resolveChatID(ctx, req.ChatID)
	if err != nil {
		return oas.Result{}, errors.Wrap(err, "resolve chatID")
	}
	if err := p.SetDescription(ctx, req.Description.Value); err != nil {
		return oas.Result{}, err
	}
	return resultOK(true), nil
}

// SetChatPermissions implements oas.Handler.
func (b *BotAPI) SetChatPermissions(ctx context.Context, req oas.SetChatPermissions) (oas.Result, error) {
	return oas.Result{}, &NotImplementedError{}
}

// SetChatPhoto implements oas.Handler.
func (b *BotAPI) SetChatPhoto(ctx context.Context, req oas.SetChatPhoto) (oas.Result, error) {
	return oas.Result{}, &NotImplementedError{}
}

// SetChatStickerSet implements oas.Handler.
func (b *BotAPI) SetChatStickerSet(ctx context.Context, req oas.SetChatStickerSet) (oas.Result, error) {
	return oas.Result{}, &NotImplementedError{}
}

// SetChatTitle implements oas.Handler.
func (b *BotAPI) SetChatTitle(ctx context.Context, req oas.SetChatTitle) (oas.Result, error) {
	p, err := b.resolveChatID(ctx, req.ChatID)
	if err != nil {
		return oas.Result{}, errors.Wrap(err, "resolve chatID")
	}
	if err := p.SetTitle(ctx, req.Title); err != nil {
		return oas.Result{}, err
	}
	return resultOK(true), nil
}

// LeaveChat implements oas.Handler.
func (b *BotAPI) LeaveChat(ctx context.Context, req oas.LeaveChat) (oas.Result, error) {
	p, err := b.resolveChatID(ctx, req.ChatID)
	if err != nil {
		return oas.Result{}, errors.Wrap(err, "resolve chatID")
	}
	if !p.Left() {
		if err := p.Leave(ctx); err != nil {
			return oas.Result{}, err
		}
	}
	return resultOK(true), nil
}