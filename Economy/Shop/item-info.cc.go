{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(item-?info|view-?item)(\s+|\z)`

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

{{/* Item information */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" (print .Guild.Name " Store") "icon_url" $icon) "timestamp" currentTime "color" $errorColor}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{if not .CmdArgs}}
	{{$embed.Set "description" (print "No `Item` argument provided.\nSyntax is `" .Cmd " <Item:Name>`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$info := or (dbGet 0 "store").Value (sdict "items" sdict)}}
{{$items := ($info.Get "items")}}
{{$name := (index .CmdArgs 0)}}
{{if not ($items.Get $name)}}
	{{$embed.Set "description" (print "This item doesn't exist")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$item := $items.Get $name}}
{{$price := $item.Get "price"}}
{{$desc := $item.Get "desc"}}
{{$qty := toInt $item.quantity}}
{{if not $qty}}
	{{$qty = "Infinite"}}
{{else}}
	{{$qty = humanizeThousands $qty}}
{{end}}
{{$role := $item.Get "role-given"}}
{{if $role}}
	{{$role = print "<@&" $role ">"}}
{{else}}
	{{$role = "None"}}
{{end}}
{{$reply := $item.replyMsg}}
{{$exp := $item.expiry}}
{{if $exp}}
	{{$exp = humanizeDurationSeconds (mult $exp .TimeSecond)}}
{{else}}
	{{$exp = "Never"}}
{{end}}
{{$user := $item.ID}}
{{$embed.Set "title" (print "**Item info**")}}
{{$embed.Set "fields" (cslice 
	(sdict "name" "Name" "value" (print $name) "inline" true)
	(sdict "name" (print "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀") "value" (print "⠀⠀") "inline" true)
	(sdict "name" "Price" "value" (print $symbol (humanizeThousands $price)) "inline" true)
	(sdict "name" "Description" "value" (print $desc) "inline" true)
	(sdict "name" (print "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀") "value" (print "⠀⠀") "inline" true)
	(sdict "name" "Shop quantity" "value" (print $qty) "inline" true)
	(sdict "name" "Role given" "value" $role "inline" true) 
	(sdict "name" "Reply message" "value" $reply)
	(sdict "name" "Inventory expiry" "value" $exp)
)}}
{{if $user}}
	{{$embed.Set "footer" (sdict "text" (print "On market from"))}}
{{end}}
{{$embed.Set "color" $successColor}}
{{sendMessage nil (cembed $embed)}}