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
{{$prefix := .ServerPrefix }}

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
		{{if ($info.Get "items")}}
			{{$items = sdict ($info.Get "items")}}
			{{with $.CmdArgs}}
				{{$name := (index . 0)}}
				{{if $items.Get $name}}
					{{$item := $items.Get (index . 0)}}
					{{$price := $item.Get "price"}}
					{{$desc := $item.Get "desc"}}
					{{$qty := $item.quantity}}
					{{if not $qty}}
						{{$qty = "inf"}}
					{{else}}
						{{$qty = humanizeThousands $qty}}
					{{$role := "none"}}
					{{if ($item.Get "role-given")}}
						{{$role = ($item.Get "role-given")}}
					{{end}}
					{{$reply := $item.replyMsg}}
					{{$exp := $item.expiry}}
					{{if (toDuration $exp)}}
						{{$exp = humanizeDurationSeconds (mult $exp $.TimeSecond)}}
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
						(sdict "name" "Role given" "value" (print $role) "inline" true) 
						(sdict "name" "Reply message" "value" $reply)
						(sdict "name" "Inventory expiry" "value" $exp)
					)}}
					{{if $user}}
						{{$fields := $embed.fields}}
						{{$fields = $fields.Append (sdict "name" "On market from" "value" (print "<@!" $user ">"))}}
						{{$embed.Set "fields" $fields}}
					{{end}}
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
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "server-set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}