{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(remove(-?item)?)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix }}

{{/* Buy item */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
	{{$a := sdict .Value}}
	{{with (dbGet 0 "store")}}
		{{$info := sdict .Value}}
		{{$items := sdict}}
		{{if ($info.Get "items")}}
			{{$items = sdict ($info.Get "items")}}
			{{with $.CmdArgs}}
				{{$name := (index . 0)}}
				{{if $items.Get $name}}
                    {{$items.Del $name}}
					{{$info.Set "items" $items}}
					{{dbSet 0 "store" $info}}
					{{$embed.Set "description" (print "Removed " $name " successfully!")}}
					{{$embed.Set "color" $successColor}}
				{{else}}
					{{$embed.Set "description" (print "This item doesn't exist :( Create it with `" $prefix "create-item " $name "`\n\nTo view all items, run the `" $prefix "shop` command.")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{else}}
				{{$embed.Set "description" (print "No item argument provided :(\nSyntax is `" $.Cmd " <Name>`\n\nTo view all items, run the `" $prefix "shop` command.")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else}}
			{{$embed.Set "description" (print "There are no items :(\nAdd some items with `" $prefix "create-item`")}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "server-set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}