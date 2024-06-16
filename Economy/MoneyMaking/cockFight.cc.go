{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)((cock|chicken)-?fight)(\s+|\z)`

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

{{/* Cock fight */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" User.Username "icon_url" (User.AvatarURL "1024")) "timestamp" currentTime "color" $errorColor}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{$betMax := $economySettings.betMax | toInt}}
{{$incomeCooldown := $economySettings.incomeCooldown | toInt}}
{{if not .CmdArgs}}
    {{$embed.Set "description" (print "No `bet` argument provided.\nSyntax is `" .Cmd " <Bet:Amount>`")}}
	{{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
{{$bal := toInt (dbGet $userID "cash").Value}}
{{$bet := index .CmdArgs 0 | str | lower}}
{{if not (or (toInt $bet) (eq $bet "all" "max"))}}
    {{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" .Cmd " <Bet:Amount>`")}}
    {{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
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
{{$econData := or (dbGet $userID "userEconData").Value (sdict "settings" sdict "inventory" sdict "streaks" (sdict "daily" 0 "weekly" 0 "monthly" 0) "cfWinChance" 50)}}
{{$inventory := $econData.inventory}}
{{if not ($inventory.Get "chicken")}}
	{{$embed.Set "description" (print "You can't bet without a chicken! Buy one from the shop with `" $prefix "buy-item chicken`")}}
	{{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
{{if ($cooldown := dbGet $userID "cockFightCooldown")}}
	{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
	{{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
{{$cfWC := or ($econData.Get "cfWinChance") 50}}
{{$chance := randInt 1 101}}
{{if le $chance $cfWC}}
	{{if not (eq $cfWC 70)}}
		{{$cfWC = $cfWC | add 1}}
	{{end}}
	{{$bal = add $bal $bet}}
	{{$embed.Set "description" (print "Your chicken won the fight and got you " $symbol $bet "!\nPlay again with `" Cmd " <Bet:Amount>`")}}
	{{$embed.Set "color" $successColor}}
{{else}}
	{{$cfWC = 50}}
	{{$embed.Set "description" (print "Your chicken lost the fight, " $symbol $bet " and died :(\nBuy a new one to play again!")}}
	{{$bal = sub $bal $bet}}
	{{$inventory.Del "chicken"}}
	{{$econData.Set "inventory" $inventory}}
{{end}}
{{$econData.Set "cfWinChance" $cfWC}}
{{dbSet $userID "cash" $bal}}
{{dbSet $userID "userEconData" $econData}}
{{dbSetExpire $userID "cockFightCooldown" "cooldown" $incomeCooldown}}
{{sendMessage nil (cembed $embed)}}