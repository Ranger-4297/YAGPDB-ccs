{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(russian-?roulette|rr)(\s+|\z)`

	©️ RhykerWells 2020-Present
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
{{$embed := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "1024")) "timestamp" currentTime "color" $errorColor}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{$betMax := $economySettings.betMax}}
{{if not (or .CmdArgs (eq (str .ExecData) "start" "cancel" "cancel2"))}}
    {{$embed.Set "description" (print "No `bet` argument provided.\nSyntax is `" .Cmd " <Bet:Amount>`")}}
	{{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
{{$russianRoulette := (dbGet 0 "russianRoulette")}}
{{if not $russianRoulette}}
    {{$embed.Set "description" (print "No `russianRoulette` database found.\nPlease set it up with the default values using `" $prefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
{{$russianRoulette = $russianRoulette.Value}}
{{$cost := $russianRoulette.cost}}
{{$game := $russianRoulette.game}}
{{$cost := $game.cost}}
{{$players := $game.players}}
{{if and .ExecData (not (eq (str .ExecData) "start"))}}
	{{if eq .ExecData "cancel"}}
		{{$embed.Set "description" (print "Not enough players joined for Russian-roulette")}}
	{{else if eq .ExecData "cancel2"}}
		{{$embed.Set "description" (print "The host took too long to start the game. Please start a new one.")}}
	{{end}}
	{{sendMessage nil (cembed $embed)}}
	{{$storage := or $russianRoulette.storage sdict}}
	{{range $game.players}}
		{{$storageAmt := add (or ($storage.Get (toString .)) 0) $game.cost}}
		{{$storage.Set (toString .) $storageAmt}}
	{{end}}
	{{$russianRoulette.Set "storage" $storage}}
	{{$russianRoulette.Del "game"}}
	{{dbSet 0 "russianRoulette" $russianRoulette}}
	{{cancelScheduledUniqueCC .CCID "rr-game"}}
	{{return}}
{{end}}
{{if not $russianRoulette.game}}
	{{$bal := toInt (dbGet $userID "cash").Value}}
	{{$bet := index .CmdArgs 0 | str | lower}}
	{{if not (or (toInt $bet) (eq $bet "all" "max" "collect"))}}
		{{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" .Cmd " <Bet:Amount>`")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{if eq $bet "all"}}{{$bet = $bal}}{{else if eq $bet "max"}}{{$bet = $betMax}}{{$bet = str $bet}}{{end}}
	{{if and (not (eq $bet "collect")) (le (toInt $bet) 0)}}
		{{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" .Cmd " <Bet:Amount>`")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{if and (not (eq $bet "collect")) (gt (toInt $bet) $bal)}}
		{{$embed.Set "description" (print "You can't bet more than you have!")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{if and (not (eq $bet "collect")) (gt (toInt $bet) $betMax)}}
		{{$embed.Set "description" (print "You can't bet more than " $symbol $betMax)}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{if $bet = toInt $bet}}
		{{$embed.Set "description" (print "A new game of Russian roulette has been started!\n\nTo join use the command `" Cmd " " $bet "` (1/6)\nTo start this game use the command `" Cmd " start`")}}
		{{$embed.Set "color" $successColor}}
		{{sendMessage nil (cembed $embed)}}
		{{$bal = sub $bal $bet}}
		{{dbSet $userID "cash" $bal}}
		{{$russianRoulette.Set "game" (sdict "cost" $bet "players" (cslice $userID) "owner" $userID)}}
		{{dbSet 0 "russianRoulette" $russianRoulette}}
		{{scheduleUniqueCC CCID nil 300 "rr-game" "cancel"}}
		{{return}}
	{{else}}
		{{$storageDB := $russianRoulette.storage}}
		{{if $storageAmt :=  $storageDB.Get (toString $userID)}}
			{{$embed.Set "description" (print "You've collected " $symbol $storageAmt)}}
			{{$embed.Set "color" $successColor}}
			{{sendMessage nil (cembed $embed)}}
			{{$storageDB.Del (toString $userID)}}
			{{dbSet 0 "russianRoulette" (sdict "storage" $storageDB)}}
			{{$bal = add $bal $storageAmt}}
			{{dbSet $userID "cash" $bal}}
			{{return}}
		{{else}}
			{{$embed.Set "description" (print "You had no winning to collect!")}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
	{{end}}
{{else}}
	{{$rouletteStart := false}}
	{{$bal := toInt (dbGet $userID "cash").Value}}
	{{$bet := or (index .CmdArgs 0 | str | lower) .ExecData}}
	{{if not (or (toInt $bet) (eq $bet "all" "max" "start" "collect"))}}
		{{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" .Cmd " <Bet:Amount>`")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{if eq $bet "all"}}{{$bet = $bal}}{{else if eq $bet "max"}}{{$bet = $betMax}}{{end}}
	{{if and (not (eq $bet "start" "collect")) (le (toInt $bet) 0)}}
		{{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" .Cmd " <Bet:Amount>`")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{if and (not (eq $bet "start" "collect")) (gt (toInt $bet) $bal)}}
		{{$embed.Set "description" (print "You can't bet more than you have!")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{if and (not (eq $bet "start" "collect")) (gt (toInt $bet) $cost)}}
		{{$embed.Set "description" (print "You can't bet more than " $symbol $cost)}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{if eq $bet "start"}}
		{{if or .ExecData (not (eq (toString $userID) (toString $game.owner)))}}
			{{$embed.Set "description" (print "You cannot start the game")}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{cancelScheduledUniqueCC .CCID "rr-game"}}
		{{if not (gt (len $players) 1)}}
			{{$embed.Set "description" (print "Not enough players to start the match :(\nStart a new one with `" Cmd " <bet>`")}}
			{{sendMessage nil (cembed $embed)}}
			{{dbSet 0 "russianRoulette" sdict}}
			{{return}}
		{{end}}
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
			{{$loser = (userArg $p)}}
			{{sendMessage nil (cembed (sdict "description" (print "**" $loser "** pulled the trigger and dies") "color" 0xA25D2D))}}
			{{break}}
		{{end}}
		{{range $players}}
			{{- if ne . $loser.ID}}
				{{- $winners = $winners.Append (userArg .).Mention}}
			{{- end -}}
		{{end}}
		{{$payout := div $game.cost (len $winners)}}
		{{$fields := cslice}}
		{{$storageDB := $russianRoulette.storage}}
		{{if not $storageDB}}
			{{$storageDB = sdict}}
		{{end}}
		{{range $winners}}
			{{$storageAmt := $storageDB.Get (toString (userArg .).ID)}}
			{{if $storageAmt}}
				{{$storageAmt = add $storageAmt $payout}}
			{{else}}
				{{$storageAmt = $payout}} 
			{{end}}
			{{- $fields = $fields.Append (sdict "Name" (print (userArg .)) "value" . "inline" false) -}}
			{{$storageDB.Set (toString (userArg .).ID) $storageAmt}}
		{{end}}
		{{$embed.Set "title" "Winners"}}
		{{$embed.Set "description" (print "payout is: " $symbol $payout " per-person")}}
		{{$embed.Set "fields" $fields}}
		{{$embed.Set "color" $successColor}}
		{{sendMessage nil (cembed $embed)}}
		{{dbSet 0 "russianRoulette" (sdict "storage" $storageDB)}}
	{{else if eq $bet "collect"}}
		{{$embed.Set "description" (print "You cannot collect during a game. Please wait till it is over to collect any owed money.")}}
		{{sendMessage nil (cembed $embed)}}
	{{else if ($bet = toInt $bet)}}
		{{if in $players $userID}}
			{{$embed.Set "description" (print "You're already in this game")}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{$players = $players.Append $userID}}
		{{$game.Set "players" $players}}
		{{dbSet 0 "russianRoulette" $russianRoulette}}
		{{$bal = sub $bal $bet}}
		{{$embed.Set "description" (print "You've joined this game of russian roulette with a bet of " $symbol $bet)}}
		{{$embed.Set "footer" (sdict "text" (print "Players: " (len $players) "/6"))}}
		{{$embed.Set "color" 0x00ff7b}}
		{{if eq (len $players) 6}}
			{{cancelScheduledUniqueCC .CCID "rr-game"}}
			{{execCC .CCID nil 0 "start"}}
		{{else}}
			{{scheduleUniqueCC .CCID nil 60 "rr-game" "cancel2"}}
		{{end}}
		{{sendMessage nil (cembed $embed)}}
		{{dbSet $userID "cash" $bal}}
	{{end}}
{{end}}