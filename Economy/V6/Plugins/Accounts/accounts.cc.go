{{/*
        Made by Ranger (765316548516380732)

        Trigger Type: `Regex`
        Trigger: ``

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

{{/* Buy item */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
    {{$a := sdict .Value}}
    {{$symbol := $a.symbol}}
    {{with $db := (dbGet 0 "accounts")}}
        {{$a := sdict .Value}}
        {{$currentAccounts := cslice}}
        {{if $db}}
            {{range $k,$v:= $a}}
                {{- $currentAccounts = $currentAccounts.Append $k -}}
            {{end}}
        {{end}}
        {{with .CmdArgs}}
            {{if index $.CmdArgs 0}}
                {{$option := (index $.CmdArgs 0) | lower}}
                {{if eq $option "create"}}
                    {{dbSet "accounts" (sdict (toString $.User.ID) (sdict "accountSettings" (sdict "whitelistedUsers" (cslice) "withdrawLimit" 500 "accountExpiry" 604800 "transferCost" 300) "accountBalance" 500))}}
                {{else if eq $option "set"}}
                    {{if gt (len $.CmdArgs) 1}}
                        {{$val := (lower (toString (index $.CmdArgs)))}}
                        {{if eq $val "wl" "whitelist"}}
                            {{if gt (len $.CmdArgs) 2}}
                                {{$ID := (joinStr " " (slice $.CmdArgs 2))}}
                                {{$users := cslice}}
                                {{range (reFindAll \d{17,19} $ID)}}
                                    {{- $users = $users.Append .}}
                                {{- end}} 
                                {{$account := $a.Get (toString $.User.ID)}}
                                {{$settings := $account.Get "accountSettings"}}
                                {{$settings.Set "whitelistedUsers" ($whitelist.AppendSlice $users)}}
                                {{$account.Set "accountSettings" $settings}}
                                {{$a.Set (toString $.User.ID) $account}}
                                {{dbSet 0 "accounts" $a}}
                                {{$embed.Set "description" (print "Users now added to whitelist.")}}
                                {{$embed.Set "color" $successColor}}
                            {{else}}
                                {{$embed.Set "description" (print "No user provided.\nPlease ensure that you use UserIDs")}}
                                {{$embed.Set "color" $errorColor}}
                            {{end}}
                        {{else if eq $val "withdrawlimit" "withlimit" "wl"}}
                            {{if gt (len $.CmdArgs) 2}}
                                {{if (toInt $.CmdArgs 2)}}
                                    {{$account := $a.Get (toString $.User.ID)}}
                                    {{$settings := $account.Get "accountSettings"}}
                                    {{$settings.Set "withdrawLimit" (toInt $.CmdArgs 2)}}
                                    {{$account.Set "accountSettings" $settings}}
                                    {{$a.Set (toString $.User.ID) $account}}
                                    {{dbSet 0 "accounts" $a}}
                                    {{$embed.Set "description" (print "New withdraw limit added.")}}
                                    {{$embed.Set "color" $successColor}}
                                {{else}}
                                {{$embed.Set "description" (print "Invalid amount provided.")}}
                                {{$embed.Set "color" $errorColor}}
                            {{end}}
                            {{else}}
                                {{$embed.Set "description" (print "No amount provided.")}}
                                {{$embed.Set "color" $errorColor}}
                            {{end}}
                        {{else}}
                            {{$embed.Set "description" (print "Invalid option provided.")}}
                            {{$embed.Set "color" $errorColor}}
                        {{end}}
                    {{else}}
                        {{$msg.Set "description" (print "No option argument passed.\nSyntax is: `" $.Cmd " <Option> <Values>`" $syntax)}}
                        {{$embed.Set "color" $errorColor}}
                    {{end}}
                {{else if eq $option "list"}}
                    {{if }}
            {{end}}
        {{else}}
            {{$msg.Set "description" (print "No option argument passed.\nSyntax is: `" $.Cmd " <Option> <Values>`" $syntax)}}
            {{$embed.Set "color" $errorColor}}
        {{end}}
    {{end}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}