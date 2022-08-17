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
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}

{{/* Configures economy settings */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a := sdict .Value}}
    {{$symbol := $a.symbol}}
	{{with (dbGet 0 "store")}}
		{{$info := sdict .Value}}
		{{$items := sdict}}
		{{if ($info.Get "Items")}}
			{{$items = sdict ($info.Get "Items")}}
			{{$entry := cslice}}
			{{if $items}}
				{{range $k,$v := $items}}
					{{$item := $k}}
					{{$desc := $v.desc }}
					{{$price := $v.price | humanizeThousands}}
					{{$qty := ""}}
					{{if $v.qty}}
						{{$qty = $v.qty}}
						{{if not (reFind "infinite" (lower (toString $qty)))}}
							{{$qty = toInt $qty | humanizeThousands}}
						{{else}}
							{{$qty = "Infinite"}}
						{{end}}
					{{else}}
						{{$qty = "Infinite"}}
					{{end}}
					{{$entry = $entry.Append (sdict "Name" $item "value" (joinStr "\n" (print "Description: " $desc) (print "Price: " $symbol $price) (print "Quantity:" $qty)) "inline" false)}}
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
				{{$embed.Set "fields" $entry}}
				{{$embed.Set "color" $successColor}}
				{{$embed.Set "footer" (sdict "text" (print "Page: " $page))}}
			{{end}}
		{{end}}
	{{else}}
		{{$embed.Set "description" (print "The shop is empty :(\nPlease add some items with ``")}}
		{{$embed.Set "color" $errorColor}}
	{{end}}
{{else}}
    {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}