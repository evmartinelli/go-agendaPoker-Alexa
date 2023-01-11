package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/evmartinelli/go-agendaPoker-Alexa/alexa"
)

// func Handler(request alexa.Request) (alexa.Response, error) {
// 	return IntentDispatcher(request), nil
// }

func main() {
	lambda.Start(Handler)
}

func Handler() (string, error) {
	return fmt.Sprintf("Hello World"), nil
}

func HandleFrontpageDealIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Frontpage Deals", "Frontpage deal data here")
}

func HandlePopularDealIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Popular Deals", "Popular deal data here")
}

func HandleHelpIntent(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder
	builder.Say("Aqui estão algumas coisas que você pode me perguntar:")
	builder.Pause("1000")
	builder.Say("Qual o torneio de hoje na H2?")
	builder.Pause("1000")
	builder.Say("Qual o horário do intervalo do torneio de hoje?")
	return alexa.NewSSMLResponse("H2 Ajuda", builder.Build())
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
