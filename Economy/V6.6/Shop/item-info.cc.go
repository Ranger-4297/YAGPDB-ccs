{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(item-?info|view-?item)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}

{{/* Item information */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{with (dbGet 0 "store")}}
		{{$info := sdict .Value}}
		{{$items := sdict}}
		{{if ($info.Get "Items")}}
			{{$items = sdict ($info.Get "Items")}}
			{{with $.CmdArgs}}
				{{$name := (index . 0)}}
				{{if $items.Get $name}}
					{{$item := $items.Get (index . 0)}}
					{{$price := $item.Get "price"}}
					{{$role := "none"}}
					{{if ($item.Get "role-given")}}
						{{$role = (print "<@&" ($item.Get "role") ">")}}
					{{end}}
					{{$qty := ""}}
					{{if ($item.Get "qty")}}
						{{$qty = ($item.Get "quantity")}}
						{{if not (reFind "infinite" (lower (toString $qty)))}}
							{{$qty = toInt $qty | humanizeThousands}}
						{{else}}
							{{$qty = "Infinite"}}
						{{end}}
					{{else}}
						{{$qty = "Infinite"}}
					{{end}}
					{{$desc := $item.Get "desc"}}
					{{$embed.Set "title" (print "**Item info**")}}
					{{$embed.Set "fields" (cslice (sdict "name" "Name" "value" (print $name) "inline" true) (sdict "name" (print "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀") "value" (print "⠀⠀") "inline" true) (sdict "name" "Price" "value" (print $symbol (humanizeThousands $price)) "inline" true) (sdict "name" "Description" "value" (print $desc) "inline" true) (sdict "name" (print "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀") "value" (print "⠀⠀") "inline" true) (sdict "name" "Quantity" "value" (print $qty) "inline" true) (sdict "name" "Role given" "value" (print $role)))}}
					{{$embed.Set "color" $successColor}}
				{{else}}
					{{$embed.Set "description" (print "Invalid item argument provided :(\nSyntax is `" $.Cmd " <Name>`")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{else}}
				{{$embed.Set "description" (print "No item argument provided :(\nSyntax is `" $.Cmd " <Name>`")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else}}
			{{$embed.Set "description" (print "There are no items :(\nAdd some items with `" $prefix "create-item <Name:Word> <Price:Int> <Quantity:Int> <Description:String>`")}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}