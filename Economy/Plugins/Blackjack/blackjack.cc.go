{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)()(\s+|\z)`

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

{{/* Blackjack */}}

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
{{if .ExecData}}
    {{$data := (dbGet 0 "bj").Value}}
	{{$embed := structToSdict (index (getMessage nil $data.embed).Embeds 0) }}
	{{range $k, $v := $embed }}
		{{- if eq (kindOf $v true) "struct"}}
			{{- $embed.Set $k (structToSdict $v)}}
		{{- end -}}
	{{end}}
	{{if $embed.Author}}
		{{$embed.Author.Set "Icon_URL" $embed.Author.IconURL}}
	{{end}}
	{{$embed.Set "description" (print "Loss: " $symbol $data.bet)}}
    {{$embed.Set "color" $errorColor}}
    {{$components := (cslice (cbutton "label" "Hit" "custom_id" "bj_hit" "style" "primary" "disabled" true) (cbutton "label" "Stand" "custom_id" "bj_stand" "style" "success" "disabled" true) (cbutton "label" "Double Down" "custom_id" "bj_double" "style" "secondary" "disabled" true))}}
	{{editMessage nil $data.embed (complexMessageEdit "embed" (cembed $embed) "components" $components)}}
    {{return}}
{{end}}
{{if (dbGet 0 "bj")}}
    {{$embed.Set "description" (print "There is currently a blackjack game running. Please wait till the dealer isn't busy.")}}
	{{sendMessage nil (cembed $embed)}}
    {{return}}
{{end}}
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
{{if eq $bet "all"}}{{$bet = $bal}}{{else if eq $bet "max"}}{{$bet = $betMax}}{{end}}
{{if lt ($bet = toInt $bet) 0}}
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
{{define "cardSetup"}}
	{{$availableSuits := cslice "H" "D" "S" "C"}}
	{{$availableCards := cslice "A" "2" "3" "4" "5" "6" "7" "8" "9" "10" "J" "Q" "K"}}
    {{$return := sdict "player" (sdict "cardValue" 0 "cards" cslice "cardPrint" cslice) "dealer" (sdict "cardValue" 0 "cards" cslice "cardPrint" cslice "cardValueReal" 0 "cardPrintReal" cslice) "used" cslice}}
    {{$selectedCards := cslice}}
    {{$selectedEmojiCards := cslice}}
    {{range 4}}
    {{$returnValue := $return.player}}{{if eq . 2}}{{$returnValue = $return.dealer}}{{$selectedCards = cslice}}{{$selectedEmojiCards = cslice}}{{else if eq . 3}}{{$returnValue = $return.dealer}}{{end}}
        {{$card := ""}}
        {{while or (eq $card "") (in $return.used $card)}}
            {{$card = (print (index $availableCards (randInt (len $availableCards))) (index $availableSuits (randInt (len $availableSuits))))}}
        {{end}}
        {{- $selectedCards = $selectedCards.Append $card -}}
        {{$value := $returnValue.cardValue | toInt}}
        {{$cardValue := index (split $card "") 0}}
        {{if in (cslice "J" "Q" "K" "1") $cardValue}}{{$cardValue = 10}}{{end}}
        {{if eq "A" (str $cardValue)}}
            {{$cardValue = 11}}
            {{if eq . 0 2}}
                {{$returnValue.Set "aceStart" true}}
            {{end}}
            {{if eq (len $selectedCards) 2}}
                {{$cardOne := index (split (index $selectedCards 0) "") 0}}
                {{$cardTwo := index (split (index $selectedCards 1) "") 0}}
                {{if and (eq $cardOne "A") (eq $cardTwo "A")}}{{$cardValue = 1}}{{end}}
            {{end}}
        {{end}}
        {{$value = add $value $cardValue}}
        {{$cardGames := (dbGet 0 "cardGames").Value}}
        {{$emojiCard := $cardGames.cards.Get $card}}
        {{- $selectedEmojiCards = $selectedEmojiCards.Append $emojiCard}}
        {{if eq (len $selectedEmojiCards) 2}}
            {{$returnValue.Set "cardPrint" (print (index $selectedEmojiCards 0) " " (index $selectedEmojiCards 1))}}
        {{end}}
        {{if eq . 3}}
            {{$returnValue.Set "cardPrintReal" (print (index $selectedEmojiCards 0) " " (index $selectedEmojiCards 1))}}
            {{$returnValue.Set "cardValueReal" $value}}
            {{$returnValue.Set "cardPrint" (print (index $selectedEmojiCards 0) " " ($cardGames.cards.Get "CB"))}}
            {{$value = sub $value $cardValue}}
        {{end}}
        {{- $returnValue.Set "cards" $selectedCards -}}
        {{- $returnValue.Set "cardValue" $value -}}
        {{- $return.Set "used" ($return.used.Append $card) -}}
    {{end}}
    {{return $return}}
{{end}}
{{$bal = sub $bal $bet}}
{{$cardSelection := execTemplate "cardSetup"}}
{{$cardOne := index (split (index $cardSelection.player.cards 0) "") 0}}
{{$cardTwo := index (split (index $cardSelection.player.cards 1) "") 0}}
{{$components := cslice}}
{{$embed.Set "fields" (cslice (sdict "name" "Your hand" "value" $cardSelection.player.cardPrint "inline" true) (sdict "name" "⠀⠀" "value" "⠀⠀" "inline" true) (sdict "name" "Dealers hand" "value" $cardSelection.dealer.cardPrint "inline" true) (sdict "name" "⠀⠀" "value" (print "Value: " $cardSelection.player.cardValue ) "inline" true) (sdict "name" "⠀⠀" "value" "⠀⠀" "inline" true) (sdict "name" "⠀⠀" "value" (print "Value: " $cardSelection.dealer.cardValue ) "inline" true))}}
{{$dealerValue := $cardSelection.dealer.cardValueReal}}
{{if eq $cardSelection.player.cardValue 21}}
    {{if eq $dealerValue 21}}
        {{$dealerValue = "Blackjack"}}
        {{$embed.Set "description" (print "Result: Push, money back")}}
        {{$embed.Set "color" 0xA25D2D}}
        {{$bal = add $bal $bet}}
    {{else}}
        {{$embed.Set "description" (print "Result: Win " $symbol $bet)}}
        {{$embed.Set "color" $successColor}}
        {{$bal = add $bal (mult $bet 2.5)}}
    {{end}}
    {{$embed.Set "fields" (cslice (sdict "name" "Your hand" "value" $cardSelection.player.cardPrint "inline" true) (sdict "name" "⠀⠀" "value" "⠀⠀" "inline" true) (sdict "name" "Dealers hand" "value" $cardSelection.dealer.cardPrintReal "inline" true) (sdict "name" "⠀⠀" "value" (print "Value: Blackjack") "inline" true) (sdict "name" "⠀⠀" "value" "⠀⠀" "inline" true) (sdict "name" "⠀⠀" "value" (print "Value: " $dealerValue ) "inline" true))}}
    {{sendMessage nil (cembed $embed)}}
    {{dbSet $userID "cash" $bal}}
    {{return}}a
{{else}}
    {{$desc := print "`hit` - take another card\n`stand` - end your turn"}}
    {{$components = cslice (cbutton "label" "Hit" "custom_id" "bj_hit" "style" "primary") (cbutton "label" "Stand" "custom_id" "bj_stand" "style" "success")}}
    {{if and (gt $bal $bet)}}
        {{$desc = print $desc "\n`double down` - double your bet, hit once and stand"}}
        {{$components = $components.Append (cbutton "label" "Double Down" "custom_id" "bj_double" "style" "secondary")}}
    {{end}}
    {{$embed.Set "description" $desc}}
{{end}}
{{$embed.Set "color" $successColor}}
{{$x := sendMessageRetID nil (complexMessage "embed" (cembed $embed) "components" $components)}}
{{scheduleUniqueCC $.CCID nil 120 "cancel" "bj"}}
{{dbSetExpire $userID "bj" (sdict "embed" $x "player" $cardSelection.player "dealer" $cardSelection.dealer "usedCards" $cardSelection.used "bet" $bet "user" $userID "ccID" .CCID) 125}}
{{dbSet $userID "cash" $bal}}