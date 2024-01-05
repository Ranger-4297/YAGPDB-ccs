{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Command`
	Trigger: `ldb3`

	©️ Ranger 2020-Present
	GNU, GPLV3 License

	Made with love, support me using https://ko-fi.com/rhykerwells
*/}}

{{$languageDB := (dbGet 0 "languageDB").Value}}
{{$languages := sdict
"turkish" (dict
	1 "İttifak yok mu? \"skip\" yazın"
	2 "Mükemmel! <@!> Sunucu Etiketinizi yeni aldınız! Şimdi <#>'a ilerleyin <1/5>"
	3 "<@!> İttifak etiketinizi buraya girin <2/5>"
	4 "Bu havalı! <@!> artık bir İttifak Etiketine sahipsiniz! Devam edin ve <#>'a ilerleyin <2/5>"
	5 "<@!> Karakter Oyunu Adınızı buraya yazın! <3/5>"
	6 "Mükemmel! <@!>, görünen adınızı Karakter Oyunu Adınızla güncellediniz! İki adım daha, <#>'a ilerleyin <3/5>"
	7 "<@!> oyun içi ittifak rütbenizi seçin <4/5>"
	8 "Botu engelleyen kullanıcıya tepki eklenemez. Takma ad güncellendi"
	9 "Lütfen 3-4 haneli bir etiket girin"
	10 "Lütfen sayısal bir etiket girin"
	11 "Lütfen 3-4 karakterlik bir ittifak etiketi girin"
	12 "Lütfen özel karakterler kullanmayın"
	13 "Lütfen 3-15 karakterlik bir kullanıcı adı girin"
	14 "Hadi başlayalım <@!>, <#> <0/5>'e ilerleyin"
	15 "Pekala <@!>, sizi katılmaya ve doğrulanmaya hazırlayalım!"
	16 "İşte başlıyoruz <@!>, daha yeni başladık! <0/5>"
	17 "Nihayet!! <@!> Buradasınız!! Diğer tarafa son adım!! <5/5>"
	18 "Hahaha! Başardın! <@!> Diğer tarafa hoş geldiniz!"
	19 "Tebrikler! <@!> şimdi son adıma hazırlanın⁠ <#> <5/5>"
	20 "### KG World - WOR'a <@!> hoş geldiniz 🌀\n> Lütfen <@!765316548516380732> tarafından kodlanan 5 Doğrulama Adımını izleyin\n> - Bu bir dakikadan az sürecektir.\n> - Değiştirmek için dil 👉 <#L>\n\n~ *Diğer taraf sizi bekliyor* 💫\n### **👇 Başlamak için ✅'a tıklayın.**"
	21 "**<@!> Este é o WOR, um KG World Server privado.**\n - Respeitamos a concorrência.\n - Damos as boas-vindas a todos.\n - Seja respeitoso.\n - Mais importante ainda, vamos nos divertir.\n### **Estamos esperando por você do outro lado** 🌀\n\n### **Reaja com um ✅ para prosseguir.**"
	22 "<@!> Mevcut İttifak Sıralamanızı seçin.\n\n### 👉 <#>"
	23 "<@!> Bu reaksiyon için bir bekleme süresine tabi tutuldunuz.\nBekleme süresi **<T>** ile sona eriyor"
)
"portugese" (dict
	1 "Sem aliança? Digite \"skip\""
	2 "Incrível! <@!> Você acabou de receber sua etiqueta de servidor! Agora prossiga para <#> <1/5>"
	3 "<@!> Digite sua tag de aliança aqui <2/5>"
	4 "Isso é legal! <@!> agora você tem uma etiqueta de aliança! Continue e prossiga para <#> <2/5>"
	5 "<@!> digite o nome do jogo do seu personagem aqui! <3/5>"
	6 "Perfeito! <@!> você atualizou seu nome de exibição para o nome do jogo do seu personagem! Mais duas etapas, prossiga para <#> <3/5>"
	7 "<@!> selecione a classificação da sua aliança no jogo. <4/5>"
	8 "Não é possível adicionar reação ao usuário que bloqueou o bot. Apelido atualizado"
	9 "Insira uma etiqueta de 3 a 4 dígitos"
	10 "Por favor insira uma tag numérica"
	11 "Por favor, insira uma tag de aliança de 3 a 4 caracteres"
	12 "Por favor, não use caracteres especiais"
	13 "Insira um nome de usuário de 3 a 15 caracteres"
	14 "Vamos começar <@!>, prossiga para <#> <0/5>"
	15 "Tudo bem <@!>, vamos prepará-lo para ser integrado e verificado!"
	16 "Aqui está <@!>, acabamos de começar! <0/5>"
	17 "Finalmente!! <@!> Você está aqui!! Passo final para o outro lado!! <5/5>"
	18 "Hahaha! Você conseguiu! <@!> Bem-vindo ao outro lado!"
	19 "Bom trabalho! <@!> agora prepare-se para a etapa final⁠ <#> <5/5>"
	20 "### Bem-vindo <@!> ao KG World - WOR 🌀\n> Siga as 5 etapas de verificação, codificadas por <@!765316548516380732>\n> - Isso levará menos de um minuto.\n> - Para alterar idioma 👉 <#L>\n\n~ *O outro lado espera por você* 💫\n### **👇 Clique em ✅ para começar.**"
	21 "**<@!> Este é o WOR, um KG World Server privado.**\n - Respeitamos a concorrência.\n - Damos as boas-vindas a todos.\n - Seja respeitoso.\n - Mais importante ainda, vamos nos divertir.\n### **Estamos esperando por você do outro lado** 🌀\n\n### **Reaja com um ✅ para prosseguir.**"
	22 "<@!> Selecione sua classificação de aliança atual.\n\n### 👉 <#>"
	23 "<@!> Você foi colocado em um tempo de espera para esta reação.\nO tempo de espera termina em **<T>**"
)
}}
{{range $k, $v := $languages}}
	{{$languageDB.Set $k $v}}
{{end}}
{{dbSet 0 "languageDB" $languageDB}}