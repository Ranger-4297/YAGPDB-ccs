{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(russian-?roulette(rr)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1}}

{{/* Russian Roulette */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{$betMax := $a.betMax}}
	{{$incomeCooldown := $a.incomeCooldown | toInt}}
	{{$userDB := (dbGet $userID "EconomyInfo")}}
	{{if not $db}}
		{{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
	{{end}}
	{{with $db}}
		{{$a = sdict .Value}}
		{{$b := sdict .Value}}
		{{$bal := $a.cash}}
		{{with $.CmdArgs}}
			{{with dbGet 0 "russianRoulette"}}
				{{$a = sdict .Value}}
				{{if $a.game}}
					{{$game := $a.game}}
					{{$cost := $game.cost}}
					{{$players := $game.players}}
					{{$bet := (index . 0)}}
					{{$continue := false}}
					{{if $bet | toInt}}
						{{if lt $bet $betMax}}
							{{if gt (toInt $bet) 0}}
								{{if le (toInt $bet) (toInt $bal)}}
									{{$continue = true}}
								{{else}}
									{{$embed.Set "description" (print "You can't bet more than you have!")}}
									{{$embed.Set "color" $errorColor}}
								{{end}}
							{{else}}
								{{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
								{{$embed.Set "color" $errorColor}}
							{{end}}
						{{else}}
							{{$embed.Set "description" (print "You can't bet more than " $symbol $betMax)}}
							{{$embed.Set "color" $errorColor}}
						{{end}}
					{{else}}
						{{$bet = $bet | lower}}
						{{if eq $bet "all"}}
							{{if ne (toInt $bal) 0}}
								{{if lt $bet $betMax}}
									{{$continue = true}}
									{{$bet = $bal}}
								{{else}}
									{{$embed.Set "description" (print "You can't bet more than " $symbol $betMax)}}
									{{$embed.Set "color" $errorColor}}
								{{end}}
							{{else}}
								{{$embed.Set "description" (print "You had no money to bet.")}}
								{{$embed.Set "color" $errorColor}}
							{{end}}
						{{else if eq $bet "start"}}
							{{if eq $game.owner $userID}}
								{{$rr := true}}
							{{else}}
								{{$embed.Set "description" (print "You cannot start the game")}}
								{{$embed.Set "color" $errorColor}}
							{{end}}
						{{else if eq $bet "retrieve" "collect"}}
							{{$storageDB (dbGet 0 "rouletteStorage")}}
							{{if $storageDB.Get (toString $userID)}}
								{{$winnings :=  $storageDB.Get (toString $userID).amount}}
								{{$embed.Set "description" (print "You've collected " $winnings)}}
								{{$embed.Set "color" $successColor}}
								{{$storageDB.Del (toString $userID )}}
								{{dbSet 0 "rouletteStorage" $storageDB}}
								{{$b.Set "cash" (add $bal $winnings)}}
								{{dbSet $userID $b}}
							{{else}}
								{{$embed.Set "description" (print "You had no winning to collect!")}}
								{{$embed.Set "color" $errorColor}}
							{{end}}
						{{else}}
							{{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
							{{$embed.Set "color" $errorColor}}
						{{end}}
					{{end}}
					{{if $continue}}
						{{if not eq (len $players) 6}}
							{{if in $players $userID}}
								{{$embed.Set "description" (print "You're already in this game")}}
								{{$embed.Set "color" $errorColor}}
							{{else}}
								{{$players = $players.Append $userID}}
								{{$game.Set "players" $players}}
								{{dbSet 0 "russianRoulette" (sdict "game" $game)}}
								{{$b.Set "cash" (sub $bal $bet)}}
								{{dbSet $userID $b}}
								{{$embed.Set "description" (print "You've joined this game of russian roulette with a bet of " $symbol $bet)}}
								{{$embed.Set "footer" (sdict "name" (print (len $players) "/6"))}}
								{{$embed.Set "color" $successColor}}
							{{end}}
						{{else}}
							{{$embed.Set "description" (print "The maximum number of players have already joined ;( Sorry")}}
							{{$embed.Set "color" $errorColor}}
						{{end}}
					{{end}}
					{{if $rr}}
						{{sendMessage nil (cembed (sdict "title" "The russian roulette game has begun!" "color" 0x0088CC))}}
						{{$winners := cslice}}
						{{$loser := ""}}
						{{$n := randInt (len $players)}}
						{{range $i, $p := $players}}
							{{- if ne $i $n -}}
								{{sendMessage nil (cembed (sdict "description" (print "**" (userArg $p) "** pulled the trigger and survived") "color" 0x0088CC))}}
								{{sleep 1}}
								{{- continue -}}
							{{- end -}}
							{{$loser := (userArg $p)}}
							{{sendMessage nil (cembed (sdict "description" (print "**" $loser "** pulled the trigger and dies") "color" 0xFF5F1F))}}
							{{break}}
						{{end}}
						{{range $players}}
							{{- if ne . $loser.ID}}
								{{- $winners = $winners.Append (userArg .).Mention}}
							{{- end -}}
						{{end}}
						{{$payout := (div $game.cost (len $winners))}}
						{{$entry := cslice}}
						{{$storedUsers := cslice}}
						{{$storageDB := (dbGet 0 "rouletteStorage")}}
						{{if $storageDB}}
							{{$storedUsers = $storageDB}}
						{{end}}
						{{range $winners}}
							{{$amount := ($storedUsers.Get (toString $userID)).amount}}
							{{if $amount}}
								{{$bet = add $amount $bet}} 
							{{end}}
							{{- $entry = $entry.Append (sdict "Name" (print .) "value" .Mention "inline" false) -}}
							{{- $storedUsers = $storedUsers.Append (sdict "user" .ID "amount" $bet) -}}
						{{end}}
						{{$embed.Set "title" "Winners"}}
						{{$embed.Set "description" (print "payout is: " $payout " per-person")}}
						{{$embed.Set "fields" $entry}}
						{{$embed.Set "color" $successColor}}
						{{dbSet 0 "rouletteStorage" $storedUsers}}
						{{end}}
					{{end}}
				{{else}}
					{{if (index . 0) | toInt}}
						{{$bet := (index . 0) | toInt}}
						{{if gt (toInt $bet) 0}}
							{{if le (toInt $bet) (toInt $bal)}}
								{{dbSet 0 "russianRoulette" (sdict "game" (sdict "cost" $bet "players" (cslice $userID)))}}
								{{$embed.Set "description" (print "A new game of Russian roulette has been started!\n\nTo join use the command " $.Cmd " " $bet " (1/6)\nTo start this game use the command " $.Cmd "start")}}
								{{$embed.Set "color" $successColor}}
							{{else}}
								{{$embed.Set "description" (print "You can't bet more than you have!")}}
								{{$embed.Set "color" $errorColor}}
							{{end}}
						{{else}}
							{{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
							{{$embed.Set "color" $errorColor}}
						{{end}}
					{{end}}
				{{end}}
			{{end}}
		{{else}}
			{{$embed.Set "description" (print "No `bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}