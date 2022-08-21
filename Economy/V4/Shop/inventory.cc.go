{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(inv(entory)?)(\s+|\z)`

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

{{/* Server shop */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a := sdict .Value}}
    {{$symbol := $a.symbol}}
	{{with (dbGet $userID "EconomyInfo")}}
		{{$info := sdict .Value}}
		{{$items := sdict}}
		{{if ($info.Get "inventory")}}
			{{$items = sdict ($info.Get "inventory")}}
			{{$entry := cslice}}
			{{$field := cslice}}
			{{if $items}}
				{{range $k,$v := $items}}
					{{$item := $k}}
					{{$desc := $v.desc }}
					{{$qty := ""}}
					{{$qty = $v.qty}}
					{{$entry = $entry.Append (sdict "Name" $item "value" (joinStr "\n" (print "Description: " $desc) (print "Quantity: " (humanizeThousands $qty))) "inline" false)}}
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
					{{end}}
				{{else}}
					{{$embed.Set "description" (print "This page is empty")}}
				{{end}}
                {{$embed.Set "title" (print "Inventory")}}
				{{$embed.Set "fields" $field}}
				{{$embed.Set "color" $successColor}}
				{{$embed.Set "footer" (sdict "text" (print "Page: " $page))}}
			{{else}}
				{{$embed.Set "description" (print "Your inventory is empty :(\nGet some items from the shop!")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else}}
			{{$embed.Set "description" (print "Your inventory is empty :(\nGet some items from the shop!")}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
	{{end}}
{{else}}
    {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}