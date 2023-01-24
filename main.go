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
	builder.Say("Aqui estão algumas coisas que você pode me perguntar:")
	builder.Pause("200")
	builder.Say("Qual o torneio de hoje na H2?")
	builder.Pause("200")
	builder.Say("Qual o horário do intervalo do torneio de hoje?")
	return alexa.NewSSMLResponse("H2 Ajuda", builder.Build())
}

func HandleAboutIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Popular Deals", "Popular deal data here")
}

func HandleDailyTournament(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder
	builder.Say("5K TRIPLE SHOT DIA UNICO")
	builder.Pause("200")
	builder.Say("Início 17:00")
	builder.Pause("100")
	builder.Say("Inscrições até 19:20")
	builder.Pause("100")
	return alexa.NewSSMLResponse("H2 Ajuda", builder.Build())
}

func HandleBreakTournament(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder
	builder.Say("5K TRIPLE SHOT DIA UNICO")
	builder.Pause("200")
	builder.Say("Inscrições até 19:20")
	builder.Pause("100")
	return alexa.NewSSMLResponse("H2 Ajuda", builder.Build())
}

func IntentDispatcher(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "AboutIntent":
		response = HandleAboutIntent(request)
	case "DailyTournament":
		response = HandleDailyTournament(request)
	case "BreakTournament":
		response = HandleBreakTournament(request)
	default:
		response = HandleAboutIntent(request)
	}
	return response
}
