{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(work|job|get-?paid|(commit-?)?crime|rob|steal)(\s+|\z)`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := (index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 )}}

{{/* Work, crime, rob */}}

{{/*
If the users aren't in the economy database 
It'll automatically add them
--
If there is no setting values
You'll be asked to set it up with default values
You can change these later
*/}}
{{/* Set DB Call counter at 0 */}}

{{/* Resoponse */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a := sdict .Value}}
    {{$min := $a.min}}
    {{$max := $a.max}}
    {{$symbol := $a.symbol}}
    {{$robCooldown := $a.robCooldown | toInt}}
    {{$crimeCooldown := $a.crimeCooldown | toInt}}
    {{if not (dbGet $userID "EconomyInfo")}}
		{{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
        {{$embed.Set "description" (print "You were not in the economy database....adding you now\nPlease try again")}}
        {{$embed.Set "color" $errorColor}}
	{{else}}
        {{with $.Cmd}}
            {{$cmd := $.Cmd | toString | lower}}
            {{if (reFind `(work|job|get-?paid|labor)` $cmd)}}
                {{with (dbGet $userID "EconomyInfo")}}
                    {{$a = sdict .Value}}
                    {{$cash := $a.cash}}
                    {{$workPay := randInt $min $max}}
                    {{$newCashBalance := $cash | add $workPay}}
                    {{$embed.Set "description" (print "You decided to work today! You got paid a hefty " $symbol $workPay)}}
                    {{$embed.Set "color" 0x00ff7b}}
                    {{$sdict := (dbGet $userID "EconomyInfo").Value}}
                    {{$sdict.Set "cash" $newCashBalance}}
                    {{dbSet $userID "EconomyInfo" $sdict}}
                {{end}}
            {{else if (reFind `(commit-?)?crime` $cmd)}}
                {{if $cooldown := dbGet $.User.ID "crimeCooldown"}}
                    {{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
                    {{$embed.Set "color" $errorColor}}
                {{else}}
                    {{with (dbGet $userID "EconomyInfo")}}
                        {{$a = sdict .Value}}
                        {{$cash := $a.cash}}
                        {{$amount := randInt $min $max}}
                        {{$newCash := ""}}
                        {{$int := randInt 1 3}}
                        {{if eq $int 1}}
                            {{$newCash = $cash | add $amount}}
                            {{$embed.Set "description" (print "You broke the law for a pretty penny! You made " $symbol $amount " in your crime spree today")}}
                            {{$embed.Set "color" $successColor}}
                        {{else}}
                            {{$newCash = $amount | sub $cash}}
                            {{$embed.Set "description" (print "You broke the law trying to commit a felony! You were arrested and lost " $symbol $amount " due to your bail.")}}
                            {{$embed.Set "color" $errorColor}}
                        {{end}}
                        {{$sdict := (dbGet $userID "EconomyInfo").Value}}
                        {{$sdict.Set "cash" $newCash}}
                        {{dbSet $userID "EconomyInfo" $sdict}}
                    {{end}}
                {{end}}
            {{else if (reFind `(rob|steal|mug|con)` $cmd)}}
                {{with $.CmdArgs}}
                    {{if (index . 0)}}
                        {{if (index . 0) | getMember}}
                            {{$user :=  (index . 0) | getMember}}
                            {{$victim := $user.User.ID}}
                            {{if eq $victim $userID}}
                                {{$embed.Set "description" (print "You can't rob yourself. Please specify a valid user")}}
                                {{$embed.Set "color" $errorColor}}
                            {{else}}
                                {{if $cooldown := dbGet $.User.ID "robCooldown"}}
                                    {{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
                                    {{$embed.Set "color" $errorColor}}
                                {{else}}
                                    {{dbSetExpire $.User.ID "robCooldown" "cooldown" $robCooldown}}
                                    {{with (dbGet $victim "EconomyInfo")}}
                                        {{$a = sdict .Value}}
                                        {{$victimsCash := $a.cash | toInt}}
                                        {{with (dbGet $userID "EconomyInfo")}}
                                            {{$a = sdict .Value}}
                                            {{$yourCash := $a.cash | toInt}}
                                            {{if eq $victimsCash 0}}
                                                {{$embed.Set "description" (print "You're unable to rob <@!" $victim "> to a value below `0`")}}
                                                {{$embed.Set "color" 0x00ff8b}}
                                            {{else}}
                                                {{$amount := (randInt $victimsCash)}} {{/* Amount stolen from victim */}}
                                                {{$victimsNewCash := $amount | sub $victimsCash}} {{/* Amout victim will have after being robbed */}}
                                                {{$yourNewCash := $amount | add $yourCash}} {{/* Amount you will have after robbing vitim */}}
                                                    {{$embed.Set "description" (print "You robbed " $symbol $amount " from <@!" $victim ">")}}
                                                    {{$embed.Set "color" $successColor}}
                                                {{$sdict := (dbGet $victim "EconomyInfo").Value}}
                                                {{$sdict.Set "cash" $victimsNewCash}}
                                                {{dbSet $victim "EconomyInfo" $sdict}}
                                                {{$sdict = (dbGet $userID "EconomyInfo").Value}}
                                                {{$sdict.Set "cash" $yourNewCash}}
                                                {{dbSet $userID "EconomyInfo" $sdict}}
                                            {{end}}
                                        {{end}}
                                    {{else}}
                                        {{dbSet $victim "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
                                        {{$embed.Set "description" (print "<@!" $victim "> doesn't have any money for you to rob!")}}
                                        {{$embed.Set "color" $errorColor}}
                                    {{end}}
                                {{end}}
                            {{end}}
                        {{else}}
                            {{$embed.Set "description" (print "Invalid `User` argument provided.\nSyntaxt is `" $.Cmd " <User:Mention/ID>`")}}
                            {{$embed.Set "color" $errorColor}}
                        {{end}}
                    {{end}}
                {{else}}
                    {{$embed.Set "description" (print "No `User` argument provided.\nSyntaxt is `" $.Cmd " <User:Mention/ID>`")}}
                    {{$embed.Set "color" $errorColor}}
                {{end}}
            {{end}}
        {{end}}
    {{end}}
{{else}}
    {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}