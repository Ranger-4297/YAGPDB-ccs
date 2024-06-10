{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Component`
	Trigger: `\Abj_`

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

{{/* Blackjack component registry */}}

{{/* Response */}}
{{$data := (dbGet $userID "bj").Value}}
{{if not $data}}{{return}}{{end}}
{{if not (eq $data.user .User.ID)}}{{return}}{{end}}
{{$embed := structToSdict (index (getMessage nil $data.embed).Embeds 0)}}{{range $k, $v := $embed }}{{- if eq (kindOf $v true) "struct"}}{{- $embed.Set $k (structToSdict $v)}}{{- end -}}{{end}}
{{$embed.Author.Set "Icon_URL" (.User.AvatarURL "1024")}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}{{sendMessage nil (cembed $embed)}}{{return}}{{end}}
{{$symbol := $economySettings.symbol}}
{{$bal := toInt (dbGet (toInt $userID) "cash").Value}}
{{define "hit"}}
	{{$userData := .data.Get .user}}
	{{$userCards := $userData.cards}}
	{{$userPrintCards := cslice}}
	{{$suits := cslice "H" "D" "S" "C"}}
	{{$cards := cslice "A" "2" "3" "4" "5" "6" "7" "8" "9" "10" "J" "Q" "K"}}
	{{$card := ""}}
	{{$used := .data.usedCards}}
	{{while or (eq $card "") (in $used $card)}}{{$card = (print (index $cards (randInt (len $cards))) (index $suits (randInt (len $suits))))}}{{end}}
	{{$userCards = $userCards.Append $card}}
	{{$used = $used.Append $card}}
	{{$value := $userData.cardValue | toInt}}
	{{$cardValue := index (split $card "") 0}}
	{{if eq $cardValue "J" "Q" "K" "1"}}{{$cardValue = 10}}{{end}}
	{{if eq "A" (str $cardValue)}}{{$cardValue = 11}}{{end}}
	{{if and (eq (toInt $cardValue) 11) (gt (add $value $cardValue) 21)}}{{$cardValue = 1}}{{end}}
	{{if and (eq (len $userCards) 3) (or (eq (index (split (index $userCards 0) "") 0) "A") (eq (index (split (index $userCards 1) "") 0) "A"))}}{{$cardValue = sub $cardValue 10}}{{end}}
	{{$value = add $value $cardValue}}
	{{$emojiCards := (dbGet 0 "cardGames").Value.cards}}
	{{range $userCards}}{{$userPrintCards = $userPrintCards.Append ($emojiCards.Get .)}}{{end}}
	{{return (sdict "cardValue" $value "cards" $userCards "cardPrint" $userPrintCards "used" $used)}}
{{end}}
{{define "splitHit"}}
	{{return .hand}}
{{end}}
{{define "standCondition"}}
	{{$embed := structToSdict (index (getMessage nil .data.embed).Embeds 0)}}{{range $k, $v := $embed}}{{- if eq (kindOf $v true) "struct"}}{{- $embed.Set $k (structToSdict $v)}}{{- end -}}{{end}}
	{{$playerPrintCards := joinStr " " .data.player.cardPrint}}{{$playerValue := .data.player.cardValue}}
	{{$dealerPrintCards := joinStr " " .data.dealer.cardPrint}}{{$dealerValue := .data.dealer.cardValue}}
	{{$embed.Set "fields" (cslice (sdict "name" "Your hand" "value" $playerPrintCards "inline" true) (sdict "name" "⠀⠀" "value" "⠀⠀" "inline" true) (sdict "name" "Dealers hand" "value" $dealerPrintCards "inline" true) (sdict "name" "⠀⠀" "value" (print "Value: " $playerValue) "inline" true) (sdict "name" "⠀⠀" "value" "⠀⠀" "inline" true) (sdict "name" "⠀⠀" "value" (print "Value: " (reReplace `21` (str $dealerValue) "Blackjack")) "inline" true))}}
	{{if or (and (lt $playerValue 21) (gt $dealerValue $playerValue) (lt $dealerValue 22))}}
		{{$embed.Set "description" (print "Result: Loss -" .symbol .data.bet)}}
		{{$embed.Set "color" 0xFF0000}}
	{{else if gt $playerValue 21}}
		{{$embed.Set "description" (print "Result: Bust -" .symbol .data.bet)}}
		{{$embed.Set "color" 0xFF0000}}
	{{else if or (gt $dealerValue 21) (and (gt $playerValue $dealerValue ) (lt $playerValue 22))}}
		{{$embed.Set "description" (print "Result: Win " .symbol .data.bet)}}
		{{dbSet .data.user "cash" (add .bal (mult .data.bet 2))}}
	{{else if (eq $dealerValue $playerValue)}}
		{{$embed.Set "description" (print "Result: Push, money back")}}
		{{$embed.Set "color" 0xA25D2D}}
		{{dbSet .data.user "cash" (add .bal .data.bet)}}
	{{end}}
	{{updateMessage (complexMessageEdit "embed" (cembed $embed))}}
	{{cancelScheduledUniqueCC .data.ccID "cancel"}}{{dbDel .data.user "bj"}}
	{{return}}
{{end}}
{{if eq .StrippedID "hit"}}
	{{$hit := execTemplate "hit" (sdict "data" $data "user" "player")}}
	{{$player := $data.player}}{{$data.Set "usedCards" $hit.Used}}{{$hit.Del "used"}}{{$player = $hit}}{{$data.Set "player" $player}}
	{{$playerPrintCards := joinStr " " $data.player.cardPrint}}
	{{$embed.Set "fields" (cslice (sdict "name" "Your hand" "value" $playerPrintCards "inline" true) (sdict "name" "⠀⠀" "value" "⠀⠀" "inline" true) (sdict "name" "Dealers hand" "value" $data.dealer.cardPrint "inline" true) (sdict "name" "⠀⠀" "value" (print "Value: " $data.player.cardValue) "inline" true) (sdict "name" "⠀⠀" "value" "⠀⠀" "inline" true) (sdict "name" "⠀⠀" "value" (print "Value: " $data.dealer.cardValue ) "inline" true))}}
	{{if eq $hit.cardValue 21}}
		{{$dealer := $data.dealer}}
		{{$value := $dealer.cardValue | toInt}}
		{{$data.dealer.Set "cardValue" $data.dealer.cardValueReal}}{{$data.dealer.Del "cardValueReal"}}{{$data.dealer.Set "cardPrint" $data.dealer.cardPrintReal}}{{$data.dealer.Del "cardPrintReal"}}
		{{while lt $value 17}}{{$hit = execTemplate "hit" (sdict "data" $data "user" "dealer")}}{{$data.Set "usedCards" $hit.Used}}{{$hit.Del "used"}}{{$dealer = $hit}}{{$data.Set "dealer" $dealer}}{{$value = add $value $hit.cardValue}}{{end}}
		{{$dealerPrintCards := joinStr " " $data.dealer.cardPrint}}
		{{$embed.Set "description" (print "Result: Win " $symbol $data.bet)}}
		{{$embed.Set "fields" (cslice (sdict "name" "Your hand" "value" $playerPrintCards "inline" true) (sdict "name" "⠀⠀" "value" "⠀⠀" "inline" true) (sdict "name" "Dealers hand" "value" $dealerPrintCards "inline" true) (sdict "name" "⠀⠀" "value" (print "Value: " $data.player.cardValue) "inline" true) (sdict "name" "⠀⠀" "value" "⠀⠀" "inline" true) (sdict "name" "⠀⠀" "value" (print "Value: " $data.dealer.cardValue ) "inline" true))}}
		{{updateMessage (complexMessageEdit "embed" (cembed $embed))}}
		{{cancelScheduledUniqueCC $data.ccID "cancel"}}{{dbDel $userID "bj"}}{{dbSet $data.user "cash" (add .bal (mult .data.bet 2))}}
		{{return}}
	{{else if gt $hit.cardValue 21}}
		{{$embed.Set "description" (print "Result: bust -" $symbol $data.bet)}}
		{{$embed.Set "color" $errorColor}}
		{{updateMessage (complexMessageEdit "embed" (cembed $embed))}}
		{{cancelScheduledUniqueCC $data.ccID "cancel"}}{{dbDel $userID "bj"}}
		{{return}}
	{{end}}
	{{$components := (cslice (cbutton "label" "Hit" "custom_id" "bj_hit" "style" "primary") (cbutton "label" "Stand" "custom_id" "bj_stand" "style" "success") (cbutton "label" "Double Down" "custom_id" "bj_double" "style" "secondary" "disabled" true) (cbutton "label" "Split" "custom_id" "bj_split" "style" "secondary" "disabled" true))}}
	{{updateMessage (complexMessageEdit "embed" (cembed $embed) "components" $components)}}
	{{scheduleUniqueCC $data.ccID nil 120 "cancel" "bj"}}{{dbSetExpire $userID "bj" (sdict "embed" $data.embed "player" $data.player "dealer" $data.dealer "usedCards" $data.usedCards "bet" $data.bet "user" .User.ID "ccID" $data.ccID) 120}}
{{else if eq .StrippedID "stand"}}
	{{$data.dealer.Set "cardValue" $data.dealer.cardValueReal}}{{$data.dealer.Del "cardValueReal"}}{{$data.dealer.Set "cardPrint" $data.dealer.cardPrintReal}}{{$data.dealer.Del "cardPrintReal"}}
	{{$dealer := $data.dealer}}
	{{$value := $dealer.cardValue | toInt}}
	{{while lt $value 17}}{{$hit := execTemplate "hit" (sdict "data" $data "user" "dealer")}}{{$data.Set "usedCards" $hit.Used}}{{$hit.Del "used"}}{{$dealer = $hit}}{{$data.Set "dealer" $dealer}}{{$value = add $value $hit.cardValue}}{{end}}
	{{template "standCondition" (sdict "data" $data "symbol" $symbol "bal" $bal)}}
{{else if eq .StrippedID "double"}}
	{{$data.Set "bet" (mult $data.bet 2)}}
	{{$hit := execTemplate "hit" (sdict "data" $data "user" "player")}}
	{{$player := $data.player}}{{$data.Set "usedCards" $hit.Used}}{{$hit.Del "used"}}{{$player = $hit}}{{$data.Set "player" $player}}
	{{if gt $data.player.cardValue 21}}
		{{template "standCondition" (sdict "data" $data "symbol" $symbol "bal" $bal)}}
		{{return}}
	{{end}}
	{{$data.dealer.Set "cardValue" $data.dealer.cardValueReal}}{{$data.dealer.Del "cardValueReal"}}{{$data.dealer.Set "cardPrint" $data.dealer.cardPrintReal}}{{$data.dealer.Del "cardPrintReal"}}
	{{$dealer := $data.dealer}}
	{{$value := $dealer.cardValue | toInt}}
	{{while lt $value 17}}{{$hit := execTemplate "hit" (sdict "data" $data "user" "dealer")}}{{$data.Set "usedCards" $hit.Used}}{{$hit.Del "used"}}{{$dealer = $hit}}{{$data.Set "dealer" $dealer}}{{$value = add $value $hit.cardValue}}{{end}}
	{{template "standCondition" (sdict "data" $data "symbol" $symbol "bal" $bal)}}
{{end}}