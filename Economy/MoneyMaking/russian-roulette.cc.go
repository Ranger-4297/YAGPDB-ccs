{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(russian-?roulette|rr)(\s+|\z)`

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
{{$em := sdict}}
{{$em.Set "timestamp" currentTime}}
{{if not .ExecData}}
	{{with dbGet 0 "EconomySettings"}}
		{{$a := sdict .Value}}
		{{$symbol := $a.symbol}}
		{{$betMax := $a.betMax}}
		{{$incomeCooldown := $a.incomeCooldown | toInt}}
		{{$userDB := (dbGet $userID "EconomyInfo")}}
		{{if not $userDB}}
			{{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
		{{end}}
		{{with $userDB}}
			{{$a = sdict .Value}}
			{{$b := sdict .Value}}
			{{$bal := $a.cash}}
			{{with $.CmdArgs}}
				{{with dbGet 0 "russianRoulette"}}
					{{$a = sdict .Value}}
					{{if $a.game}}
						{{$cost := $a.cost}}
						{{$game := $a.game}}
						{{$cost := $game.cost}}
						{{$players := $game.players}}
						{{$bet := (index $.CmdArgs 0)}}
						{{$continue := false}}
						{{$rr := false}}
						{{if $bet | toInt}}
							{{if lt (toInt $bet) (toInt $betMax)}}
								{{if gt (toInt $bet) 0}}
									{{if le (toInt $bet) (toInt $bal)}}
										{{if le (toInt $bet) (toInt $cost)}}
											{{$continue = true}}
										{{else}}
											{{$em.Set "description" (print "You can't bet more than " $cost)}}
											{{$em.Set "color" $errorColor}}
										{{end}}
									{{else}}
										{{$em.Set "description" (print "You can't bet more than you have!")}}
										{{$em.Set "color" $errorColor}}
									{{end}}
								{{else}}
									{{$em.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet>`")}}
									{{$em.Set "color" $errorColor}}
								{{end}}
							{{else}}
								{{$em.Set "description" (print "You can't bet more than " $symbol $betMax)}}
								{{$em.Set "color" $errorColor}}
							{{end}}
						{{else}}
							{{$bet = $bet | lower}}
							{{if eq $bet "all"}}
								{{if ne (toInt $bal) 0}}
									{{if lt $bet $betMax}}
										{{$continue = true}}
										{{$bet = $bal}}
									{{else}}
										{{$em.Set "description" (print "You can't bet more than " $symbol $betMax)}}
										{{$em.Set "color" $errorColor}}
									{{end}}
								{{else}}
									{{$em.Set "description" (print "You had no money to bet.")}}
									{{$em.Set "color" $errorColor}}
								{{end}}
							{{else if eq $bet "start"}}
								{{if eq (toString $userID) (toString $game.owner)}}
									{{$rr = true}}
								{{else}}
									{{$em.Set "description" (print "You cannot start the game")}}
									{{$em.Set "color" $errorColor}}
								{{end}}
							{{else if eq $bet "retrieve" "collect"}}
								{{$sDB := (dbGet 0 "rouletteStorage")}}
								{{if $sDB.Get (toString $userID)}}
									{{$winnings :=  $sDB.Get (toString $userID).amount}}
									{{$em.Set "description" (print "You've collected " $winnings)}}
									{{$em.Set "color" $successColor}}
									{{$sDB.Del (toString $userID )}}
									{{dbSet 0 "rouletteStorage" $sDB}}
									{{$b.Set "cash" (add $bal $winnings)}}
									{{dbSet $userID $b}}
								{{else}}
									{{$em.Set "description" (print "You had no winning to collect!")}}
									{{$em.Set "color" $errorColor}}
								{{end}}
							{{else}}
								{{$em.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
								{{$em.Set "color" $errorColor}}
							{{end}}
						{{end}}
						{{if $continue}}
							{{if not (eq (len $players) 6)}}
								{{if in $players $userID}}
									{{$em.Set "description" (print "You're already in this game")}}
									{{$em.Set "color" $errorColor}}
								{{else}}
									{{$players = $players.Append $userID}}
									{{$game.Set "players" $players}}
									{{dbSet 0 "russianRoulette" (sdict "game" $game)}}
									{{$b.Set "cash" (sub $bal $bet)}}
									{{dbSet $userID "EconomyInfo" $b}}
									{{$em.Set "description" (print "You've joined this game of russian roulette with a bet of " $symbol $bet)}}
									{{$em.Set "footer" (sdict "text" (print "Players: " (len $players) "/6"))}}
									{{$em.Set "color" $successColor}}
									{{cancelScheduledUniqueCC $.CCID "rr-game"}}
									{{scheduleUniqueCC $.CCID nil 420 "rr-game-2" "cancel-2"}}
								{{end}}
							{{else}}
								{{$em.Set "description" (print "The maximum number of players have already joined ;( Sorry")}}
								{{$em.Set "color" $errorColor}}
							{{end}}
						{{end}}
						{{if $rr}}
							{{sendMessage nil (cembed (sdict "title" "The russian roulette game has begun!" "color" 0x0088CC))}}
							{{$winners := cslice}}
							{{$loser := ""}}
							{{if gt (len $players) 1}}
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
							{{else}}
								{{$em.Set "description" (print "Not enough players to start the match :(\nStart a new one with " $.Cmd " <bet>")}}
								{{$em.Set "color" $errorColor}}
								{{dbSet 0 "russianRoulette" sdict}}
								{{cancelScheduledUniqueCC $.CCID "rr-game"}}
							{{end}}
							{{$payout := (div $game.cost (len $winners))}}
							{{$fields := cslice}}
							{{$sU := cslice}}
							{{$sDB := (dbGet 0 "rouletteStorage")}}
							{{if $sDB}}
								{{$sU = $sDB}}
							{{end}}
							{{range $winners}}
								{{$amount := ($sU.Get (toString .ID)).amount}}
								{{if $amount}}
									{{$amount = add $amount $payout}} 
								{{else}}
									{{$amount = $payout}}
								{{end}}
								{{- $fields = $fields.Append (sdict "Name" (print .) "value" .Mention "inline" false) -}}
								{{- $sU = $sU.Append (sdict "user" .ID "amount" $amount) -}}
							{{end}}
							{{$em.Set "title" "Winners"}}
							{{$em.Set "description" (print "payout is: " $payout " per-person")}}
							{{$em.Set "fields" $fields}}
							{{$em.Set "color" $successColor}}
							{{dbSet 0 "rouletteStorage" $sU}}
							{{dbSet 0 "russianRoulette" sdict}}
							{{cancelScheduledUniqueCC $.CCID "rr-game"}}
						{{end}}
					{{else}}
						{{if (index $.CmdArgs 0) | toInt}}
							{{$bet := (index $.CmdArgs 0) | toInt}}
							{{if gt (toInt $bet) 0}}
								{{if le (toInt $bet) (toInt $bal)}}
									{{if $a.game}}
										{{$em.Set "description" (print "There is already a current game. Join it with `" $.Cmd " " $a.game.cost "`")}}
										{{$em.Set "color" $errorColor}}
									{{else}}
										{{dbSet 0 "russianRoulette" (sdict "game" (sdict "cost" $bet "players" (cslice $userID) "owner" $userID))}}
										{{$bal = sub $bal $bet}}
										{{$b.Set "cash" $bal}}
										{{dbSet $userID "EconomyInfo" $b}}
										{{$em.Set "description" (print "A new game of Russian roulette has been started!\n\nTo join use the command `" $.Cmd " " $bet "` (1/6)\nTo start this game use the command `" $.Cmd " start`")}}
										{{$em.Set "color" $successColor}}
										{{scheduleUniqueCC $.CCID nil 300 "rr-game" "cancel"}}
									{{end}}
								{{else}}
									{{$em.Set "description" (print "You can't bet more than you have!")}}
									{{$em.Set "color" $errorColor}}
								{{end}}
							{{else}}
								{{$em.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet>`")}}
								{{$em.Set "color" $errorColor}}
							{{end}}
						{{end}}
					{{end}}
				{{end}}
			{{else}}
				{{$em.Set "description" (print "No `bet` argument provided.\nSyntax is `" $.Cmd " <Bet>`")}}
				{{$em.Set "color" $errorColor}}
			{{end}}
		{{end}}
	{{else}}
		{{$em.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
		{{$em.Set "color" $errorColor}}
	{{end}}
{{else}}
	{{if eq .ExecData "cancel"}}
		{{dbSet 0 "russianRoulette" sdict}}
		{{$em.Set "description" (print "Not enough players joined for Russian-roulette")}}
		{{$em.Set "color" $errorColor}}
		{{cancelScheduledUniqueCC .CCID "rr-game"}}
	{{else}}
		{{dbSet 0 "russianRoulette" sdict}}
		{{$em.Set "description" (print "The host took too long to start the game. Please start a new one.")}}
		{{$em.Set "color" $errorColor}}
		{{cancelScheduledUniqueCC .CCID "rr-game-2"}}
		{{with dbGet 0 "russianRoulette"}}
			{{$a := sdict .Value}}
			{{$cost := $a.cost}}
			{{$sU := cslice}}
			{{$sDB := (dbGet 0 "rouletteStorage")}}
			{{if $sDB}}
				{{$sU = $sDB}}
			{{end}}
			{{range $a.players}}
				{{$amount := ($sU.Get (toString .ID)).amount}}
				{{if $amount}}
					{{$amount = add $amount $cost}} 
				{{else}}
					{{$amount = $cost}}
				{{end}}
				{{- $sU = $sU.Append (sdict "user" .ID "amount" $amount) -}}
			{{end}}
			{{dbSet 0 "rouletteStorage" $sU}}
		{{end}}
	{{end}}
{{end}}
{{sendMessage nil (cembed $em)}}