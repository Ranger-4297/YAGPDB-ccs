{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Reaction`
	Trigger: `Added reaction only`

	©️ Ranger 2020-Present
	GNU, GPLV3 License

	Made with love, support me using https://ko-fi.com/rhykerwells
*/}}


{{/* Configuration values start */}}
{{$verifyingRole := 1191469975416864818}}		{{/* roleID of VERIFYING role */}}
{{$enrollingRole := 1190603830065373224}}		{{/* roleID of ENROLLING role */}}
{{$onboardingRole := 1189413709454520360}}		{{/* roleID of onboarding role */}}
{{$serverRole := 1185703179254505582}}			{{/* roleID of SERVER role */}}
{{$allianceRole := 1185727965397516379}}		{{/* roleID of ALLIANNCE role */}}
{{$nameRole := 1185703261613863013}}			{{/* roleID of NAME role */}}
{{$rankRole := 1190609302218620938}}			{{/* roleID of RANK role */}}
{{$verifyRole := 1185731311927828540}}			{{/* roleID of VERIFIED role */}}
{{$languageChannel := 1187971295397281862}}		{{/* channelID of LANGUAGE channel*/}}
{{$guideChannel := 1189408420714004611}}		{{/* channelID of GUIDE channel*/}}
{{$serverChannel := 1185342810312945804}}		{{/* channelID of SERVER channel*/}}
{{$rankChannel := 1185723603086491648}}			{{/* channelID of RANK channel */}}
{{$rulesChannel := 1185730875644706826}}        {{/* channelID of RULES channel */}}
{{$publicChat := 1191231631726493706}}			{{/* channelID of PUBLIC CHAT channel*/}}
{{$r1Role := 1185724333654540339}}				{{/* roleID of R1 role */}}
{{$r2Role := 1185724406002094190}}				{{/* roleID of R2 role */}}
{{$r3Role := 1185724444698751067}}				{{/* roleID of R3 role */}}
{{$r4Role := 1185724479138173049}}				{{/* roleID of R4 role */}}
{{$r5Role := 1185724508322148363}}				{{/* roleID of R5 role */}}
{{/* Configuration values end */}}

{{/* response */}}
{{sleep 1}}
{{$languages := cslice "english" "spanish" "french" "russian" "chinese" "arabic" "korean" "german" "vietnamese" "japanese" "turkish" "portugese" "malaysian" "fillipino" "ukranian" "indonesian" "greek" "dutch" "italian" "romanian" "danish" "polish" "hebrew"}}
{{$language := "english"}}
{{range (getMember .User.ID).Roles}}
	{{if in $languages (getRole .).Name}}
		{{$language = (getRole .).Name}}
		{{break}}
	{{end}}
{{end}}
{{$language = (dbGet 0 "languageDB").Value.Get $language}}
{{if eq .Channel.ID $languageChannel}}
	{{if $cd := dbGet .User.ID "languageCooldownTime"}}
		{{if $cd := dbGet .User.ID "languageCooldown"}}
			{{$m := sendMessageNoEscapeRetID $languageChannel (print (reReplace `<@!>` (reReplace `<T>` ($language.Get 23) (print "<t:" $cd.ExpiresAt.Unix ":R>")) .User.Mention))}}
			{{deleteMessage $languageChannel $m 60}}
			{{return}}
		{{end}}
		{{$x := dbIncr .User.ID "languageReactionCount" 1}}
		{{if eq (toInt (dbGet .User.ID "languageReactionCount").Value) 3}}
			{{dbDel .User.ID "languageReactionCount"}}
			{{dbSetExpire .User.ID "languageCooldown" true 900}}
		{{end}}
	{{else}}
		{{dbSet .User.ID "languageReactionCount" 1}}
		{{dbSetExpire .User.ID "languageCooldownTime" true 600}}
	{{end}}
	{{addRoleID $enrollingRole}}
	{{sleep 1}}
	{{$gm := sendMessageNoEscapeRetID $guideChannel (print (reReplace "<@!>" (reReplace `<#L>` (reReplace `<#S>` ($language.Get 20) (printf "<#%d>" $serverChannel)) (printf "<#%d>" 1187971295397281862)) .User.Mention))}}
	{{addMessageReactions $guideChannel $gm "✅"}}
	{{$m := sendMessageNoEscapeRetID $guideChannel (print (reReplace `<@!>` ($language.Get 15) .User.Mention))}}
	{{deleteMessage $guideChannel $m 60}}
{{else if eq .Channel.ID $guideChannel}}
	{{if $cd := dbGet .User.ID "guideCooldownTime"}}
		{{if $cd := dbGet .User.ID "guideCooldown"}}
			{{$m := sendMessageNoEscapeRetID $guideChannel (print (reReplace `<@!>` (reReplace `<T>` ($language.Get 23) (print "<t:" $cd.ExpiresAt.Unix ":R>")) .User.Mention))}}
			{{deleteMessage $guideChannel $m 60}}
			{{return}}
		{{end}}
		{{$x := dbIncr .User.ID "guideReactionCount" 1}}
		{{if eq (toInt (dbGet .User.ID "guideReactionCount").Value) 3}}
			{{dbDel .User.ID "guideReactionCount"}}
			{{dbSetExpire .User.ID "guideCooldown" true 900}}
		{{end}}
	{{else}}
		{{dbSet .User.ID "guideReactionCount" 1}}
		{{dbSetExpire .User.ID "guideCooldownTime" true 600}}
	{{end}}
	{{addRoleID $onboardingRole}}
	{{$m := sendMessageNoEscapeRetID $guideChannel (print (reReplace `<#>` (reReplace `<@!>` ($language.Get 14) .User.Mention) (printf "<#%d>" $serverChannel)))}}
	{{deleteMessage $guideChannel $m 60}}
	{{$m2 := sendMessageNoEscapeRetID $serverChannel (print (reReplace `<@!>` ($language.Get 16) .User.Mention))}}
	{{deleteMessage $serverChannel $m2 60}}
{{else if eq .Channel.ID $rankChannel}}
	{{if $cd := dbGet .User.ID "rankCooldownTime"}}
		{{$m := sendMessageNoEscapeRetID $rankChannel (print (reReplace `<@!>` (reReplace `<T>` ($language.Get 23) (print "<t:" $cd.ExpiresAt.Unix ":R>")) .User.Mention))}}
		{{deleteMessage $rankChannel $m 5}}
		{{return}}
	{{else}}
		{{dbSetExpire .User.ID "rankCooldownTime" true 5}}
	{{end}}
	{{$rankedRole := ""}}
	{{with .Reaction.Emoji.Name}}
		{{if eq . "1️⃣"}}
			{{$rankedRole = $r1Role}}
		{{else if eq . "2️⃣"}}
			{{$rankedRole = $r2Role}}
		{{else if eq . "3️⃣"}}
			{{$rankedRole = $r3Role}}
		{{else if eq . "4️⃣"}}
			{{$rankedRole = $r4Role}}
		{{else if eq . "5️⃣"}}
			{{$rankedRole = $r5Role}}
		{{end}}
	{{end}}
	{{$rankRoles := cslice $r1Role $r2Role $r3Role $r4Role $r5Role}}
	{{range $rankRoles}}
		{{if in (getMember $.User.ID).Roles .}}
			{{removeRoleID .}}
		{{end}}
	{{end}}
	{{addRoleID $rankedRole}}
	{{addRoleID $rankRole}}
	{{if $cd := dbGet .User.ID "Rrcooldown3"}}
		{{return}}
	{{end}}
	{{$m := sendMessageNoEscapeRetID $rankChannel (print (reReplace `<#>` (reReplace `<@!>` ($language.Get 19) .User.Mention) (printf "<#%d>" $rulesChannel)))}}
	{{deleteMessage $rankChannel $m 60}}
	{{$m2 := sendMessageNoEscapeRetID $rulesChannel (print (reReplace `<@!>` ($language.Get 21) .User.Mention))}}
	{{addMessageReactions $rulesChannel $m2 "✅"}}
{{else if eq .Channel.ID $rulesChannel}}
	{{if $cd := dbGet .User.ID "rulesCooldownTime"}}
		{{if $cd := dbGet .User.ID "rulesCooldown"}}
			{{$m := sendMessageNoEscapeRetID $rulesChannel (print (reReplace `<@!>` (reReplace `<T>` ($language.Get 23) (print "<t:" $cd.ExpiresAt.Unix ":R>")) .User.Mention))}}
			{{deleteMessage $rulesChannel $m}}
			{{return}}
		{{end}}
		{{$x := dbIncr .User.ID "rulesReactionCount" 1}}
		{{if eq (toInt (dbGet .User.ID "rulesReactionCount").Value) 2}}
			{{dbDel .User.ID "rulesReactionCount"}}
			{{dbSetExpire .User.ID "rulesCooldown" true 900}}
		{{end}}
	{{else}}
		{{dbSet .User.ID "rulesReactionCount" 1}}
		{{dbSetExpire .User.ID "rulesCooldownTime" true 600}}
	{{end}}
	{{addRoleID $verifyRole}}
	{{range (cslice $verifyingRole $serverRole $allianceRole $nameRole $enrollingRole $onboardingRole $rankRole)}}
		{{if (hasRoleID .)}}
			{{removeRoleID .}}
		{{end}}
	{{end}}
	{{$roleMap := dict 1191468021676523640 1191468021676523640 1187973833257390140 1191153131237683240 1189306777494368306 1191468395435135117 1191889113335480413 1191889229014388809 1189336979406589952 1191468814727118868 1191889805676650606 1191889843349893120 1189258793230536766 1191468603111911464 1189258827930009640 1191468616290402324 1191889900254015598 1191889903521374368 1189314049129861220 1191468620350509156 1191890501301977149 1191890505290756116 1191890508969152572 1191890512056164475}}
	{{range $old, $new := $roleMap}}
		{{- if (hasRoleID $old)}}
			{{- removeRoleID $old}}
			{{- addRoleID $new}}
		{{- end -}}
	{{end}}
	{{sendMessageNoEscape $publicChat (print (reReplace `<@!>` ($language.Get 18) .User.Mention))}}
{{end}}