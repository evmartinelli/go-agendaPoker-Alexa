package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/evmartinelli/go-agendaPoker-Alexa/alexa"
)

func Handler(request alexa.Request) (alexa.Response, error) {
	return IntentDispatcher(request), nil
}

func main() {
	lambda.Start(Handler)
}

func HandleFrontpageDealIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Frontpage Deals", "Frontpage deal data here")
}

func HandlePopularDealIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Popular Deals", "Popular deal data here")
}

func HandleHelpIntent(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder
	builder.Say("Here are some of the things you can ask:")
	builder.Pause("1000")
	builder.Say("Give me the frontpage deals.")
	builder.Pause("1000")
	builder.Say("Give me the popular deals.")
	return alexa.NewSSMLResponse("Slick Dealer Help", builder.Build())
}

func HandleAboutIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("About", "Slick Dealer was created by Nic Raboy in Tracy, California as an unofficial Slick Deals application.")
}

func IntentDispatcher(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "FrontpageDealIntent":
		response = HandleFrontpageDealIntent(request)
	case "PopularDealIntent":
		response = HandlePopularDealIntent(request)
	case alexa.HelpIntent:
		response = HandleHelpIntent(request)
	case "AboutIntent":
		response = HandleAboutIntent(request)
	default:
		response = HandleAboutIntent(request)
	}
	return response
}
