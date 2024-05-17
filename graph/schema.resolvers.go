package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.46

import (
	"context"
	"gatepass/graph/generated"
	"gatepass/graph/model"
	"gatepass/src/codegenerator"
	"gatepass/src/codetophone"
	"gatepass/src/phonesender"
)

func (r *mutationResolver) GetCodeFromPhone(ctx context.Context, input model.Phone) (*model.ResponseGetCode, error) {
	code, err := codegenerator.GenerateCode()
	if err != nil {
		return &model.ResponseGetCode{
			Message: "Error While Generating Code",
		}, nil
	}
	isOk, err := phonesender.SendMessage(code, input.Phone)
	if err != nil {
		return &model.ResponseGetCode{
			Message: "Error While Sending Message",
		}, nil
	}
	if isOk {
		return &model.ResponseGetCode{
			Message: "Code sent to " + input.Phone,
		}, nil
	}
	message := "Failed to send code"
	return &model.ResponseGetCode{
		Message: message,
	}, nil
}

func (r *mutationResolver) VerifyCode(ctx context.Context, input model.Verification) (*model.ResponseVerifyCode, error) {
	if codetophone.VerifyCodeToPhone(input.Phone, input.Code) {
		message := "Correct Verification Code"
		return &model.ResponseVerifyCode{
			Authenticated: true,
			Message:       &message,
			Phone:         &input.Phone,
		}, nil
	}
	errMessage := "Invalid code"
	return &model.ResponseVerifyCode{
		Authenticated: false,
		Error:         &errMessage,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }