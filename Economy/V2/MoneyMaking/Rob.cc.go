{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(rob|steal|mug|con)(\s+|\z)`

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

{{/* Rob a user */}}

{{/*
If the user isn't in the economy database 
It'll automatically add them
--
If there is no setting values
You'll be asked to set it up with default values
You can change these later
*/}}

{{with (dbGet 0 "EconomySettings")}}
    {{$a := sdict .Value}}
    {{$failRate := $a.failRate}}
    {{$symbol := $a.symbol}}
    {{with $.CmdArgs}}
        {{if index . 0 }}
            {{if index . 0 | getMember}}
                {{$user := getMember (index . 0)}}
                {{$victim := $user.User.ID}}
                {{if eq $victim $userID}}
                    {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're can't rob yourself. Please specify a valid user")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                {{else}}
                    {{if not (dbGet $victim "EconomyInfo")}}
                        {{dbSet $victim "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
                    {{end}}
                    {{with (dbGet $victim "EconomyInfo")}}
                        {{$a = sdict .Value}}
                        {{$victimsCash := (toInt $a.cash)}}
                        {{if not (dbGet $userID "EconomyInfo")}}
                            {{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
                        {{end}}
                        {{with (dbGet $userID "EconomyInfo")}}
                            {{$a = sdict .Value}}
                            {{$yourCash := (toInt $a.cash)}}
                            {{if eq $victimsCash 0}}
                                {{sendMessage nil (cembed
                                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                                            "description" (print "You're unable to rob <@!" $victim "> to a value below `0`")
                                            "color" 0x00ff8b
                                            "timestamp" currentTime
                                            )}}
                            {{else}}
                                {{$amount := (randInt $victimsCash)}} {{/* Amount stolen from victim */}}
                                {{$victimsNewCash := $amount | sub $victimsCash}} {{/* Amout victim will have after being robbed */}}
                                {{$yourNewCash := $amount | add $yourCash}} {{/* Amount you will have after robbing vitim */}}
                                {{sendMessage nil (cembed
                                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))
                                            "description" (print "You robbed " $symbol $amount " from <@!" $victim ">")
                                            "color" $successColor
                                            "timestamp" currentTime
                                            )}}
                                {{$sdict := (dbGet $victim "EconomyInfo").Value}}
                                {{$sdict.Set "cash" $victimsNewCash}}
                                {{dbSet $victim "EconomyInfo" $sdict}}
                                {{$sdict = (dbGet $userID "EconomyInfo").Value}}
                                {{$sdict.Set "cash" $yourNewCash}}
                                {{dbSet $userID "EconomyInfo" $sdict}}
                            {{end}}
                        {{end}}
                    {{end}}
                {{end}}
            {{end}}
        {{else}}
            {{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "Invalid `User` argument provided.\nSyntaxt is ` " $.Cmd " <User:Mention/ID>`")
            "color" $errorColor
            "timestamp" currentTime
            )}}
        {{end}}
    {{else}}
        {{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "No `User` argument provided.\nSyntaxt is `" $.Cmd " <User:Mention/ID>`")
            "color" $errorColor
            "timestamp" currentTime
            )}}
    {{end}}
{{else}}
    {{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")
            "color" $errorColor
            "timestamp" currentTime
            )}}
{{end}}