{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)((reset|remove)-?user)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix }}

{{/* Reset user */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{$perms := split (index (split (exec "viewperms") "\n") 2) ", "}}
{{if or (in $perms "Administrator") (in $perms "ManageServer")}}
	{{with (dbGet 0 "EconomySettings")}}
		{{$a := sdict .Value}}
		{{$symbol := $a.symbol}}
		{{with $.CmdArgs}}
            {{if index . 0 | getMember}}
                {{$user := getMember (index . 0)}}
				{{$user = $user.User}}
                {{if (dbGet $user.ID "cash")}}
                    {{dbDel $user.ID "cash"}}
                {{end}}
                {{if (dbGet $user.ID "userEconData")}}
                    {{dbDel $user.ID "userEconData"}}
                {{end}}
                {{$bank := (dbGet 0 "bank").Value.Get (toString $user.ID)}}
                {{if $bank}}
                    {{$bankDB := (dbGet 0 "bank").Value}}
                    {{$bankDB.Del (toString $user.ID)}}
                {{end}}
                {{$embed.Set "description" (print "Successfully removed the user from the economy system.")}}
                {{$embed.Set "color" $successColor}}
            {{else}}
                {{$embed.Set "description" (print "Invalid `User` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID>`")}}
                {{$embed.Set "color" $errorColor}}
            {{end}}
        {{else}}
			{{$embed.Set "description" (print "No `User` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID>`")}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
    {{else}}
        {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "server-set default`")}}
        {{$embed.Set "color" $errorColor}}
    {{end}}
{{else}}
	{{$embed.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}