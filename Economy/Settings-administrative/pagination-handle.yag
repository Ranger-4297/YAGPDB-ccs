{{/*
Made by ranger_4297 (765316548516380732)

Trigger Type: `Component`
Trigger: `\Aeconomy_`

©️ RhykerWells 2020-Present
GNU, GPLV3 License
Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$ex := or (and (reFind "a_" .Guild.Icon) "gif") "png"}}
{{$icon := print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon "." $ex "?size=1024"}}

{{/* Pagination */}}

{{/* Response */}}
{{if ne $userID .Message.ReferencedMessage.Author.ID}}
	{{return}}
{{end}}
{{$author := ""}}
{{$embed := sdict "author" $author "timestamp" currentTime "color" $errorColor}}
{{$buttons := cslice (sdict "label" "previous" "style" "danger" "custom_id" "economy_back" "disabled" true) (sdict "label" "next" "style" "success" "custom_id" "economy_forward" "disabled" true)}}
{{$symbol := (dbGet 0 "EconomySettings").Value.symbol}}
{{$paginationType := reFind `(leaderboard|shop|inventory)$` (index .Interaction.Message.Embeds 0).Author.Name | lower}}
{{$page := reFind `\d+` (index .Interaction.Message.Embeds 0).Footer.Text}}
{{if eq .StrippedID "forward"}}{{$page = add $page 1}}{{else}}{{$page = sub $page 1}}{{end}}
{{$display := ""}}
{{$len := ""}}
{{$stop := ""}}
{{if eq $paginationType "shop"}}
	{{$author = sdict "name" (print .Guild.Name " shop") "icon_url" $icon}}
	{{$items := (dbGet 0 "store").Value.items}}
	{{$entry := cslice}}
	{{$fields := cslice}}
	{{range $k, $v := $items}}
		{{$qty := $v.quantity}}
		{{if $qty}}
			{{$qty = humanizeThousands $qty}}
		{{else}}
			{{$qty = "Infinite"}}
		{{end}}
		{{$entry = $entry.Append (sdict "Name" (print $k " - "  $symbol (humanizeThousands $v.price) " - " $qty) "Value"  $v.desc)}}
	{{end}}
	{{$start := mult 10 (sub $page 1)}}
	{{$stop = mult $page 10}}
	{{if ge $stop (len $entry)}}
		{{$stop = (len $entry)}}
	{{end}}
	{{range (seq $start $stop)}}
		{{$fields = $fields.Append (index $entry .)}}
	{{end}}
	{{$display = (print "Buy an item with `buy-item <Name> [Quantity:Int]`\nFor more information on an item use `item-store <Name>`.")}}
	{{$embed.Set "fields" $fields}}
	{{$len = len $items}}
{{else if eq $paginationType "inventory"}}
	{{$user := .User}}
	{{$userCheck := split .Message.ReferencedMessage.Content " "}}
	{{if gt (len $userCheck) 1}}
		{{if $newUser := (index $userCheck 1) | userArg}}
			{{$user = $newUser}}
		{{end}}
	{{end}}
	{{$author = sdict "name" (print $user.Username " inventory") "icon_url" ($user.AvatarURL "1024")}}
	{{$userdata := or (dbGet $user.ID "userEconData").Value (sdict "inventory" sdict)}}
	{{$inventory := $userdata.inventory}}
	{{$entry := cslice}}
	{{$fields := cslice}}
	{{range $k, $v := $inventory}}
		{{$role := $v.role}}
		{{if $role}}
			{{$role = print "<@&" $role ">"}}
		{{else}}
			{{$role = "none"}}
		{{end}}
		{{$expiry := $v.expiry}}
		{{$expires := "never"}}
		{{if $expiry}}
			{{$timeSeconds := toDuration (humanizeDurationSeconds (mult $expiry .TimeSecond))}}
			{{$expires = (print "<t:" (currentTime.Add $timeSeconds).Unix ":f>")}}
		{{end}}
		{{$entry = $entry.Append (sdict "Name" $k "Value" (print "Description: " $v.desc "\nQuantity: " (humanizeThousands $v.quantity) "\nRole given: " $role "\nExpiry: " $expires))}}
	{{end}}
	{{$start := mult 10 (sub $page 1)}}
	{{$stop = mult $page 10}}
	{{if ge $stop (len $entry)}}
		{{$stop = (len $entry)}}
	{{end}}
	{{range (seq $start $stop)}}
		{{$fields = $fields.Append (index $entry .)}}
	{{end}}
	{{$embed.Set "fields" $fields}}
	{{$len = len $inventory}}
{{else}}
	{{$author = sdict "name" (print .Guild.Name " leaderboard") "icon_url" $icon}}
	{{$rank := mult (sub $page 1) 10}}
	{{$len = dbCount "cash"}}
	{{$displayUsers := dbTopEntries "cash" 10 $rank}}
	{{$pos := dict 1 "🥇" 2 "🥈" 3 "🥉"}}
	{{$dRank := $rank}}
	{{range $displayUsers}}
		{{$cash := humanizeThousands (toInt .Value)}}
		{{$rank = add $rank 1}}
		{{$dRank = $rank}}
		{{if in (cslice 1 2 3) $rank}}
			{{- $dRank = $pos.Get $rank -}}
		{{else}}
			{{$dRank = print "  " $rank "."}}
		{{end}}
		{{$display = (print $display "**" $dRank "** " .User.String  " **•** " $symbol $cash "\n")}}
	{{end}}
	{{$stop = $rank}}
{{end}}
{{if gt $len $stop}}
	{{(index $buttons 1).Set "disabled" false}}
{{end}}
{{if ne $page 1}}
	{{(index $buttons 0).Set "disabled" false}}
{{end}}
{{$embed.Set "author" $author}}
{{$embed.Set "description" $display}}
{{$embed.Set "footer" (sdict "text" (print "Page: " $page))}}
{{$embed.Set "color" $successColor}}
{{updateMessage (complexMessageEdit "embed" $embed "buttons" $buttons)}}