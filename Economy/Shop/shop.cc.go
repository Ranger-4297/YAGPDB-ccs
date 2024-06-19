{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(s(tore|hop))(\s+|\z)`

	©️ RhykerWells 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix }}
{{$ex := or (and (reFind "a_" .Guild.Icon) "gif") "png"}}
{{$icon := print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon "." $ex "?size=1024"}}

{{/* Server shop */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" (print .Guild.Name " shop") "icon_url" $icon) "timestamp" currentTime "color" $errorColor}}
{{$buttons := cslice (sdict "label" "previous" "style" "danger" "custom_id" "economy_back" "disabled" true) (sdict "label" "next" "style" "success" "custom_id" "economy_forward" "disabled" true)}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{$store := (dbGet 0 "store").Value}}
{{$items := $store.items}}
{{if not $items}}
	{{$embed.Set "description" (print "The shop is empty\nAdd some items with `" $prefix "create-item`")}}
	{{sendMessage nil (complexMessage "embed" $embed "reply" .Message.ID)}}
	{{return}}
{{end}}
{{$entry := cslice}}
{{$display := ""}}
{{range $item, $itemValue := $items}}
	{{$qty := $itemValue.quantity}}
	{{if $qty}}
		{{$qty = humanizeThousands $qty}}
	{{else}}
		{{$qty = "Infinite"}}
	{{end}}
	{{$entry = $entry.Append (print "**" $item " - " $symbol (humanizeThousands $itemValue.price) " - " $qty "**\n" $itemValue.desc)}}
{{end}}
{{$page := 1}}
{{if .CmdArgs}}
	{{$page = (index .CmdArgs 0) | toInt}}
	{{if lt $page 1}}
		{{$page = 1}}
	{{end}}
{{end}}
{{$start := mult 10 (sub $page 1)}}
{{$stop := mult $page 10}}
{{if ge $stop (len $entry)}}
	{{$stop = (len $entry)}}
{{end}}
{{if not (and (le $start $stop) (ge (len $entry) $start) (le $stop (len $entry)))}}
	{{$embed.Set "description" (print "There are no items on this page\nAdd some with `create-item`")}}
{{else}}
	{{range (seq $start $stop)}}
		{{$display = (print $display (index $entry .) "\n")}}
	{{end}}
	{{$embed.Set "description" (print "Buy an item with `buy-item <Name> [Quantity:Int]`\nFor more information on an item use `item-store <Name>`.\n\n" $display)}}
	{{$embed.Set "color" $successColor}}
{{end}}
{{if gt (len $items) $stop}}
	{{(index $buttons 1).Set "disabled" false}}
{{end}}
{{if ne $page 1}}
	{{(index $buttons 0).Set "disabled" false}}
{{end}}
{{$embed.Set "footer" (sdict "text" (print "Page: " $page))}}
{{sendMessage nil (complexMessage "reply" .Message.ID "embed" $embed "buttons" $buttons)}}