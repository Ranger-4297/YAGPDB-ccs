{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(roulette)(\s+|\z)`

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

{{/* Roulette */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "1024")) "timestamp" currentTime "color" $errorColor}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{$betMax := toInt $economySettings.betMax}}
{{$bal := or (dbGet $userID "cash").Value 0 | toInt}}
{{$roulette := (dbGet 0 "roulette").Value}}
{{if not $roulette}}
	{{$embed.Set "description" (print "No `roulette` database found.\nPlease set it up with the default values using `" $prefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$game := or $roulette.game sdict}}
{{$numbers := seq 1 37}}
{{$d1 := seq 1 13}}
{{$d2 := seq 13 25}}
{{$d3 := seq 25 37}}
{{$c1 := cslice 1 4 7 10 13 16 19 22 25 28 31 34}}
{{$c2 := cslice 2 5 8 11 14 17 20 23 26 29 32 35}}
{{$c3 := cslice 3 6 9 12 15 18 21 24 27 30 33 36}}
{{$h1 := seq 1 19}}
{{$h2 := seq 19 37}}
{{$red := cslice 1 3 5 7 9 12 14 16 18 19 21 23 25 27 30 32 34 36}}
{{$black := cslice 2 4 6 8 10 11 13 15 17 20 22 24 26 28 29 31 33 35}}
{{$even := cslice 2 4 6 8 10 12 14 16 18 20 22 24 26 28 30 32 34 36}}
{{$odd := cslice 1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35}}
{{if .ExecData}}
	{{cancelScheduledUniqueCC .CCID "r-game"}}
	{{$winners := sdict}}
	{{$land := randInt 37 | str}}
	{{range $k, $v := $game}}
		{{$pay := 0}}
		{{range $v.bets}}
			{{if eq .space $land}}
				{{$pay = add (mult .bet 35) $pay}}
			{{else if (or (and (eq (str .space) "1-12") (in $d1 $land)) (and (eq (str .space) "13-24") (in $d2 $land)) (and (eq (str .space) "25-36") (in $d3 $land)) (and (eq (str .space) "1st") (in $c1 $land)) (and (eq (str .space) "2nd") (in $c2 $land)) (and (eq (str .space) "3rd") (in $c3 $land)))}}
				{{$pay = add (mult .bet 3) $pay}}
			{{else if (or (and (eq (str .space) "1-18") (in $h1 $land)) (and (eq (str .space) "19-36") (in $c2 $land)) (and (eq (str .space) "even") (in $even $land)) (and (eq (str .space) "odd") (in $odd $land)) (and (eq (str .space) "red") (in $red $land)) (and (eq (str .space) "black") (in $black $land)))}}
				{{$pay = add (mult .bet 2) $pay}}
			{{end}}
			{{if $pay}}
				{{$winners.Set $k $pay}}
			{{end}}
		{{end}}
	{{end}}
	{{$space := ""}}
	{{if in $black $land}}
		{{$space = "black"}}
	{{else if in $red $land}}
		{{$space = "red"}}
	{{end}}
	{{$fields := cslice}}
	{{$storageDB := $roulette.storage}}
	{{if not $storageDB}}
		{{$storageDB = sdict}}
	{{end}}
	{{range $k, $v := $winners}}
		{{$roulettemt := $storageDB.Get (toString $k)}}
		{{if $roulettemt}}
			{{$roulettemt = add $roulettemt $v}}
		{{else}}
			{{$roulettemt = $v}} 
		{{end}}
		{{- $fields = $fields.Append (sdict "name" "⠀⠀" "value" (print (userArg $k) " won " $symbol $v)) -}}
		{{- $storageDB.Set (toString $k) $roulettemt -}}
	{{end}}
	{{$embed.Set "color" $successColor}}
	{{if not $fields}}
		{{- $fields = cslice (sdict "name" (print "⠀⠀") "value" (print "**No winners**")) -}}
		{{- $embed.Set "color" $errorColor -}}
	{{end}}
	{{$embed.Set "description" (print "The ball landed on **" $space " " $land "**\n**Winners:**")}}
	{{$embed.Set "fields" $fields}}
	{{dbSet 0 "roulette" (sdict "storage" $storageDB "game" sdict)}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{if not .CmdArgs }}
	{{$embed.Set "description" (print "No `Space` argument provided.\nSyntax is `" $.Cmd " <Space> <Bet:Amount>`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$side := index .CmdArgs 0}}
{{if or (in $numbers (toInt $side)) (eq (str $side) "1-12") (eq (str $side) "13-24") (eq (str $side) "25-36") (eq (lower (str $side)) "even") (eq (lower (str $side)) "odd") (eq (lower (str $side)) "red") (eq (lower (str $side)) "black") (eq (lower (str $side)) "1st") (eq (lower (str $side)) "2nd") (eq (lower (str $side)) "3rd") (eq (str $side) "1-18") (eq (str $side) "19-36")}}
	{{if lt (len .CmdArgs 2)}}
		{{$embed.Set "description" (print "No `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount> <Space>`")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{$bet := index .CmdArgs 1 | str | lower}}
	{{if eq $bet "all"}}
		{{$bet = $bal}}
	{{else if eq $bet "max"}}
		{{$bet = $betMax}}
	{{end}}
	{{if le ($bet = toInt $bet) 0}}
		{{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" .Cmd " <Bet:Amount>`")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{if gt $bet $bal}}
		{{$embed.Set "description" (print "You can't bet more than you have!")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{if gt $bet $betMax}}
		{{$embed.Set "description" (print "You can't bet more than " $symbol $betMax)}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{with $game.Get (toString $userID)}}
		{{.Set "betNo" (add .betNo 1)}}
		{{.bets.Set .betNo (sdict "bet" $bet "space" $side)}}
	{{else}}
		{{if not $game}}
			{{scheduleUniqueCC $.CCID nil 30 "r-game" "start"}}
			{{$embed.Set "footer" (sdict "text" (print "The game will start in 30s"))}}
		{{end}}
		{{$game.Set (toString $userID) (sdict "betNo" 1 "bets" (dict 1 (sdict "bet" $bet "space" $side)))}}
		{{$bal = sub $bal $bet}}
		{{$roulette.Set "game" $game}}
	{{end}}
	{{$embed.Set "description" (print "You placed a bet of " $symbol $bet " on " $side "!")}}
	{{$embed.Set "color" $successColor}}
	{{dbSet 0 "roulette" $roulette}}
{{else if eq $side "info"}}
	{{$embed.Set "description" (print "**Payout:**\n[x35] Single number\n[x3] Dozens (1-12, 13-24, 25-36)\n[x3] Columns (1st, 2nd, 3rd)\n[x2] Halves (1-18, 19-36)\n[x2] Odd/Even\n[2x] Colours (red, black)")}}
	{{$embed.Set "image" (sdict "url" "https://raw.githubusercontent.com/Ranger-4297/YAGPDB-ccs/main/Economy/Plugins/Roulette/roulette-board.png")}}
	{{$embed.Set "color" $successColor}}
{{else if eq $side "collect"}}
	{{$storageDB := $roulette.storage}}
	{{$pay := $storageDB.Get (toString $userID)}}
	{{if $pay}}
		{{$embed.Set "description" (print "You've collected " $symbol $pay)}}
		{{$embed.Set "color" $successColor}}
		{{$storageDB.Del (toString $userID )}}
		{{dbSet 0 "roulette" (sdict "storage" $storageDB)}}
		{{$bal = add $bal $pay}}
	{{else}}
		{{$embed.Set "description" (print "You had no winning to collect!")}}
		{{$embed.Set "color" $errorColor}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "Invalid `Space` argument provided.\nSyntax is `" $.Cmd " <Space> <Bet:Amount>`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{dbSet $userID "cash" $bal}}
{{sendMessage nil (cembed $embed)}}