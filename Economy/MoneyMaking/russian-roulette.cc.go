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
{{$prefix := .ServerPrefix}}

{{/* Russian Roulette */}}

{{/* Response */}}
{{$em := sdict}}
{{$em.Set "timestamp" currentTime}}
{{if not (eq (toString .ExecData) "start")}}
	{{with dbGet 0 "EconomySettings"}}
		{{$a := sdict .Value}}
		{{$symbol := $a.symbol}}
		{{$betMax := $a.betMax | toInt}}
		{{$bal := or (dbGet $userID "cash").Value 0 | toInt}}
		{{with or $.CmdArgs (eq (toString $.ExecData) "start")}}
			{{with dbGet 0 "russianRoulette"}}
				{{$a := sdict .Value}}
				{{if $a.game}}
					{{$cost := $a.cost}}
					{{$game := $a.game}}
					{{$cost := $game.cost}}
					{{$players := $game.players}}
					{{$bet := (index $.CmdArgs 0)}}
					{{$ct := false}}
					{{$rr := false}}
					{{if eq ($bet | toString) "all"}}
						{{$bet = $bal}}
					{{else if and $betMax (eq (toString $bet) "max")}}
						{{$bet = $betMax}}
					{{end}}
					{{if toInt $bet}}
						{{$bet = toInt $bet}}
						{{if gt $bet 0}}
							{{if le $bet $bal}}
								{{if le $bet $cost}}
									{{$ct = true}}
									{{if $betMax}}
										{{if gt $bet $betMax}}
											{{$em.Set "description" (print "You can't bet more than " $symbol $betMax)}}
											{{$em.Set "color" $errorColor}}
											{{$ct = false}}
										{{end}}
									{{end}}
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
						{{$bet = lower (toString $bet)}}
						{{if eq $bet "start"}}
							{{if eq (toString $userID) (toString $game.owner)}}
								{{$rr = true}}
							{{else}}
								{{$em.Set "description" (print "You cannot start the game")}}
								{{$em.Set "color" $errorColor}}
							{{end}}
						{{else if eq $bet "collect"}}
							{{$em.Set "description" (print "You cannot collect during a game. Please wait till it is over to collect any owed money.")}}
							{{$em.Set "color" $errorColor}}
						{{else if not $.ExecData}}
							{{$em.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
							{{$em.Set "color" $errorColor}}
						{{end}}
					{{end}}
					{{if $ct}}
						{{if not (eq (len $players) 6)}}
							{{if in $players $userID}}
								{{$em.Set "description" (print "You're already in this game")}}
								{{$em.Set "color" $errorColor}}
							{{else}}
								{{$players = $players.Append $userID}}
								{{$game.Set "players" $players}}
								{{dbSet 0 "russianRoulette" (sdict "game" $game)}}
								{{$bal = sub $bal $bet}}
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
					{{if or $rr (eq (toString $.ExecData) "start")}}
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
								{{$loser = (userArg $p)}}
								{{sendMessage nil (cembed (sdict "description" (print "**" $loser "** pulled the trigger and dies") "color" 0xFF5F1F))}}
								{{break}}
							{{end}}
							{{range $players}}
								{{- if ne . $loser.ID}}
									{{- $winners = $winners.Append (userArg .).Mention}}
								{{- end -}}
							{{end}}
							{{$payout := (div $game.cost (len $winners))}}
							{{$fields := cslice}}
							{{$sDB := $a.storage}}
							{{if not $sDB}}
								{{$sDB = sdict}}
							{{end}}
							{{range $winners}}
								{{$amt := ($sDB.Get (toString (userArg .).ID))}}
								{{if $amt}}
									{{$amt = add $amt $payout $cost}}
								{{else}}
									{{$amt = add $payout $cost}} 
								{{end}}
								{{- $fields = $fields.Append (sdict "Name" (print (userArg .)) "value" . "inline" false) -}}
								{{$sDB.Set (toString (userArg .).ID) $amt}}
							{{end}}
							{{$em.Set "title" "Winners"}}
							{{$em.Set "description" (print "payout is: " $symbol $payout " per-person")}}
							{{$em.Set "fields" $fields}}
							{{$em.Set "color" $successColor}}
							{{dbSet 0 "russianRoulette" (sdict "storage" $sDB)}}
							{{cancelScheduledUniqueCC $.CCID "rr-game"}}
						{{else}}
							{{$em.Set "description" (print "Not enough players to start the match :(\nStart a new one with " $.Cmd " <bet>")}}
							{{$em.Set "color" $errorColor}}
							{{dbSet 0 "russianRoulette" sdict}}
							{{cancelScheduledUniqueCC $.CCID "rr-game"}}
						{{end}}
					{{end}}
				{{else}}
					{{if (index $.CmdArgs 0) | toInt}}
						{{$bet := (index $.CmdArgs 0) | toInt}}
						{{if gt (toInt $bet) 0}}
							{{if le (toInt $bet) (toInt $bal)}}
								{{$continue := true}}
								{{if $betMax}}
									{{if gt $bet $betMax}}
										{{$em.Set "description" (print "You can't bet more than " $symbol $betMax)}}
										{{$em.Set "color" $errorColor}}
										{{$continue = false}}
									{{end}}
								{{end}}
								{{if $continue}}
									{{dbSet 0 "russianRoulette" (sdict "game" (sdict "cost" $bet "players" (cslice $userID) "owner" $userID))}}
									{{$bal = sub $bal $bet}}
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
					{{else if eq (index $.CmdArgs 0) "collect"}}
						{{$sDB := $a.storage}}
						{{$amt := $sDB.Get (toString $userID)}}
						{{if $amt}}
							{{$em.Set "description" (print "You've collected " $symbol $amt)}}
							{{$em.Set "color" $successColor}}
							{{$sDB.Del (toString $userID )}}
							{{dbSet 0 "russianRoulette" (sdict "storage" $sDB)}}
							{{$bal = add $bal $amt}}
						{{else}}
							{{$em.Set "description" (print "You had no winning to collect!")}}
							{{$em.Set "color" $errorColor}}
						{{end}}
					{{end}}
				{{end}}
			{{else}}
				{{$em.Set "description" (print "No `russianRoulette` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
				{{$em.Set "color" $errorColor}}
			{{end}}
		{{else}}
			{{$em.Set "description" (print "No `bet` argument provided.\nSyntax is `" $.Cmd " <Bet>`")}}
			{{$em.Set "color" $errorColor}}
		{{end}}
		{{dbSet $userID "cash" $bal}}
	{{else}}
		{{$em.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
		{{$em.Set "color" $errorColor}}
	{{end}}
{{else}}
	{{if eq .ExecData "cancel"}}
		{{$em.Set "description" (print "Not enough players joined for Russian-roulette")}}
		{{$em.Set "color" $errorColor}}
		{{cancelScheduledUniqueCC .CCID "rr-game"}}
	{{else}}
		{{$em.Set "description" (print "The host took too long to start the game. Please start a new one.")}}
		{{$em.Set "color" $errorColor}}
	{{end}}
	{{with dbGet 0 "russianRoulette"}}
		{{$a := sdict .Value}}
		{{$cost := $a.game.cost}}
		{{$sDB := $a.storage}}
		{{if not $sDB}}
			{{$sDB = sdict}}
		{{end}}
		{{range $a.game.players}}
			{{$amt := ($sDB.Get (toString .))}}
			{{if $amt}}
				{{$amt = add $amt $cost}} 
			{{else}}
				{{$amt = $cost}}
			{{end}}
			{{$sDB.Set (toString .) $amt}}
		{{end}}
		{{dbSet 0 "russianRoulette" (sdict "storage" $sDB)}}
	{{end}}
	{{cancelScheduledUniqueCC .CCID "rr-game-2"}}
{{end}}
{{sendMessage nil (cembed $em)}}