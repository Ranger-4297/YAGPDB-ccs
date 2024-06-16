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
{{$embed := sdict "author" (sdict "name" (print .Guild.Name " Store")) "timestamp" currentTime "color" $errorColor}}
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
	{{$embed.Set "description" (print "The shop is empty :(\nAdd some items with `" $prefix "create-item`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$entry := cslice}}
{{$field := cslice}}
{{range $k,$v := $items}}
	{{$item := $k}}
	{{$desc := $v.desc }}
	{{$price := print $symbol (humanizeThousands $v.price)}}
	{{$qty := ""}}
	{{if $v.quantity}}
		{{$qty = $v.quantity}}
		{{if not (reFind "inf" (lower (toString $qty)))}}
			{{$qty = toInt $qty | humanizeThousands}}
		{{else}}
			{{$qty = "Infinite"}}
		{{end}}
	{{else}}
		{{$qty = "Infinite"}}
	{{end}}
	{{if $user := $v.ID}}
		{{$entry = $entry.Append (sdict "Name" (print $item " - " $price " - " $qty " - " $user ) "value"  $desc "inline" false)}}
	{{else}}
		{{$entry = $entry.Append (sdict "Name" (print $item " - " $price " - " $qty) "value"  $desc "inline" false)}}
	{{end}}
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
{{if and (le $start $stop) (ge (len $entry) $start) (le $stop (len $entry))}}
	{{range (seq $start $stop)}}
		{{$field = $field.Append (index $entry .)}}
		{{$embed.Set "description" (print "Buy an item with `buy-item <Name> [Quantity:Int]`\nFor more information on an item use `item-store <Name>`.")}}
		{{$embed.Set "color" $successColor}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "There are no items on this page\nAdd some with `create-item <Name:Word> <Price:Int> <Quantity:Int> <Description:String>`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{$embed.Set "fields" $field}}
{{$embed.Set "footer" (sdict "text" (print "Page: " $page))}}
{{sendMessage nil (cembed $embed)}}