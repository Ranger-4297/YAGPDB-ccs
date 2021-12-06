{{/*
        Made by Ranger (765316548516380732)

	Trigger Type: `Command`
	Trigger: `setMin`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$args := parseArgs 1 "setMin <value>" (carg "int" "Value")}}
{{$newMin := $args.Get 0}}

{{$max := ""}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a = sdict .Value}}
    {{$max = $a.max}}
    {{if gt $newMin $max}}
		{{$errorEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1020"))
            "description" (print "You cannot set `min` to a value above `max`.\nYour max is set to `£" $max "`")
            "color" 0x00ff8b
            "timestamp" currentTime
            )}}
        {{sendMessage nil $errorEmbed}}
    {{else}}
		{{$updateEmbed := (cembed
			"author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1020"))
			"description" (print "Successfully set `min` to" $min)
			"color" 0x00ff8b
			"timestamp" currentTime
		)}}
		{{sendMessage nil $updateEmbed}}
		{{$sdict := (dbGet 0 "EconomySettings").Value}}
		{{$sdict.Set "min" $newMin}}
		{{dbSet 0 "EconomySettings" $sdict}}
	{{end}}
{{end}}
