{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(use(-?item)?)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}

{{/* Use Item */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
	{{$a := sdict .Value}}
	{{with dbGet $userID "EconomyInfo"}}
		{{$info := sdict .Value}}
		{{if ($info.Get "inventory")}}
			{{$items := sdict ($info.Get "inventory")}}
			{{with $.CmdArgs}}
				{{$name := (index . 0)}}
				{{if $items.Get $name}}
					{{$item := $items.Get (index . 0)}}
					{{$qty := $item.Get "qty"}}
					{{$nqty := (sub (toInt $qty) 1)}}
					{{if eq (toInt $nqty) 0}}
						{{$items.Del $name}}
						{{dbSet $userID "EconomyInfo" $info}}
					{{else}}
						{{$item.Set "qty" $nqty}}
						{{$items.Set $name $item}}
						{{$info.Set "inventory" $items}}
						{{dbSet $userID "EconomyInfo" $info}}
					{{end}}
					{{$embed.Set "description" (print "You've just used " $name "!\nYou had " (humanizeThousands $qty) " and now have " (humanizeThousands $nqty))}}
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
			{{$embed.Set "description" (print "You have no items in your inventory, purchase some from the shop!")}}
		{{end}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}