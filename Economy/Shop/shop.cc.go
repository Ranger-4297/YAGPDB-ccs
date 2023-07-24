{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(s(tore|hop))(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix }}
{{$ex := or (and (reFind "a_" .Guild.Icon) "gif") "png" }}
{{$icon := print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon "." $ex "?size=1024" }}

{{/* Server shop */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" (print .Guild.Name " store") "icon_url" $icon)}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{with (dbGet 0 "store")}}
		{{$info := sdict .Value}}
		{{$items := sdict}}
		{{if ($info.Get "items")}}
			{{$items = sdict ($info.Get "items")}}
			{{$entry := cslice}}
			{{$field := cslice}}
			{{if $items}}
				{{range $k,$v := $items}}
					{{$item := $k}}
					{{$desc := $v.desc }}
					{{$price := (print $symbol ($v.price))}}
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
					{{$user := $v.user}}
					{{if $user}}
						{{$entry = $entry.Append (sdict "Name" (print $item " - " $price " - " $qty " - <@!" $user ">") "value"  $desc "inline" false)}}
					{{else}}
						{{$entry = $entry.Append (sdict "Name" (print $item " - " $price " - " $qty) "value"  $desc "inline" false)}}
					{{end}}
				{{end}}
				{{$page := ""}}
				{{if $.CmdArgs}}
					{{$page = (index $.CmdArgs 0) | toInt}}
					{{if lt $page 1}}
						{{$page = 1}}
					{{end}}
				{{else}}
					{{$page = 1}}
				{{end}}
				{{$start := (mult 10 (sub $page 1))}}
				{{$stop := (mult $page 10)}}
				{{if ge $stop (len $entry)}}
					{{$stop = (len $entry)}}
				{{end}}
				{{if and (le $start $stop) (ge (len $entry) $start) (le $stop (len $entry))}}
					{{range (seq $start $stop)}}
						{{$field = $field.Append (index $entry .)}}
						{{$embed.Set "description" (print "Buy an item with `buy-item <Name> [Quantity:Int]`\nFor more information on an item use `item-info <Name>`.")}}
						{{$embed.Set "color" $successColor}}
					{{end}}
				{{else}}
					{{$embed.Set "description" (print "There are no items on this page\nAdd some with `create-item <Name:Word> <Price:Int> <Quantity:Int> <Description:String>`")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
				{{$embed.Set "fields" $field}}
				{{$embed.Set "footer" (sdict "text" (print "Page: " $page))}}
			{{else}}
				{{$embed.Set "description" (print "The shop is empty :(\nAdd some items with `create-item <Name:Word> <Price:Int> <Quantity:Int> <Description:String>`")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else}}
			{{$embed.Set "description" (print "The shop is empty :(\nAdd some items with `" $prefix "create-item <Name:Word> <Price:Int> <Quantity:Int> <Description:String>`")}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "server-set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}