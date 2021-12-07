{{/* Add this to your join message to set the default starter-balance */}}

{{$startBalance := ""}}
{{$a :=""}}
{{with (dbGet 0 "EconomySettings")}}
	{{$a = sdict .Value}}
	{{$startBalance = (toInt $a.startBalance)}}
	{{dbSet .User.ID "EconomyInfo" (sdict "cash" $startBalance "bank" 0}}
{{end}}
